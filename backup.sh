#!/bin/bash
CONTAINER=workapp_mysql
DB_NAME=workApplication
FILENAME=~/backup/workApplication/$(date "+%Y%m%d_%H%M%S").sql
expect << EOF
spawn docker exec -it ${CONTAINER} sh -c "mysqldump -u root -p ${DB_NAME}" > ${FILENAME}
expect {
    "*Enter password:*" {
        send "siou0722\r"
        send "exit\r"
    }
}
expect eof
EOF
echo "success" > &2