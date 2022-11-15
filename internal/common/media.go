package common

import (
	"context"
	"github.com/litsoftware/litmedia/internal/ent"
	"github.com/litsoftware/litmedia/internal/ent/media"
	"github.com/litsoftware/litmedia/internal/pkg/config/orm"
)

func GetMediaById(ctx context.Context, id string, app *ent.App) (*ent.Media, error) {
	return orm.GetClient().Media.Query().
		Where(media.Sn(id)).
		Where(media.AppID(app.ID)).
		First(ctx)
}
