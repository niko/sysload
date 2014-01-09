package main

import (
	"os"
	"os/exec"
	"fmt"
	"strings"
	"strconv"
)

func main() {
	loadavg, err := exec.Command("cat", "/proc/loadavg").Output()
	if err != nil { fmt.Printf("%s", err) ; os.Exit(1) }

	cpuinfo, err := exec.Command("cat", "/proc/cpuinfo").Output()
	if err != nil { fmt.Printf("%s", err) ; os.Exit(10) }

	s := strings.Split(string(loadavg), " ")[0]
	load, err := strconv.ParseFloat(s, 64)
	if err != nil { fmt.Printf("%s", err) ; os.Exit(100) }

	numprocs := float64(strings.Count(string(cpuinfo), "processor"))

	fmt.Printf("%.2f\n", load/numprocs)
}
