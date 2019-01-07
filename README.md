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

