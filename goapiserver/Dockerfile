# Start by building the application.
# Criando build em goapiserver com distroless
FROM golang:1.11.2 as builder

RUN go get -u github.com/didip/tollbooth
RUN go get -u github.com/dgrijalva/jwt-go
RUN go get -u github.com/jeffotoni/goapiwebserver/goapiserver
WORKDIR /go/src/goapiserver

ENV GO111MODULE=on

COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

RUN ls -lh 

# RUN go build -o goapiserver
# Now copy it into our base image.
FROM gcr.io/distroless/base

# copy bin diretory
COPY --from=builder /go/bin/goapiserver /

CMD ["/goapiserver"]
