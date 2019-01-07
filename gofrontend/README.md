# Golang / Server Front-End
A server Go to render html and serve the front end to communicate with an API in the back end.
The server is an example of using libs html / template in Go, go-bindata was used to generate a .go with the templates in html that was created and thus we maintain the legibility of the templates in native Html and optimizing even more our server .

We removed the lib fmt from our front-end server, we used the lib bufio by typing straight into stdout.

The front-end server is independent and communicates with an API on the back end.

## Technologies 

	- Front End Golang: version 1.11.2
	  - Golang Render HTML
	  - Html
	  - Js
	  - Css
	  - img

	- Back End  
	  - Golang


## Structure Go Front-end

	- config
	- handler
	- mocks
	- pkg
		- assets
		- request
		- session
		- util
	- repo
	- templates
	- web
		- static
			- css
				- admin
				- fonts
			- img
			- js

## Docker With the static files in the image

This image has the static files css, js and img.
The disadvantage of this image is that if our web client growth, our deploy, backup etc will take a long time.
One more and more efficient way would be to leave the static files on a CDN.

```docker

$ docker run -d -p 5001:5001 -e HOST_SERVER=0.0.0.0 --restart=always --name gofronted jeffotoni/gofrontend
$ docker logs -f gofrontend

```

## Docker With files without the static files in the image

We can serve static files using docker without having to upload them to image, we can use volume option with -v $(pwd)/web:/web

```docker

$ git clone http://github.com/jeffotoni/goapiwebserver
$ cd goapiwebserver/gofrontend
$ docker run -d -p 5001:5001 -e HOST_SERVER=0.0.0.0 $(pwd)/web:/web --restart=always --name gofronted jeffotoni/gofrontend

```

## Install build

You can download the project in github and do the build, so golang version 1.11.2 has that is installed on your machine, you can check here https://golang.org/dl/

```go

$ git clone http://github.com/jeffotoni/goapiwebserver
$ cd goapiwebserver/gofrontend
$ make build

```