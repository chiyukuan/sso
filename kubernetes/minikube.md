# Mini kube

## Installation
```
curl -Lo minikube https://storage.googleapis.com/minikube/releases/latest/minikube-linux-amd64 \
  && chmod +x minikube

sudo mkdir -p /opt/k8s/bin
sudo mv minikube /opt/k8s/bin/
sudo ln -s /opt/k8s/bin/minikube /usr/local/bin
```

## Install minikube driver
kvm2, Linux KVM driver

- Install KVM
Once KVM is installed, run the <code>virt-host-validate</code> to validate it. It there is the fuse module unloaded error, run <code>sudo modprobe fuse</code>

- Driver installation:
``` bash
curl -LO https://storage.googleapis.com/minikube/releases/latest/docker-machine-driver-kvm2 \
  && sudo install docker-machine-driver-kvm2 /usr/local/bin/
```

## Cleanup local state

```
minikube start --vm-driver=kvm2
```

To make kvm as default driver <code>minikube config set vm-driver kvm2</ode>

**Issue**
```
sfx_sw@k8s ~]$ minikube start --vm-driver=kvm2
* minikube v1.3.1 on Centos 7.6.1810
* Downloading VM boot image ...
minikube-v1.3.0.iso.sha256: 65 B / 65 B [--------------------] 100.00% ? p/s 0s
minikube-v1.3.0.iso: 131.07 MiB / 131.07 MiB [-------] 100.00% 6.15 MiB p/s 21s
* Creating kvm2 VM (CPUs=2, Memory=2000MB, Disk=20000MB) ...
E0902 18:39:32.127571   12502 start.go:723] \
  StartHost: create: Error creating machine: Error in driver during machine creation: \
  creating network: getting libvirt connection: error connecting to libvirt socket.: \
  virError(Code=94, Domain=60, Message='authentication unavailable: \
  no polkit agent available to authenticate action 'org.libvirt.unix.manage'')
*
X Unable to start VM
* Error: [KVM_CONNECTION_ERROR] \
  create: Error creating machine: Error in driver during machine creation: creating network: \
  getting libvirt connection: error connecting to libvirt socket.: \
  virError(Code=94, Domain=60, Message='authentication unavailable: \
  no polkit agent available to authenticate action 'org.libvirt.unix.manage'')
* Suggestion: Have you set up libvirt correctly?
* Documentation: https://minikube.sigs.k8s.io/docs/reference/drivers/kvm2/
*
* If the above advice does not help, please let us know:
  - https://github.com/kubernetes/minikube/issues/new/choose
```