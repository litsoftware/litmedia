package cmd

import (
	"github.com/litsoftware/litmedia/cmd/http"
	"github.com/spf13/cobra"
)

var startUpCmd = &cobra.Command{
	Use:   "start",
	Short: "启动服务。 gRPC、API-Gateway",
	Run: func(cmd *cobra.Command, args []string) {
		startup()
	},
}

func startup() {
	doMigrate()

	RunServer([]ServerInterface{
		http.NewGinServer(),
	})
}
