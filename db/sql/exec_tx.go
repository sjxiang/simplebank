package db

import (
	"context"
	"fmt"
	"database/sql"
)

// ExecTx executes a function within a database transaction
func (data *Data) execTx(ctx context.Context, fn func(*sql.Tx) error) error {
	// 开启事务
	tx, err := data.db.BeginTx(ctx, &sql.TxOptions{
		Isolation: sql.LevelDefault,
		ReadOnly:  false,
	})
	if err != nil {
		return err 
	}

	err = fn(tx)
	if err != nil {
		// 事务失败，回滚失败
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		// 事务失败，回滚成功
		return err
	}

	// 提交事务
	return tx.Commit()
}
