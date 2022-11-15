package config

import (
	"fmt"
	"github.com/apaxa-go/helper/osh"
	"github.com/litsoftware/litmedia/pkg/path"
	"github.com/litsoftware/litmedia/pkg/random"
	"log"
)

func check() {
	if v.GetString("app.app_public_key") == "" {
		log.Fatalf("rsa public key not set")
	}

	fmt.Println(path.RootPathWithPostfix(v.GetString("app.app_public_key")))
	if v.GetString("app.app_public_key") == "" || !osh.Exists(path.RootPathWithPostfix(v.GetString("app.app_public_key"))) {
		log.Fatalf("rsa public key not set")
	}

	if v.GetString("app.app_key") == "" {
		log.Println("here is an example: ", random.String(32))
		log.Fatalf("app key not set")
	}
}
