if [ $# -eq 0 ]
then
  echo "Call ./build-with-gui.sh <version>"
  exit
fi
echo ">Build Client"
cd mailserver-configurator-client
npm install
npm run build
cd ..
echo "> Build Interface"
cd mailserver-configurator-interface
rm -r -f public
mkdir public
cd public
cp -r ../../mailserver-configurator-client/dist/* ./
cd ..
~/go/bin/statik -f -src=./public
go build ./
cd ..
mv mailserver-configurator-interface/mailserver-configurator-interface go-mail-admin-with-gui-$1
