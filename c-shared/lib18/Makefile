all: lib18.so

%.so: %.go
	go build -buildmode=c-shared -o $@ .

.PHONY: all clean

clean:
	$(RM) -v lib18.so lib18.h
