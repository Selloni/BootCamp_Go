NAME=compareDB
SOURS=cmd/app.go

all: build

build: clean
	go build -o $(NAME) $(SOURS)
	./compareDB --old dock/original_database.json --new dock/stolen_database.xml

clean:
	rm -rf $(NAME)
	rm -rf dock_for_parsing/stolen_database.json
	rm -rf dock_for_parsing/stolen_database.xml