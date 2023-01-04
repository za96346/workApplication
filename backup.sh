#!/bin/bash/expect
set CONTAINER "workapp_mysql"
set DB_NAME "workApplication"
set FILENAME ~/backup/workApplication/"$(date "+%Y%m%d_%H%M%S")".sql
set timeout 10s

spawn docker exec -i $CONTAINER sh -c "mysqldump -u root -p $DB_NAME" > $FILENAME
expect {
    "*Enter password:*" {
        send "siou0722\r"
        send "exit\r"
    }
}
