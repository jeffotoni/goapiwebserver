# Makefile

.EXPORT_ALL_VARIABLES:	

GO111MODULE=on
SESSION_KEY = "gofrontend"
PORT_SERVER = 5000
AUTHORIZATION = "MTIzNEFVVEhPUklaQVRJT04="
X_KEY = "MTIzNA=="

all:
	echo $$GOPATH
	##go get -d
	##go run *.go

test:
	@echo "\n"
	@echo "\033[0;33m################################## go tests ##################################\033[0m"
	go test pkg/request/* 

build:
	@echo "\n"
	@echo "--------------------------------------------------------------------------------------"
	@echo "\033[0;33m################################## build gofrontend ##################################\033[0m"
	go build -o gofrontend && ./gofrontend
	@echo "\033[0;32mfrontend generated\033[0m \033[0;33m[ok]\033[0m \n"

prod:
	@echo "\033[0;33m################################## build go-bindata ##################################\033[0m"
	go-bindata -o pkg/assets/binfrontend.go --pkg=assets web/templates/... web/static/...
	@echo "\033[0;32mbindata generated\033[0m  \033[0;33m[ok]\033[0m "
	@echo "\n"
	@echo "--------------------------------------------------------------------------------------"
	@echo "\033[0;33m################################## build gofrontend ##################################\033[0m"
	go build -o gofrontend && ./gofrontend
	@echo "\033[0;32mfrontend generated\033[0m \033[0;33m[ok]\033[0m \n"
