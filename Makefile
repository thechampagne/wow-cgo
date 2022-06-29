CC := gcc
CFLAGS := -Wall -Wextra -std=c99 -pthread
CGO := go build
STATIC := -buildmode=c-archive
SHARED := -buildmode=c-shared
LIBS := build/wow.a build/wow.so

.PHONY: all
all: $(LIBS)

build/wow.a: wow.go
	$(CGO) $(STATIC) -o build/wow.a $<

build/wow.so: wow.go
	$(CGO) $(SHARED) -o build/wow.so $<

.PHONY: test
test: all
	$(CC) $(CFLAGS) test/main.c -o test/main -Ibuild -Lbuild -l:wow.a
	./test/main
	rm test/main

.PHONY: clean
clean:
	find build -type f \( -name '*.h' -o -name '*.so' -o -name '*.a' \) -delete
