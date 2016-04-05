package node

const (
	authorizedKeyName = "authorized-key-"
	knownKeyName = "known-host-"
)

// Abbreviated version of node data. Using the real k8s.io restclient
// results in binaries that are 30+ MB, so do this ourselves
type NodeMeta struct {
	Name string `json:"name,omitempty"`
	Annotations map[string]string `json:"annotations,omitempty"`
}

type Node struct {
	NodeMeta `json:"metadata,omitempty"`
}

