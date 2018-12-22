TOPDIR=$(shell pwd)

init:
	GO111MODULE=on \
	go mod init

all:
	go build -o $(TOPDIR)/bin/goconcurrency main.go

barrycenter_datasets:
	mkdir -p res && \
	go run main.go generateBarrycenter 1000 > res/1000k-barrycenter.txt && \
	go run main.go generateBarrycenter 10000 > res/10000k-barrycenter.txt