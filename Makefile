TOPDIR=$(shell pwd)
SRCDIR:=$(shell echo `pwd`/src)
CDDIR=cd $(SRCDIR)

main:
	$(CDDIR) && go build -o $(TOPDIR)/bin/main main/main.go

all: main
	./bin/main

