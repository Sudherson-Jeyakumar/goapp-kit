package log

import (
	"bytes"
	"os"
	"strings"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_BaseLogger(t *testing.T) {
	Convey("Intialize an empty Writer in logger", t, func() {
		_, err := InitBaseLogger(nil, "hello")

		So(err, ShouldNotBeNil)
	})

	Convey("Intialize a non empty Writer in logger", t, func() {
		_, err := InitBaseLogger(os.Stdout, "hello")

		So(err, ShouldBeNil)
	})

	var w bytes.Buffer
	Convey("With Buffer as writer", t, func() {
		s, err := InitBaseLogger(&w, "hello")
		So(err, ShouldBeNil)

		Convey("Writer an info", func() {
			s.Infof("%s", "Writing first String")
		})

		So(strings.TrimRight(w.String(), "\n"), ShouldEndWith, "INFO <hello> : Writing first String")
	})
}

func TestBaseLogger_ChangePrefix(t *testing.T) {
	var w bytes.Buffer
	Convey("Given, With Buffer as writer", t, func() {
		s, err := InitBaseLogger(&w, "hello")
		So(err, ShouldBeNil)

		s.Infof("%s", "Writing first String")
		So(strings.TrimRight(w.String(), "\n"), ShouldEndWith, "INFO <hello> : Writing first String")

		Convey("When prefix is changed, content should match", func() {
			s.PrefixLogType("info", "changed")
			s.Infof("%s", "Writing second String")
			So(strings.TrimRight(w.String(), "\n"), ShouldEndWith, "INFO <changed> : Writing second String")

		})

	})
}
