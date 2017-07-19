package cfg

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_SimpleConfig(t *testing.T) {
	Convey("Given, an empty Config", t, func() {
		s := Simple()

		Convey("Check the value of the data structure", func() {
			So(len(s), ShouldBeZeroValue)
		})
	})

	Convey("Check, when a value is set in empty path", t, func() {
		s := Simple()

		s.Set("hello")
		So(len(s), ShouldBeZeroValue)
	})

	Convey("Check, when a value is set in non empty path", t, func() {
		s := Simple()

		s.Set("I am here", "hello")
		So(len(s), ShouldEqual, 1)

		So(s.Value("hello"), ShouldEqual, "I am here")
	})

	Convey("Check, when multiple value is set in non empty path", t, func() {
		s := Simple()
		s.Set([]string{"this", "is", "an", "array", "test"}, "firstlevel", "secondlevel", "thirdlevel")
		So(len(s), ShouldEqual, 1)

		So(fmt.Sprintf("%v", s.Value("firstlevel", "secondlevel", "thirdlevel")), ShouldEqual, fmt.Sprintf("%v", []string{"this", "is", "an", "array", "test"}))
	})

	Convey("Check, when multiple value is set in non empty path", t, func() {
		s1 := Simple()

		s1.Set([]string{"this", "is", "an", "array", "test"}, "firstlevel", "secondlevel", "thirdlevel")
		So(len(s1), ShouldEqual, 1)

		Convey("Lookup for a non existence path", func() {
			So(s1.Value("firstlevel", "thirdlevel"), ShouldBeNil)
		})

		Convey("Lookup for an empty path", func() {
			So(s1.Value(), ShouldBeNil)
		})

		Convey("Lookup for overwritten path", func() {
			s1.Set("I am overwritten here", "firstlevel", "secondlevel")
			So(s1.Value("firstlevel", "secondlevel", "thirdlevel"), ShouldBeNil)
			So(s1.Value("firstlevel", "secondlevel"), ShouldEqual, "I am overwritten here")
		})
	})

}
