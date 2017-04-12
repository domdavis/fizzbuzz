#!/usr/bin/env bash

docker network create --driver bridge fbn

docker run -d -p 1201:8000 --net=fbn --name fizz fizzbuzz -service fizz
docker run -d -p 1202:8000 --net=fbn --name buzz fizzbuzz -service buzz
docker run -d -p 1203:8000 --net=fbn --name number fizzbuzz -service number
docker run -d -p 8000:8000 --net=fbn --name fizzbuzz fizzbuzz \
    -service fizzbuzz \
    -fizz http://fizz:8000/fizz \
    -buzz http://buzz:8000/buzz \
    -number http://number:8000/number
