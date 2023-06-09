NAME=compareDB
SOURS=cmd/app.go

all: build

build: clean
	go build -o $(NAME) $(SOURS)
	./compareDB --old dock/original_database.json --new dock/original_database.xml

clean:
	rm -rf $(NAME)
