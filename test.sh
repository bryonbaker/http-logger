#! /bin/bash
echo "Testing using port $1"

curl -X PUT http://localhost:$1/ \
     -H "Content-Type: application/json" \
     -d '{"message":"Hello, PUT request!"}'

