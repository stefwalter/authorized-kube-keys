package node

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"io"

	"github.com/stefwalter/authorized-kube-keys/pkg/client"
)

func PushAuthorized(cli *client.Client, data string) error {
	node, err := fetchNode(cli)
	if err != nil {
		return err
	}

	return PushAuthorizedNode(cli, *node, data)
}

func PushAuthorizedNode(cli *client.Client, node Node, data string) error {
	uuid := make([]byte, 16)
	n, err := io.ReadFull(rand.Reader, uuid)
	if n != len(uuid) || err != nil {
		return err
	}

	key := authorizedKeyName + hex.EncodeToString(uuid)

	if node.NodeMeta.Annotations == nil {
		node.NodeMeta.Annotations = make(map[string]string)
	}
	node.NodeMeta.Annotations[key] = data

	// Empty these before marshalling a patch
	name := node.NodeMeta.Name
	node.NodeMeta.Name = ""

	// TODO: Move away from strategic merge to avoid round tripping data
	patch, err := json.Marshal(node)
	if err != nil {
		return err
	}

	_, err = cli.Patch("node", name, string(patch))
	return err
}

