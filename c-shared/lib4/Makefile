all: lib4.so

%.so: %.go
	go build -buildmode=c-shared -o $@ .

.PHONY: all clean

clean:
	$(RM) -v lib4.so lib4.h
