//go:build ignore

package main

import (
	"context"
	"net/http"
)

func HandlerOK1() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		println("something first")

		err := doSomething(r.Context())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		println("something later")
	}
}

func HandlerOK2() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		println("something first")

		go func(ctx context.Context) {
			err := doSomething(ctx)
			if err != nil {
				println(err.Error())
			}
		}(context.Background())

		ctx := context.Background()
		go func(ctx context.Context) {
			err := doSomething(ctx)
			if err != nil {
				println(err.Error())
			}
		}(ctx)

		println("something later")
	}
}

func Handler1() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		println("something first")

		err := doSomething(context.Background())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		println("something later")
	}
}

func Handler2() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.Background()

		println("something first")

		err := doSomething(ctx)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		println("something later")
	}
}

type H struct{}

func (h H) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	println("something first")

	err := doSomething(context.Background())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	println("something later")
}

func doSomething(ctx context.Context) error {
	panic("not implelemented")
}
