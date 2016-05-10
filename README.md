<img src="client/resources/img/icon.png" width="353"><br />
[![GoDoc](https://godoc.org/github.com/Geodan/gost?status.svg)](https://godoc.org/github.com/Geodan/gost)
[![Build Status](http://beta.drone.io/api/badges/drone/drone/status.svg)](https://drone.io/github.com/Geodan/gost/latest)
[![Go Report Card](https://goreportcard.com/badge/geodan/gost)](https://goreportcard.com/report/geodan/gost)
[![Coverage Status](https://coveralls.io/repos/github/Geodan/gost/badge.svg?branch=master)](https://coveralls.io/github/Geodan/gost?branch=master)
[![Join the chat at https://gitter.im/Geodan/gost](https://badges.gitter.im/Geodan/gost.svg)](https://gitter.im/Geodan/gost?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)<br />

GOST (Go-SensorThings) is a sensor server written in Go. It implements the [OGC SensorThings API] (http://ogc-iot.github.io/ogc-iot-api/api.html) standard.

## Disclaimer

GOST is alpha software and is not considered appropriate for customer use. Feel free to help development :-)

## License

GOST licensed under [MIT](https://opensource.org/licenses/MIT).

## Getting started

1] Install GoLang (https://golang.org/)

2] Install Postgresql (http://www.postgresql.org/)

3] Clone code

```sh
git clone https://github.com/Geodan/gost.git
```
4] Get dependencies

```sh
go get -t
```
5] Change connection to database

Edit config.yaml or set enviroment settings

6] Create database

```sh
go run main.go -install ./scripts/createdb.sql
```

7] Start

```sh
go run main.go
```

8] Go in browser to http://localhost:8080

## Sample requests

For sample requests (setting up sensors/datastreams/things and adding observations) see the tests in the [playground](test/playground_tests.md). 
For a complete collection of working requests install Postman and import the [Postman file](test/GOST.json.postman_collection) 

## Startup flags

-config : specify the config file (default config.yaml)

-install database_script_file: creates the database schema

## Configuration

Default file: config.yaml

- port: port of webserver
- externalUri: External uri for links in responses
- database.host: name or ip of database server
- database.port: port of database server
- database.user: username of database login
- database.password: password of database login
- database.database: database name
- database.schema: database schema
- database.ssl: use ssl flag

The following configuration parameters can be overruled 
from the following environment variables:
gost_db_host, gost_db_port, gost_db_user, gost_db_password

Example setting Gost enviroment variable on Windows:

set gost_db_host=192.168.10.40

Example setting Gost enviroment variable on Mac/Linux:

export gost_db_host=192.168.10.40

## Dependencies

[yaml v2](https://github.com/go-yaml/yaml)<br />
[pq](https://github.com/lib/pq)<br />
[mux](https://github.com/gorilla/mux)<br />
[Paho](https://github.com/eclipse/paho.mqtt.golang)<br />

## Roadmap

- Complete implementation of the OGC SensorThings spec
- Tests!
- MQTT
- Frontend
- Benchmarks
- Different storage providers such as MongoDB (Now using PostgreSQL)

## TODO

[See wiki](https://github.com/Geodan/gost/wiki/TODO)
