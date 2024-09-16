# Load environment variables from .env file
ifneq (,$(wildcard ./.env))
    include .env
    export
endif

run: build
	@./bin/app

build:
	@go build -o bin/app .

css:
	tailwindcss -i views/css/app.css -o public/styles.css --watch   

templ:
	templ generate --watch --proxy="http://localhost$(LISTEN_ADDR)" --open-browser=true


# dev: 
# 	npx nodemon --signal SIGTERM -e "templ go" -x "templ generate && go run main.go serve" -i "**/*_templ.go"

# generate: 
# 	templ generate

# build: generate
# 	go build

# run: generate
# 	go run main.go serve