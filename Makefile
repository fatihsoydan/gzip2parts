# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean

ifeq '$(findstring ;,$(PATH))' ';'
TMP=C:\\Temp
BUILD_NO=$(shell type .buildno)
OUTPUT=gzip2parts.exe
BCMD=$(GOBUILD) $(LDFLAGS) -o bin/$(OUTPUT)
else 
TMP=/tmp
BUILD_NO=$(shell cat ./.buildno)
LDFLAGS=-ldflags "-X main.BuildNumber=$(BUILD_NO) "
OUTPUT=gzip2parts-$(shell uname -s)
BCMD=$(GOBUILD) $(LDFLAGS) -o bin/$(OUTPUT) && let "BUILD_NO=$(BUILD_NO)+1" && echo $$BUILD_NO > ./.buildno		
endif



LDFLAGS=-ldflags "-X main.BuildNumber=$(BUILD_NO) "


build:
		$(BCMD)
		
buildwindows:
		$(GOBUILD) $(LDFLAGS) -o bin/$(OUTPUT) 

clean:
			$(GOCLEAN)

run:	
			$(GOBUILD) $(LDFLAGS) -o $(TMP)/$(OUTPUT) -v *.go 
			$(TMP)/$(OUTPUT) -c -i=testContent -o=$(TMP)/out
			$(TMP)/$(OUTPUT) -x -i=$(TMP)/out -o=$(TMP)/out2