#include <unistd.h>

#include "lib1.h"
#include "lib2.h"
#include "lib3.h"
#include "lib4.h"
#include "lib5.h"
#include "lib6.h"
#include "lib7.h"
#include "lib8.h"
#include "lib9.h"

int main() {
	lib1_f();
	lib2_f();
	lib3_f();
	lib4_f();
	lib5_f();
	lib6_f();
	lib7_f();
	lib8_f();
	lib9_f();
	sleep(50000);
	return 42;
}
