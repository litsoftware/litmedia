package cmd

import (
	"context"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/litsoftware/litmedia/internal/ent/operator"
	"github.com/litsoftware/litmedia/internal/pkg/config/orm"
	"github.com/litsoftware/litmedia/pkg/hash"
	"github.com/litsoftware/litmedia/pkg/stringsh"
	"github.com/spf13/cobra"
)

var paramCmd string
var paramUsername string
var paramPassword string
var paramEmail string
var paramPhone string

var operatorCmd = &cobra.Command{
	Use:   "operator",
	Short: "操作员",
	Run: func(cmd *cobra.Command, args []string) {
		switch paramCmd {
		case "make": // 创建信息用户
			makeOperator()
		case "ls": // 列表所有的用户
			listOperator()
		default:
			err := cmd.Usage()
			if err != nil {
				return
			}
		}
	},
}

func makeOperator() {
	if paramUsername == "" || paramEmail == "" || paramPhone == "" {
		printLine("参数不完整")
		return
	}

	if paramPassword == "" {
		paramPassword = stringsh.RandStringRunes(20)
	}

	// check phone and email
	c, err := orm.GetClient().Operator.Query().
		Where(operator.Phone(paramPhone)).
		Where(operator.Email(paramEmail)).
		Count(context.Background())
	if err != nil {
		printLine("查询失败，请稍后再试")
		return
	}

	if c > 0 {
		printLine("手机号或邮箱已经存在")
		return
	}

	pwd, err := hash.PasswordHash(paramPassword)
	if err != nil {
		printLine("密码加密失败")
		return
	}

	o, err := orm.GetClient().Operator.Create().SetEmail(paramEmail).SetPhone(paramPhone).SetNickname(paramUsername).
		SetPassword(pwd).Save(context.Background())
	if err != nil {
		printLine("创建失败: ", err)
		return
	}

	printLine("创建成功: ", o.ID)
	printLine("Done!")
}

func listOperator() {
	items, err := orm.GetClient().Operator.Query().All(context.Background())
	if err != nil {
		printLine("查询出错了: ", err)
		return
	}

	printTable([]interface{}{"#", "NickName", "Phone", "Email"}, func(t table.Writer) table.Writer {
		for _, item := range items {
			t.AppendRows([]table.Row{
				{item.ID, item.Nickname, item.Phone, item.Email},
			})
			t.AppendSeparator()
		}
		return t
	}, []interface{}{"", "", "总记录数", len(items)})

	printLine("Done!")
}

func init() {
	operatorCmd.Flags().StringVarP(&paramCmd, "cmd", "", "", "操作命令。 make: 创建信息用户； ls: 列表所有的用户")
	operatorCmd.Flags().StringVarP(&paramUsername, "username", "", "", "新用户的用户名")
	operatorCmd.Flags().StringVarP(&paramPassword, "password", "", "", "新用户的密码")
	operatorCmd.Flags().StringVarP(&paramPhone, "phone", "", "", "新用户的联系电话")
	operatorCmd.Flags().StringVarP(&paramEmail, "email", "", "", "新用户的邮箱")
}
