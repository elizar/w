# Makefile
.PHONY: dev
dev:
	cd tmpl && go-bindata -debug -ignore '\.go' -pkg tmpl . && cd .. &&  up start
