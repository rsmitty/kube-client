### Kube-Client

A Kubernetes toy client written with golang.

##### Motivation

This toy was created to familiarize myself with the Kubernetes golang client, as well as with various other packages used heavily around the Kubernetes community. Namely cobra and logrus. I have implemented two pretty basic Kubernetes API operations, pod list and pod delete. 

##### Building

This client can be easily built using the provided Makefile. Simply issue one of the following:
- `make darwin` - Builds 64-bit Mac client in the bin/ directory
- `make linux` - Builds 64-bit Linux client in the bin/ directory
- `make all` - Builds both of the above in the bin/ directory

##### Usage and Assumptions

This client assumes a few things:
- You have a pre-existing kube config file that can be used
- You're going to use this as an "out of cluster" tool. This simply means we're not going to expect to use a service account token for talking to the API right now. Related to the above point.
- You're using minikube or GKE. I've only really tested with these two auth methods.

Using the client is simple. There's only two commands:
- `kube-client pod list`
- `kube-client pod delete $POD_NAME`

There are two flags that are shared between these commands. There are defaults set, but it may be worth seeing if you need to override them:
- `--kubeconfig string   A path to your kubeconfig file. (default "$HOME/.kube/config")`
- `--namespace string    A kubernetes namespace you would like to interact with. (default "default")`

Since this client uses the kube config file to authenticate to a cluster, it will use your current context. You can switch that to point to different clusters using `kubectl config use-context $WHATEVER` or by editing the `.kube/config` yaml file directly.

There's also a `--log-level` flag that I implemented while playing with logrus. It is accepted at the root command level and flows down into the subcommands. That said, there's really only fatal logs used, so you will pretty much see any relevant logs anyways.

##### Examples
Examples of usage of the client. Note that I don't actually need the flags shown since the defaults are fine for me, but they're specified just to show they are there.
List pods:
```bash
$ ./kube-client-darwin pod list --namespace default --kubeconfig ~/.kube/config

Here are the pods in default namespace:
#0: alpine
#1: etcd0
#2: etcd1
#3: etcd2
#4: get-python-1265444380-brd7m
#5: hello-1055340253-cjvs5
#6: idle-rat-mariadb-4117317837-flb1b
#7: idle-rat-wordpress-456697286-twb7j
#8: kube-ops-view-4153059266-zxxpf
#9: simple-api
#10: virulent-dingo-gitlab-ce-3545565135-p2p0p
#11: virulent-dingo-postgresql-96911118-b4v3f
#12: virulent-dingo-redis-3663642221-fbtxg
```

Delete pod:
```bash
$ ./kube-client-darwin pod delete kube-ops-view-4153059266-zxxpf --namespace default --kubeconfig ~/.kube/config
Successfully submitted deletion for pod kube-ops-view-4153059266-zxxpf
```
