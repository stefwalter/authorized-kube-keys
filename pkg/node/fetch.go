package node

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/stefwalter/authorized-kube-keys/pkg/client"
)

func PrintAuthorized(cli *client.Client, name string) (error) {
	output, err := cli.Get("node", name)
	if err != nil {
		return err
	}

	var n Node
	err = json.Unmarshal(output, &n)
	if err != nil {
		return err
	}

	for k, v := range n.NodeMeta.Annotations {
		if strings.HasPrefix(k, "authorized-key-") {
			fmt.Println(v)
		}
	}
	return nil
}
