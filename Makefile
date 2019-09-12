UNAME_S := $(shell uname -s)
ifeq ($(UNAME_S),Linux)
	BUILDJUDYCFLAGS= "-fno-strict-aliasing -fno-tree-ccp -fno-tree-dominator-opts -fno-tree-copy-prop -fno-tree-vrp"
endif
all: thirdparty/opt/judy/lib/libJudy.a judy1.go judyl.go judyhs.go 
thirdparty/opt/judy/lib/libJudy.a:thirdparty/src/judy-1.0.5
	cd thirdparty/src/judy-1.0.5 && patch -p1 < ../Judy-1.0.5-undefined-behavior.patch && CFLAGS=$(BUILDJUDYCFLAGS) ./configure --prefix=${CURDIR}/thirdparty/opt/judy && make && make install
	#cd thirdparty/src/judy-1.0.5 && ./configure --prefix=${CURDIR}/thirdparty/opt/judy && make && make install
thirdparty/src/judy-1.0.5: thirdparty/src/Judy-1.0.5.tar.gz
	cd thirdparty/src && tar xf Judy-1.0.5.tar.gz
judy1.go:judy1.go.in config_path.py
	python config_path.py $< > $@
judyl.go:judyl.go.in config_path.py
	python config_path.py $< > $@
judyhs.go:judyhs.go.in config_path.py
	python config_path.py $< > $@
clean:
	rm -rf judy1.go judyl.go judyhs.go thirdparty/opt/ thirdparty/src/judy-1.0.5
