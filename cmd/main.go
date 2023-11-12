package main

import (
	"bytes"
	"fmt"
)

func debugReader(r *bytes.Reader) {
    fmt.Printf(
        `Len: %d\n
         Values: %v\n`, 
         r.Len(), r,
     )

}

func main() {
	arr := [1024]byte{}
    reader := bytes.NewReader(arr[:])
    debugReader(reader)
    
}
