#!/usr/bin/env sh

r=$(cat /dev/urandom | head -n 20 | md5sum | head -c 20)
rm config.yaml

# shellcheck disable=SC2129
echo "WebServer:" >> config.yaml
echo "  Address: $WEB_ADDRESS" >> config.yaml
echo "  Port: $WEB_PORT" >> config.yaml
echo "  JwtSecret: $r" >> config.yaml
echo "Database:" >> config.yaml
echo "  Type: $DB_TYPE" >> config.yaml
echo "  Address: $MYSQL_ADDRESS" >> config.yaml
echo "  Port: $MYSQL_PORT" >> config.yaml
echo "  User: $MYSQL_USER" >> config.yaml
echo "  Password: $MYSQL_PASSWORD" >> config.yaml
echo "  Database: $MYSQL_DATABASE" >> config.yaml
echo "  Prefix: $MYSQL_TABLE_PREFIX" >> config.yaml
echo "Mode: release" >> config.yaml

echo "=====start====="
./server s
echo "======end======"