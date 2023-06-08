NAME=compareDB
SOURS=cmd/app.go

all: build

build: clean
	go build -o $(NAME) $(SOURS)
	./compareDB --old dock/original_database.json --new dock/stolen_database.json

clean:
	rm -rf $(NAME)
