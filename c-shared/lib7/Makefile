all: lib7.so

%.so: %.go
	CGO_ENABLED=1 GOARCH=386 go build -buildmode=c-shared -o $@ .

.PHONY: all clean

clean:
	$(RM) -v lib7.so lib7.h
