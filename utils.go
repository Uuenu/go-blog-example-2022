package main

import (
	"crypto/rand"
	"fmt"
)

//генерация рандомного id
func GenerateId() string {
	b := make([]byte, 16)
	rand.Read(b)
	return fmt.Sprint("%x", b)
}
