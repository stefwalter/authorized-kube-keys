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
	DefaultNode string
}

func (cl *Client) nodeName() (string, error) {
	if cl.DefaultNode == "" {
		name, err := exec.Command("uname", "-n").Output()
		if err != nil {
			return "", err
		}
		cl.DefaultNode = strings.TrimSpace(string(name))
	}
	return cl.DefaultNode, nil
}

func (cl *Client) execute(args []string) ([]byte, error) {
	if cl.MasterUrl != "" {
		args = append(args, "--server", cl.MasterUrl)
	}
	if cl.KubeConfig != "" {
		args = append(args, "--kubeconfig", cl.KubeConfig)
	}

	command := exec.Command("kubectl", args...)

	// Yup this is how to retrieve stderr (see below)
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

func (cl *Client) Get(kind string, name string) ([]byte, error) {
	if name == "" {
		var err error
		name, err = cl.nodeName()
		if err != nil {
			return nil, err
		}
	}
	args := []string{"get", kind, name, "--output=json"}
	return cl.execute(args)
}

func (cl *Client) Patch(kind string, name string, patch string) ([]byte, error) {
	if name == "" {
		var err error
		name, err = cl.nodeName()
		if err != nil {
			return nil, err
		}
	}
	args := []string{"patch", kind, name, "-p", patch}
	return cl.execute(args)
}
