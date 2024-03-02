---
title: "交叉编译的一些记录"
date: 2024-03-02T11:06:35+08:00
draft: false
tags: [C, C++]
---

任务的目标是在 Host Ubuntu Focal(20.04) X86 上交叉编译 Target 为 Ubuntu arm64 的 binary

首先是如何安装交叉编译器的问题，有两个选择，一个是自己编译出交叉编译器（一般使用 cross-ng），另一种使用 APT 仓库中的，`apt install gcc-aarch64-linux-gnu g++-aarch64-linux-gnu` 就可以。如要要安装交叉编译所需要的库文件，需要添加 Multi-Arch 支持，比如 `apt -y install libssl-dev:arm64`

比较奇怪的一点就是，安装后的交叉编译器的 `sysroot` 是根目录，而没有在一个单独的子目录（不像 Cross-ng 编译出来的），感觉这样就很容易跟别的库文件混淆。
sysroot 也可以把目标系统的跟目录挂载过来, 作为 sysroot，这样所有的头文件，库文件，都可以从 sysroot 中查找

```
$ aarch64-linux-gnu-gcc --print-sysroot
/
```

比如安装的 libssl-dev:arm64 是安装在如下的目录的。通用的 .h 文件直接就安装在了 `/usr/include` 下，只有和架构有关的头文件才安装在 `/usr/include/aarch64-linux-gun` 下。

```
dpkg -L libssl-dev:arm64
/.
/usr
/usr/include
/usr/include/aarch64-linux-gnu
/usr/include/aarch64-linux-gnu/openssl
/usr/include/aarch64-linux-gnu/openssl/opensslconf.h
/usr/include/openssl
/usr/include/openssl/aes.h
/usr/include/openssl/asn1.h
/usr/include/openssl/asn1_mac.h
/usr/include/openssl/asn1err.h
/usr/include/openssl/asn1t.h
/usr/include/openssl/async.h
/usr/include/openssl/asyncerr.h
/usr/include/openssl/bio.h
/usr/include/openssl/bioerr.h
/usr/include/openssl/blowfish.h
/usr/include/openssl/bn.h
/usr/include/openssl/bnerr.h
/usr/include/openssl/buffer.h
/usr/include/openssl/buffererr.h
/usr/include/openssl/camellia.h
/usr/include/openssl/cast.h
/usr/include/openssl/cmac.h
/usr/include/openssl/cms.h
/usr/include/openssl/cmserr.h
/usr/include/openssl/comp.h
/usr/include/openssl/comperr.h
/usr/include/openssl/conf.h
/usr/include/openssl/conf_api.h
/usr/include/openssl/conferr.h
/usr/include/openssl/crypto.h
/usr/include/openssl/cryptoerr.h
/usr/include/openssl/ct.h
/usr/include/openssl/cterr.h
/usr/include/openssl/des.h
/usr/include/openssl/dh.h
/usr/include/openssl/dherr.h
/usr/include/openssl/dsa.h
/usr/include/openssl/dsaerr.h
/usr/include/openssl/dtls1.h
/usr/include/openssl/e_os2.h
/usr/include/openssl/ebcdic.h
/usr/include/openssl/ec.h
/usr/include/openssl/ecdh.h
/usr/include/openssl/ecdsa.h
/usr/include/openssl/ecerr.h
/usr/include/openssl/engine.h
/usr/include/openssl/engineerr.h
/usr/include/openssl/err.h
/usr/include/openssl/evp.h
/usr/include/openssl/evperr.h
/usr/include/openssl/hmac.h
/usr/include/openssl/idea.h
/usr/include/openssl/kdf.h
/usr/include/openssl/kdferr.h
/usr/include/openssl/lhash.h
/usr/include/openssl/md2.h
/usr/include/openssl/md4.h
/usr/include/openssl/md5.h
/usr/include/openssl/mdc2.h
/usr/include/openssl/modes.h
/usr/include/openssl/obj_mac.h
/usr/include/openssl/objects.h
/usr/include/openssl/objectserr.h
/usr/include/openssl/ocsp.h
/usr/include/openssl/ocsperr.h
/usr/include/openssl/opensslv.h
/usr/include/openssl/ossl_typ.h
/usr/include/openssl/pem.h
/usr/include/openssl/pem2.h
/usr/include/openssl/pemerr.h
/usr/include/openssl/pkcs12.h
/usr/include/openssl/pkcs12err.h
/usr/include/openssl/pkcs7.h
/usr/include/openssl/pkcs7err.h
/usr/include/openssl/rand.h
/usr/include/openssl/rand_drbg.h
/usr/include/openssl/randerr.h
/usr/include/openssl/rc2.h
/usr/include/openssl/rc4.h
/usr/include/openssl/rc5.h
/usr/include/openssl/ripemd.h
/usr/include/openssl/rsa.h
/usr/include/openssl/rsaerr.h
/usr/include/openssl/safestack.h
/usr/include/openssl/seed.h
/usr/include/openssl/sha.h
/usr/include/openssl/srp.h
/usr/include/openssl/srtp.h
/usr/include/openssl/ssl.h
/usr/include/openssl/ssl2.h
/usr/include/openssl/ssl3.h
/usr/include/openssl/sslerr.h
/usr/include/openssl/stack.h
/usr/include/openssl/store.h
/usr/include/openssl/storeerr.h
/usr/include/openssl/symhacks.h
/usr/include/openssl/tls1.h
/usr/include/openssl/ts.h
/usr/include/openssl/tserr.h
/usr/include/openssl/txt_db.h
/usr/include/openssl/ui.h
/usr/include/openssl/uierr.h
/usr/include/openssl/whrlpool.h
/usr/include/openssl/x509.h
/usr/include/openssl/x509_vfy.h
/usr/include/openssl/x509err.h
/usr/include/openssl/x509v3.h
/usr/include/openssl/x509v3err.h
/usr/lib
/usr/lib/aarch64-linux-gnu
/usr/lib/aarch64-linux-gnu/libcrypto.a
/usr/lib/aarch64-linux-gnu/libssl.a
/usr/lib/aarch64-linux-gnu/pkgconfig
/usr/lib/aarch64-linux-gnu/pkgconfig/libcrypto.pc
/usr/lib/aarch64-linux-gnu/pkgconfig/libssl.pc
/usr/lib/aarch64-linux-gnu/pkgconfig/openssl.pc
/usr/share
/usr/share/doc
/usr/share/doc/libssl-dev
/usr/lib/aarch64-linux-gnu/libcrypto.so
/usr/lib/aarch64-linux-gnu/libssl.so
/usr/share/doc/libssl-dev/changelog.Debian.gz
/usr/share/doc/libssl-dev/changelog.gz
/usr/share/doc/libssl-dev/copyright
```

而目录 `/usr/aarch64-linux-gnu/lib/` 下面又有一堆的库文件，所以很难设置一个比较独立的 sysroot. 一个 workaround 就是，把这些文件都拷贝到一个目录作为 sysroot
, 比如 `cp -fr  /usr/aarch64-linux-gnu/lib/* /opt/arm64/lib`，另外所有的依赖库都自己编译，然后安装到 `/opt/arm64`这样可以得到一个比较独立的 sysroot.

```
vscode@fc893c6d29d7:/tmp/libssh2-1.10.0/build$ ls -l /usr/aarch64-linux-gnu/lib/
total 48244
-rw-r--r-- 1 root root      496 Apr 21  2022 Mcrt1.o
-rw-r--r-- 1 root root     1704 Apr 21  2022 Scrt1.o
-rw-r--r-- 1 root root     1872 Apr 21  2022 crt1.o
-rw-r--r-- 1 root root     1440 Apr 21  2022 crti.o
-rw-r--r-- 1 root root     1016 Apr 21  2022 crtn.o
-rw-r--r-- 1 root root     2664 Apr 21  2022 gcrt1.o
-rwxr-xr-x 1 root root   145320 Apr 21  2022 ld-2.31.so
lrwxrwxrwx 1 root root       10 Apr 21  2022 ld-linux-aarch64.so.1 -> ld-2.31.so
-rw-r--r-- 1 root root     6240 Apr 21  2022 libBrokenLocale-2.31.so
-rw-r--r-- 1 root root     1982 Apr 21  2022 libBrokenLocale.a
lrwxrwxrwx 1 root root       20 Apr 21  2022 libBrokenLocale.so -> libBrokenLocale.so.1
lrwxrwxrwx 1 root root       23 Apr 21  2022 libBrokenLocale.so.1 -> libBrokenLocale-2.31.so
-rw-r--r-- 1 root root    10400 Apr 21  2022 libSegFault.so
-rw-r--r-- 1 root root    14840 Apr 21  2022 libanl-2.31.so
-rw-r--r-- 1 root root    21500 Apr 21  2022 libanl.a
lrwxrwxrwx 1 root root       11 Apr 21  2022 libanl.so -> libanl.so.1
lrwxrwxrwx 1 root root       14 Apr 21  2022 libanl.so.1 -> libanl-2.31.so
lrwxrwxrwx 1 root root       16 Oct 25  2022 libasan.so.5 -> libasan.so.5.0.0
-rw-r--r-- 1 root root 15230984 Oct 25  2022 libasan.so.5.0.0
lrwxrwxrwx 1 root root       18 Jul 10  2023 libatomic.so.1 -> libatomic.so.1.2.0
-rw-r--r-- 1 root root    38736 Jul 10  2023 libatomic.so.1.2.0
-rwxr-xr-x 1 root root  1446928 Apr 21  2022 libc-2.31.so
-rw-r--r-- 1 root root  4616522 Apr 21  2022 libc.a
-rw-r--r-- 1 root root      317 Apr 21  2022 libc.so
lrwxrwxrwx 1 root root       12 Apr 21  2022 libc.so.6 -> libc-2.31.so
-rw-r--r-- 1 root root    22972 Apr 21  2022 libc_nonshared.a
-rw-r--r-- 1 root root    14560 Apr 21  2022 libdl-2.31.so
-rw-r--r-- 1 root root    13930 Apr 21  2022 libdl.a
lrwxrwxrwx 1 root root       10 Apr 21  2022 libdl.so -> libdl.so.2
lrwxrwxrwx 1 root root       13 Apr 21  2022 libdl.so.2 -> libdl-2.31.so
-rw-r--r-- 1 root root     1242 Apr 21  2022 libg.a
-rw-r--r-- 1 root root    80072 Jul 10  2023 libgcc_s.so.1
lrwxrwxrwx 1 root root       16 Jul 10  2023 libgomp.so.1 -> libgomp.so.1.0.0
-rw-r--r-- 1 root root   253328 Jul 10  2023 libgomp.so.1.0.0
lrwxrwxrwx 1 root root       15 Jul 10  2023 libitm.so.1 -> libitm.so.1.0.0
-rw-r--r-- 1 root root    96560 Jul 10  2023 libitm.so.1.0.0
lrwxrwxrwx 1 root root       16 Jul 10  2023 liblsan.so.0 -> liblsan.so.0.0.0
-rw-r--r-- 1 root root  3268752 Jul 10  2023 liblsan.so.0.0.0
-rw-r--r-- 1 root root   628896 Apr 21  2022 libm-2.31.so
-rw-r--r-- 1 root root  1597840 Apr 21  2022 libm.a
lrwxrwxrwx 1 root root        9 Apr 21  2022 libm.so -> libm.so.6
lrwxrwxrwx 1 root root       12 Apr 21  2022 libm.so.6 -> libm-2.31.so
-rw-r--r-- 1 root root     1552 Apr 21  2022 libmcheck.a
-rw-r--r-- 1 root root    18696 Apr 21  2022 libmemusage.so
-rw-r--r-- 1 root root    93088 Apr 21  2022 libnsl-2.31.so
-rw-r--r-- 1 root root   198012 Apr 21  2022 libnsl.a
lrwxrwxrwx 1 root root       11 Apr 21  2022 libnsl.so -> libnsl.so.1
lrwxrwxrwx 1 root root       14 Apr 21  2022 libnsl.so.1 -> libnsl-2.31.so
-rw-r--r-- 1 root root    35536 Apr 21  2022 libnss_compat-2.31.so
lrwxrwxrwx 1 root root       18 Apr 21  2022 libnss_compat.so -> libnss_compat.so.2
lrwxrwxrwx 1 root root       21 Apr 21  2022 libnss_compat.so.2 -> libnss_compat-2.31.so
-rw-r--r-- 1 root root    22808 Apr 21  2022 libnss_dns-2.31.so
lrwxrwxrwx 1 root root       15 Apr 21  2022 libnss_dns.so -> libnss_dns.so.2
lrwxrwxrwx 1 root root       18 Apr 21  2022 libnss_dns.so.2 -> libnss_dns-2.31.so
-rw-r--r-- 1 root root    51640 Apr 21  2022 libnss_files-2.31.so
lrwxrwxrwx 1 root root       17 Apr 21  2022 libnss_files.so -> libnss_files.so.2
lrwxrwxrwx 1 root root       20 Apr 21  2022 libnss_files.so.2 -> libnss_files-2.31.so
-rw-r--r-- 1 root root    18736 Apr 21  2022 libnss_hesiod-2.31.so
lrwxrwxrwx 1 root root       18 Apr 21  2022 libnss_hesiod.so -> libnss_hesiod.so.2
lrwxrwxrwx 1 root root       21 Apr 21  2022 libnss_hesiod.so.2 -> libnss_hesiod-2.31.so
-rw-r--r-- 1 root root    47544 Apr 21  2022 libnss_nis-2.31.so
lrwxrwxrwx 1 root root       15 Apr 21  2022 libnss_nis.so -> libnss_nis.so.2
lrwxrwxrwx 1 root root       18 Apr 21  2022 libnss_nis.so.2 -> libnss_nis-2.31.so
-rw-r--r-- 1 root root    59808 Apr 21  2022 libnss_nisplus-2.31.so
lrwxrwxrwx 1 root root       19 Apr 21  2022 libnss_nisplus.so -> libnss_nisplus.so.2
lrwxrwxrwx 1 root root       22 Apr 21  2022 libnss_nisplus.so.2 -> libnss_nisplus-2.31.so
-rw-r--r-- 1 root root     6192 Apr 21  2022 libpcprofile.so
-rwxr-xr-x 1 root root   160704 Apr 21  2022 libpthread-2.31.so
-rw-r--r-- 1 root root  6428186 Apr 21  2022 libpthread.a
lrwxrwxrwx 1 root root       15 Apr 21  2022 libpthread.so -> libpthread.so.0
lrwxrwxrwx 1 root root       18 Apr 21  2022 libpthread.so.0 -> libpthread-2.31.so
-rw-r--r-- 1 root root    80624 Apr 21  2022 libresolv-2.31.so
-rw-r--r-- 1 root root   116866 Apr 21  2022 libresolv.a
lrwxrwxrwx 1 root root       14 Apr 21  2022 libresolv.so -> libresolv.so.2
lrwxrwxrwx 1 root root       17 Apr 21  2022 libresolv.so.2 -> libresolv-2.31.so
-rw-r--r-- 1 root root    55926 Apr 21  2022 librpcsvc.a
-rw-r--r-- 1 root root    31584 Apr 21  2022 librt-2.31.so
-rw-r--r-- 1 root root    78978 Apr 21  2022 librt.a
lrwxrwxrwx 1 root root       10 Apr 21  2022 librt.so -> librt.so.1
lrwxrwxrwx 1 root root       13 Apr 21  2022 librt.so.1 -> librt-2.31.so
lrwxrwxrwx 1 root root       19 Jul 10  2023 libstdc++.so.6 -> libstdc++.so.6.0.28
-rw-r--r-- 1 root root  1915656 Jul 10  2023 libstdc++.so.6.0.28
-rw-r--r-- 1 root root    35584 Apr 21  2022 libthread_db-1.0.so
lrwxrwxrwx 1 root root       17 Apr 21  2022 libthread_db.so -> libthread_db.so.1
lrwxrwxrwx 1 root root       19 Apr 21  2022 libthread_db.so.1 -> libthread_db-1.0.so
lrwxrwxrwx 1 root root       16 Jul 10  2023 libtsan.so.0 -> libtsan.so.0.0.0
-rw-r--r-- 1 root root  9207488 Jul 10  2023 libtsan.so.0.0.0
lrwxrwxrwx 1 root root       17 Jul 10  2023 libubsan.so.1 -> libubsan.so.1.0.0
-rw-r--r-- 1 root root  3078480 Jul 10  2023 libubsan.so.1.0.0
-rw-r--r-- 1 root root    14672 Apr 21  2022 libutil-2.31.so
-rw-r--r-- 1 root root    15288 Apr 21  2022 libutil.a
lrwxrwxrwx 1 root root       12 Apr 21  2022 libutil.so -> libutil.so.1
lrwxrwxrwx 1 root root       15 Apr 21  2022 libutil.so.1 -> libutil-2.31.so
```

在有了一个确定的 sysroot 后，就方便些出 toolchain file 了

```cmake
set(CMAKE_SYSTEM_NAME Linux)
set(CMAKE_SYSTEM_PROCESSOR aarch64)

set(CMAKE_C_COMPILER "/usr/bin/aarch64-linux-gnu-gcc")
set(CMAKE_CXX_COMPILER "/usr/bin/aarch64-linux-gnu-g++")
set(CMAKE_SYSROOT "/opt/arm64")

set(CMAKE_FIND_ROOT_PATH_MODE_PROGRAM NEVER)
set(CMAKE_FIND_ROOT_PATH_MODE_LIBRARY ONLY)
set(CMAKE_FIND_ROOT_PATH_MODE_INCLUDE ONLY)
set(CMAKE_FIND_ROOT_PATH_MODE_PACKAGE ONLY)
```

> CMake find*\* commands will look in the sysroot, and the CMAKE_FIND_ROOT_PATH entries by default in all cases, as well as looking in the host system root prefix. Although this can be controlled on a case-by-case basis, when cross-compiling, it can be useful to exclude looking in either the host or the target for particular artifacts. Generally, includes, libraries and packages should be found in the target system prefixes, whereas executables which must be run as part of the build should be found only on the host and not on the target. This is the purpose of the CMAKE_FIND_ROOT_PATH_MODE*\* variables.

find\_\*\* 会在 sysroot 和 `CMAKE_FIND_ROOT_PATH` 中查找头文件和库文件

```
NEVER: CMake will not search in paths specified by CMAKE_FIND_ROOT_PATH.
ONLY: CMake will only search in paths specified by CMAKE_FIND_ROOT_PATH.
BOTH: CMake will search in both the paths specified by CMAKE_FIND_ROOT_PATH and the system paths.
```

> 另外在使用 find_xxx 函数的时候可以打开 debug-find 这样 cmake 会打印出每一个查找的路径

```
set(CMAKE_FIND_DEBUG_MODE TRUE)
或者
--debug-find
# New in version 3.17.
```

system path 只 host 上默认的路径。

## 一些 Debug 的方法

configure 的时候可以使用 `--trace-expand` 来打印详细的 log

```
cmake --trace-expand ..
```

在编译的时候也可以使用 `--verbose` 打印出具体的编译或者链接命令

```
cmake --build . --verbose
```
