SERVICES := account shortener api

test-account:
	make -C account test

test-shortener:
	cd shortener
	make test
	cd ../

test-api:
	cd api
	make test
	cd ../

test-all:
	for service in ${SERVICES}; do \
	  make -C $$service test; \
	done

.PHONY: build proto

