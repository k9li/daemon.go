package daemongo

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
)

func init() {
	d := flag.Bool("d", false, "run app as a daemon with -d=true.")
	flag.Parse()

	if *d {
		cmd := exec.Command(os.Args[0], flag.Args()...)
		if e := cmd.Start(); e != nil {
			fmt.Printf("start %s failed, error: %v\n", os.Args[0], e)
			os.Exit(1)
		}
		fmt.Printf("%s [PID] %d running...\n", os.Args[0], cmd.Process.Pid)
		os.Exit(0)
	}
}
