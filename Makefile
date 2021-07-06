client-build:
	rm -f -r  ./mailserver-configurator-client/dist/
	cd ./mailserver-configurator-client; npm install
	cd ./mailserver-configurator-client; npm run build

goreleaser-vue:
	mkdir ./mailserver-configurator-interface/public/
	cp -r ./mailserver-configurator-client/dist/* ./mailserver-configurator-interface/public/

interface-copy-client:
	rm -r ./mailserver-configurator-interface/public/*
	cp -r ./mailserver-configurator-client/dist/* ./mailserver-configurator-interface/public/

interface-install-deps:
	go get github.com/go-sql-driver/mysql
	go get github.com/go-chi/cors
	go get github.com/unrolled/render
	go get github.com/go-chi/chi
	go get github.com/rakyll/statik
	go get github.com/99designs/basicauth-go
	go get github.com/gomodule/redigo/redis
	go get -u github.com/rs/zerolog/log
	go get gopkg.in/unrolled/render.v1

interface-build:
	cd ./mailserver-configurator-interface; ~/go/bin/statik -f -src=./public
	cd ./mailserver-configurator-interface; go build -o ../go-mail-admin-with-gui ./

statik:
	cd ./mailserver-configurator-interface; ~/go/bin/statik -f -src=./public

interface-compile:
	echo "Compiling for every OS and Platform"
	cd ./mailserver-configurator-interface; ~/go/bin/statik -f -src=./public
	rm -f ./bin/*
	cd ./mailserver-configurator-interface; CGO_ENABLED=0 GOOS=linux GOARCH=386 go build -o ../bin/go-mail-admin-with-gui-linux-386 ./
	cd ./mailserver-configurator-interface; CGO_ENABLED=0 GOOS=linux GOARM=7 GOARCH=arm go build -o ../bin/go-mail-admin-with-gui-linux-arm ./
	cd ./mailserver-configurator-interface; CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o ../bin/go-mail-admin-with-gui-linux-arm64 ./
	cd ./mailserver-configurator-interface; CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ../bin/go-mail-admin-with-gui-linux-amd64 ./

deb-i386:
	cp ./bin/go-mail-admin-with-gui-linux-386 ./resources/debpkg-i386/opt/go-mail-admin/go-mail-admin-with-gui-linux-386
	cd ./resources/; dpkg-deb --build ./debpkg-i386
	mv ./resources/debpkg-i386.deb ./bin/go-mail-admin-i386.deb

deb-amd64:
	cp ./bin/go-mail-admin-with-gui-linux-amd64 ./resources/debpkg-amd64/opt/go-mail-admin/go-mail-admin-with-gui-linux-amd64
	cd ./resources/; dpkg-deb --build ./debpkg-amd64
	mv ./resources/debpkg-amd64.deb ./bin/go-mail-admin-amd64.deb

deb-armhf:
	cp ./bin/go-mail-admin-with-gui-linux-arm ./resources/debpkg-armhf/opt/go-mail-admin/go-mail-admin-with-gui-linux-arm
	cd ./resources/; dpkg-deb --build ./debpkg-armhf
	mv ./resources/debpkg-armhf.deb ./bin/go-mail-admin-armhf.deb

changeVersion:
	sed -i -E "s/Version:?.*/\Version: $(VERSION)/" resources/debpkg-amd64/DEBIAN/control
	sed -i -E "s/Version:?.*/\Version: $(VERSION)/" resources/debpkg-armhf/DEBIAN/control
	sed -i -E "s/Version:?.*/\Version: $(VERSION)/" resources/debpkg-i386/DEBIAN/control

version:
	mv bin/go-mail-admin-amd64.deb bin/go-mail-admin-amd64-$(VERSION).deb
	mv bin/go-mail-admin-armhf.deb bin/go-mail-admin-armhf-$(VERSION).deb
	mv bin/go-mail-admin-i386.deb bin/go-mail-admin-i386-$(VERSION).deb
	mv bin/go-mail-admin-with-gui-linux-amd64 bin/go-mail-admin-with-gui-linux-amd64-$(VERSION)
	mv bin/go-mail-admin-with-gui-linux-arm64 bin/go-mail-admin-with-gui-linux-arm64-$(VERSION)
	mv bin/go-mail-admin-with-gui-linux-arm bin/go-mail-admin-with-gui-linux-arm-$(VERSION)
	mv bin/go-mail-admin-with-gui-linux-386 bin/go-mail-admin-with-gui-linux-386-$(VERSION)

run:
	GOMAILADMIN_DB="vmail:vmailpassword@tcp(127.0.0.1:3306)/vmail" go run ./mailserver-configurator-interface

init-test:
	docker-compose down
	docker-compose rm
	docker-compose up -d
	sleep 10
test:
	GOMAILADMIN_DB="vmail:vmailpassword@tcp(127.0.0.1:3306)/vmail" go test ./mailserver-configurator-interface


build: client-build interface-copy-client interface-install-deps interface-build
compile: client-build interface-copy-client interface-install-deps interface-compile
all: compile deb-i386 deb-amd64 deb-armhf
release: changeVersion all version
