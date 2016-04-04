package cmd

import (
	"fmt"
	"os"
	"strings"

	"k8s.io/kubernetes/pkg/api"

	"github.com/spf13/cobra"
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
	client, err := RestClient()
	if err != nil {
		return err
	}

	name, err := NodeName()
	if err != nil {
		return err
	}

	result := client.Verb("GET").
		AbsPath("api", "v1", "nodes", name).
		Do()

	node := &api.Node{}
	err = result.Into(node)
	if err != nil {
		return err
	}

	for k, v := range node.ObjectMeta.Annotations {
		if strings.HasPrefix(k, "authorized-key-") {
			fmt.Println(v)
		}
	}
	return nil
}
