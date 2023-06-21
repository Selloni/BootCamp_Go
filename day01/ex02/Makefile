NAME=compareFS
SOURS=cmd/app.go
VALUE = --old dock/snapshot1.txt --new dock/snapshot2.txt

all: build

build: clean
	go build -o $(NAME) $(SOURS)
	./$(NAME) $(VALUE)

clean:
	rm -rf $(NAME)