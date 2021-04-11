#!/bin/bash
# please run with root or sudo

mysql_user="buaashow"
mysql_password="rott@123"
mysql_host="127.0.0.1"
db_name="buaashow"


mysql -uroot -h${mysql_host} -e "DROP DATABASE IF EXISTS ${db_name};"
mysql -uroot -h${mysql_host} -e "CREATE DATABASE ${db_name} CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;"
mysql -uroot -h${mysql_host} -e "create user IF NOT EXISTS ${mysql_user}@'${mysql_host}' identified by \"${mysql_password}\";"
mysql -uroot -h${mysql_host} -e "grant all privileges on ${db_name}.* to ${mysql_user}@'${mysql_host}' with grant option;"