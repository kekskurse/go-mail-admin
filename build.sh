if [ $# -eq 0 ]
then
  echo "Call ./build-with-gui.sh <version>"
  exit
fi
cd mailserver-configurator-client
npm install
npm build
cd ..
cd mailserver-configurator-interface
mkdir public
cd public
cp -r ../../mailserver-configurator-client/dist/* ./
cd ..
~/go/bin/statik -src=./public
go build ./
cd ..
mv mailserver-configurator-interface/mailserver-configurator-interface go-mail-admin-with-gui-$1