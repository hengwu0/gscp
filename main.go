package main

import (
	"fmt"
	"gscp/scp"
	"os"
)

func main() {
	err := scp.ScpCli(os.Args)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
