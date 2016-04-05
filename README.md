Proof of concept for storing SSH authorized-keys in Kubernetes Node info

### Building

Then just ```go install``` and copy the resulting binary to ```/usr/local/bin```
on your nodes.

### Running

Make sure to have ```kubectl``` installed.

Add the following lines to your ```/etc/sshd/sshd_config``` on each Node:

```
AuthorizedKeysCommand /usr/local/bin/authorized-kube-keys --kubeconfig=/var/lib/kubelet/kubeconfig
AuthorizedKeysCommandUser root
```

Restart sshd.

### Examples

Run the command like this in order to push a key for a the current node:

```
/usr/local/bin/authorized-kube-keys push --kubeconfig=/var/lib/kubelet/kubeconfig /path/to/id_rsa.pub
```

Use ```--hostname-override``` to push to a specific node name.

This proof of concept places keys in the ```metadata.annotations``` Node data. Some further work here: https://github.com/kubernetes/kubernetes/pull/23811
