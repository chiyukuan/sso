# KeyCloak

## installation

Download link: <https://www.keycloak.org/downloads.html>

``` bash
 wget https://downloads.jboss.org/keycloak/6.0.1/keycloak-6.0.1.tar.gz
 tar xvfz keycloak-6.0.1.tar.gz

 yum install -y java-1.8.0-openjdk-devel
 systemctl stop firewalld
 systemctl disable firewalld
```

## start

``` bash
 cd bin
 ./standalone.sh
```

## Initial setup add admin user

  ./add-user-keycloak.sh -u root -p password

- Admin console
  start keycloak -b 0.0.0.0 port: 8888

  <http://sso:8888/auth/admin>

  - create Realm
    - From Master drop-down menu, click "Add Realm"
    - Name: demo
    - Click: "Create"

  - create Roles
    - From Master drop-down menu, select "demo" Realm
    - click "Roles"
    - click "Add Role"
    - Rolename: user
    - click "Save"
  - Add User under demo Realm
    - From Master drop-down menu, select "demo" Realm
    - click "Users"
    - click "Add User"
    - Username: foo
    - click "Save"

    - Click Credentials, enter password and Click "Reset password"
    - Click "Role Mappings", add "user" to assigned Roles

- User Account Service @ demo realm
  <http://sso:8888/auth/realms/demo/account>
  
  
## OpenID
First we need to create a client that can be used to obtain the token. Go to the Keycloak admin console again and create a new client. This time give it the Client ID curl and select public for access type. Under Valid Redirect URIs enter http://localhost.

As we are going to manually obtain a token and invoke the service let's increase the lifespan of tokens slightly. In production access tokens should have a relatively low timeout, ideally less than 5 minutes. To increase the timeout go to the Keycloak admin console again. This time click on Realm Settings then on Tokens. Change the value of Access Token Lifespan to 15 minutes. That should give us plenty of time to obtain a token and invoke the service before it expires.

Now we're ready to get our first token using CURL. To do this run:
```
RESULT=`curl --data "grant_type=password&client_id=curl&username=user&password=password" http://localhost:8180/auth/realms/master/protocol/openid-connect/token`
```
This is a bit cryptic and luckily this is not how you should really be obtaining tokens. Tokens should be obtained by web applications by redirecting to the Keycloak login page. We're only doing this so we can test the service as we don't have an application that can invoke the service yet. Basically what we are doing here is invoking Keycloaks OpenID Connect token endpoint with grant type set to password which is the Resource Owner Credentials flow that allows swapping a username and a password for a token.
Take a look at the result by running:

```
echo $RESULT
```
The result is a JSON document that contains a number of properties. There's only one we need for now though so we need to parse this output to retrieve only the value we want. To do this run:

```
TOKEN=`echo $RESULT | sed 's/.*access_token":"//g' | sed 's/".*//g'`
```
This command uses sed to strip out everything before and after the value of the access token property. I'm afraid these instructions will only work on Linux or iOS so Windows users will have to figure out how to do this themselves or install Cygwin. If anyone knows how to do this on Windows I'd appreciate if you'd add the instructions as a comment to this post.
Now that we have the token we can invoke the secured service. To do this run:

```
curl http://localhost:8080/service/secured -H "Authorization: bearer $TOKEN"
```
