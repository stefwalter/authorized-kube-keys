package client

import (
	"bytes"
	"errors"
	"os/exec"
	"strings"
)

type Client struct {
	KubeConfig string
	MasterUrl string
}

func (cl Client) Get(kind string, name string) ([]byte, error) {
	args := []string{"get", "--output=json"}
	if cl.MasterUrl != "" {
		args = append(args, "--server", cl.MasterUrl)
	}
	if cl.KubeConfig != "" {
		args = append(args, "--kubeconfig", cl.KubeConfig)
	}
	args = append(args, kind, name)
	command := exec.Command("kubectl", args...)
	var e bytes.Buffer
	command.Stderr = &e
	output, err := command.Output()
	if err != nil {
		data := strings.TrimSpace(string(e.Bytes()))
		if data != "" {
			return nil, errors.New(data)
		}
	}
	return output, err
}
