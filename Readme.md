# Go-Mail-Admin
HTTP-Interface with a small GUI for a Mailserver based on the [Tutorial "Own mail server based on Dovecot, Postfix, MySQL, Rspamd and Debian 9 Stretch"](https://thomas-leister.de/en/mailserver-debian-stretch/).
The Project connect to the MYSQL-Database to read and write the config. 

You just need to download one binary file to the server and run it, no other dependencies. Anyway you may want to add an nginx with SSL in front of it, or make it just available in a VPN.
# Installation
Download the last binary from the [Release Page](https://github.com/kekskurse/go-mail-admin/releases) to your Ubuntu/Debian Mailserver. 

Set the Enviroment-Varieables to configure the Go-Mail-Admin
```
export GOMAILADMIN_DB="vmail:vmailpassword@tcp(127.0.0.1:3306)/vmail"
export GOMAILADMIN_APIKEY=abc
export GOMAILADMIN_APISECRET=abc
```

After it you can start the Go-Mail-Admin with the following command
```
./go-mail-admin-with-gui-<VERSION>
```

# Usage
## Config
The script can be config with environment Variables. The following Settings are possible:

| Key | Required | Notice |
| --- | ---      | --- |
| GOMAILADMIN_DB | Yes | Database connection string like 'username:password@tcp(127.0.0.1:3306)/database' |
| GOMAILADMIN_APIKEY | No | API Key for HTTP-Basic-Auth (just use if APISECRET  is set too)  |
| GOMAILADMIN_APISECRET | No | API Secret for HTTP-Basic-Auth (just use if APIKEY is set too) |

## API
### Domains
Domain Action can be triggert by a http call to /api/v1/domain. Parameter can be transmitted as JSON-Body

| Method | Notice |
| ---    | ---    |
| GET    | List all Domains |
| POST   | Create a new Domain, required Parameters: domain |
| DELETE | Delete a Domain, required Parameters: domain |

### Aliases
Alias Action can be triggert by a http call to /api/v1/alias. Parameter can be transmitted as JSON-Body

| Method | Notice |
| ---    | ---    |
| GET    | List all Aliases |
| POST   | Create a new Alias, required Parameters: source_username, source_domain, destination_username, destination_domain, enabled |
| DELETE | Delete a Domain, required Parameters: id |
| PUT    | Apdate a Alias, required Parameters: source_username, source_domain, destination_username, destination_domain, enabled, id |

### Account
Alias Action can be triggert by a http call to /api/v1/account. Parameter can be transmitted as JSON-Body

| Method | Notice |
| ---    | ---    |
| GET    | List all Accounts |
| POST   | Create a new Account, required Parameters: username, domain, password, quota, enabled, sendonly |
| DELETE | Delete a Account, required Parameters: id |
| PUT    | Apdate a Alias, required Parameters: quota, enabled, sendonly, id |

To update the Password make a http call to /api/v1/account/password with the Json-Body Parameter: id, password

### TLSpolicies
TLSPolicys can be changed by http call to /api/v1/tlspolicy

| Method | Notice |
| ---    | ---    |
| GET    | List all Policys |
| POST   | Create a new Policy, required Parameters: domain, policy, params |
| DELETE | Delete a Policy, required Parameters: id |
| PUT    | Apdate a Policy, required Parameters: domain, policy, params, id |

## Frontend
There is a VueJS Frontend to see and configure your Mailserver. You can access it by open http://yourmailserver.de:3001, if you set an API Key and Secret you need to performe a HTTP-Basic-Auth.

### Screenshots
*Domainliste*
![Domainlist](statik/domains.png)
*Aliases*
![Domainlist](statik/aliases.png)
*Aliases Edit*
![Domainlist](statik/aliases-edit.png)
*Accounts*
![Domainlist](statik/accounts.png)
*Accounts Edit*
![Domainlist](statik/account-edit.png)
*TLS Policy*
![Domainlist](statik/tlspolicy.png)
*TLS Policy Edit*
![Domainlist](statik/tlspolicy-edit.png)

# Dev
## Project structure
In the Folder "mailserver-configurator-interface" is the GO-Project which provides a HTTP-API for the Mailserver-Config, the "mailserver-configurator-client" is a VUEJS interface for that HTTP-API.

## Compile
### Mailserver-Configurator-Client
Go in the Project folder and build the Vue.js project

```
npm install
npm run build
```
### Mailserver-Configurator-Interface
Go in the Project folder. Make sure you have all dependencies

```
go get github.com/go-sql-driver/mysql
go get github.com/go-chi/cors
go get gopkg.in/unrolled/render.v1
go get github.com/go-chi/chi
go get github.com/rakyll/statik
```

Create a "public" folder and copy the dist folder from the Client into it

```
mkdir public
cd public
cp -r ../../mailserver-configurator-client/dist/* ./
cd ..
```

Include the Static files (Vue-App) to Go to make it available in the binary.
```
statik -f -src=./public
```
oder
```
~/go/bin/statik -f -src=./public
```
Than build the GO Project
```
go build ./
```

In the build-with-gui.sh in the project root is a shell script to build the client and interface. You can call it with a name
```
./build-with-gui.sh dev
```
will create a go-mail-admin-with-gui-dev in the project root.