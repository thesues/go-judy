all: thirdparty/opt/jemalloc/lib/libjemalloc.a thirdparty/opt/judy/lib/libJudy.a
thirdparty/opt/jemalloc/lib/libjemalloc.a: thirdparty/src/jemalloc-5.2.0
	cd thirdparty/src/jemalloc-5.2.0 && ./configure --disable-debug --with-jemalloc-prefix= --prefix=${CURDIR}/thirdparty/opt/jemalloc && make -j8 && make install
thirdparty/src/jemalloc-5.2.0: thirdparty/src/jemalloc-5.2.0.tar.bz2
	cd thirdparty/src && tar xf jemalloc-5.2.0.tar.bz2
thirdparty/opt/judy/lib/libJudy.a:thirdparty/src/judy-1.0.5
	cd thirdparty/src/judy-1.0.5 && ./configure --prefix=${CURDIR}/thirdparty/opt/judy && make && make install
thirdparty/src/judy-1.0.5: thirdparty/src/Judy-1.0.5.tar.gz
	cd thirdparty/src && tar xf Judy-1.0.5.tar.gz
