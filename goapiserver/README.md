# Go Server Back-end

This API is a server that plays the role of back-end. It is responsible for user authentication, application access release, modules, user profile and various other features.


docker run -p 5002:5002 -e HOST_SERVER=0.0.0.0 --rm --name goapiserver jeffotoni/goapiserver

go mod init github.com/jeffotoni/goapiwebserver/goapiserver

for ((i=1;i<=10000000;i++)); do curl  -X POST "localhost:5002/api/v1/ping"; done

-H "Content-Type: application/json\
-H "Accept: application/json\

wrk -t2 -c100 -d30s -R5000 http://localhost:5002/api/v1/ping

$ curl -X POST  localhost:5002/api/v1/ping

$ curl -X POST -H "Authorization: Basic MTIzNDU2IzIwMjA=:MTIzNDU2YWplZmZvdG9uaTIwMjA=" localhost:5002/api/v1/token



