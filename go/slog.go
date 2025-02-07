//go:build ignore

package main

import (
	"errors"

	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
)

func main() {
	logger := log.Default

	err := errors.New("blah")
	level.Info(logger).Log("msg", "hello", "user", "user-id", "err", err)
}
