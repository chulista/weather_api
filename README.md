# weather_api
this API let you get information about the weather in different locations

Doc links for this method: 
https://pkg.go.dev/github.com/chulista/weather_api?tab=overview 

https://godoc.org/github.com/chulista/weather_api


Design
application
Write business logic
domain
Define interface
repository interface for infrastructure
Define struct
Entity struct that represent mapping to data model
infrastructure
Implements repository interface
Solves backend technical topics
e.x. message queue, persistence with RDB
interfaces
Write HTTP handler and middleware
References:
https://speakerdeck.com/mercari/ja-golang-package-composition-for-web-application-the-case-of-mercari-kauru
http://pospome.hatenablog.com/entry/2017/10/11/023848
https://medium.com/@timakin/go%E3%81%AE%E3%83%91%E3%83%83%E3%82%B1%E3%83%BC%E3%82%B8%E6%A7%8B%E6%88%90%E3%81%AE%E5%A4%B1%E6%95%97%E9%81%8D%E6%AD%B4%E3%81%A8%E7%8F%BE%E7%8A%B6%E7%A2%BA%E8%AA%8D-fc6a4369337



##Docker
###commands:

    docker version 
    docker ps // running process
    docker ps -a // ran process
    docker images // check images
    docker run docker/whalesay cowsay boo //download docker image


    docker build -t elchulito/golang-docker .
    docker run -d -p 8080:8080 elchulito/golang-docker


####remove images
    docker rmi -f CONTAINERID
    
####docker push container
    docker push elchulito/golang-docker:latest

##VirtualBox
### install docker-machine (reference: https://github.com/docker/machine/releases)
    curl -L https://github.com/docker/machine/releases/download/v0.16.2/docker-machine-`uname -s`-`uname -m` >/usr/local/bin/docker-machine && \
    chmod +x /usr/local/bin/docker-machine

### create docker image on virtual box (reference: https://docs.docker.com/machine/reference/create/)
    docker-machine create --driver virtualbox weatherapi
    
To see how to connect your Docker Client to the Docker Engine running on this virtual machine, run: 

    docker-machine env weatherapi


##How To Upload changes on the API
####1º step: push github repo changes    
    git push origin ...
####2º step: build docker image
    docker build -t elchulito/golang-docker .
####3º step: push docker changes
     docker push elchulito/golang-docker:latest
####4º step: go into the virtual box image and pull docker images changes
    docker pull elchulito/golang-docker:latest
####5º step: run DockerImage on VirtualBox
    docker run -p 8080:8080 elchulito/golang-docker
####8º step: get DockerPort:
    docker-machine ls
    NAME         ACTIVE   DRIVER       STATE     URL                         SWARM   DOCKER      ERRORS
    weatherapi   *        virtualbox   Running   tcp://192.168.99.100:2376           v19.03.12
####7º step: test API with Postman
    http method: GET
    url: 192.168.99.100:8080
    