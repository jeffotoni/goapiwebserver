# Makefile

.EXPORT_ALL_VARIABLES:	

GO111MODULE=on
SESSION_KEY = "gofrontend"
PORT_SERVER = 5001
AUTHORIZATION = "MTIzNEFVVEhPUklaQVRJT04="
X_KEY = "MTIzNA=="

all:
	echo $$GOPATH
	##go get -d
	##go run *.go

test:
	@echo "\033[0;33m################################## go tests ##################################\033[0m"
	go test pkg/request/* 

build:
	@echo "--------------------------------------------------------------------------------------"
	@echo "\033[0;33m################################## build gofrontend ##################################\033[0m"
	go build -o gofrontend && ./gofrontend
	@echo "\033[0;32mfrontend generated\033[0m \033[0;33m[ok]\033[0m \n"

prod:
	@echo "\033[0;33m################################## build go-bindata ##################################\033[0m"
	go-bindata -o pkg/assets/binfrontend.go --nocompress --pkg=assets templates/... web/static/...
	@echo "\033[0;32mbindata generated\033[0m  \033[0;33m[ok]\033[0m "
	@echo "--------------------------------------------------------------------------------------"
	@echo "\033[0;33m################################## build gofrontend ##################################\033[0m"
	go build -o gofrontend && ./gofrontend
	@echo "\033[0;32mfrontend generated\033[0m \033[0;33m[ok]\033[0m \n"
