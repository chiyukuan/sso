- installation

{{{
 Download link: https://www.keycloak.org/downloads.html

 wget https://downloads.jboss.org/keycloak/6.0.1/keycloak-6.0.1.tar.gz
 tar xvfz keycloak-6.0.1.tar.gz

 yum install -y java-1.8.0-openjdk-devel
 systemctl stop firewalld
 systemctl disable firewalld


 cd bin
 ./standalone.sh

}}}

- Initial setup add admin user
  ./add-user-keycloak.sh -u root
  Password: password

- Admin console
  start keycloak -b 0.0.0.0 port: 8888

  sso 192.168.146.128
  http://sso:8888/auth/admin


  - create Realm
    From Master drop-down menu, click "Add Realm"
    Name: demo
    Click: "Create"

  - Add User under demo Realm
    From Master drop-down menu, select "demo" Realm
    click "Users"
    click "Add User"
    Username: foo
    click "Save"

    Click Credentials, enter password and Click "Reset password"

- User Account Service @ demo realm
  http://sso:8888/auth/realm/demo/account

