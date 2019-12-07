package main

import (
	"fmt"
	"github.com/pkg/errors"
)


func main() {
	err := errors.New("错误发生地方")
	err = errors.Wrap(err, "错误2")
	err = errors.Wrap(err, "错误3")
	err = errors.Wrap(err, "错误4")

	err = errors.WithMessage(err, "message")
	//错误堆栈
	fmt.Println(err)

	//错误来源
	fmt.Println(errors.Cause(err))
}
