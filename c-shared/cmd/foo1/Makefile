all: foo1

foo1: foo1.c lib1.so lib2.so lib3.so lib4.so lib5.so
	gcc -I../../lib1 -I../../lib2 -I../../lib3 -I../../lib4 -I../../lib5 -o $@ $^

lib1.so:
	make -C ../../lib1
	cp ../../lib1/lib1.so $@

lib2.so:
	make -C ../../lib2
	cp ../../lib2/lib2.so $@

lib3.so:
	make -C ../../lib3
	cp ../../lib3/lib3.so $@

lib4.so:
	make -C ../../lib4
	cp ../../lib4/lib4.so $@

lib5.so:
	make -C ../../lib5
	cp ../../lib5/lib5.so $@

.PHONY: all clean

clean:
	$(RM) -v foo1
	$(RM) -v *.so
