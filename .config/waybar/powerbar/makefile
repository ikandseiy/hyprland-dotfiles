app=powerbar
PKGS:=$(shell go list ./... | grep -v vendor)
OUTDIR=.bin
BINARY=$(OUTDIR)/$(app)
GOBIN=$(GOPATH)/bin
LINTBIN=$(GOBIN)/golangci-lint
VERSION:=$(shell git describe --tags --all HEAD)
clean_list=$(OUTDIR)

.PHONY: build clean default dev lint install-local publish test pkgcheck

default: lint test build dev

$(OUTDIR):
	@mkdir -p $@

$(BINARY): . $(OUTDIR)
	go build -o $@ $<

$(LINTBIN):
	@GO111MODULE=off go get github.com/golangci/golangci-lint/cmd/golangci-lint

aur/PKGBUILD: 
	@mkdir -p ./aur/
	@cp PKGBUILD.template aur/PKGBUILD

pkg: aur/PKGBUILD
	cd aur && \
		makepkg -f && \
		makepkg --printsrcinfo > .SRCINFO

pkgcheck: $(OUTDIR)/PKGBUILD
	cd $(OUTDIR) && makepkg -f

dev: clean $(BINARY)
	LB_DEBUG=true $(BINARY) $(FPATH)

build: clean $(BINARY)

lint: $(LINTBIN)
	go mod tidy -v
	$(LINTBIN) run -p bugs -p format -p performance -p unused

test: lint
	go test -v -cover $(PKGS)

clean: $(clean_list)
	rm -rf $<  aur/src aur/pkg aur/$(app) aur/*.xz aur/PKGBUILD aur/.SRCINFO

install-local: build
	cp $(BINARY) $${HOME}/.config/waybar/
