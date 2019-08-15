 WildFly is bundled with keycload.
 - run a separate instance on same machine as the Keycloak server to run your java servlet application.
 - Run keycload at port 8888 and WildFly at port 8080, jboss.socket.binding.port-offset

- Install wildfly
  Download: wget https://download.jboss.org/wildfly/17.0.1.Final/wildfly-17.0.1.Final.tar.gz
  tar xvfz wildfly-17.0.1.Final.tar.gz -C /opt


- Installing the Client Adapter
  - OPENID Connection: https://downloads.jboss.org/keycloak/6.0.1/adapters/keycloak-oidc/keycloak-wildfly-adapter-dist-6.0.1.tar.gz

  cd /opt/wildfly-17.0.1.Final
  tar xvfz ~/Download/keycloak-wildfly-adapter-dist-6.0.1.tar.gz
  cd bin
  ./jboss-cli.sh --file=adapter-elytron-install-offline.cli

  - SAML 2.0: https://downloads.jboss.org/keycloak/6.0.1/adapters/saml/keycloak-saml-wildfly-adapter-dist-6.0.1.tar.gz
  cd /opt/wildfly-17.0.1.Final
  tar xvfz ~/Download/keycloak-saml-wildfly-adapter-dist-6.0.1.tar.gz
  cd bin
  ./jboss-cli.sh --file=adapter-elytron-install-offline.cli


- Deploy keycloak quickstarts app
  - Install maven
```
    wget https://www-us.apache.org/dist/maven/maven-3/3.6.0/binaries/apache-maven-3.6.0-bin.tar.gz -P /tmp
    sudo tar xf /tmp/apache-maven-3.6.0-bin.tar.gz -C /opt
    sudo ln -s /opt/apache-maven-3.6.0 /opt/maven
    sudo ln -s /opt/maven/bin/mvn /usr/bin
```
  -
```
  git clone https://github.com/keycloak/keycloak-quickstarts
  cd keycloak-quickstarts/app-profile-jee-vanilla
  mvn clean wildfly:deploy
```


- Error: The required mechanism ‘BASIC’ is not available in mechanisms [KEYCLOAK] from the HttpAuthenticationFactory.

  Edit /opt/wildfly-11.0.0.Final/standalone/configuration/standalone.xml and change KEYCLOACK TO BASIC

  <mechanism mechanism-name="KEYCLOACK"> ==>  <mechanism mechanism-name="BASIC">

- Creating and Registering the Client

  Login admin console: 8180/auth/admin

    select demo realm
    click "Clients"

    click "Create"

      Client ID: vanilla
      click "Save"

    click "Installation"

      Select Format: Keycloak OIDC JBoss Subsystem XML

      copy this "text section"

  Edit standalone/configuration/standalone.xml

```
    <subsystem xmlns="urn:jboss:domain:keycloak:1.1"/>
```

    to

```
<subsystem xmlns="urn:jboss:domain:keycloak:1.1">
  <secure-deployment name="WAR MODULE NAME.war">
    <realm>demo</realm>
    <auth-server-url>http://localhost:8180/auth</auth-server-url>
    <public-client>true</public-client>
    <ssl-required>EXTERNAL</ssl-required>
    <resource>vanilla</resource>
  </secure-deployment>
</subsystem>
```

Change "WAR MODULE NAME.war" to "vanilla.war"


