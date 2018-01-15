/*
 * gocmd
 * For the full copyright and license information, please view the LICENSE.txt file.
 */

package table_test

import (
	"testing"

	"github.com/devfacet/gocmd/table"
	. "github.com/smartystreets/goconvey/convey"
)

func TestNew(t *testing.T) {
	Convey("should create a new table", t, func() {
		t, err := table.New(table.Options{
			Data: [][]string{
				[]string{0: "foo", 1: "bar"},
				[]string{0: "1", 1: "2"},
			},
		})

		So(err, ShouldBeNil)
		So(t, ShouldNotBeNil)

		var d = t.Data()
		So(d[0][0], ShouldEqual, "foo")
		So(d[0][1], ShouldEqual, "bar")
		So(d[1][0], ShouldEqual, "1")
		So(d[1][1], ShouldEqual, "2")
	})
}

func TestTable_Data(t *testing.T) {
	Convey("should return the given data", t, func() {
		t, err := table.New(table.Options{})

		So(err, ShouldBeNil)
		So(t, ShouldNotBeNil)

		t.SetData(1, 1, "foo")
		t.SetData(1, 2, "bar")
		t.SetData(2, 1, "1")
		t.SetData(2, 2, "2")

		var d = t.Data()
		So(d[0][0], ShouldEqual, "foo")
		So(d[0][1], ShouldEqual, "bar")
		So(d[1][0], ShouldEqual, "1")
		So(d[1][1], ShouldEqual, "2")
	})
}

func TestTable_SetData(t *testing.T) {
	Convey("should set the given data", t, func() {
		t, err := table.New(table.Options{})

		So(err, ShouldBeNil)
		So(t, ShouldNotBeNil)

		So(t.SetData(1, 1, "foo"), ShouldBeNil)
		So(t.SetData(1, 2, "bar"), ShouldBeNil)
		So(t.SetData(2, 1, "1"), ShouldBeNil)
		So(t.SetData(2, 2, "2"), ShouldBeNil)
	})

	Convey("should fail to set the given data", t, func() {
		t, err := table.New(table.Options{})

		So(err, ShouldBeNil)
		So(t, ShouldNotBeNil)

		So(t.SetData(0, 1, "foo"), ShouldBeError, "invalid row")
		So(t.SetData(1, 0, "foo"), ShouldBeError, "invalid column")
	})
}

func TestTable_SetRow(t *testing.T) {
	Convey("should set the given row", t, func() {
		t, err := table.New(table.Options{})

		So(err, ShouldBeNil)
		So(t, ShouldNotBeNil)
		So(t.SetRow(1, "foo", "bar"), ShouldBeNil)
		So(t.SetRow(2, "1", "2"), ShouldBeNil)
		So(t.SetRow(-1, "foo", "bar"), ShouldBeError, "invalid row")
	})
}

func TestTable_AddRow(t *testing.T) {
	Convey("should add the given row", t, func() {
		t, err := table.New(table.Options{})

		So(err, ShouldBeNil)
		So(t, ShouldNotBeNil)
		So(t.AddRow("foo", "bar"), ShouldBeNil)
		So(t.AddRow("1", "2"), ShouldBeNil)
	})
}

func TestTable_FormattedData(t *testing.T) {
	Convey("should return the formatted table data", t, func() {
		t, err := table.New(table.Options{})
		So(err, ShouldBeNil)
		So(t, ShouldNotBeNil)
		So(t.AddRow("foo", "bar"), ShouldBeNil)
		So(t.FormattedData(), ShouldEqual, "foo	bar\t\n")
	})
}
