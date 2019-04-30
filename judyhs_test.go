package judy

import (
	"math/rand"
	"testing"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandStringBytes(n int) string {
    b := make([]byte, n)
    for i := range b {
        b[i] = letterBytes[rand.Intn(len(letterBytes))]
    }
    return string(b)
}

func TestEmptyJudyHSArray(t *testing.T) {

	j := JudyHS{}
	r := j.Free()

	if r != 0 {
		t.Errorf("Free should return 0, returned %v", r)
	}
}

type Case struct {
	s string
	v uint64
}

func TestJudyHSBasic(t *testing.T) {
	j := JudyHS{}
	cases := []Case{
		Case{"asdf", 10},
		Case{"asdf2", 11},
		Case{"asdf4", 12},
		Case{"hello world", 90},
		Case{"中文", 190},
		Case{"test", 93430},
		Case{"guess\n\tasdf", 91110},
		Case{"you", 10},
	}

	//insert data
	for _, i := range cases {
		j.Insert([]byte(i.s), i.v)
	}


	//verifiy data
	for _, i := range cases {
		data, _ := j.Get([]byte(i.s))
		if data != i.v {
			t.Errorf("expected is %v, but get data %d\n", i, data)
		}
	}

	for _, i := range cases {
		if ok := j.Delete([]byte(i.s)); !ok {
			t.Errorf("should have %v to be deleted\n", i)
		}
	}

	ok := j.Delete([]byte("中文"));
	if ok == true {
		t.Errorf("should be false")
	}

	j.Free()
}
