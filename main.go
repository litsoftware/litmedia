package main

import (
	"github.com/litsoftware/litmedia/cmd"
	_ "github.com/litsoftware/litmedia/internal/pkg/auto"
	"runtime"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	cmd.Main()
}
