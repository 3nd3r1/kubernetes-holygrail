#!/bin/sh

wget -r $WEBSITE_URL -P /usr/share/nginx/html -nH
nginx -g 'daemon off;'
