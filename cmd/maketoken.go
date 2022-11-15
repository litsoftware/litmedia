package cmd

import (
	"context"
	"github.com/litsoftware/litmedia/internal/common"
	"github.com/litsoftware/litmedia/internal/pkg/config/orm"
	"github.com/spf13/cobra"
)

var maketoken = &cobra.Command{
	Use:   "token",
	Short: "生成token",
	Run: func(cmd *cobra.Command, args []string) {
		app, err := orm.GetClient().App.Query().First(context.Background())
		if err != nil {
			printLine("error ", err.Error())
			return
		}

		err = common.DecryptAppInfo(app)
		if err != nil {
			printLine("error ", err.Error())
			return
		}

		token, err := common.GenerateStsTokenByApp(app)
		printLine("token: ")
		printLine(token)
		if err != nil {
			printLine("error ", err.Error())
		}
	},
}
