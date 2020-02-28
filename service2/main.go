package main

import (
	"fmt"
	"service2/localpackage"

	"github.com/IliaRusin/multimodule/libs/lib1"
	"github.com/IliaRusin/multimodule/libs/lib2"
)

func main() {
	fmt.Println(lib2.Lib2())
	fmt.Println(lib2.Lib1())
	fmt.Println(lib1.Lib1())

	localpackage.MyAwesomeLocalPackageFunction()
}
