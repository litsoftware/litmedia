package path

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

var root string

func init() {
	pwd := os.Getenv("PWD")

	for i := 0; i < 10; i = i + 1 {
		if strings.HasSuffix(pwd, "media-service") {
			break
		}

		pwd = filepath.Dir(pwd)
	}

	root = pwd
}

func TestRootPath(t *testing.T) {
	if root != RootPath() {
		t.Errorf("the return of RootPath is not correct path")
	}
}

func TestRootPathWithPostfix(t *testing.T) {
	if fmt.Sprintf("%s/%s", root, "a") != RootPathWithPostfix("a") {
		fmt.Println("left: ", fmt.Sprintf("%s/%s", root, "a"))
		fmt.Println("right: ", RootPathWithPostfix("a"))
		t.Errorf("the return of RootPathWithPostfix is not correct path")
	}
}

func TestStoragePath(t *testing.T) {
	if fmt.Sprintf("%s/%s", root, "storage") != StoragePath() {
		t.Errorf("the return of StoragePath is not correct path")
	}
}

func TestStoragePathWithPostfix(t *testing.T) {
	if fmt.Sprintf("%s/%s/%s", root, "storage", "a") != StoragePathWithPostfix("a") {
		t.Errorf("the return of StoragePathWithPostfix is not correct path")
	}
}
