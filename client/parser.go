package client

import (
	"fmt"
	"os"
	"strings"
)

type Classmates []string

func (i *Classmates) String() string {
	return fmt.Sprint(*i)
}

func (i *Classmates) Set(value string) error {
	for _, dt := range strings.Split(value, ",") {
		*i = append(*i, dt)
	}
	return nil
}

var Mates Classmates

func HomeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE") // windows
}
