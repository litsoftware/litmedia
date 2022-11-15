package test

import (
	"github.com/litsoftware/litmedia/internal/pkg/config"
	_ "github.com/litsoftware/litmedia/pkg/runtime"
	"testing"
)

func TestInit(t *testing.T) {
	if config.GetString("public.app_env") != "test" {
		t.Errorf("evn file is not wrong")
	}
}
