package postgres

import (
	"context"
	"github.com/jackc/pgx/v5"
)

// WithInTransaction - открывает транзакцию
func (s *DB) WithInTransaction(ctx context.Context, tFunc func(ctx context.Context) error) error {
	tx, err := s.Pool.Begin(ctx)
	if err != nil {
		s.logger.Errorf("Postgres (s.Pool.Begin): %v", nil, err)

		if err := tx.Rollback(ctx); err != nil {

			s.logger.Errorf("Postgres (tx.Rollback): %v", nil, err)
			return err
		}

		return err
	}

	err = tFunc(injectTx(ctx, tx))
	if err != nil {
		s.logger.Errorf("Postgres (tFunc(injectTx(ctx, tx))): %v", nil, err)

		if err := tx.Rollback(ctx); err != nil {
			s.logger.Errorf("Postgres (tx.Rollback): %v", nil, err)
			return err
		}

		return err
	}

	if err := tx.Commit(ctx); err != nil {
		s.logger.Errorf("Postgres (s.Pool.Begin): %v", nil, err)

		return err
	}

	return nil
}

// injectTx - сохраняет результат транзакции в ctx
func injectTx(ctx context.Context, tx pgx.Tx) context.Context {
	return context.WithValue(ctx, tx, tx)
}
