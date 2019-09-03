# kubectl

Kubectl is a command line interface of running commands against Kubernetes clusters. See <https://kubernetes.io/docs/reference/kubectl/overview/> for detail.

## Installation
```
curl -LO https://storage.googleapis.com/kubernetes-release/release/`curl \
      -s https://storage.googleapis.com/kubernetes-release/release/stable.txt`/bin/linux/amd64/kubectl

chmod +x ./kubectl
sudo mkdir -p /opt/k8s/bin
sudo mv ./kubectl /opt/k8s/bin
sudo ln -s /opt/k8s/bin/kubectl /usr/local/bin

kubectl version
```

## Examples


## Syntax

```
kubectl [command] [Resource-TYPE] [resource-NAME] [flags]

command: create, get, describe, delete
flags: specifies optional flags.
```

