package main

import (
	"encoding/json"
	"fmt"
	"github.com/litsoftware/litmedia/pkg/arraysh"
	"github.com/litsoftware/litmedia/pkg/file"
	"github.com/litsoftware/litmedia/pkg/jsonh"
	"github.com/litsoftware/litmedia/pkg/path"
	"github.com/spf13/cobra"
	"os"
	"runtime"
	"strings"
)

var (
	rootCmd = &cobra.Command{
		Short: "初始化数据",
		Run: func(cmd *cobra.Command, args []string) {
			printLine("开始格式化 swagger.json")

			swFile := path.RootPathWithPostfix("docs/openapi/trade.swagger.json")
			// check file exist
			if file.IsNotExist(swFile) {
				printLine("swagger.json 文件不存在")
				return
			}

			// load file content
			content, err := os.ReadFile(swFile)
			if err != nil {
				printLine("读取 swagger.json 文件失败")
				return
			}

			cm := map[string]interface{}{}
			err = json.Unmarshal(content, &cm)
			if err != nil {
				printLine("解析 swagger.json 文件失败")
				return
			}

			printLine("合并 tags")
			tags := make([]map[string]string, 0)
			var tagList []string
			for _, v := range cm["paths"].(map[string]interface{}) {
				for _, vv := range v.(map[string]interface{}) {
					tag := vv.(map[string]interface{})["tags"].([]interface{})[0].(string)
					// if not contains
					ok, _ := arraysh.InArray(tag, tagList)
					if !ok {
						tagList = append(tagList, tag)
					}
				}
			}

			if len(tagList) > 0 {
				for _, v := range tagList {
					tags = append(tags, map[string]string{
						"name": v,
					})
				}
			}

			cm["tags"] = tags

			printLine("重新写入")
			// write file
			content2, err := json.Marshal(cm)
			if err != nil {
				printLine("重新写入失败")
				return
			}

			prettyString, err := jsonh.PrettyString(content2)
			if err != nil {
				printLine("重新写入失败")
				return
			}

			err = os.WriteFile(swFile, []byte(prettyString), 0644)
			if err != nil {
				printLine("重新写入失败")
				return
			}

			printLine("down!")
		},
	}
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func printLine(args ...interface{}) {
	if len(args) > 1 {
		if strings.Contains(args[0].(string), "%") {
			fmt.Println(fmt.Sprintf(args[0].(string), args[1:]...))
		} else {
			fmt.Println(args...)
		}
	} else {
		fmt.Println(args[0])
	}
}

func main() {
	_ = rootCmd.Execute()
}
