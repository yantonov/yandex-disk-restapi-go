#!/bin/sh

SCRIPT_DIR=$(cd `dirname $0` && pwd)

CLIENT_ID=$1
CLIENT_SECRET=$2

if [ -z $CLIENT_ID ]; then
    echo "[ ERROR ] client id is not defined"    
fi

if [ -z $CLIENT_SECRET ]; then
    echo "[ ERROR ] client secret is not defined"    
fi

if [ -z $CLIENT_ID ] || [ -z $CLIENT_SECRET ]; then
    echo "\nUsage : run_oauth_example.sh <client_id> <client_secret>\n"
    echo "\t where client id and secret can be configured at http://oauth.yandex.ru"
    exit 1
fi

go run $SCRIPT_DIR/../examples/oauth_example.go -id=$CLIENT_ID -secret=$CLIENT_SECRET
