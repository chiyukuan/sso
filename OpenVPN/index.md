# OpenVPN

## installation

``` bash
yum install epel-release -y
yum install openvpn -y
```

## Install Certificate

Example assume we are using the easy-rsa locally. **see [easy-rsa](../Certificate/easy-rsa.md) for generating CA and keys**

- Copy Server Key and Certificate
``` bash
cp /usr/share/easy-rsa/3/pki/ca.crt /etc/openvpn/server/
cp /usr/share/easy-rsa/3/pki/issued/openvpn-server.crt /etc/openvpn/server/
cp /usr/share/easy-rsa/3/pki/private/openvpn-server.key /etc/openvpn/server/
```

- Copy client key and Certificate
```
cp /usr/share/easy-rsa/3/pki/ca.crt /etc/openvpn/client/
cp /usr/share/easy-rsa/3/pki/issued/openvpn-client.crt /etc/openvpn/client/
cp /usr/share/easy-rsa/3/pki/private/openvpn-client.key /etc/openvpn/client/
```

- Copy DH and CRL Key
```
cp /usr/share/easy-rsa/3/pki/dh.pem /etc/openvpn/server/
cp /usr/share/easy-rsa/3/pki/crl.pem /etc/openvpn/server/
```

## Keys
```
Filename	Needed By	              Purpose	                Secret
ca.crt	    server + all clients	  Root CA certificate	    NO
ca.key	    key signing machine only  Root CA key	            YES
dh{n}.pem	server only	              Diffie Hellman parameters	NO
server.crt	server only	              Server Certificate	    NO
server.key	server only	              Server Key	            YES
client1.crt	client1 only	          Client1 Certificate	    NO
client1.key	client1 only	          Client1 Key	            YES
```

## Configure OpenVPN

Edit the [/etc/openvpn/server.conf](./server.conf)

## Configure Port-Forwarding and Routing Firewalld

- configure
``` bash
echo 'net.ipv4.ip_forward = 1' >> /etc/sysctl.conf
sysctl -p

firewall-cmd --permanent --add-service=openvpn
firewall-cmd --permanent --zone=trusted --add-interface=tun0

firewall-cmd --permanent --zone=trusted --add-masquerade

SERVERIP=$(ip route get 84.200.69.80 | awk 'NR==1 {print $(NF-2)}') firewall-cmd --permanent --direct --passthrough ipv4 -t nat -A POSTROUTING -s  10.10.1.0/24 -o $SERVERIP -j MASQUERADE

firewall-cmd --reload

systemctl start openvpn@server
systemctl enable openvpn@server
```

- check it
``` bash
netstat -plntu
systemctl status openvpn@server
```

## Setup OpenVPN Client

Edit the [/etc/openvpn/client/openvpn-client.ovpn](./openvpn-client.ovpn)
Client need client.crt client.key ca.crt client.ovpn 4 files for VPN connections. **Need to verify it**

- pack client files at openvpn-server
``` bash
cd /etc/openvpn/
tar -cvfz client.tar.gz client/*
```

- unpack client files at openvpn-client
``` bash
scp root@openvpn-server:/etc/openvpn/client.tar.gz .
tar -xvfz client.tar.gz
```

- Installation OpenVPN client
``` bash
sudo yum install -y openvpn NetworkManager-openvpn NetworkManager-openvpn-gnome -y
```

- Connect to OpenVPN server
``` bash
openvpn --config openvpn-client.ovpn
```

- Check connection __Net Yet__
``` bash
curl ifconfig.io
```

# MISC

## X.509 certificate

  tmimestamp indicates when a centificate becomes valid and when it expires

## Authentication and plugins

- options
  - --auth-iser-pass:  require username and passwords
  - --auth-user-pass-verify script is the last in a long chain of scripts that are run.
  - --client-config-dir: client specific configuration and routing
  - --ccd-exclusive: prevent client connections from clients who do not have a file in the client-config directory.
  - When this option is present, even an empty file named to match the CN is sufficient to maet this constraint
  -  --username-as-common-name
  - --client-cert-not-required
  - --verify-client-cert

## installation

<https://www.howtoforge.com/tutorial/how-to-install-openvpn-server-and-client-with-easy-rsa-3-on-centos-7/>

- windows client
  <https://github.com/OpenVPN/openvpn-gui>
