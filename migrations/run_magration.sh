#!/bin/bash

# Укажите полный путь к local.yaml
CONFIG_FILE="../config/local.yaml"

# Чтение данных из local.yaml
HOST=$(yq e '.database.host' "$CONFIG_FILE")
PORT=$(yq e '.database.port' "$CONFIG_FILE")
DBNAME=$(yq e '.database.dbname' "$CONFIG_FILE")
USER=$(yq e '.database.user' "$CONFIG_FILE")
PASSWORD=$(yq e '.database.password' "$CONFIG_FILE")
SSLMODE=$(yq e '.database.sslmode' "$CONFIG_FILE")

# Формирование строки подключения
DSN="host=$HOST port=$PORT dbname=$DBNAME user=$USER password=$PASSWORD sslmode=$SSLMODE"

# Выполнение миграций
~/go/bin/goose -dir ./migrations postgres "$DSN" up