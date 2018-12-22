TOPDIR=$(shell pwd)

init:
	GO111MODULE=on \
	go mod init

all:
	go build -o $(TOPDIR)/bin/goconcurrency main.go