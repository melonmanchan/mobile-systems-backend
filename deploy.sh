#!/usr/bin/env bash
set -e

if [ $# -ne 1 ]; then
    echo $0: Missing deployment host
    exit 1
fi

DEPLOY_HOST=$1

echo "Building binary..."
env GOOS=linux GOARCH=amd64 go build -o server-deploy .

echo "Packaging assets..."
mkdir -p build
cp server-deploy build
cp -r migrations build

echo "Stopping server..."
ssh $DEPLOY_HOST "source /etc/profile; pkill server-deploy 2>/dev/null"

echo "Copying assets..."
scp -r build $DEPLOY_HOST:/root

echo "Restarting server..."
ssh "$DEPLOY_HOST" "source /etc/profile; pkill server-deploy; nohup /root/build/server-deploy > tutee.out 2> tutee.err.out < /dev/null &"

rm -r build
echo "Done!"
