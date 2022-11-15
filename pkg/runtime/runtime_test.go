package runtime

import (
	"os"
	"testing"
)

func TestInit(t *testing.T) {
	if os.Getenv("TESTING_MODE") != "1" {
		t.Errorf("missing testing argument")
	}
}
