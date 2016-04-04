Proof of concept for storing SSH authorized-keys in Kubernetes Node info

### Building

Make sure to either set your $GOPATH to point to a checked out
```kubernetes/Godeps/_workspace``` directory. At least until the following
kubernetes pull requests are included:

 * https://github.com/kubernetes/kubernetes/pull/23789
 * https://github.com/kubernetes/kubernetes/pull/23632

Then just ```go install``` and copy the resulting binary to ```/usr/local/bin```
on your nodes.

### Running

Add the following lines to your ```/etc/sshd/sshd_config``` on each Node:

```
AuthorizedKeysCommand /usr/local/bin/authorized-kube-keys --kubeconfig=/var/lib/kubelet/kubeconfig
AuthorizedKeysCommandUser root
```

Restart sshd.

### Examples

The tool doesn't yet have a way to add/remove keys. Look in the examples directory
for some ```kubectl patch``` JSON you can use.
