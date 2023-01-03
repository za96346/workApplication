#!/bin/bash
CONTAINER=workapp_mysql
DB_NAME=workApplication
FILENAME=/backup/workApplication/${DB_NAME}_$(date "+%Y%m%d_%H%M%S").sql

docker exec ${CONTAINER} sh -c "mysqldump -u root ${DB_NAME}" > ${FILENAME}