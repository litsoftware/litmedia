package orm

import (
	"context"
	"github.com/litsoftware/litmedia/internal/ent"
)

var client *ent.Client

func GetClient() *ent.Client {
	return client
}

func GetTx(ctx context.Context) (*ent.Tx, error) {
	return client.Tx(ctx)
}
