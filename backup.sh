#!/bin/bash
CONTAINER=workapp_mysql
DB_NAME=workApplication
FILENAME=~/backup/workApplication/"$(date "+%Y%m%d_%H%M%S")".sql
expect <<EOF
spawn echo "fuck you" > $FILENAME
expect {
    "*Enter password*" {
        send "siou0722\r"
    }
}
expect eof
EOF