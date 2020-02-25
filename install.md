# Installation
Easy installation to run go-mail-admin

## Create User
To run the script permanent you need to create a new user. Do it by running the following command
```
sudo useradd -m -c "GoMailAdmin" -s /bin/false gomailadmin
```

## Download binary
Download the go-mail-admin binary and save it in the home folder of the new user (check if the link is the newest version at the release page)

```
wget https://github.com/kekskurse/go-mail-admin/releases/download/v0.1.1/go-mail-admin-with-gui-0.1.1 -O /home/gomailadmin/go-mail-admin-with-gui-0.1.1 
```
Change the owner of the binary
```
chown gomailadmin:gomailadmin /home/gomailadmin/go-mail-*
```

Create a link to the newest version
```
sudo ln -s /home/gomailadmin/go-mail-admin-with-gui-0.1.1 /home/gomailadmin/go-mail-admin
```

Make the binary executable

```
chmod u+x /home/gomailadmin/go-mail-admin-*
```

## Create Service file
Create a new file called /etc/systemd/system/go-mail-admin.service
```
sudo nano /etc/systemd/system/go-mail-admin.service
```
copy the content from the [sample file](https://github.com/kekskurse/go-mail-admin/blob/master/go-mail-admin.service) and change the 3 Environment settings!

## Enabled and Start Service
Run the following command to enable and start the service. After it the service will autostart always.

```
systemctl enable go-mail-admin.service
systemctl start go-mail-admin.service
```