package cmd

import (
	"fmt"
	"github.com/jedib0t/go-pretty/v6/table"
	"os"
	"strings"
)

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

func printTable(header []interface{}, fn func(t table.Writer) table.Writer, footer []interface{}) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(append(table.Row{}, header...))
	t = fn(t)
	t.AppendFooter(append(table.Row{}, footer...))
	t.Render()
}
