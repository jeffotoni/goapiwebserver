# Start by building the application.
# Criando build em gofrontend com distroless
FROM golang:1.11.2 as builder

RUN go get -u github.com/gorilla/sessions
RUN go get -u github.com/gorilla/mux
RUN go get -u github.com/jeffotoni/goapiwebserver/gofrontend

WORKDIR /go/src/gofrontend

ENV GO111MODULE=on

COPY . .

RUN go get -d -v ./...

RUN go install -v ./...

RUN ls -lh 

# RUN go build -o gofrontend

# Now copy it into our base image.
FROM gcr.io/distroless/base

# copy static diretory
# COPY --from=builder /go/src/gofrontend/web /web

# copy bin diretory
COPY --from=builder /go/bin/gofrontend /

CMD ["/gofrontend"]
