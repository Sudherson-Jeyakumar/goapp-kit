package os

import (
	"bytes"
	"context"
	"strings"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_SeriallyExecute(t *testing.T) {

	Convey("Given a simple Command executor", t, func() {
		l := getSimpleExecutor()
		l.command = "echo"
		l.args = []string{"First Command Output"}

		l2 := getSimpleExecutor()
		l2.command = "echo"
		l2.args = []string{"Second Command Output"}

		var outbuf bytes.Buffer
		var errbuf bytes.Buffer

		Convey("Execute both the commands serially", func() {
			err := SeriallyExecute(context.Background(), &outbuf, &errbuf, l, l2)
			So(err, ShouldBeNil)

			So(strings.Replace(outbuf.String(), "\n", "", -1), ShouldEqual, "First Command OutputSecond Command Output")
			So(errbuf.String(), ShouldEqual, "")
		})
	})
}

func Test_PipeExecute(t *testing.T) {

	Convey("Given a simple Command executor", t, func() {
		l := getSimpleExecutor()
		l.command = "echo"
		l.args = []string{"First Command Output"}

		l2 := getSimpleExecutor()
		l2.command = "sed"
		l2.args = []string{"s/First/Second/g"}

		var outbuf bytes.Buffer
		var errbuf bytes.Buffer

		Convey("Execute both the commands in Pipe", func() {
			err := PipeExecute(context.Background(), &outbuf, &errbuf, l, l2)
			So(err, ShouldBeNil)

			So(strings.Replace(outbuf.String(), "\n", "", -1), ShouldEqual, "Second Command Output")
			So(errbuf.String(), ShouldEqual, "")
		})
	})
}
