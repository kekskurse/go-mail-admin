client-build:
	rm -f -r  ./mailserver-configurator-client/dist/
	cd ./mailserver-configurator-client; npm install
	cd ./mailserver-configurator-client; npm run build

interface-copy-client:
	rm -r ./mailserver-configurator-interface/public/*
	cp -r ./mailserver-configurator-client/dist/* ./mailserver-configurator-interface/public/

interface-install-deps:
	go get github.com/rakyll/statik

interface-build:
	cd ./mailserver-configurator-interface; ~/go/bin/statik -f -src=./public
	cd ./mailserver-configurator-interface; go build -o ../go-mail-admin-with-gui ./

statik:
	cd ./mailserver-configurator-interface; ~/go/bin/statik -f -src=./public

run:
	GOMAILADMIN_DB="vmail:vmailpassword@tcp(127.0.0.1:3306)/vmail" GOMAILADMIN_V3="on" GOMAILADMIN_AUTH_Username="test" GOMAILADMIN_AUTH_Password="test"  go run ./mailserver-configurator-interface

gorelease-vue:
	go get github.com/rakyll/statik
	rm -f -r  ./mailserver-configurator-client/dist/
	cd ./mailserver-configurator-client; npm install
	cd ./mailserver-configurator-client; npm run build
	mkdir ./mailserver-configurator-interface/public/
	cp -r ./mailserver-configurator-client/dist/* ./mailserver-configurator-interface/public/
	cd ./mailserver-configurator-interface; ~/go/bin/statik -f -src=./public

init-test:
	docker-compose down
	docker-compose rm
	docker-compose up -d
	sleep 10

test:
	GOMAILADMIN_DB="vmail:vmailpassword@tcp(127.0.0.1:3306)/vmail" GOMAILADMIN_V3="on" go test ./mailserver-configurator-interface


build: client-build interface-copy-client interface-install-deps interface-build
