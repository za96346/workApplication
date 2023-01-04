#!/bin/bash
CONTAINER=workapp_mysql
DB_NAME=workApplication
FILENAME=~/backup/workApplication/hi.sql
expect <<EOF
spawn docker exec -i $CONTAINER sh -c "mysqldump -u root -p $DB_NAME" > ~/backup/workApplication/hi.sql
expect {
    "*Enter password*" {
        send "siou0722\r"
        send "exit\r"
    }
}
expect eof
EOF
echo $FILENAME
echo $DB_NAME
echo $CONTAINER