package config

import (
	"fmt"
	"github.com/litsoftware/litmedia/pkg/jsonh"
	"github.com/litsoftware/litmedia/pkg/path"
	_ "github.com/litsoftware/litmedia/pkg/runtime"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
	"strings"
)

var (
	err error
	v   *viper.Viper
)

func init() {
	v = viper.New()
	v.SetConfigType("yaml")
	v.SetConfigName("app")

	p := fmt.Sprintf("%s/%s", path.RootPath(), "configs")
	v.AddConfigPath(p)

	initForTestMode()
	initForCiTestMode()
	findConfigFiles()

	err = v.ReadInConfig()
	if err != nil {
		fmt.Printf("parse configuration file err %#v\n", err)
		os.Exit(1)
	}

	if v.GetBool("watch") {
		v.WatchConfig()
	}

	check()
}

func findConfigFiles() {
	pwd := os.Getenv("PWD")

	for i := 0; i < 10; i = i + 1 {
		if strings.HasSuffix(pwd, "media-service") {
			v.AddConfigPath(fmt.Sprintf("%s/%s", pwd, "configs"))
			break
		}

		pwd = filepath.Dir(pwd)
	}
}

func initForTestMode() {
	if os.Getenv("TESTING_MODE") == "1" {
		v.SetConfigName("app_test")
	}
}

func initForCiTestMode() {
	if os.Getenv("CI_TESTING_MODE") == "1" {
		v.SetConfigName("app_test_ci")
	}
}

func SetDefault(key string, value interface{}) {
	v.SetDefault(key, value)
}

func GetNormal(key string) interface{} {
	return v.Get(key)
}

func GetBool(key string) bool {
	return v.GetBool(key)
}

func GetString(key string) string {
	return v.GetString(key)
}

func GetStringDefault(key string, d string) string {
	s := v.GetString(key)
	if s == "" {
		return d
	}

	return s
}

func GetStringSlice(key string) []string {
	return v.GetStringSlice(key)
}

func GetInt(key string) int {
	return v.GetInt(key)
}

func GetFloat64(key string) float64 {
	return v.GetFloat64(key)
}

func GetInt64(key string) int64 {
	return v.GetInt64(key)
}

func GetInt32(key string) int32 {
	return v.GetInt32(key)
}

func GetMap(key string) map[string]interface{} {
	return v.GetStringMap(key)
}

func GetStruct(key string, obj interface{}) {
	m := v.GetStringMap(key)
	jsonh.ConvertTo(m, obj)
}

func GetEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func GetConfig() *viper.Viper {
	return v
}
