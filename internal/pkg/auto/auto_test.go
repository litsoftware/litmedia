package auto

import (
	"github.com/litsoftware/litmedia/internal/pkg/config"
	"testing"
)

func TestInit(t *testing.T) {
	if config.GetString("app.app_env") != "test" {
		t.Errorf("evn file is not wrong")
	}
}
