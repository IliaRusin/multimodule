package logics

import (
	"fmt"
	"github.com/IliaRusin/multimodule/libs/lib1"
	"github.com/IliaRusin/multimodule/libs/lib2"
)

func Service1() string {
	return fmt.Sprintf("lib1: %s\nlib2: %s\nlib2 imported lib1: %s\n",
		lib1.Lib1(), lib2.Lib2(), lib2.Lib2())
}
