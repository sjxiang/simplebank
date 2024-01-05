package db

import (
	"context"
	"fmt"
)

func (store *SQLStore) execTx(ctx context.Context, fn func(*Queries) error) error {
	// 开启事务
	tx, err := store.connPool.Begin(ctx)
	if err != nil {
		return err
	}

	q := New(tx)

	err = fn(q)
	if err != nil {
		// 事务失败，回滚失败
		if rbErr := tx.Rollback(ctx); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		// 事务失败，回滚成功
		return err
	}

	// 提交事务
	return tx.Commit(ctx)
}
