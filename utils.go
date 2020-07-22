// MIT License
// Copyright (c) 2020 ysicing <i@ysicing.me>

package main

import (
	"bytes"

	ps "github.com/mitchellh/go-ps"
	"k8s.io/klog/v2"
	"os/exec"
)

// Scmd 路径
func Scmd(name string, arg ...string) string {
	buf := new(bytes.Buffer)
	klog.Info("[os]exec cmd is : ", name, arg)
	cmd := exec.Command(name, arg[:]...)
	cmd.Stdout = buf
	err := cmd.Run()
	if err != nil {
		klog.Error(err)
	}
	return buf.String()
}

// ShellCmd shell string
//func ShellCmd(name string) string {
//	if runtime.GOOS == "linux" {
//		// return fmt.Sprintf("%v -c pkill -SIGTERM -f %v", Scmd("which", "bash"), name)
//		return fmt.Sprintf("%v -c pkill -SIGTERM -f %v", "bash", name)
//
//	}
//	return fmt.Sprintf("pkill -SIGTERM -f %v", name)
//}

// IsRunning returns true if a process with the name "bash" is found
func IsRunning() bool {
	processes, _ := ps.Processes()
	for _, p := range processes {
		klog.Info(p.Executable())
	}
	return false
}
