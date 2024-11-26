package txer

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type (
	RunFunc             func(tx *pgx.Tx) error
	RunValueFunc[T any] func(tx *pgx.Tx) (res T, err error)
)

func noValueFunc(runFunc RunFunc) RunValueFunc[struct{}] {
	return func(tx *pgx.Tx) (res struct{}, err error) {
		return res, runFunc(tx)
	}
}

func InTx(ctx context.Context, db *pgxpool.Pool, run RunFunc) error {
	_, err := InTxWithValue(ctx, db, noValueFunc(run))
	return err
}

func InTxWithOptions(
	ctx context.Context, db *pgxpool.Pool,
	options *pgx.TxOptions,
	run RunFunc,
) error {
	_, err := InTxWithValueAndOptions(ctx, db, options, noValueFunc(run))
	return err
}

func InTxWithValue[T any](
	ctx context.Context,
	db *pgxpool.Pool,
	run RunValueFunc[T],
) (T, error) {
	return InTxWithValueAndOptions(ctx, db, &pgx.TxOptions{
		IsoLevel:   pgx.Serializable,
		AccessMode: pgx.ReadWrite,
	}, run)
}

func InTxWithValueAndOptions[T any](
	ctx context.Context,
	db *pgxpool.Pool,
	options *pgx.TxOptions,
	run RunValueFunc[T],
) (T, error) {
	var empty T

	if options == nil {
		options = &pgx.TxOptions{}
	}
	options.IsoLevel = pgx.Serializable
	tx, err := db.BeginTx(ctx, *options)
	if err != nil {
		return empty, fmt.Errorf("begin tx: %w", err)
	}

	res, err := run(&tx)
	if err != nil {
		if rErr := tx.Rollback(ctx); rErr != nil {
			return empty, errors.Join(err, rErr)
		}
		return empty, err
	}

	return res, tx.Commit(ctx)
}
