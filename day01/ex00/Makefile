NAME=readDB
SOURS=cmd/app.go

all: build

build: clean
	go build -o $(NAME) $(SOURS)

clean:
	rm -rf $(NAME)
	rm -rf dock_for_parsing/stolen_database.json
	rm -rf dock_for_parsing/stolen_database.xml