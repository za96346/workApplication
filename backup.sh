#!/bin/bash
CONTAINER=workapp_mysql
DB_NAME=workApplication
FILENAME=~/backup/workApplication/"$(date "+%Y%m%d_%H%M%S")".sql
set timeout 10s
expect << EOF
spawn docker exec -i ${CONTAINER} sh -c "mysqldump -u root -p ${DB_NAME}" > ${FILENAME}
expect {
    "*Enter password:*" {
        send "siou0722\r"
    }
}
expect eof
EOF
echo "success" >&2