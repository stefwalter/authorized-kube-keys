package cmd

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/spf13/cobra"

	"github.com/stefwalter/authorized-kube-keys/pkg/node"
)

var publicFile string

func init() {
	pushCmd.PersistentFlags().StringVar(&publicFile, "public-key", "", "SSH public key file")
	RootCmd.AddCommand(pushCmd)
}

var pushCmd = &cobra.Command{
	Use:   "push",
	Short: "Push an authorized key into kubernetes",
	Long: `Push SSH authorized key for a kubernetes Node`,
	Run: func(cmd *cobra.Command, args []string) {
		err := performPush(args)
		if err != nil {
			fmt.Fprintln(os.Stderr, "authorized-kube-keys:", err)
			os.Exit(1)
		}
	},
}

func performPush(pubfiles []string) error {
	if len(pubfiles) == 0 {
		return errors.New("specify one or more public key files")
	}
	client := Client()
	for _, pubfile := range pubfiles {
		data, err := ioutil.ReadFile(pubfile)
		if err != nil {
			return err
		}

		err = node.PushAuthorized(client, strings.TrimSpace(string(data)))
		if err != nil {
			return err
		}
	}
	return nil
}
