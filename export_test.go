package main

import "testing"
import . "github.com/smartystreets/goconvey/convey"

func TestAdd(t *testing.T) {
	Convey("test add", t, func() {
		a := 3
		b := 1
		c := Add(a, b)
		So(c, ShouldEqual, 4)
	})
}
