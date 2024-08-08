package main

import (
	"fmt"
	"io"
)

type writer struct{}

func (w *writer) Write(p []byte) (n int, err error) {
	return 0, nil
}

// reference: https://go.dev/blog/errors-are-values
type errWriter struct {
	w   io.Writer
	err error
}

func (ew *errWriter) write(buf []byte) {
	// if an error is occurs do doing and exit
	if ew.err != nil {
		return
	}
	_, ew.err = ew.w.Write(buf)
}

func writeString(s string) (err error) {
	p := []byte(s)
	ew := &errWriter{w: &writer{}}
	ew.write(p[0:1])
	ew.write(p[1:2])
	ew.write(p[3:4])
	// and so on
	if ew.err != nil {
		return ew.err
	}
	return nil
}

func main() {
	err := writeString("hello world")
	fmt.Printf("err: %v", err)
}

// the bufio package’s Writer is actually an implementation of the errWriter idea.
// The Write method of bufio.Writer behaves just like our errWriter.write method above,
// with Flush reporting the error
// so our example could be written like this:

// b := bufio.NewWriter(fd)
// b.Write(p0[a:b])
// b.Write(p1[c:d])
// b.Write(p2[e:f])
// if b.Flush() != nil {
//     return b.Flush()
// }
