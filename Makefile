TOPDIR=$(shell pwd)

init:
	GO111MODULE=on \
	go mod init

all:
	go build -o $(TOPDIR)/bin/goconcurrency main.go

barycenter_datasets:
	mkdir -p res && \
	go run main.go generateBarycenter 1000 > res/1000k-barrycenter.txt && \
	go run main.go generateBarycenter 10000 > res/10000k-barrycenter.txt