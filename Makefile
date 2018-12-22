TOPDIR=$(shell pwd)

init:
	GO111MODULE=on \
	go mod init

all:
	go build -o $(TOPDIR)/bin/goconcurrency main.go

barycenter_datasets:
	mkdir -p res && \
	go run main.go generateBarycenter 1000 > res/1k-barycenter.txt && \
	go run main.go generateBarycenter 10000 > res/10k-barycenter.txt && \
	go run main.go generateBarycenter 1000000 > res/1mil-barycenter.txt