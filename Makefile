GO ?= go

# start development server
dev: tmpl/tmpl_bindata.go
	@gin
.PHONY: dev

# up start
dev.up: tmpl/tmpl_bindata.go
	@export UP=1
	@up start

# install dev dependencies
install.deps:
	@echo "  --> Installing dependencies"
	@$(GO) get -u github.com/codegangsta/gin
	@$(GO) get -u github.com/zendesk/go-bindata/...
.PHONY: install.deps

# compile template binaries
tmpl/tmpl_bindata.go:
	@echo "  --> Compiling template binaries"
	@cd tmpl && go-bindata -debug -ignore '\.go' -pkg tmpl -o tmpl_bindata.go .
