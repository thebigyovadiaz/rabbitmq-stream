package util

import (
	"fmt"
	"os"
)

func LogFailOnError(err error) {
	if err != nil {
		fmt.Printf("%s ", err)
		os.Exit(1)
	}
}

func LogSuccessful(msg string) {
	fmt.Println(msg)
}
