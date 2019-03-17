all: lib5.so

%.so: %.go
	go build -buildmode=c-shared -o $@ .

.PHONY: all clean

clean:
	$(RM) -v lib5.so lib5.h
