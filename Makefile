.PHONY: all build build-server local-server

BINARY_NAME=danktoolz

TEMPL_CMD=templ
ifdef CI
	TEMPL_CMD := go run github.com/a-h/templ/cmd/templ@v0.3.1020
endif

all: build-server

build: init build-server

build-server: generate
	@go build -ldflags="-w -s" -o ${BINARY_NAME} ./cmd/http/main.go

init: htmx-init amber-theme-init go-init # tailwindcss-init

generate:
	@${TEMPL_CMD} generate -path views/

go-init:
	@go mod tidy

htmx-init:
	@mkdir -p static/assets/js/htmx && \
	wget https://unpkg.com/hyperscript.org@0.9.14/dist/_hyperscript.min.js -O static/assets/js/htmx/hyperscript.min.js &&\
	wget https://unpkg.com/htmx-ext-json-enc@2.0.2/dist/json-enc.min.js -O static/assets/js/htmx/json-enc.js &&\
	wget https://unpkg.com/htmx-ext-loading-states@2.0.1/dist/loading-states.min.js -O static/assets/js/htmx/loading-states.js &&\
	wget https://unpkg.com/htmx.org@2.0.4/dist/htmx.min.js -O static/assets/js/htmx/htmx.min.js

tailwindcss-init:
	@mkdir -p static/assets/css &&\
	npm i &&\
	npx @tailwindcss/cli -i static/assets/css/style.css -o static/assets/css/tailwind.css -m

amber-theme-init:
	-mkdir ./static/assets/fonts/amber/

	@git clone https://github.com/DutchDiederik/AmberConsole &&\
	cd AmberConsole &&\
	git reset --hard 425c43046efc7e446df208c7c3399ba4f4a30c87 &&\
	git clean -df;

	@cp AmberConsole/dist/amber-console.min.css ./static/assets/css/ &&\
	cp -r AmberConsole/fonts/* ./static/assets/fonts/amber/;

	@sed -i 's|../fonts|/static/fonts/amber|g' ./static/assets/css/amber-console.min.css;

	@rm -rf AmberConsole;

tailwindcss-build:
	@npx @tailwindcss/cli -i static/assets/css/style.css -o static/assets/css/tailwind.css

tailwindcss-server:
	@npx @tailwindcss/cli -i static/assets/css/style.css -o static/assets/css/tailwind.css --watch

dev:
	@air -v > /dev/null
	@if [ $$? != 0 ]; then \
		echo "air was not found, installing it..."; \
		go install github.com/cosmtrek/air@v1.51.0; \
	fi
	export `cat .env | xargs` && air

local-server:
	@./${BINARY_NAME}

clean:
	@go clean
