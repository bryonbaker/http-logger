#! /bin/bash
echo "Testing using url $1"

# curl -k -X PUT https://http-logger-route-http-logger.apps.sno.bakerapps.net/ -H "Content-Type: application/json" -d '{"message":"Hello, PUT request!"}'

curl -k -X PUT $1 -H "Content-Type: application/json" -d '{"message":"Hello, PUT request!"}'
