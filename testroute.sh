#!/bin/bash
echo "==== 1. GET all events ===="
curl -s http://localhost:8080/events | jq
echo -e "\n"
