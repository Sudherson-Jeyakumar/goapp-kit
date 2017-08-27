package os

import (
	"bytes"
	"context"
	"io"
	"io/ioutil"
	"strings"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_SimpleCommandExecutor(t *testing.T) {
	Convey("Intialize a Simple Command Executor", t, func() {
		l := getSimpleExecutor()
		_, err := l.Execute(nil)
		So(err, ShouldNotBeNil)
	})

	Convey("Given a simple command executor", t, func() {
		l := getSimpleExecutor()
		l.command = "echo"
		l.args = []string{"Print data"}
		f, err := l.Execute(nil)
		So(err, ShouldBeNil)

		Convey("Try Executing echo command", func() {
			err := f(context.Background(), ioutil.Discard, ioutil.Discard)
			So(err, ShouldBeNil)
		})
	})

	Convey("Given a simple Command executor", t, func() {
		l := getSimpleExecutor()
		l.command = "echo"
		l.args = []string{"Print data"}
		f, err := l.Execute(nil)
		So(err, ShouldBeNil)

		var outbuf bytes.Buffer
		var errbuf bytes.Buffer

		Convey("Try Executing echo Command with Output and Error Buffer", func() {
			err := f(context.Background(), &outbuf, &errbuf)
			So(err, ShouldBeNil)

			So(strings.Replace(outbuf.String(), "\n", "", -1), ShouldEqual, "Print data")
			So(errbuf.String(), ShouldEqual, "")
		})
	})

	Convey("Given a simple Command executor with Arguments", t, func() {
		l := getSimpleExecutor()
		l.command = "sed"
		l.args = []string{"s/l/L/g"}

		var outbuf bytes.Buffer
		var errbuf bytes.Buffer
		var inbuf bytes.Buffer

		io.WriteString(&inbuf, "HelloWorld")

		f, err := l.Execute(&inbuf)
		So(err, ShouldBeNil)

		Convey("Try Executing echo Command with Output and Error Buffer", func() {
			err := f(context.Background(), &outbuf, &errbuf)
			So(err, ShouldBeNil)

			So(strings.Replace(outbuf.String(), "\n", "", -1), ShouldEqual, "HeLLoWorLd")
			So(errbuf.String(), ShouldEqual, "")
		})
	})
}
