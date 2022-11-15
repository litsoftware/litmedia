package common

import (
	"context"
	"github.com/litsoftware/litmedia/internal/ent"
	app2 "github.com/litsoftware/litmedia/internal/ent/app"
	"github.com/litsoftware/litmedia/internal/pkg/config"
	"github.com/litsoftware/litmedia/internal/pkg/config/orm"
	"github.com/litsoftware/litmedia/pkg/aes"
)

func DecryptAppInfo(app *ent.App) error {
	key := []byte(config.GetString("app.system_aes"))
	appPub, err := aes.Decrypt(key, app.EncryptedAppPublicKey)
	if err != nil {
		return err
	}

	appPriv, err := aes.Decrypt(key, app.EncryptedAppPrivateKey)
	if err != nil {
		return err
	}

	rsaPub, err := aes.Decrypt(key, app.EncryptedOperatorRsaPublicKey)
	if err != nil {
		return err
	}

	app.EncryptedAppPublicKey = appPub
	app.EncryptedAppPrivateKey = appPriv
	app.EncryptedOperatorRsaPublicKey = rsaPub

	return nil
}

func EncryptAppInfo() {

}

func GetAppById(ctx context.Context, id int) (*ent.App, error) {
	app, err := orm.GetClient().App.Query().Where(app2.ID(id)).First(ctx)
	if err != nil {
		return nil, err
	}

	err = DecryptAppInfo(app)
	if err != nil {
		return nil, err
	}

	return app, nil
}

func GetAppByAppId(ctx context.Context, AppId string) (*ent.App, error) {
	app, err := orm.GetClient().App.Query().Where(app2.AppID(AppId)).First(ctx)
	if err != nil {
		return nil, err
	}

	err = DecryptAppInfo(app)
	if err != nil {
		return nil, err
	}

	return app, nil
}
