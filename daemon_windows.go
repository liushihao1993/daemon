package daemon

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func init() {
	daemon := flag.Bool("daemon", false, "to run it as a full daemon. only support for linux and mac os")
	var argNoDaemon []string
	for _, a := range os.Args {
		if strings.Contains(a, "-daemon") {
			*daemon = true
			continue
		}
		argNoDaemon = append(argNoDaemon, a)
	}
	if *daemon {
		cmd := exec.Command("powershell", "-c", strings.Join(argNoDaemon, " "))
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Start(); err != nil {
			fmt.Printf("start %s failed, error: %v\n", os.Args[0], err)
			os.Exit(1)
		}
		fmt.Printf("%+v PID: %d running...\n", argNoDaemon, cmd.Process.Pid)
		os.Exit(0)
	}
}
