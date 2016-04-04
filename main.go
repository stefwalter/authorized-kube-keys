package main

import (
	"fmt"
	"os"

	"github.com/stefwalter/authorized-kube-keys/cmd"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(2)
	}
}
