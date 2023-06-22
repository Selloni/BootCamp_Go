NAME=myFind
SOURS=ex00/cmd/app.go
VALUE = /Users/grandpat/golang_learn/day01
FLAG = -d
all: build

build: clean
	go build -o $(NAME) $(SOURS)
	./$(NAME) $(FLAG) $(VALUE)

clean:
	rm -rf $(NAME)
