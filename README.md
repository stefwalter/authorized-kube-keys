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

The tool doesn't yet have a way to add/remove keys. Look in the examples directory
for some ```kubectl patch``` JSON you can use.

This proof of concept places keys in the ```metadata.annotations``` Node data. Some further work here: https://github.com/kubernetes/kubernetes/pull/23811
