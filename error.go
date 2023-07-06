package main

import (
	"errors"
	"fmt"
)

func main() {
	err := funcA()
	if err != nil {
		panic(err)
	}
	fmt.Println("All functions Success!")
}

func funcA() error {
	err := funcB()
	if err != nil {
		return err
	}

	err = failedFunc()
	if err != nil {
		return err
	}

	err = funcC()
	if err != nil {
		return err
	}

	return nil
}

func funcB() error {
	return nil
}

func funcC() error {
	return nil
}

func failedFunc() error {
	return errors.New("error has occurred")
}
