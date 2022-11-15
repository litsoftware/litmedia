package pid

import (
	"bufio"
	"fmt"
	"github.com/litsoftware/litmedia/pkg/file"
	"os"
	"strconv"
	"testing"
)

var pidfile = fmt.Sprintf("test_%s", "pidfile.pid")

func TestRemoveLocalDirPidFile(t *testing.T) {
	if file.IsNotExist(pidfile) {
		f, err := file.Create(pidfile, ".")
		if err != nil {
			t.Errorf("create pid file faild, err: %#v", err)
		}
		defer f.Close()

		buf := bufio.NewWriter(f)
		_, _ = buf.Write([]byte("10111"))
		buf.Flush()
	}

	RemovePidFile(pidfile)

	if !file.IsNotExist(pidfile) {
		t.Errorf("remove pid file faild")
	}
}

func TestCheckAndCreatePidFile(t *testing.T) {
	pid := os.Getpid()
	CheckAndCreatePidFile(pidfile)

	pidStr, err := file.GetContent(pidfile)
	if err != nil {
		t.Errorf("write pid faild")
	}

	if pidStr != strconv.Itoa(pid) {
		t.Errorf("write pid faild")
	}

	RemovePidFile(pidfile)
}
