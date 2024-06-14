################################################################################
#### INSTALLATION VARS
################################################################################
PREFIX=$(HOME)

################################################################################
#### BUILD VARS
################################################################################
BIN=mbscan
BINDIR=bin
RELEASEDIR=release
HEAD=$(shell git describe --dirty --long --tags 2> /dev/null  || git rev-parse --short HEAD)
COHASH=$(shell git rev-parse --short HEAD)
TIMESTAMP=$(shell TZ=UTC date '+%FT%T (%Z)')
TEST_COVER_FILE=$(BIN)-test-coverage.out
# TIMESTAMP=$(shell date '+%Y-%m-%dT%H:%M:%S %z (%Z)')

LDFLAGS="-X 'main.BinName=$(BIN)' -X 'main.BuildVersion=$(HEAD)' -X 'main.BuildTimestamp=$(TIMESTAMP)' -X 'main.CompiledBy=$(shell go version)'"

all: local

.PHONY: version
version:
	@printf "\n\n%s\n\n" $(HEAD)

################################################################################
#### HOUSE CLEANING
################################################################################

.PHONY: dep
dep:
	go mod tidy

.PHONY: check
check:
	goimports -w ./
	go vet

################################################################################
#### UN/INSTALL
################################################################################

.PHONY: _setup
_setup:
	mkdir -p $(BINDIR)
	mkdir -p $(RELEASEDIR)

.PHONY: clean
clean:
	rm -f $(BIN) $(BIN)-* $(BINDIR)/$(BIN) $(BINDIR)/$(BIN)-*

.PHONY: install
install: local
	mkdir -p $(PREFIX)/$(BINDIR)
	mv $(BINDIR)/$(BIN) $(PREFIX)/$(BINDIR)/$(BIN)
	@echo "\ninstalled $(BIN) to $(PREFIX)/$(BINDIR)\n"


.PHONY: uninstall
uninstall:
	rm -f $(PREFIX)/$(BINDIR)/$(BIN)

################################################################################
#### TESING
################################################################################

.PHONY: test
test: dep check
	go test -tags memory -covermode=count ./...

.PHONY: test-cover
test-cover: dep check
	go test -tags memory -covermode=count -coverprofile $(TEST_COVER_FILE) ./...
	go tool cover -html=$(TEST_COVER_FILE)

################################################################################
#### BUILDS
################################################################################

.PHONY: debug
debug: dep check _setup
	go build -ldflags $(DBGLDFLAGS) -o $(BINDIR)/$(BIN)

.PHONY: local
local: dep check _setup
	GOWORK=off go build -ldflags $(LDFLAGS) -o $(BINDIR)/$(BIN)

.PHONY: release
release: dep check _setup
	go build -ldflags $(LDFLAGS) -o $(BINDIR)/$(BIN)
	tar -czf $(RELEASEDIR)/$(BIN)-$(COHASH)-$(GOOS)-$(GOARCH).tgz -C $(BINDIR) $(BIN)
