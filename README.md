
# T3chFest 2018: Crea una API REST con Go [![Build Status](https://travis-ci.org/mresti/go_rest_api.svg?branch=master)](https://travis-ci.org/mresti/go_rest_api) [![Code Report](http://goreportcard.com/badge/mresti/go_rest_api)](https://goreportcard.com/report/mresti/go_rest_api)

Este repositorio es el material usado en el taller del T3chFest 2018.

Más información sobre la charla, en la web oficial del evento [T3chFest 2018](https://t3chfest.uc3m.es/2018/programa/crea-una-api-rest-con-go/)

# Compilar API server

    $ make
    
    $ go build -o bin/api src/main.go
    
*Nota: El makefile genera binarios para la plataforma linux, macos y windows.*
    
# Ejecutar API server

    $ make run
    
    $ go run src/main.go
    
    $ ./bin/api -port 3000        

# Mostrar parámetros del binario API server
    
    $ ./bin/api -help
    
    $ ./bin/api -h

# Ejecutar los tests del API server

    $ make test
    
    $ go test -v -race ./src/...

# Testear los HTTP Endpoints del API server

## Chequear el root path

    $ curl -i localhost:3000/

## Chequear el stats path

    $ curl -i localhost:3000/stats
    
# Desplegar la API en un docker

## Usando docker-compose


### Contruir la infraestructura docker-compose
    
    $ docker-compose build

### Ejecutar la infraestructura docker-compose

    $ docker-compose up -d

### Parar la infraestructura docker-compose

    $ docker-compose down
