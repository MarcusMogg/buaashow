#!/bin/bash
# please run with root or sudo

mysql_user="buaashow"
mysql_password="rott@123"
db_name="buaashow"

mysql -uroot -hlocalhost -e "DROP DATABASE IF EXISTS ${db_name};"
mysql -uroot -hlocalhost -e "CREATE DATABASE ${db_name} CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;"
mysql -uroot -hlocalhost -e "create user IF NOT EXISTS ${mysql_user}@'127.0.0.1' identified by \"${mysql_password}\";"
mysql -uroot -hlocalhost -e "create user IF NOT EXISTS ${mysql_user}@'localhost' identified by \"${mysql_password}\";"
mysql -uroot -hlocalhost -e "grant all privileges on ${db_name}.* to ${mysql_user}@'127.0.0.1' with grant option;"
mysql -uroot -hlocalhost -e "grant all privileges on ${db_name}.* to ${mysql_user}@'localhost' with grant option;"