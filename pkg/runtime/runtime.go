package runtime

import (
	"os"
	"strings"
)

func init() {
	for _, arg := range os.Args {
		if strings.HasSuffix(arg, "testing") {
			os.Setenv("TESTING_MODE", "1")
		}

		if strings.HasSuffix(arg, "ci") {
			os.Setenv("CI_TESTING_MODE", "1")
		}
	}
}
