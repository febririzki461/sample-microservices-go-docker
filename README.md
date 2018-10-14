## Introduction
This is Sample Microservices go docker

## Table of contents
<!--ts-->
  * [Install](#install)
  * [Docker-build](#docker-build)
    * [Logins](#logins)
    * [Users](#users)
    * [Port](#dockerfile)
  * [Docker](#docker-run)
  * [Push-docker-hub](#push-docker-hub)
  * [Testing-curl](#testing-curl)
    * [Testing-logins](#testing-logins)
    * [Testing-users](#testing-users)
<!--te-->

## Install

Please Clone your server
```bash
git clone git@github.com:febririzki461/sample-microservices-go-docker.git
```

## Docker-build
this sample microservices have 2 service : 

## Logins
```bash
cd logins

docker build -t logins .
```

## Users
```bash
cd users

docker build -t users .
```
