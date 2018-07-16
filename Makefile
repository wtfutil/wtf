.PHONY: contrib_check dependencies install run size

install:
	bash installer.sh

run: build
	bin/wtf

size:
	loc --exclude vendor/ _sample_configs/ _site/ docs/ Makefile *.md *.toml
