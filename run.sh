#!/bin/bash

source env.sh

echo "Run migration"
if ! command -v migrate &> /dev/null; then
    echo "Migrate is not found!"
    echo "Please install migrate first!"
    echo "Tutorial for Installation (https://github.com/golang-migrate/migrate/tree/master/cmd/migrate)"
    exit 1
fi
migrate -database "mysql://$MYSQL_USERNAME:$MYSQL_PASSWORD@tcp($MYSQL_HOST:$MYSQL_PORT)/$MYSQL_DATABASE_NAME" -path migrations/mysql up

echo "Building application"
go build -o ./bin/

echo "Application is running!"
./bin/golang-rest-clean-architecture