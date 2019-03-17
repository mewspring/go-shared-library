all: lib3.so

%.so: %.go
	go build -buildmode=c-shared -o $@ .

.PHONY: all clean

clean:
	$(RM) -v lib3.so lib3.h
