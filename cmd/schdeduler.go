package cmd

import (
	"context"
	"github.com/jasonlvhit/gocron"
	"github.com/spf13/cobra"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var scheduleCmd = &cobra.Command{
	Use:   "schedule",
	Short: "schedule",
	Run: func(cmd *cobra.Command, args []string) {
		startScheduler()

		quit := make(chan os.Signal)
		signal.Notify(quit, os.Interrupt, os.Kill, syscall.SIGINT)
		<-quit
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		stopScheduler(ctx)
	},
}

func startScheduler() {
	// cron jobs

	go func() {
		<-gocron.Start()
	}()
}

func stopScheduler(ctx context.Context) {
	gocron.Clear()
}
