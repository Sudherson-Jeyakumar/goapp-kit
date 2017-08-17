package log

import (
	"bytes"
	"io/ioutil"
	"os"
	"strings"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_BaseLogger(t *testing.T) {
	Convey("Intialize an empty Writer in logger", t, func() {
		l := initBaseLogger(nil, "hello")

		So(l, ShouldNotBeNil)
	})

	var w bytes.Buffer
	Convey("With Buffer as writer", t, func() {
		s := initBaseLogger(&w, "hello")
		So(s, ShouldNotBeNil)

		Convey("Writer an info", func() {
			s.Infof("%s", "Writing first String")
		})

		So(strings.TrimRight(w.String(), "\n"), ShouldEndWith, "INFO <hello> : Writing first String")
	})
}

func TestBaseLogger_ChangePrefix(t *testing.T) {
	var w bytes.Buffer
	Convey("Given, With Buffer as writer", t, func() {
		s := initBaseLogger(&w, "hello")
		So(s, ShouldNotBeNil)

		s.Infof("%s", "Writing first String")
		So(strings.TrimRight(w.String(), "\n"), ShouldEndWith, "INFO <hello> : Writing first String")

		Convey("When prefix is changed, content should match", func() {
			s.PrefixLogType("info", "changed")
			s.Infof("%s", "Writing second String")
			So(strings.TrimRight(w.String(), "\n"), ShouldEndWith, "INFO <changed> : Writing second String")

		})

		Convey("When common prefix is changed, content should match", func() {
			s.Prefix("common prefix changed")
			s.Infof("%s", "Writing third String")
			So(strings.TrimRight(w.String(), "\n"), ShouldEndWith, "INFO <common prefix changed> : Writing third String")

			s.Errorf("%s", "Writing fourth String")
			So(strings.TrimRight(w.String(), "\n"), ShouldEndWith, "ERROR <common prefix changed> : Writing fourth String")

		})
	})
}

func TestBaseLogger_Enabler(t *testing.T) {
	var w bytes.Buffer
	Convey("Given, With Buffer as writer", t, func() {
		s := initBaseLogger(&w, "hello")
		So(s, ShouldNotBeNil)

		s.Enabler(false)

		s.Infof("%s", "Writing first String")
		So(strings.TrimRight(w.String(), "\n"), ShouldEqual, "")
	})
}

func TestFileLogger_ChangePrefix(t *testing.T) {

	os.Remove("./test/log.txt")

	Convey("Given, With File as writer", t, func() {
		Convey("Wrong FileName is specified", func() {
			_, err := initLoggerWithFile("/invaliddir/test/log.txt")
			So(err, ShouldNotBeNil)
		})

		Convey("Given that a file is prepared", func() {
			s, err := initLoggerWithFile("./test/log.txt")
			So(s, ShouldNotBeNil)
			So(err, ShouldBeNil)

			s.Infof("%s", "Writing first String")

			Convey("When prefix is changed, content should match", func() {
				s.PrefixLogType("info", "changed")
				s.Infof("%s", "Writing second String")
				s.Close()

				allContent, err := ioutil.ReadFile("./test/log.txt")
				So(err, ShouldBeNil)
				So(strings.Trim(string(allContent), "\n"), ShouldEndWith, "INFO <changed> : Writing second String")
			})
		})

		Convey("When an empty writer is given as input", func() {
			_, err := initLoggerWithFile(" ")
			So(err, ShouldNotBeNil)
		})
	})
}
