#!/usr/bin/env bash
docker run --name questionnaire_container -e TZ=Asia/Almaty -d -p 1000:8000 questionnaire_image
