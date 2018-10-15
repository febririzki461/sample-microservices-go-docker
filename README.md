<p align="center"><img src="https://www.cbncloud.co.id/wp-content/uploads/2017/02/logo-header-fix.png"></p>

<p align="center">
<a href="https://travis-ci.org/febririzki461/sample-microservices-go-docker.svg?branch=master"><img src="https://travis-ci.org/febririzki461/sample-microservices-go-docker.svg?branch=master" alt="Build Status"></a>
</p>

## Introduction
Microservice architecture, or simply microservices, is a distinctive method of developing software systems that tries to focus on building single-function modules with well-defined interfaces and operations. The trend has grown popular in recent years as Enterprises look to become more Agile and move towards a DevOps and continuous testing. Microservices can help create scalable, testable software that can be delivered weekly, not yearly.

## Docker 
Docker is a tool designed to make it easier to create, deploy, and run applications by using containers. Containers allow a developer to package up an application with all of the parts it needs, such as libraries and other dependencies, and ship it all out as one package. By doing so, thanks to the container, the developer can rest assured that the application will run on any other Linux machine regardless of any customized settings that machine might have that could differ from the machine used for writing and testing the code.

## Container
A container is a standard unit of software that packages up code and all its dependencies so the application runs quickly and reliably from one computing environment to another. A Docker container image is a lightweight, standalone, executable package of software that includes everything needed to run an application: code, runtime, system tools, system libraries and settings.

## Table of contents
<!--ts-->
  * [Install](#install)
  * [Docker-build](#docker-build)
    * [Logins](#logins)
    * [Users](#users)
  * [Docker](#docker)
  * [Docker-port](#docker-port)
  * [Push-docker-hub](#push-docker-hub)
  * [Testing](#testing)
    * [Testing-curl-logins](#testing-curl-logins)
    * [Testing-curl-users](#testing-curl-users)
<!--te-->

## Install

Please Clone your server
```bash
git clone git@github.com:febririzki461/sample-microservices-go-docker.git
```

## Docker-build
The docker build command builds Docker images from a Dockerfile and a “context”. A build’s context is the set of files located in the specified PATH or URL. The build process can refer to any of the files in the context. For example, your build can use a COPY instruction to reference a file in the context.so this sample microservices have 2 service : 

## Logins
```bash
cd logins

docker build -t logins .
```
in docker images you can see <b>logins</b> in images docker

## Users
```bash
cd users

docker build -t users .
```
in docker images you can see <b>Users</b> in images docker

## Docker
in docker you can start container, example :
```bash
docker run -d logins
```
noted: 
<b> by default follow the installation port in the docker image </b>

## Docker-port
same with Docker, but you can change docker port, example :
```bash
docker run -p 5000:9080 -d logins
```
