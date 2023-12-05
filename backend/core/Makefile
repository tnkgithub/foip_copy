BINDIR:=bin
ROOT_PACKAGE:=$(shell go list .)
BINARY_NAME:=server
BINARIES:=bin/$(BINARY_NAME)

GOCMD:=@go
GORUN:=$(GOCMD) run
GOBUILD:=$(GOCMD) build
GOINSTALL:=$(GOCMD) install
GOGET:=$(GOCMD) get

SWAGCMD:=@swag
SWAGFMT:=$(SWAGCMD) fmt
SWAGINIT:=$(SWAGCMD) init

RMCMD:=@rm

DOCKERCMD:=@docker
DOCKERBUILD:=$(DOCKERCMD) build
DOCKERRUN:=$(DOCKERCMD) run
IMAGETAG:=$(shell go run main.go -v | sed 's/ version /\//')

export GO111MODULE := on

#-----    Go    -----#
all: build

.PHONY: build
build: $(BINARIES)

$(BINARIES): swag
	$(GOGET);
	$(GOBUILD) -o $@ $(ROOT_PACKAGE);

.PHONY: swag
swag: $(SWAG_INSTALL)
	$(SWAGFMT);
	$(SWAGINIT) \
		--output ./docs \
		--outputTypes go,yaml \
		--generatedTime \
		--quiet;

$(SWAG_INSTALL):
	$(GOINSTALL) github.com/swaggo/swag/cmd/swag@latest;

.PHONY: clean
clean:
	$(RMCMD) -r $(BINDIR);

#-----  Docker  -----#
# https://christina04.hatenablog.com/entry/2017/06/01/125203
.PHONY: image
image:
	$(DOCKERBUILD) \
		-f Dockerfile \
		--tag backend/$(IMAGETAG) .

image.local.run:
	$(DOCKERRUN) \
		-p 8080:8080 \
		--name local -td backend/$(IMAGETAG)
