# Dev
## Project structure
In the folder "mailserver-configurator-interface" you'll find the GO project which provides an HTTP API for the mailserver config. The "mailserver-configurator-client" is a VUEJS interface for that HTTP API.

## Compile
### Mailserver-Configurator-Client
Open the project folder and build the Vue.js project

```
make client-build
```

### Mailserver-Configurator-Interface
Open the project folder. Make sure you have all dependencies

```
make interface-install-deps
```

Create a "public" folder and copy the dist folder from the client into it

```
make interface-copy-client
```

Then copy the static files to the go project and build the GO project
```
make interface-build
```

# Release a new Version
```
make release VERSION=0.1.3
```

