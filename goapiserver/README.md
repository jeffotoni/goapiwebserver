# Go Server Back-end

This API is a server that plays the backend role. It is responsible for user authentication via jwt and access to the other apis.

The backend was created in standard library, we used some libs like jwt, and tollbooth.

We created a standard for backend development using standard library, we create our own and well-defined middlwares.

A standard was born, which has yet to improve, but what has been implemented is already very easy to scale.

The great advantage of this model is its portability and scalability, we leave the limit of requests above 10k req / s and connections by clients above 15k.

We run the wrk to do some stress tests and the curl in a for.

## Technologies 
	- Back-end Golang: version 1.11.2

## Pacakges

I'm using the go mod, so do not need to do go get, but in case you need to.

```sh

$ go get -u github.com/dgrijalva/jwt-go
$ go get -u github.com/didip/tollbooth
$ go get -u github.com/go-sql-driver/mysql

```

## Golang mod

The repository is already all correct, but in case you want to test go mod just remove in the root the api go.sum and go.mod.
And run the go mod:

```sh

$ go mod init github.com/jeffotoni/goapiwebserver/goapiserver
$ go build
$ go vendor

```

## Generate the keys

To generate the JWT token we are using openssl, do not need to generate if you do not want to, we take the keys to memory to avoid opening the file for generating the keys.
We did with the configuration files in the same way nothing on disk everything in environment variables.

```

$ openssl genrsa -out private.rsa 1024
$ openssl rsa -in private.rsa -pubout > public.rsa.pub

```

## Structure Go Front-end

	- cert
	- config
	- handler
		- createlogin.go
		- endpoint.go
		- home.go
		- ping.go
		- route.go
		- token.go
	- logg
	- middleare
		- middleware-adapter.go
		- middleware-auth-jwt.go
		- middleware-limit.go
		- middleware-maxlicent.go
		- middleware-notify.go
	- mocks
		- ping
		- mysql
	- models
	- pkg
		- assets
		- email
		- gjwt
		- mysql
		- template
		- util
	- repo
	- repo
	- templates

## Docker

So that goapiserver works correctly you will need to set all environment variables

```docker

$ docker run -d -p 5002:5002 \
-e REQUEST_SEC=10000 \
-e DB_HOST=xxxxxxxxx \
-e DB_NAME=goapiserver \
-e DB_USER=goapiserver \
-e DB_PASSWORD=xxxxxxx \
-e DB_PORT=3306 \
-e DB_SORCE=mysql \
-e HOST_SERVER="0.0.0.0" \
--restart=always \
--name goapiserver jeffotoni/goapiserver

$ docker -f logs goapiserver

```

## Install build

You can download the project in github and do the build, so golang version 1.11.2 has that is installed on your machine, you can check here https://golang.org/dl/

```go

$ git clone http://github.com/jeffotoni/goapiwebserver
$ cd goapiwebserver/goapiserver
$ make build

```

## stress test

```sh

$ for ((i=1;i<=10000000;i++)); do curl  -X POST "localhost:5002/api/v1/ping"; done

```

```sh

wrk -t2 -c300 -d30s -R10000 http://localhost:5002/api/v1/ping

```

## mocks and simulate

Here we will find bench connection tests and ping token tests
For the mocks to run the goapiserver needs to be running.

you can choose the way you want to start in the application, or using docker or doing build.
Remembering that the mysql database connection is all done via environment variables.

```go

$ git clone http://github.com/jeffotoni/goapiwebserver
$ cd goapiwebserver/goapiserver/mocks/mysql
$ go run mysql

$ cd goapiwebserver/goapiserver/mocks/ping
$ go run ping

```

## curl simulation

### Get Token JWT

```sh

$ curl -H "Authorization: Basic <keyuser>=:<keyaccess>" localhost:5002/api/v1/token

```

## out json token

```json
{
	"token":"xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
	"expires":"2019-02-06",
	"message":"use the token to access the endpoints"
}

```

### Auth JWT Login

```sh

$ curl -X POST -H "Content-Type: application/x-www-form-urlencoded" \
	-H "Authorization: Bearer xxxxxxxxxxxxx" \
	-d "email=putnamehere&password=passhere" \
	localhost:5002/api/v1/login

```

### Create New User

```sh

$ curl -X POST localhost:5002/api/v1/user \
-H "Content-Type: application/x-www-form-urlencoded" \
-H "Authorization: Bearer xxxxxxxxxxxxxx" \
-d 'email=jeff.otoni@gmail.com' \
-d 'firstname=jefferson' \
-d 'lastname=otoni lima' \
-d 'phone=987338383893' \
-d 'password=12334'

```

```json

{
	"status":"success",
	"message":"create user new here."
}

```
### It's User Exist

```sh

$ curl -X GET localhost:5002/api/v1/login/<email> \
-H "Content-Type: application/x-www-form-urlencoded" \
-H "Authorization: Bearer xxxxxxxxxxxxxxxxxxx"

```

