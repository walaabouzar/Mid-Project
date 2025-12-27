!/bin/bash
echo "==== 1. GET all events ===="
curl -s http://localhost:8090/events | jq
echo -e "\n"

