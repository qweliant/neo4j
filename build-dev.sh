#!/bin/bash
docker-compose -f docker-compose.yml up -d --build
# docker tag tfclient_client qweliant/tfneo4j_objmap:latest
# docker push qweliant/tfneo4j_objmap