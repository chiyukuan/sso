# easy-rsa: Simple Shell based CA utility

## Resource
- [easy-rsa cmd options](https://github.com/OpenVPN/easy-rsa/issues/73)

## installation

``` bash
yum install epel-release -y
yum install easy-ras -y
```

## Configuration

This configuration is based on the easy-rsa 3

- Edit the /usr/share/easy-rsa/3/vars as following

``` bash
sso /usr/share/easy-rsa/3]$ cat /usr/share/easy-rsa/3/vars

set_var EASYRSA                 "$PWD"
set_var EASYRSA_PKI             "$EASYRSA/pki"
set_var EASYRSA_DN              "cn_only"
set_var EASYRSA_REQ_COUNTRY     "ID"
set_var EASYRSA_REQ_PROVINCE    "Jakarta"
set_var EASYRSA_REQ_CITY        "Jakarta"
set_var EASYRSA_REQ_ORG         "hakase-labs CERTIFICATE AUTHORITY"
set_var EASYRSA_REQ_EMAIL       "openvpn@hakase-labs.io"
set_var EASYRSA_REQ_OU          "HAKASE-LABS EASY CA"
set_var EASYRSA_KEY_SIZE        2048
set_var EASYRSA_ALGO            rsa
set_var EASYRSA_CA_EXPIRE       7500
set_var EASYRSA_CERT_EXPIRE     365
set_var EASYRSA_NS_SUPPORT      "no"
set_var EASYRSA_NS_COMMENT      "HAKASE-LABS CERTIFICATE AUTHORITY"
set_var EASYRSA_EXT_DIR         "$EASYRSA/x509-types"
set_var EASYRSA_SSL_CONF        "$EASYRSA/openssl-1.0.cnf"
set_var EASYRSA_DIGEST          "sha256"

sso /usr/share/easy-rsa/3]$ sudo chmod u+x /usr/share/easy-rsa/3/vars

```

## Build OpenVPN Keys (CA key, server and client keys)

- brief
  - CA key: init-pki and build-ca
  
  ``` bash
  sso ~]$ cd /usr/share/easy-rsa/3/
  sso /usr/share/easy-rsa/3]$ sudo ./easyrsa init-pki
  sso /usr/share/easy-rsa/3]$ sudo ./easyrsa build-ca
  sso /usr/share/easy-rsa/3]$ sudo find pki/
  pki/
  pki/private
  pki/private/ca.key
  pki/reqs
  pki/safessl-easyrsa.cnf
  pki/issued
  pki/certs_by_serial
  pki/revoked
  pki/revoked/certs_by_serial
  pki/revoked/private_by_serial
  pki/revoked/reqs_by_serial
  pki/renewed
  pki/renewed/certs_by_serial
  pki/renewed/private_by_serial
  pki/renewed/reqs_by_serial
  pki/index.txt
  pki/serial
  pki/ca.crt
  ```

  **Building CA need the password.**

  - Server Key: gen, sign, and verify
  ``` bash
  sso ~]$ cd /usr/share/easy-rsa/3/
  sso /usr/share/easy-rsa/3]$ sudo ./easyrsa gen-req openvpn-server nopass
  sso /usr/share/easy-rsa/3]$ sudo ./easyrsa sign-req server openvpn-server
  sso /usr/share/easy-rsa/3]$ sudo openssl verify pki/ca.crt pki/issued/openvpn-server.crt
  ```

  **Signing key need CA's password.**

  - client key: gen, sign and verify
  ``` bash
  sso ~]$ cd /usr/share/easy-rsa/3/
  sso /usr/share/easy-rsa/3]$ sudo ./easyrsa gen-req openvpn-client nopass
  sso /usr/share/easy-rsa/3]$ sudo ./easyrsa sign-req client openvpn-client
  sso /usr/share/easy-rsa/3]$ sudo openssl verify pki/ca.crt pki/issued/openvpn-client.crt
  ```

  - Build Diffie-Hellman key, under pki directory
  ``` bash
  sso ~]$ cd /usr/share/easy-rsa/3/
  sso /usr/share/easy-rsa/3]$ sudo ./easyrsa gen-dh
  ```

  - Gen the CRL (Certificate Revoking List) key
  ``` bash
  sso ~]$ cd /usr/share/easy-rsa/3/
  sso /usr/share/easy-rsa/3]$ sudo ./easyrsa revoke <someone>
  sso /usr/share/easy-rsa/3]$ sudo ./easyrsa gen-crl
  ```

  **Do I have to reset the ?revoke-list? at beginning?**

- detail, CA and Server key only

``` bash
"sso ~]$" cd /usr/share/easy-rsa/3/
"sso /usr/share/easy-rsa/3]$" sudo ./easyrsa init-pki

Note: using Easy-RSA configuration from: ./vars

init-pki complete; you may now create a CA or requests.
Your newly created PKI dir is: /usr/share/easy-rsa/3.0.6/pki

sso /usr/share/easy-rsa/3]$ sudo ./easyrsa build-ca

Note: using Easy-RSA configuration from: ./vars

Using SSL: openssl OpenSSL 1.0.2k-fips  26 Jan 2017

Enter New CA Key Passphrase:
Re-Enter New CA Key Passphrase:
Generating RSA private key, 2048 bit long modulus
...........+++
............................................................................+++
e is 65537 (0x10001)
You are about to be asked to enter information that will be incorporated
into your certificate request.
What you are about to enter is what is called a Distinguished Name or a DN.
There are quite a few fields but you can leave some blank
For some fields there will be a default value,
If you enter '.', the field will be left blank.
-----
Common Name (eg: your user, host, or server name) [Easy-RSA CA]:sso

CA creation complete and you may now import and sign cert requests.
Your new CA certificate file for publishing is at:
/usr/share/easy-rsa/3.0.6/pki/ca.crt

"sso /usr/share/easy-rsa/3]$" sudo ./easyrsa gen-req openvpn-server nopass

Note: using Easy-RSA configuration from: ./vars

Using SSL: openssl OpenSSL 1.0.2k-fips  26 Jan 2017
Generating a 2048 bit RSA private key
.............+++
........+++
writing new private key to ''/usr/share/easy-rsa/3.0.6/pki/private/openvpn-server.key.sAcxvusC4f''
-----
You are about to be asked to enter information that will be incorporated
into your certificate request.
What you are about to enter is what is called a Distinguished Name or a DN.
There are quite a few fields but you can leave some blank
For some fields there will be a default value,
If you enter '.', the field will be left blank.
-----
Common Name (eg: your user, host, or server name) [openvpn-server]:

Keypair and certificate request completed. Your files are:
req: /usr/share/easy-rsa/3.0.6/pki/reqs/openvpn-server.req
key: /usr/share/easy-rsa/3.0.6/pki/private/openvpn-server.key

"sso /usr/share/easy-rsa/3]$" sudo ./easyrsa sign-req server openvpn-server

Note: using Easy-RSA configuration from: ./vars

Using SSL: openssl OpenSSL 1.0.2k-fips  26 Jan 2017


You are about to sign the following certificate.
Please check over the details shown below for accuracy. Note that this request
has not been cryptographically verified. Please be sure it came from a trusted
source or that you have verified the request checksum with the sender.

Request subject, to be signed as a server certificate for 365 days:

subject=
    commonName                = openvpn-server


Type the word ''yes'' to continue, or any other input to abort.
  Confirm request details: yes
Using configuration from /usr/share/easy-rsa/3.0.6/pki/safessl-easyrsa.cnf
Enter pass phrase for /usr/share/easy-rsa/3.0.6/pki/private/ca.key:
Check that the request matches the signature
Signature ok
The Subject''s Distinguished Name is as follows
commonName            :ASN.1 12:''openvpn-server''
Certificate is to be certified until Aug 30 17:31:07 2020 GMT (365 days)

Write out database with 1 new entries
Data Base Updated

Certificate created at: /usr/share/easy-rsa/3.0.6/pki/issued/openvpn-server.crt

"sso /usr/share/easy-rsa/3]$" sudo openssl verify pki/ca.crt pki/issued/openvpn-server.crt
pki/ca.crt: CN = sso
error 18 at 0 depth lookup:self signed certificate
OK
pki/issued/openvpn-server.crt: CN = openvpn-server
error 20 at 0 depth lookup:unable to get local issuer certificate
```
