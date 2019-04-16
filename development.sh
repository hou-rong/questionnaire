#!/usr/bin/env bash
docker stop questionnaire_development_container
docker rm questionnaire_development_container
docker rmi questionnaire_development_image
docker build -t questionnaire_development_image .
docker run --name questionnaire_development_container -e TZ=Asia/Almaty -d -p 1000:8000 questionnaire_development_image