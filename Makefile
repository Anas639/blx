GOPATH:=$(shell go env GOPATH)
APP_NAME = blx
INSTALL_DIR = $(GOPATH)/bin

build:
	go build --tags "fts5" -o $(APP_NAME) .

install: build
	mv $(APP_NAME) $(INSTALL_DIR)/$(APP_NAME)

uninstall:
	rm $(INSTALL_DIR)/$(APP_NAME)

clean:
	rm -f $(APP_NAME)
