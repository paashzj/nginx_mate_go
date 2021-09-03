#!/bin/bash

cp $OPENRESTY_HOME/mate/conf/nginx.conf $OPENRESTY_HOME/nginx/conf/nginx.conf
mkdir $OPENRESTY_HOME/nginx/conf/tcp
mkdir $OPENRESTY_HOME/mate/storage
mkdir $OPENRESTY_HOME/mate/storage/static-tcp-route
nohup $OPENRESTY_HOME/mate/nginx_mate >>$OPENRESTY_HOME/nginx_mate.stdout.log 2>>$OPENRESTY_HOME/nginx_mate.stderr.log
tail -f /dev/null
