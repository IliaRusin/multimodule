package lib2

import "github.com/IliaRusin/multimodule/libs/lib1"

func Lib2() string {
	return "Lib1"
}

func GetLib1Value() string {
	return lib1.Lib1()
}
