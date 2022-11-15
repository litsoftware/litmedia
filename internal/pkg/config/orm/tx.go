package orm

import (
	"context"
	"github.com/litsoftware/litmedia/internal/ent"
	"github.com/pkg/errors"
)

func WithTx(ctx context.Context, client *ent.Client, fn func(tx *ent.Tx) error) (err error) {
	tx, err := client.Tx(ctx)
	if err != nil {
		return
	}

	defer func() {
		if v := recover(); v != nil {
			err = tx.Rollback()
			if err != nil {
				return
			}
		}
	}()

	err = fn(tx)

	if err != nil {
		if rerr := tx.Rollback(); rerr != nil {
			err = errors.Wrapf(err, "rolling back transaction: %v", rerr)
		}
		return
	}

	if err = tx.Commit(); err != nil {
		return errors.Wrapf(err, "committing transaction: %v", err)
	}

	return
}
