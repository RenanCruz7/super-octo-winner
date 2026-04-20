package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func check3(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	path := filepath.Join(os.TempDir(), "dat")
	dat, err := os.ReadFile(path)
	check3(err)
	fmt.Print(string(dat))

	f, err := os.Open(path)
	check3(err)

	b1 := make([]byte, 5)
	n1, err := f.Read(b1)
	check3(err)
	fmt.Printf("%d bytes: %s\n", n1, string(b1[:n1]))

	o2, err := f.Seek(6, io.SeekStart)
	check3(err)
	b2 := make([]byte, 2)
	n2, err := f.Read(b2)
	check3(err)
	fmt.Printf("%d bytes @ %d: ", n2, o2)
	fmt.Printf("%v\n", string(b2[:n2]))

	_, err = f.Seek(2, io.SeekCurrent)
	check3(err)

	_, err = f.Seek(-4, io.SeekEnd)
	check3(err)

	o3, err := f.Seek(6, io.SeekStart)
	check3(err)
	b3 := make([]byte, 2)
	n3, err := io.ReadAtLeast(f, b3, 2)
	check3(err)
	fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))

	_, err = f.Seek(0, io.SeekStart)
	check3(err)

	r4 := bufio.NewReader(f)
	b4, err := r4.Peek(5)
	check3(err)
	fmt.Printf("5 bytes: %s\n", string(b4))

	f.Close()
}
