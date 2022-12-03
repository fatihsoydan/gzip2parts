# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
BUILD_NO=$(shell cat ./.buildno)
LDFLAGS=-ldflags "-X main.BuildNumber=$(BUILD_NO) "
OUTPUT=gzip2parts-$(shell uname -s)
TMP=/tmp


build:
			$(GOBUILD) $(LDFLAGS) -o bin/$(OUTPUT) -v *.go && let "BUILD_NO=$(BUILD_NO)+1" && echo $$BUILD_NO > ./.buildno

clean:
			$(GOCLEAN)

run:	
			$(GOBUILD) $(LDFLAGS) -o $(TMP)/$(OUTPUT) -v *.go 
	 		/tmp/$(OUTPUT) -c -i=testContent -o=/tmp/out
			/tmp/$(OUTPUT) -x -i=/tmp/out -o=/tmp/out2