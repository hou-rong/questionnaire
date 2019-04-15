#!/usr/bin/env bash
docker stop questionnaire_container
docker rm questionnaire_container
docker rmi questionnaire_image