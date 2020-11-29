# Installation from Debian Package

## Download Deb
Download the go-mail-admin deb file from the Release Page and load it to your server, for debian with
```
wget https://github.com/kekskurse/go-mail-admin/releases/download/v0.1.3/go-mail-admin-amd64-0.1.3.deb
```

After it install the Debian Package

```
dpkg -i go-mail-admin-amd64-*.deb
```

Remove the deb file after it

```
rm go-mail-admin-amd64-*.deb 
```

Copy the sample config in /etc/go-mail-admin and change the values
```
cp /etc/go-mail-admin/go-mail-admin.env.sample /etc/go-mail-admin/go-mail-admin.env
```

If you want to enabled the autostart the go-mail-admin run
```
sudo systemctl enable go-mail-admin
```

than start go-mail-admin by run
```
sudo systemctl start go-mail-admin
```

# Installation from Binary
Easy installation to run go-mail-admin

## Create User
To run the script permanent you need to create a new user. Do it by running the following command
```
sudo useradd -m -c "GoMailAdmin" -s /bin/false gomailadmin
```

## Download binary
Download the go-mail-admin binary and save it in the home folder of the new user (check if the link is the newest version at the release page)

```
wget https://github.com/kekskurse/go-mail-admin/releases/download/v0.1.2/go-mail-admin-with-gui-0.1.2 -O /home/gomailadmin/go-mail-admin-with-gui-0.1.2 
```
Change the owner of the binary
```
chown gomailadmin:gomailadmin /home/gomailadmin/go-mail-*
```

Create a link to the newest version
```
sudo ln -s /home/gomailadmin/go-mail-admin-with-gui-0.1.2 /home/gomailadmin/go-mail-admin
```

Make the binary executable

```
chmod u+x /home/gomailadmin/go-mail-admin-*
```

## Run Binary
Set the environment variables to configure the Go-Mail-Admin, e.g.:
```
export GOMAILADMIN_DB="vmail:vmailpassword@tcp(127.0.0.1:3306)/vmail"
export GOMAILADMIN_AUTH_METHOD=None
export GOMAILADMIN_PORT=3001
```

Then you can start the Go-Mail-Admin with the following command
```
./go-mail-admin-with-gui-<VERSION>
```

After that you can open the gui via http at http://servername:3001 (or your specified custom port)

## Create Service file
Create a new file called /etc/systemd/system/go-mail-admin.service
```
sudo nano /etc/systemd/system/go-mail-admin.service
```
copy the content from the [sample file](https://github.com/kekskurse/go-mail-admin/blob/master/docs/go-mail-admin.service) and change the 4 environment settings!

## Enabled and Start Service
Run the following command to enable and start the service. After it the service will autostart always.

```
systemctl enable go-mail-admin.service
systemctl start go-mail-admin.service
```

After that you can call the gui via http at http://servername:3001



# Update

## With Binary
Just download the new Binary for your system and replace it with the old one, after it restart the service with
```
systemctl restart go-mail-admin.service
```

## From Binary to Deb Package 
Stop the Project if it is still running
```
systemctl stop go-mail-admin.service
```
You can remove the Binary and the "gomailadmin" if you dont use it for anything else. Remove the systemd file
```
rm /etc/systemd/system/go-mail-admin.service
```
After it install the newest Version from the Debian Package.