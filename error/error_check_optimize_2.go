package main

import (
	"fmt"
	"io"
	"os"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func CopyFile(src, dst string) (err error) {
	var r, w *os.File

	defer func() {
		if r != nil {
			err := r.Close()
			if err != nil {
				return
			}
		}
		if w != nil {
			err := w.Close()
			if err != nil {
				return
			}
		}
		if e := recover(); e != nil {
			if w != nil {
				err := os.Remove(dst)
				if err != nil {
					return
				}
			}
			err = fmt.Errorf("copy %s %s: %v", src, dst, err)
		}
	}()

	r, err = os.Open(src)
	check(err)
	w, err = os.Create(dst)
	check(err)
	_, err = io.Copy(w, r)
	check(err)

	return nil
}
