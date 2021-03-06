package cmd

import (
	"github.com/spf13/cobra"

	"github.com/stefwalter/authorized-kube-keys/pkg/client"
)

var masterUrl string
var kubeConfig string
var hostOverride string

func init() {
	RootCmd.PersistentFlags().StringVar(&masterUrl, "master", "", "Kubernetes API master URL")
	RootCmd.PersistentFlags().StringVar(&kubeConfig, "kubeconfig", "", "Kubernetes client config")
	RootCmd.PersistentFlags().StringVar(&hostOverride, "hostname-override", "",
		"If non-empty, will use this string identification instead of the actual hostname")
}

func Client() *client.Client {
	return &client.Client{KubeConfig: kubeConfig, MasterUrl: masterUrl, DefaultNode: hostOverride }
}

var RootCmd = &cobra.Command{
    Use:   "authorized-kube-keys",
    Short: "Fetch or place authorized keys from kubernetes",
    Long:  `Fetch or store SSH authorized keys from kubernetes API server`,
    Run: func(cmd *cobra.Command, args []string) {
	fetchCmd.Run(cmd, args)
    },
}
