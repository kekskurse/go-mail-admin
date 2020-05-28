# Auth
You can setup the auth method via the enviroment varieable `GOMAILADMIN_AUTH_METHOD`, the following values are available:

* None
* HTTPBasicAuth
* AdminMail


Based on that Setting some additional parameters are needed

## None
No additional parameters are required, no auth for the Admin interface, anyone who can access the Frontend can change any settings. May be used if the authorisation happens at another place.

## HTTPBasicAuth
To Access the API or the Webfrontend a HTTP-Basic-Auth authorisation is requested. To use this Setting the Username and Password for the Basic-Auth has to be set as environment Variable:

* GOMAILADMIN_AUTH_HTTPBasic_Username -> Username
* GOMAILADMIN_AUTH_HTTPBasic_Password -> Password

## AdminMail
To access the Webfrontend a user can login with a e-mail/password combination from the Mailserver. If there are no E-Mail Accounts at the Mailserver this method will not work! To access the API an API-Key can send as HTTP-Header.

To setup a login via the GUI add the following environment Variable 

* GOMAILADMIN_AUTH_AdminMail_MAIL

e.g.

```
GOMAILADMIN_AUTH_AdminMail_MAIL=admin@example.de,othermail@example.com
```

To access only the API the following environment Variable, if it is not set or empty a API-Access is not possible (or you need to login via E-Mail and Password)

* GOMAILADMIN_AUTH_AdminMail_API

```
GOMAILADMIN_AUTH_AdminMail_MAIL=onerealysecretapikey,anotherevenmoreapikey,adshfioöqweshrguiwergtzfhqw7p
```

## DontStart
Just the default, the Project will stop with a panic.

# Auth in older Versions

## 0.1.5 to *current*
If one of the old environment Variable is set (GOMAILADMIN_APIKEY, GOMAILADMIN_APISECRET) the authorisation will always be HTTPBasic. A warning in the log that the fields are deprecated will be appear. 

In the feature the Project will not start with the old environment Variable.

## 0.1.0 to 0.1.5
Only HTTP-Basic-Auth is available, you can use it by setting the following environment Variable

* GOMAILADMIN_APIKEY -> Username
* GOMAILADMIN_APISECRET -> Password
