# Golang / Api Webserver

This project is an organization model and with a practical example of how to create a client server API in golang to render html and communicate with a backend API.
We have two APIs, one for front-end and one for the back-end, the front-end API consumes back-end APIs all native to Go, without using golang and Js frameworks.

It was created APIs in the same structure to facilitate at the beginning all the work, are two strands or split into two repository or continue in a single as google a times ago presented that they use with a single repositories their projects and pass the Gigabytes. The client and back-end apis are independent and can make their own lives. 

We use the standard rEST, we use in Go standard Library and some libs to help us compose our APIS. 

The important thing is that I have been able to create a template, a standard for backend and client, of course there are still a lot of improvements and optimizations, tests, mocks, etc. But from the outset the assembled structure is very easy to scale and better using very little server resource. 

I did some stress testing and is holding 5k to 10k of req / s connections to the bakcend API and can still improve it, running on a machine with 2cores and 4G memory only on AWS. 

The best test would be with lambda, but as the time is short, as soon as I pass this correria I will mount a poc and create stress tests with lambda to see the result...

We create Dockerfile and put the services in container, then we can put in docker-compose and as it grows, with some microservices like sending emails that can turn into a microservice, such as the auth of the login that could turn another.

The Docker images were created in hub.docker, they are all public.

The two APIS have been made, and can be found at:

Note: Access to the database is just below, the sensitive data did not go up to
the repositories, it is just below, in case I need something else I am at disposal.

github (I'm always updating)
https://github.com/jeffotoni/goapiwebserver

bitbucket public

https://bitbucket.org/jeffotoni/je_ot

The two apis are in the same repository to facilitate, but are completely independent.

The service has been hosted on AWS and can be viewed working at this link: https://frontend.s3apis.com/login

## Completed features

Creating a mini framework in Go for the backend and APIS frontend.
This scenario has much to improve.
```bash
<table>
<tr>
<td>
- New User Registration 			(100%)
- Login 							        (100%)
- Logout 							        (100%)
- Update Profile 					    (100%)
- Page admin/perfil (blocked) (100%)
- Tests (100%)
- Forgot password 					  (performed 60%)
- Update password 					  (performed 70%)
- Update password 					  (performed 70%)
- integration with google api (30%) [lacked time left last]
- google-map 						      (0%) [lacked time left last]
- address fields 					    (0%) [lacked time left last]
- swagger documentation 			(0%) [lacked time left last]
</td>
</tr>
</table>
```

## INSTALLATION
Note: In the **README** of each API is documenting all the steps and details.

## Install and build (requires go1.11.2 linux / amd64)
Will download, enter the api and do make, the command will compile and run

## Installing APISERVER (BACKEND)

```bash
$ git clone https://github.com/jeffotoni/goapiwebserver.git
$ cd goapiserver
$ make build
```
## Installing APIFRONTEND (FRONTEND)
$ cd gofrontend
$ make build

## DOCKER
If you prefer you can run in docker but remember that to leave on the same machine would 
have to configure the network or the domain of the backend.
You can use nginx to serve as a proxy if you have more services running on the same machine.

## API - BACKEND

```bash
$ docker run -p 5002:5002  \
-e REQUEST_SEC=10000 \
-e DB_HOST="xxxx" \
-e DB_NAME=xxxxx  \
-e DB_USER=xxxxx \
-e DB_PASSWORD=xxxxx \
-e and DB_PORT=xxxx \
-e DB_SORCE=mysql \
-e HOST_SERVER=0.0.0.0 \
--rm --name goapiserver jeffotoni / goapiserver
```

## API - FRONTEND

```bash
$ docker run -p 5001:5001 \
-e API_SCHEME="https" \
-e API_URL=xxxxx \
-e HOST_SERVER=0.0.0.0 \
-e API_PORTA="" \
--restart=always --name gofronted jeffotoni/gofrontend
```

## Simulations were created using cURL, everything is in the README of each API.

In the BACK-END API, there was no use of the lib to manage
the routes and creating the endpoints was all done native standard lib.
Some middlewares have been created.
I used go mod to manage the packages so in go.mod you will be able to see
the packages that I used.

## Libs used:

```bash
jwt-go v3.2.0
tollbooth
go-sql-driver
gorilla/mux
```

In the FRONT-END API, gorillamux was used to manage the routes and sessions.
