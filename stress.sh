#!/bin/bash

# Anzahl der Durchläufe
count=8000

for i in $(seq 1 $count)
do
  curl -vk -X DELETE http://localhost:8080/v1/beleg/${i} -H "Content-Type: application/json" -d '{"price": 21.99, "mwst": 7.8, "date": "31.12.2023", "shop": "Migros"}'
  #curl -vk http://localhost:8080/v1/beleg -H "Content-Type: application/json" -d '{"price": 21.99, "mwst": 7.8, "date": "31.12.2023", "shop": "Migros"}'
done
