---
title: "来，写个 bug  -- 读未初始化的内存"
date: 2021-04-27T22:21:31+08:00
draft: false
---

此为《深入理解计算机系统》一书上的示例程序，`matvec` 函数分配了内存，但是未初始化内存的值，因此 `y` 指向的内存的值是不确定的。在此段小程序中问题是显而易见的，在大型软件系统中可以能借助与 `Valgrind` 来检测


```C
#include <stdlib.h>
#include <stdio.h>
#define COUNT 3

int* matvec(int A[][COUNT], int* x, int n) {
	int i, j;
	int* y = (int*)malloc(n * sizeof(int));
	for (int i = 0; i < n; i++) {
		for (int j = 0; j < n; j++) {
			y[i] += A[i][j] * x[j];
		}
	}
	return y;
}

void main() {
	int a[COUNT][COUNT] = {
		{0, 1, 2},
		{2, 3, 4},
		{4, 5, 6}
	};
	
	int x[3] = {1, 2, 3};
	int* y = matvec(a, x, 3);
	for (int i = 0; i < COUNT; i++) {
		printf("%d ", y[i]);
	}
	free(y);
}
```

让我们试试使用 `Valgrind`，注意 `Valgrind` 只能检测出使用了未初始化的内存，并且*对外部环境*造成影响的情况。如果把第 26 行的 `去掉` ，则 `Valgrind` 不会检测出任何问题
```
bzhu@bzhu-Macmini:~/src$ gcc -g 9.11.2.c && valgrind --tool=memcheck --leak-check=full --track-origins=yes ./a.out
==31526== Memcheck, a memory error detector
==31526== Copyright (C) 2002-2017, and GNU GPL'd, by Julian Seward et al.
==31526== Using Valgrind-3.15.0 and LibVEX; rerun with -h for copyright info
==31526== Command: ./a.out
==31526== 
==31526== Conditional jump or move depends on uninitialised value(s)
==31526==    at 0x48D3AD8: __vfprintf_internal (vfprintf-internal.c:1687)
==31526==    by 0x48BDEBE: printf (printf.c:33)
==31526==    by 0x109334: main (9.11.2.c:26)
==31526==  Uninitialised value was created by a heap allocation
==31526==    at 0x483B7F3: malloc (in /usr/lib/x86_64-linux-gnu/valgrind/vgpreload_memcheck-amd64-linux.so)
==31526==    by 0x1091D0: matvec (9.11.2.c:7)
==31526==    by 0x1092FE: main (9.11.2.c:24)
==31526== 
==31526== Use of uninitialised value of size 8
==31526==    at 0x48B781B: _itoa_word (_itoa.c:179)
==31526==    by 0x48D36F4: __vfprintf_internal (vfprintf-internal.c:1687)
==31526==    by 0x48BDEBE: printf (printf.c:33)
==31526==    by 0x109334: main (9.11.2.c:26)
==31526==  Uninitialised value was created by a heap allocation
==31526==    at 0x483B7F3: malloc (in /usr/lib/x86_64-linux-gnu/valgrind/vgpreload_memcheck-amd64-linux.so)
==31526==    by 0x1091D0: matvec (9.11.2.c:7)
==31526==    by 0x1092FE: main (9.11.2.c:24)
==31526== 
==31526== Conditional jump or move depends on uninitialised value(s)
==31526==    at 0x48B782D: _itoa_word (_itoa.c:179)
==31526==    by 0x48D36F4: __vfprintf_internal (vfprintf-internal.c:1687)
==31526==    by 0x48BDEBE: printf (printf.c:33)
==31526==    by 0x109334: main (9.11.2.c:26)
==31526==  Uninitialised value was created by a heap allocation
==31526==    at 0x483B7F3: malloc (in /usr/lib/x86_64-linux-gnu/valgrind/vgpreload_memcheck-amd64-linux.so)
==31526==    by 0x1091D0: matvec (9.11.2.c:7)
==31526==    by 0x1092FE: main (9.11.2.c:24)
==31526== 
==31526== Conditional jump or move depends on uninitialised value(s)
==31526==    at 0x48D43A8: __vfprintf_internal (vfprintf-internal.c:1687)
==31526==    by 0x48BDEBE: printf (printf.c:33)
==31526==    by 0x109334: main (9.11.2.c:26)
==31526==  Uninitialised value was created by a heap allocation
==31526==    at 0x483B7F3: malloc (in /usr/lib/x86_64-linux-gnu/valgrind/vgpreload_memcheck-amd64-linux.so)
==31526==    by 0x1091D0: matvec (9.11.2.c:7)
==31526==    by 0x1092FE: main (9.11.2.c:24)
==31526== 
==31526== Conditional jump or move depends on uninitialised value(s)
==31526==    at 0x48D386E: __vfprintf_internal (vfprintf-internal.c:1687)
==31526==    by 0x48BDEBE: printf (printf.c:33)
==31526==    by 0x109334: main (9.11.2.c:26)
==31526==  Uninitialised value was created by a heap allocation
==31526==    at 0x483B7F3: malloc (in /usr/lib/x86_64-linux-gnu/valgrind/vgpreload_memcheck-amd64-linux.so)
==31526==    by 0x1091D0: matvec (9.11.2.c:7)
==31526==    by 0x1092FE: main (9.11.2.c:24)
==31526== 
8 20 32 ==31526== 
==31526== HEAP SUMMARY:
==31526==     in use at exit: 0 bytes in 0 blocks
==31526==   total heap usage: 2 allocs, 2 frees, 1,036 bytes allocated
==31526== 
==31526== All heap blocks were freed -- no leaks are possible
==31526== 
==31526== For lists of detected and suppressed errors, rerun with: -s
==31526== ERROR SUMMARY: 19 errors from 5 contexts (suppressed: 0 from 0)
```

另外，最简单的情况，也是读取了未初始化的内存。未初始化的内存可以来自与堆上，也可以来自于普通变量（栈上）。

```C
int main()
{
  int x;
  printf ("x = %d\n", x);
}
```

更详细描述，请参考`Valgrind` [手册](https://www.valgrind.org/docs/manual/mc-manual.html#mc-manual.uninitvals)

