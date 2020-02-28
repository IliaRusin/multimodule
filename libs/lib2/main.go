package lib2

import "github.com/IliaRusin/multimodule/libs/lib1"

func Lib1() string {
	return "Lib1"
}

func Lib2() string {
	return lib1.Lib1()
}
