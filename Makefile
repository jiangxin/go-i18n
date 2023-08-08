PROGRAM = go-i18n

# Include wrappers for gettext utils
include makefile.i18n

FIND ?= find

GO_SOURCES := $(sort $(patsubst ./%,%,$(shell git ls-files '*.go' ':!vendor/' ':!test/' ':!Documentation/' 2>/dev/null || \
	$(FIND) . \
	-name .git -prune -o \
	-name vendor -prune -o \
	-name test -prune -o \
	-name Documentation -prune -o \
	-name '*.go' -print)))

VENDOR_EXISTS=$(shell test -d vendor && echo 1 || echo 0)
ifeq ($(VENDOR_EXISTS), 1)
    GOBUILD := GO111MODULE=on CGO_ENABLED=0 go build -mod=vendor
    GOTEST := GO111MODULE=on go test -mod=vendor
else
    GOBUILD := GO111MODULE=on CGO_ENABLED=0 go build
    GOTEST := GO111MODULE=on go test
endif

all:: $(PROGRAM)

$(PROGRAM): $(GO_SOURCES)
	$(GOBUILD) -o $@

.PHONY: test
test:
	make -C test

.PHONY: clean
clean:
	rm -f $(PROGRAM)
