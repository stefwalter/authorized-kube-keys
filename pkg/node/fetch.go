package node

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/stefwalter/authorized-kube-keys/pkg/client"
)

func fetchNode(cli *client.Client) (*Node, error) {
	output, err := cli.Get("node", "")
	if err != nil {
		return nil, err
	}

	var n Node
	err = json.Unmarshal(output, &n)
	if err != nil {
		return nil, err
	}
	return &n, nil
}

func PrintAuthorized(cli *client.Client) (error) {
	n, err := fetchNode(cli)
	if err != nil {
		return err
	}
	for k, v := range n.NodeMeta.Annotations {
		if strings.HasPrefix(k, authorizedKeyName) {
			fmt.Println(v)
		}
	}
	return nil
}

func PrintKnown(cli *client.Client) (error) {
	n, err := fetchNode(cli)
	if err != nil {
		return err
	}
	for k, v := range n.NodeMeta.Annotations {
		if strings.HasPrefix(k, knownKeyName) {
			fmt.Println(v)
		}
	}
	return nil
}
