#!/bin/bash

cp $OPENRESTY_HOME/mate/conf/nginx.conf $OPENRESTY_HOME/nginx/conf/nginx.conf
mkdir -p $OPENRESTY_HOME/logs
mkdir -p $OPENRESTY_HOME/nginx/conf/tcp
mkdir -p $OPENRESTY_HOME/nginx/cert
cd $OPENRESTY_HOME/nginx/cert
openssl req -new -newkey rsa:2048 -days 365 -nodes -x509 -keyout server.key -out server.crt -subj '/C=AA/ST=AA/L=AA/O=AA Ltd/OU=AA/CN=*.ca.com/emailAddress=shoothzj@gmail.com'
nohup $OPENRESTY_HOME/mate/nginx_mate >>$OPENRESTY_HOME/logs/nginx_mate.stdout.log 2>>$OPENRESTY_HOME/logs/nginx_mate.stderr.log

