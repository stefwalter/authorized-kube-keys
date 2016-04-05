package node

import (
	"crypto/rand"
	"crypto/rsa"
	"encoding/json"
	"fmt"
	"net"
	"os"

	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/agent"

	"github.com/stefwalter/authorized-kube-keys/pkg/client"
)

type nodeList struct {
	Items []Node `json:"items,omitempty"`
}

func fetchNodes(cli *client.Client) ([]Node, error) {
	output, err := cli.List("node")
	if err != nil {
		return nil, err
	}

	var nodes nodeList
	err = json.Unmarshal(output, &nodes)
	if err != nil {
		return nil, err
	}
	return nodes.Items, nil
}

func generateKey() (*rsa.PrivateKey, string, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, "", err
	}

	publicKeyBlock, err := ssh.NewPublicKey(&privateKey.PublicKey)
	if err != nil {
		return nil, "", err
	}

	publicKeyData := ssh.MarshalAuthorizedKey(publicKeyBlock)
	return privateKey, string(publicKeyData), nil
}

func sshAgentAdd(key *agent.AddedKey) (error) {
	conn, err := net.Dial("unix", os.Getenv("SSH_AUTH_SOCK"))
	if err != nil {
		return err
	}
	defer conn.Close()

	ag := agent.NewClient(conn)
	return ag.Add(*key)
}

func AuthorizeNodes(cli *client.Client) error {
	nodes, err := fetchNodes(cli)
	if err != nil {
		return err
	}

	private, public, err := generateKey()
	if err != nil {
		return err
	}

	// TODO: Lifetime
	key := &agent.AddedKey{
		PrivateKey: private,
		Comment: "authorized-kube-keys",
	}

	err = sshAgentAdd(key)
	if err != nil {
		return err
	}


	fmt.Println(public)

	for _, node := range nodes {
		err = PushAuthorizedNode(cli, node, public)
		if err != nil {
			return err
		}
	}
	return nil
}
