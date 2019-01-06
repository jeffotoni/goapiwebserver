# Golang / Server Front-End
A server Go to render html and serve the front end to communicate with an API in the back end.
The server is an example of using libs html / template in Go, go-bindata was used to generate a .go with the templates in html that was created and thus we maintain the legibility of the templates in native Html and optimizing even more our server .

We removed the lib fmt from our front-end server, we used the lib bufio by typing straight into stdout.

The front-end server is independent and communicates with an API on the back end.

## Technologies 

	- Front End 
	  - Golang Render HTML
	  - Html
	  - Js
	  - Css

	- Back End  
	  - Golang

## Install

```go

$ go get github.com/jeffotoni/goapiwebserver
$ cd gofrontend
$ make build

```

## Docker

$ docker run -d -p 5001:5001 -e HOST_SERVER=0.0.0.0 --restart=always --name gofronted jeffotoni/gofrontend
