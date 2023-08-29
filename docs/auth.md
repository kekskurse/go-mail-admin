# Auth
You can setup the auth method via the enviroment varieable `GOMAILADMIN_AUTH_METHOD`, the following values are available:

* None
* HTTPBasicAuth
* Username

Based on that Setting some additional parameters are needed

## None
No additional parameters are required, no auth for the Admin interface, anyone who can access the Frontend can change any settings. May be used if the authorisation happens at another place.

## HTTPBasicAuth
To Access the API or the Webfrontend a HTTP-Basic-Auth authorisation is requested. To use this Setting the Username and Password for the Basic-Auth has to be set as environment Variable:

* GOMAILADMIN_AUTH_HTTPBasic_Username -> Username
* GOMAILADMIN_AUTH_HTTPBasic_Password -> Password

## Username
Access via Username/Password with the Webfronend.

Make sure to configure: 

* `GOMAILADMIN_AUTH_METHOD=Username`
* `GOMAILADMIN_AUTH_Username_Username` -> Set your username
* `GOMAILADMIN_AUTH_Username_Password` -> Set your Password

for example - in the systemd init file: 

```
...
Environment="GOMAILADMIN_AUTH_METHOD=Username"
Environment="GOMAILADMIN_AUTH_Username_Username=myusername"
Environment="GOMAILADMIN_AUTH_Username_Password=mypassword"
...
```


## DontStart
Just the default, the Project will stop with a panic.

# Auth in older Versions

## 0.2.0 to *current*
If one of the old environment Variable is set (GOMAILADMIN_APIKEY, GOMAILADMIN_APISECRET) the authorisation will always be HTTPBasic. A warning in the log that the fields are deprecated will be appear. 

In the feature the Project will not start with the old environment Variable.

## 0.1.0 to 0.1.X
Only HTTP-Basic-Auth is available, you can use it by setting the following environment Variable

* GOMAILADMIN_APIKEY -> Username
* GOMAILADMIN_APISECRET -> Password
