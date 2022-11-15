package tzinit

import (
	"github.com/golang-module/carbon"
	_ "github.com/litsoftware/litmedia/pkg/runtime"
	"time"

	"fmt"
	"github.com/litsoftware/litmedia/pkg/path"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

var (
	err error
	v   *viper.Viper
)

func initConfig() {
	v = viper.New()
	v.SetConfigType("toml")
	v.SetConfigName("app")

	p := fmt.Sprintf("%s/%s", path.RootPath(), "configs")
	v.AddConfigPath(p)

	initForTestMode()
	initForCiTestMode()

	err = v.ReadInConfig()
	if err != nil {
		fmt.Printf("parse configuration file err in test %#v\n", err)
		os.Exit(1)
	}

	if v.GetBool("watch") {
		v.WatchConfig()
	}
}

func initForTestMode() {
	if os.Getenv("TESTING_MODE") == "1" {
		v.SetConfigName("app_test")

		pwd := os.Getenv("PWD")
		for i := 0; i < 10; i = i + 1 {
			if strings.HasSuffix(pwd, "media-service") {
				v.AddConfigPath(fmt.Sprintf("%s/%s", pwd, "configs"))
				break
			}

			pwd = filepath.Dir(pwd)
		}
	}
}

func initForCiTestMode() {
	if os.Getenv("CI_TESTING_MODE") == "1" {
		v.SetConfigName("app_test_ci")

		pwd := os.Getenv("PWD")
		for i := 0; i < 10; i = i + 1 {
			if strings.HasSuffix(pwd, "srv") {
				v.AddConfigPath(fmt.Sprintf("%s/%s", pwd, "configs"))
				break
			}

			pwd = filepath.Dir(pwd)
		}
	}
}

func TestInit(t *testing.T) {
	initConfig()
	initForTestMode()
	initForCiTestMode()

	tz := v.GetString("public.timezone")
	fmt.Println("tz", tz)
	if tz != os.Getenv("TZ") {
		t.Errorf("timezone set failed")
	}

	fmt.Printf("\n time: %#v, carbon: %#v", time.Now().String(), carbon.Now().ToDateTimeString())
}
