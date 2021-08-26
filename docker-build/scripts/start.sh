#!/bin/bash

nohup $OPENRESTY_HOME/mate/nginx_mate >>$OPENRESTY_HOME/nginx_mate.stdout.log 2>>$OPENRESTY_HOME/nginx_mate.stderr.log
tail -f /dev/null
