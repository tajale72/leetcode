package main

import (
	"testing"
)

type Mock interface{}

func TestHomepage(t *testing.T) {
	c := Mock
	Homepage(c)
}
