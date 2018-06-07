
.PHONY: install run

install:
	bash installer.sh

run: build
	bin/wtf
