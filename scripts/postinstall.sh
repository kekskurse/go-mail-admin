#/bin/sh
FILE=/etc/go-mail-admin/go-mail-admin.env
if [ -f "$FILE" ]; then
    systemctl daemon-reload
    systemctl restart go-mail-admin
else
    cp /etc/go-mail-admin/go-mail-admin.env.sample /etc/go-mail-admin/go-mail-admin.env
fi
