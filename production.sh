#!/usr/bin/env bash
docker stop questionnaire_production_container
docker rm questionnaire_production_container
docker rmi questionnaire_production_image
docker build -t questionnaire_production_image .
docker run --name questionnaire_production_container -e TZ=Asia/Almaty -d -p 1001:8000 questionnaire_production_image