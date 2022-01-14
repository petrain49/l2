package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/shirou/gopsutil/v3/process"
)

func pwd() string {
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}

	return path
}

func kill(s string) int {
	pid, err := strconv.Atoi(s)
	if err != nil {
		log.Println(err)
	}

	process, err := os.FindProcess(pid)
	if err != nil {
		log.Println(err)
	}

	err = process.Kill()
	if err != nil {
		log.Println(err)
	}

	return pid
}

func ps() []string {
	proc, err := process.Processes()
	if err != nil {
		log.Println(err)
	}

	var processes []string

	for _, p := range proc {
		name, _ := p.Name()
		processes = append(processes, fmt.Sprintf("process: %s %d\n", name, p.Pid))
	}

	return processes
}
