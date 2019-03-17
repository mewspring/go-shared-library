all: lib1.so

%.so: %.go
	CGO_ENABLED=1 GOARCH=386 go build -buildmode=c-shared -o $@ .

.PHONY: all clean

clean:
	$(RM) -v lib1.so lib1.h
