package main

import (
	"fmt"
	"os"
)

func main() {
	// tinyscanner -u 20.78.06.364 -p 80
	if len(os.Args) != 5 {
		fmt.Fprintln(os.Stderr, "Usage: portscanner [OPTIONS] [HOST] [PORT]")
		fmt.Fprintln(os.Stderr, "Options:")
		fmt.Fprintln(os.Stderr, "-p               Range of ports to scan")
		fmt.Fprintln(os.Stderr, "-u               UDP scan")
		fmt.Fprintln(os.Stderr, "-t               TCP scan")
		fmt.Fprint(os.Stderr, "--help           Show this message and exit.")
		return
	}
}
