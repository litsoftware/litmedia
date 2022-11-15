package tzinit

import (
	"github.com/litsoftware/litmedia/internal/pkg/config"
	"os"
	"time"
)

func init() {
	tz := config.GetStringDefault("public.timezone", "UTC")
	_ = os.Setenv("TZ", tz)

	loc, err := time.LoadLocation(tz)
	if err != nil {
		time.Local = loc
	}
}
