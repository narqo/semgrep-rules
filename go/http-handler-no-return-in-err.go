//go:build ignore

package main

import (
	"fmt"
	"io"
	"net/http"
)

func HandlerOK1() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		println("something first")

		var err error
		if err != nil {
			writeError(w, r, err)
			return
		}

		println("something later")
	}
}

func HandlerOK2() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		println("something first")

		var ss []int
		if ss != nil {
			println("nil ss")
		}

		println("something later")
	}
}

func HandlerOK3() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		println("something first")

		go func() {
			var err error
			if err != nil {
				println(err.Error())
			}
		}()

		println("something later")
	}
}

func HandlerOK4() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		println("something first")

		if _, err := io.WriteString(w, "done"); err != nil {
			writeError(w, r, err)
		}
	}
}

func HandlerOK5() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		println("something first")

		var err error
		if err != nil {
			writeError(w, r, err)
		}
	}
}

func HandlerOK6() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		println("something first")

		var err error
		if err != nil {
			fmt.Printf("%v", err)
		}

		println("something later")
	}
}

func Handler1() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		println("something first")

		var err error
		if err != nil {
			writeError(w, r, err)
		}

		println("something later")
	}
}

func Handler2(w http.ResponseWriter, r *http.Request) {
	println("something first")

	var err error
	if err != nil {
		writeError(w, r, err)
	}

	println("something later")
}

type H struct{}

func (h H) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	println("something first")

	var err error
	if err != nil {
		writeError(w, r, err)
	}

	println("something later")
}

func writeError(w http.ResponseWriter, r *http.Request, err error) {
	// TODO
}
