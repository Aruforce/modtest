package modtest

import (
	"fmt"
	"github.com/Aruforce/hello"
	hv2 "github.com/Aruforce/hello/v2"
)

func Mod() string{
	fmt.Println(hv2.Hello())
	return hello.Hello();
}