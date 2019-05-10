#!/usr/bin/python
import string
import sys
import os


INCLUDE_PATH=os.getcwd() + "/" +"thirdparty/opt/judy/include -O3 -g"
LIBRARY_PATH=os.getcwd() + "/" +"thirdparty/opt/jemalloc/lib/libjemalloc.a" + " " + os.getcwd() + "/" + "thirdparty/opt/judy/lib/libJudy.a -ldl -lpthread"

input = sys.argv[1]
s = open(input,"r")
for line in s.read().splitlines():
   line = string.replace(line, '${INCLUDE_PATH}',INCLUDE_PATH)
   line = string.replace(line, '${LIBRARY_PATH}',LIBRARY_PATH)
   print line
s.close()
