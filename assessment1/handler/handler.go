package handler

import (
	"fmt"
	"io"
	"strings"
)

func Post(s string) error {

	m := strings.NewReader("Go is a language")
	_, err := io.ReadAll(m)
	fmt.Println(err)
	if err != nil {
		return err
	}

	return nil
}
