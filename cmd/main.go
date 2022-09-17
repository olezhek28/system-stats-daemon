package main

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

func main() {
	cmd := exec.Command("sysctl", "-n", "vm.loadavg")
	res, err := cmd.CombinedOutput()
	if err != nil {
		return
	}

	resFields := strings.Fields(string(res))
	user, err := strconv.Atoi(resFields[16])
	if err != nil {
		return
	}

	system, err := strconv.Atoi(resFields[17])
	if err != nil {
		return
	}

	idle, err := strconv.Atoi(resFields[18])
	if err != nil {
		return
	}

	fmt.Println(user, system, idle)
}
