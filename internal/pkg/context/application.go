package context

import (
	"github.com/litsoftware/litmedia/internal/pkg/config"
	"github.com/litsoftware/litmedia/pkg/logger"
	"github.com/spf13/viper"
	"time"
)

type ApplicationContext struct {
	Logger logger.ILogger
	Conf   *viper.Viper

	AppReady bool
	Opts     map[string]interface{}
}

func (a *ApplicationContext) IsRunInProductionMode() bool {
	return a.Conf.GetString("app.app_env") == "production"
}

func (*ApplicationContext) Deadline() (deadline time.Time, ok bool) {
	return
}

func (*ApplicationContext) Done() <-chan struct{} {
	return nil
}

func (*ApplicationContext) Err() error {
	return nil
}

func (*ApplicationContext) Value(key interface{}) interface{} {
	return nil
}

func NewAppContext() *ApplicationContext {
	ctx := new(ApplicationContext)

	ctx.Logger = logger.GetLogger()
	ctx.Conf = config.GetConfig()

	return ctx
}
