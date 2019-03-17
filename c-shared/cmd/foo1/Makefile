all: foo1

foo1: foo1.c lib1.so lib2.so lib3.so lib4.so lib5.so lib6.so lib7.so lib8.so lib9.so lib10.so lib11.so lib12.so lib13.so lib14.so lib15.so lib16.so lib17.so lib18.so lib19.so
	gcc -I../../lib1 -I../../lib2 -I../../lib3 -I../../lib4 -I../../lib5 -I../../lib6 -I../../lib7 -I../../lib8 -I../../lib9 -I../../lib10 -I../../lib11 -I../../lib12 -I../../lib13 -I../../lib14 -I../../lib15 -I../../lib16 -I../../lib17 -I../../lib18 -I../../lib19 -o $@ $^

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

lib6.so:
	make -C ../../lib6
	cp ../../lib6/lib6.so $@

lib7.so:
	make -C ../../lib7
	cp ../../lib7/lib7.so $@

lib8.so:
	make -C ../../lib8
	cp ../../lib8/lib8.so $@

lib9.so:
	make -C ../../lib9
	cp ../../lib9/lib9.so $@

lib10.so:
	make -C ../../lib10
	cp ../../lib10/lib10.so $@

lib11.so:
	make -C ../../lib11
	cp ../../lib11/lib11.so $@

lib12.so:
	make -C ../../lib12
	cp ../../lib12/lib12.so $@

lib13.so:
	make -C ../../lib13
	cp ../../lib13/lib13.so $@

lib14.so:
	make -C ../../lib14
	cp ../../lib14/lib14.so $@

lib15.so:
	make -C ../../lib15
	cp ../../lib15/lib15.so $@

lib16.so:
	make -C ../../lib16
	cp ../../lib16/lib16.so $@

lib17.so:
	make -C ../../lib17
	cp ../../lib17/lib17.so $@

lib18.so:
	make -C ../../lib18
	cp ../../lib18/lib18.so $@

lib19.so:
	make -C ../../lib19
	cp ../../lib19/lib19.so $@


.PHONY: all clean

clean:
	$(RM) -v foo1
	$(RM) -v *.so
