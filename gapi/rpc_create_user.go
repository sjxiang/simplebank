package gapi

import (
	"context"
	"time"

	"github.com/hibiken/asynq"
	
	db "github.com/sjxiang/simplebank/db/sqlc"
	"github.com/sjxiang/simplebank/pb"
	"github.com/sjxiang/simplebank/util"
	"github.com/sjxiang/simplebank/val"
	"github.com/sjxiang/simplebank/worker"

	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	// 校验请求字段
	violations := validateCreateUserRequest(req)
	// 至少有一个无效字段
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}

	// 密码哈希
	hashedPassword, err := util.HashPassword(req.GetPassword())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to hash password: %s", err)
	}

	// 事务（db 有误，回滚；redis 有误，回滚）
	arg := db.CreateUserTxParams{
		CreateUserParams: db.CreateUserParams{
			Username:       req.GetUsername(),  // getter
			HashedPassword: hashedPassword,
			FullName:       req.GetFullName(),
			Email:          req.GetEmail(),
		},

		// 回调，发送验证邮件
		AfterCreate: func(user db.User) error {
			taskPayload := &worker.PayloadSendVerifyEmail{
				Username: user.Username,
			}
			opts := []asynq.Option{
				asynq.MaxRetry(10),                 // 重试，最多 10 次
				asynq.ProcessIn(10 * time.Second),  // 延迟，正常提交，但执行延缓 10 秒
				asynq.Queue(worker.QueueCritical),  // 优先级
			}

			return server.taskDistributor.DistributeTaskSendVerifyEmail(ctx, taskPayload, opts...)
		},
	}

	txResult, err := server.store.CreateUserTx(ctx, arg)
	if err != nil {
		if db.ErrorCode(err) == db.UniqueViolation {
			return nil, status.Errorf(codes.AlreadyExists, err.Error())  // 对应 http 409，已经被占用
		}
		
		return nil, status.Errorf(codes.Internal, "failed to create user: %s", err)
	}

	// 响应
	rsp := &pb.CreateUserResponse{
		User: convertUser(txResult.User),
	}
	return rsp, nil
}

func validateCreateUserRequest(req *pb.CreateUserRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	// 请求是否赋值，string 如果没有赋值，则传递零值 ""；如果赋值空字符串，也是传递 ""，怎么区分
	// getter 有啥意义呢
	if err := val.ValidateUsername(req.GetUsername()); err != nil {
		violations = append(violations, fieldViolation("username", err))
	}

	if err := val.ValidatePassword(req.GetPassword()); err != nil {
		violations = append(violations, fieldViolation("password", err))
	}

	if err := val.ValidateFullName(req.GetFullName()); err != nil {
		violations = append(violations, fieldViolation("full_name", err))
	}

	if err := val.ValidateEmail(req.GetEmail()); err != nil {
		violations = append(violations, fieldViolation("email", err))
	}

	return violations
}
