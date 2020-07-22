// MIT License
// Copyright (c) 2020 ysicing <i@ysicing.me>

package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"

	"k8s.io/klog/v2"

	"github.com/spf13/cobra"
)

var name string // 进程路径

var rootCmd = &cobra.Command{
	Use:   "waitshutdown",
	Short: "wait shutdown",
	Run: func(cmd *cobra.Command, args []string) {
		sc := fmt.Sprintf("pkill -SIGTERM -f %v", name)
		cmdres := exec.Command("bash", "-c", sc)
		cmdres.Stdout = os.Stdout
		if err := cmdres.Run(); err != nil {
			klog.Errorf("unexpected error terminating: %v", err)
			os.Exit(1)
		}

		timer := time.NewTicker(time.Second * 2)
		for range timer.C {
			// TODO
			if !IsRunning() {
				timer.Stop()
				break
			}
		}
	},
}

func init() {
	rootCmd.PersistentFlags().StringVar(&name, "name", "bash", "process name")
}

func main() {
	rootCmd.Execute()
}
