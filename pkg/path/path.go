package path

import (
	"os"
	"path/filepath"
	"strings"
)

func RootPath() string {
	return Root()
}

func Root() string {
	pwd := os.Getenv("PWD")

	if strings.HasPrefix(pwd, "/go/src") {
		return "/go/src"
	}

	if strings.HasPrefix(pwd, "/srv") {
		return "/srv"
	}

	for i := 0; i < 10; i = i + 1 {
		if strings.HasSuffix(pwd, "litmedia") {
			break
		}

		pwd = filepath.Dir(pwd)
	}

	return pwd
}

func RootPathWithPostfix(p string) string {
	d := RootPath() + "/" + p
	return d
}

func StoragePath() string {
	return RootPath() + "/storage"
}

func StoragePathWithPostfix(p string) string {
	d := StoragePath() + "/" + p
	return d
}
