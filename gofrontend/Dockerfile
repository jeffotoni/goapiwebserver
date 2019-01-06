# Start by building the application.
# Criando build em gofrontend com distroless
FROM golang:1.11.2 as builder

RUN go get -u github.com/gorilla/sessions
RUN go get -u github.com/gorilla/mux
RUN go get -u github.com/jeffotoni/goapiwebserver/gofrontend

WORKDIR /go/src/gofrontend
COPY . .

# RUN go get -d -v ./...
#RUN go install -v ./...
RUN go build -o gofrontend

# Now copy it into our base image.
FROM gcr.io/distroless/base

COPY --from=builder /go/bin/gofrontend /

CMD ["/gofrontend"]
