all: lib2.so

%.so: %.go
	go build -buildmode=c-shared -o $@ .

.PHONY: all clean

clean:
	$(RM) -v lib2.so lib2.h
