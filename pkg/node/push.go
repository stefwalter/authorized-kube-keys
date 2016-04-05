package node

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"io"
	"io/ioutil"
	"strings"

	"github.com/stefwalter/authorized-kube-keys/pkg/client"
)

func PushAuthorized(cli *client.Client, pubfile string) error {
	uuid := make([]byte, 16)
	n, err := io.ReadFull(rand.Reader, uuid)
	if n != len(uuid) || err != nil {
		return err
	}

	data, err := ioutil.ReadFile(pubfile)
	if err != nil {
		return err
	}

	name := authorizedKeyName + hex.EncodeToString(uuid)

	node, err := fetchNode(cli)
	if node.NodeMeta.Annotations == nil {
		node.NodeMeta.Annotations = make(map[string]string)
	}
	node.NodeMeta.Annotations[name] = strings.TrimSpace(string(data))

	// Empty these before marshalling a patch
	node.NodeMeta.Name = ""

	// TODO: Move away from strategic merge to avoid round tripping data
	patch, err := json.Marshal(node)
	if err != nil {
		return err
	}

	_, err = cli.Patch("node", "", string(patch))
	return err
}
