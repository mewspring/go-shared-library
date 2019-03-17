all: lib4.so

%.so: %.go
	CGO_ENABLED=1 GOARCH=386 go build -buildmode=c-shared -o $@ .

.PHONY: all clean

clean:
	$(RM) -v lib4.so lib4.h
