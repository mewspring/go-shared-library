all: lib8.so

%.so: %.go
	CGO_ENABLED=1 GOARCH=386 go build -buildmode=c-shared -o $@ .

.PHONY: all clean

clean:
	$(RM) -v lib8.so lib8.h
