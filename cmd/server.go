package cmd

import (
	"fmt"
	"github.com/litsoftware/litmedia/internal/g"
	"github.com/litsoftware/litmedia/internal/pkg/context"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"
)

type ServerInterface interface {
	Start() error
	Stop()
}

var servers []ServerInterface

type exitCode struct{ Code int }

func init() {
	g.App = context.NewAppContext()
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func handleExit() {
	if e := recover(); e != nil {
		if exit, ok := e.(exitCode); ok {
			if exit.Code != 0 {
				_, _ = fmt.Fprintln(os.Stderr, "Failed at", time.Now().Format("January 2, 2006 at 3:04pm (MST)"))
			} else {
				_, _ = fmt.Fprintln(os.Stderr, "Stopped at", time.Now().Format("January 2, 2006 at 3:04pm (MST)"))
			}

			os.Exit(exit.Code)
		}
		panic(e)
	}
}

func start(exitChannel chan os.Signal) int {
	if len(servers) > 0 {
		for _, s := range servers {
			err := s.Start()
			if err != nil {
				s.Stop()
			}
		}

		for {
			s := <-exitChannel
			g.App.Logger.Info("Shutdown triggered, signal: ", s)

			for _, s := range servers {
				s.Stop()
			}

			return 0
		}
	}

	return 0
}

func RunServer(s []ServerInterface) {
	servers = s

	defer handleExit()
	exitChannel := make(chan os.Signal, 1)
	signal.Notify(exitChannel, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM)
	panic(exitCode{start(exitChannel)})
}
