GREEN="\033[00;32m"
RESTORE="\033[0m"

# make the output of the message appear green
define style_calls
	$(eval $@_msg = $(1))
	echo ${GREEN}${$@_msg}
	echo ${RESTORE}
endef

.PHONY: format spell spell-write all help

default_target: help

format:
	@$(call style_calls,"Running gofmt")
	@gofmt -e -s -d -w .

spell:
	@$(call style_calls,"Running codespell check")
	@codespell --quiet-level=2 --check-hidden --skip=./.git .

spell-write:
	@$(call style_calls,"Running codespell write")
	@codespell --quiet-level=2 --check-hidden --skip=./.git --write-changes .

all: tests lint spell

help:
	@echo "make format      - Run formatting"
	@echo "make spell       - Run spell check"
	@echo "make spell-write - Run spell check and write changes"
	@echo "make all         - Run all checks"
