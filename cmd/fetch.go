package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/stefwalter/authorized-kube-keys/pkg/node"
)

func init() {
	RootCmd.AddCommand(fetchCmd)
}

var fetchCmd = &cobra.Command{
	Use:   "fetch",
	Short: "Fetch authorized keys from kubernetes",
	Long: `Fetch SSH authorized keys from kubernetes API server
	       and print them to stdout in the usual format.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := performFetch()
		if err != nil {
			fmt.Fprintln(os.Stderr, "authorized-kube-keys:", err)
			os.Exit(1)
		}
	},
}

func performFetch() error {
	client := Client()
	return node.PrintAuthorized(client)
}
