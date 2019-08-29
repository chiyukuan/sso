# OpenVPN

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