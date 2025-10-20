#!/bin/bash

echo "==== 1. GET all alerts (before creation) ===="
curl -s http://localhost:8080/alerts | jq
echo -e "\n"
echo "==== 2. CREATE alert 1 ===="
alert1=$(curl -s -X POST http://localhost:8080/alerts \
     -H "Content-Type: application/json" \
     -d '{"dest":"student1@example.com","msg":"Exam tomorrow at 10AM","typea":"reminder","agenda":1}' | jq '.id')
echo "Alert 1 ID: $alert1"
echo -e "\n"

echo "==== 3. CREATE alert 2 ===="
alert2=$(curl -s -X POST http://localhost:8080/alerts \
     -H "Content-Type: application/json" \
     -d '{"dest":"student2@example.com","msg":"cours calculabilité prévu à 10h demain","typea":"créneau","agenda":1}' | jq '.id')
echo "Alert 2 ID: $alert2"
echo -e "\n"

echo "==== 4. GET all alerts (after creation) ===="
curl -s http://localhost:8080/alerts | jq
echo -e "\n"

echo "==== 5. GET alert by ID 1 ===="
curl -s http://localhost:8080/alerts/$alert1 | jq
echo -e "\n"

echo "==== 6. UPDATE alert ID 1 ===="
curl -s -X PUT http://localhost:8080/alerts/$alert1 \
     -H "Content-Type: application/json" \
     -d '{"dest":"student1@example.com","msg":"Exam rescheduled to 11AM","typea":"update","agenda":1}' | jq
echo -e "\n"

echo "==== 7. DELETE alert ID 2 ===="
curl -s -X DELETE http://localhost:8080/alerts/$alert2 | jq
echo -e "\n"

echo "==== 8. GET all alerts (after deleting alert 2) ===="
curl -s http://localhost:8080/alerts | jq
echo -e "\n"

echo "==== 9. DELETE all alerts ===="
curl -s -X DELETE http://localhost:8080/alerts | jq
echo -e "\n"

echo "==== 10. GET all alerts (after DELETE ALL) ===="
curl -s http://localhost:8080/alerts | jq
echo -e "\n"

