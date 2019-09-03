# Kubernetes

## Installation (learning environment)

- [Install kubectl](kubectl.md)
- Install Hypervisor
  - KVM
  - Virtual Box
- Install minikube(minikube.md)

- Install Docker Machine KVM driver
```

curl -LO https://storage.googleapis.com/minikube/releases/latest/docker-machine-driver-kvm2
chmod +x docker-machine-driver-kvm2
sudo mv docker-machine-driver-kvm2 /opt/k8s/bin
sudo ln -s /opt/k8s/bin/docker-machine-driver-kvm2 /usr/local/bin
```
