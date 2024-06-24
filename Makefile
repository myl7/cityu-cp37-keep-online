SHELL = /bin/sh
.SUFFIXES:
.SUFFIXES: .go
INSTALL = install -DT
INSTALL_PROGRAM = $(INSTALL)
INSTALL_DATA = $(INSTALL) -m 644
prefix = /usr/local
exec_prefix = $(prefix)
bindir = $(exec_prefix)/bin
sysconfdir = $(prefix)/etc

all: build
.PHONY: all build install install-deb

build: out/cp37-login
out/cp37-login: $(wildcard cmd/cp37-login/*.go) $(wildcard *.go)
	go build -o $@ $(wildcard cmd/cp37-login/*.go)

install: build
	$(INSTALL_PROGRAM) out/cp37-login $(DESTDIR)$(bindir)/cp37-login
	$(INSTALL_DATA) init/cp37-login.service $(DESTDIR)$(sysconfdir)/systemd/system/cp37-login.service
	$(INSTALL_DATA) init/cp37-login.timer $(DESTDIR)$(sysconfdir)/systemd/system/cp37-login.timer
	$(INSTALL_DATA) init/cp37-login $(DESTDIR)$(sysconfdir)/default/cp37-login
install-deb: install
	$(INSTALL_DATA) build/package/control $(DESTDIR)/DEBIAN/control

# TODO: Pack README.md and LICENSE
