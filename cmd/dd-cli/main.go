package main

import (
	"fmt"
	"os"

	"github.com/jmacksf/dd-cli/internal/cmd/root"

)

func main() {
	rootCmd := root.NewCmdRoot()
	if _, err := rootCmd.ExecuteC(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}
