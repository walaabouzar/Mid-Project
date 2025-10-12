#!/bin/bash

echo "==== 1. GET all users ===="
curl -s http://localhost:8080/users | jq
echo -e "\n"

echo "==== 2. GET all agendas (before creation) ===="
curl -s http://localhost:8080/agendas | jq
echo -e "\n"

echo "==== 3. CREATE agenda 1 ===="
agenda1=$(curl -s -X POST http://localhost:8080/agendas \
     -H "Content-Type: application/json" \
     -d '{"group_id":"G1","ical_url":"https://example.com/calendar1.ics"}' | jq '.id')
echo $agenda1
echo -e "\n"

echo "==== 4. CREATE agenda 2 ===="
agenda2=$(curl -s -X POST http://localhost:8080/agendas \
     -H "Content-Type: application/json" \
     -d '{"group_id":"G2","ical_url":"https://example.com/calendar2.ics"}' | jq '.id')
echo $agenda2
echo -e "\n"

echo "==== 5. GET all agendas (after creation) ===="
curl -s http://localhost:8080/agendas | jq
echo -e "\n"

echo "==== 6. GET agenda by ID 1 (agenda1) ===="
curl -s http://localhost:8080/agendas/$agenda1 | jq
echo -e "\n"

echo "==== 7. UPDATE agenda ID 1 (agenda1) ===="
curl -s -X PUT http://localhost:8080/agendas/$agenda1 \
     -H "Content-Type: application/json" \
     -d '{"group_id":"G1-updated","ical_url":"https://example.com/calendar1-updated.ics"}' | jq
echo -e "\n"

echo "==== 8. DELETE agenda ID 2 (agenda2) ===="
curl -s -X DELETE http://localhost:8080/agendas/$agenda2 | jq
echo -e "\n"

echo "==== 9. GET all agendas (after deletion of agenda2) ===="
curl -s http://localhost:8080/agendas | jq
echo -e "\n"

echo "==== 10. DELETE all agendas ===="
curl -s -X DELETE http://localhost:8080/agendas | jq
echo -e "\n"

echo "==== 11. GET all agendas (after DELETE ALL) ===="
curl -s http://localhost:8080/agendas | jq
echo -e "\n"
