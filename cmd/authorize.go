package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/stefwalter/authorized-kube-keys/pkg/node"
)

func init() {
	RootCmd.AddCommand(authorizeCmd)
}

var authorizeCmd = &cobra.Command{
	Use:   "authorize",
	Short: "Generate a new key and authorize it for all Kubernetes nodes",
	Long: `Generate a new key and authorize it for all Kubernetes nodes`,
	Run: func(cmd *cobra.Command, args []string) {
		err := performAuthorize()
		if err != nil {
			fmt.Fprintln(os.Stderr, "authorized-kube-keys:", err)
			os.Exit(1)
		}
	},
}

func performAuthorize() error {
	client := Client()
	return node.AuthorizeNodes(client)
}
