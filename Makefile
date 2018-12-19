TOPDIR=$(shell pwd)
SRCDIR:=$(shell echo `pwd`/src)
CDDIR=cd $(SRCDIR)

all:
	$(CDDIR) && go run main/main.go

