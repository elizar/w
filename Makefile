GO ?= go

# start development server
dev: tmpl/tmpl_bindata.go
	@up start -c 'gin -p $$PORT'
.PHONY: dev

# install dev dependencies
install.deps:
	@echo
	@echo "  ==> Installing dependencies"
	@$(GO) get -u github.com/codegangsta/gin
	@$(GO) get -u github.com/zendesk/go-bindata/...
	@curl -sf https://up.apex.sh/install | sh
.PHONY: install.deps

# compile template binaries
tmpl/tmpl_bindata.go:
	@echo
	@echo "  ==> Compiling template binaries"
	@cd tmpl && go-bindata -debug -ignore '\.go' -pkg tmpl -o tmpl_bindata.go .
