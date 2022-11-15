package cmd

import (
	"context"
	"github.com/litsoftware/litmedia/internal/ent"
	"github.com/litsoftware/litmedia/internal/pkg/config/orm"
	"github.com/spf13/cobra"
	"log"
)

var databaseMigrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "执行数据库迁移",
	Run: func(cmd *cobra.Command, args []string) {
		printLine("数据库迁移开始")

		doMigrate()

		printLine("done!")
	},
}

func doMigrate() {
	client := orm.GetClient()
	defer func(client *ent.Client) {
		err := client.Close()
		if err != nil {
			printLine("关闭数据库连接失败")
		}
	}(client)

	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}
