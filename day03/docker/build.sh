#!bin/bash
docker stop elas
docker rm elas
docker build -t grandpat:v1 .
docker run --name elas -p 9200:9200 -d grandpat:v1