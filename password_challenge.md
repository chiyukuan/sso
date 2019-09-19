# Password challenge

- How to can Web Server identify the OpenVPN client?
  shared secret + password challenge-response authentication
  client send challenge, encrypted with shared secret, to receiver
  receiver response a similarly encrypted value which is some predetermined function of the originally offered information.
  For example, client - integer:N -> server, server - integer: (N+1) -> client
  
  shared secret key: encoding( version-number, random-string )
  server's response is based on the version-number
