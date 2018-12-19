TOPDIR=$(shell pwd)
SRCDIR := $(shell echo `pwd`/src)
CDDIR =cd $(SRCDIR)

intro:
	$(CDDIR) && go build -o $(TOPDIR)/bin/intro intro/main.go

goroutines:
	$(CDDIR) && go run goroutines/main_goroutines.go

