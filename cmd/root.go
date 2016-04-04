package cmd

import (
	"errors"
	"os/exec"

	"github.com/spf13/cobra"

	"k8s.io/kubernetes/pkg/api"
	"k8s.io/kubernetes/pkg/api/unversioned"
	"k8s.io/kubernetes/pkg/client/restclient"
	"k8s.io/kubernetes/pkg/client/unversioned/clientcmd"
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

func RestConfig() (*restclient.Config, error) {
	if kubeConfig == "" && masterUrl == "" {
		return nil, errors.New("Neither --kubeconfig nor --master was specified.")
	}

	data, err := clientcmd.LoadFromFile(kubeConfig)
	if err != nil {
		return nil, err
	}

	config, err := clientcmd.NewNonInteractiveClientConfig(*data, data.CurrentContext,
		&clientcmd.ConfigOverrides{}).ClientConfig()
	if err != nil {
		return nil, err
	}

	groupVersion := unversioned.GroupVersion{"api", "v1"}
	config.ContentConfig.GroupVersion = &groupVersion
	config.Codec = api.Codecs.LegacyCodec(groupVersion)
	return config, nil
}

func RestClient() (*restclient.RESTClient, error) {
	config, err := RestConfig()
	if err != nil {
		return nil, err
	}
	return restclient.RESTClientFor(config)
}

func NodeName() (string, error) {
	if len(hostOverride) != 0 {
		return hostOverride, nil
	}

	name, err := exec.Command("uname", "-n").Output()
	if err != nil {
		return "", err
	}
	return string(name), nil
}

var RootCmd = &cobra.Command{
    Use:   "authorized-kube-keys",
    Short: "Fetch or place authorized keys from kubernetes",
    Long:  `Fetch or store SSH authorized keys from kubernetes API server`,
    Run: func(cmd *cobra.Command, args []string) {
	fetchCmd.Run(cmd, args)
    },
}
