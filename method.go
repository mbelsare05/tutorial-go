package tutorial

import (
	"bytes"
	szlib "compress/zlib"
	"github.com/klauspost/compress/zlib"
	"io/ioutil"
)

func Fuzz(data []byte) int {
	fr, err := szlib.NewReader(bytes.NewReader(data))
	if err != nil {
		return 0
	}
	data1, err := ioutil.ReadAll(fr)
	if err != nil {
		return 0
	}
	for level := 0; level <= 9; level++ {
		buf := new(bytes.Buffer)
		fw, err := zlib.NewWriterLevel(buf, level)
		if err != nil {
			panic(err)
		}
		n, err := fw.Write(data1)
		if n != len(data1) {
			panic("short write")
		}
		if err != nil {
			panic(err)
		}
		err = fw.Close()
		if err != nil {
			panic(err)
		}
		fr1, err := zlib.NewReader(buf)
		if err != nil {
			panic(err)
		}
		data2, err := ioutil.ReadAll(fr1)
		if err != nil {
			panic(err)
		}
		if bytes.Compare(data1, data2) != 0 {
			panic("not equal")
		}
	}
	return 1
}
