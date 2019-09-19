# Sirius


## Use Case

1. Sirius VPNClient installation time.
  - It generate the key-pair and user can provides an optional passphrase to protect the private key. Take easyras as example,
    
    <code>easyrsa gen-req <Client-Name> [nopass]</code> generate two files, certificate request file, <code>.req</code> and private key file, <code><.key</code>.
     The Client-Name should be unique across the PKI.
     *To protect private key, the file should be encrypted on-disk.*
  - User configure the specified [security domain](securty_domain.md).
  - key generation and configuration can be done at setting mode also
     
2. User launch Sirius VPNClient
  - Client and Server validation. VPNClient can uses the [shared secret + password challenge](password_challenge.md) to verify client and server. Once validation is done, server will send the session-id to client. The server does not associate any authentication attribute to this session-id yet.
  - If client certificate is not exist, the VPNClient will trigger the certificate download process.
    - pass the SSO, SAML and connect-id based on the security domain
    - upload the certificate request to server
    - server signed this certificate and push this certificate VPN server.
    - client download this certificate
    - client download the VPN server certificate
    - client download the openvpn client configure file 
  - VPNClient connect to the VPN Server via, 2 certificates, client private key and VPN client configure file
    - run openvpn client as daemon?
