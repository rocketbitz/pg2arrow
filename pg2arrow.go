package main

// #cgo CFLAGS: -g -Wall -I/usr/include/postgresql/server
// #include <stdlib.h>
// #include "pg2arrow.h"
import "C"
import (
	"fmt"
	"unsafe"
)

func query(sql string) ([]byte, error) {
	q := C.CString(sql)
	defer C.free(unsafe.Pointer(q))

	buf := C.query(q)

	if len(buf) == 0 {
		return nil, fmt.Errorf("failed to query postgres")
	}

	return buf, nil
}

func main() {
	buf, err := query("SELECT 1")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(buf))
}
