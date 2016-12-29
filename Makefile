SHELL := /bin/bash

.PHONY: test codecov

all: test

filename="coverage.txt"

test:
	@{ \
	set -e; \
	mode="atomic"; \
	echo "mode: $$mode" > $(filename); \
	for dir in $$(go list ./...); do \
		out=$$(go test -race -coverprofile=profile.out -covermode=$$mode $$dir); \
		echo $$out; \
	  	if [[ $$out == *FAIL* ]]; then \
			rm -f profile.out; \
    			rm -f $(filename); \
    			exit 1; \
		fi; \
		if [ -f profile.out ]; then \
			cat profile.out | grep -v "mode: $$mode" >> $(filename); \
			rm -f profile.out; \
  		fi; \
	done \
	}

codecov:
	@{ \
	set -e; \
	bash <(curl -s https://codecov.io/bash) || echo "Codecov fails"; \
	if [ -f $(filename) ]; then \
		rm -f $(filename); \
	fi; \
	}

