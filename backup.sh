#!/bin/bash
CONTAINER=workapp_mysql
DB_NAME=workApplication
FILENAME=~/backup/workApplication/"$(date "+%Y%m%d_%H%M%S")".sql
docker exec -i $CONTAINER bash -c "mysqldump -u root --password siou0722 -p $DB_NAME" > $FILENAME