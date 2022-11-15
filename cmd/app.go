package cmd

import (
	"context"
	"fmt"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/litsoftware/litmedia/internal/common"
	"github.com/litsoftware/litmedia/internal/pkg/config"
	"github.com/litsoftware/litmedia/internal/pkg/config/orm"
	"github.com/litsoftware/litmedia/pkg/aes"
	"github.com/litsoftware/litmedia/pkg/random"
	"github.com/litsoftware/litmedia/pkg/rsa"
	"github.com/spf13/cobra"
	"os"
)

var paramAppCmd string
var paramOperatorId int
var paramRsaPublicKey string
var paramAppTitle string
var paramAppDescription string
var aesKey string

var appCmd = &cobra.Command{
	Use:   "app",
	Short: "App 操作",
	Run: func(cmd *cobra.Command, args []string) {
		aesKey = config.GetString("public.system_aes")
		switch paramAppCmd {
		case "make": // 创建信息用户
			makeApp()
		case "ls": // 列表所有的用户
			listApp()
		default:
			err := cmd.Usage()
			if err != nil {
				return
			}
		}
	},
}

func makeApp() {
	if paramAppTitle == "" {
		printLine("app title is empty")
		return
	}

	if paramRsaPublicKey == "" {
		printLine("app rsa public key is empty")
		return
	}

	content, err := os.ReadFile(paramRsaPublicKey)
	if err != nil {
		printLine("read rsa public key file error: ", err.Error())
		return
	}

	rsaPub, err := aes.Encrypt([]byte(config.GetString("public.system_aes")), string(content))
	if err != nil {
		printLine("encrypt rsa public key error: ", err.Error())
		return
	}

	_, _, priv, pub, err := rsa.GenerateRsaKeyPair()
	if err != nil {
		printLine("generate rsa key pair error: ", err.Error())
		return
	}

	privStr, err := aes.Encrypt([]byte(aesKey), string(priv))
	if err != nil {
		printLine("encrypt rsa private key error: ", err.Error())
		return
	}

	pubStr, err := aes.Encrypt([]byte(aesKey), string(pub))
	if err != nil {
		printLine("encrypt rsa public key error: ", err.Error())
		return
	}

	appId := fmt.Sprintf("app_%s", random.String(10))
	app, err := orm.GetClient().App.Create().
		SetAppID(appId).
		SetAppSecret(random.String(32)).
		SetEncryptedAppPrivateKey(privStr).
		SetEncryptedAppPublicKey(pubStr).
		SetTitle(paramAppTitle).
		SetDescription(paramAppDescription).
		SetOperatorID(paramOperatorId).
		SetEncryptedOperatorRsaPublicKey(rsaPub).
		Save(context.Background())

	err = common.DecryptAppInfo(app)
	if err != nil {
		return
	}

	if err != nil {
		printLine("创建失败: ", err)
		return
	}

	printLine("创建成功: ", app.ID)
	printLine("Done!")
}

func listApp() {
	items, err := orm.GetClient().App.Query().All(context.Background())
	if err != nil {
		printLine("查询出错了: ", err)
		return
	}

	printTable([]interface{}{"#", "App ID", "Title", "Description", "Create Time"}, func(t table.Writer) table.Writer {
		for _, item := range items {
			t.AppendRows([]table.Row{
				{item.ID, item.AppID, item.Title, item.Description, item.CreatedAt.Format("2006-01-02 15:04:05")},
			})
			t.AppendSeparator()
		}
		return t
	}, []interface{}{"", "总记录数", len(items)})

	printLine("Done!")
}

func init() {
	appCmd.Flags().StringVarP(&paramAppCmd, "cmd", "", "", "操作命令. make 创建 App，ls 列表所有 App")
	appCmd.Flags().IntVarP(&paramOperatorId, "operator_id", "", 0, "操作用户ID")
	appCmd.Flags().StringVarP(&paramRsaPublicKey, "rsa_public_key", "", "", "用户的公钥")
	appCmd.Flags().StringVarP(&paramAppTitle, "title", "", "", "名称")
	appCmd.Flags().StringVarP(&paramAppDescription, "description", "", "", "描述")
}
