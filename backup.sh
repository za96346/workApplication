#!/bin/bash
CONTAINER=workapp_mysql
DB_NAME=workApplication
FILENAME=~/backup/workApplication/"$(date "+%Y%m%d_%H%M%S")".sql
pwd=siou0722
docker exec -i $CONTAINER bash -c "mysqldump -u root -p$pwd  $DB_NAME" > $FILENAME