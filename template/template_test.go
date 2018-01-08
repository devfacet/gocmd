/*
 * gocmd
 * For the full copyright and license information, please view the LICENSE.txt file.
 */

package template_test

import (
	"errors"
	"testing"
	"time"

	"github.com/devfacet/gocmd/template"
	. "github.com/smartystreets/goconvey/convey"
)

func TestNew(t *testing.T) {
	Convey("should create a new template", t, func() {
		tpl, err := template.New(template.Options{})
		So(err, ShouldBeNil)
		So(tpl, ShouldNotBeNil)

		tpl, err = template.New(template.Options{
			FilePath: "./test.txt",
		})
		So(err, ShouldBeNil)
		So(tpl, ShouldNotBeNil)

		tpl, err = template.New(template.Options{
			FilePath: "./test.txt",
			Data:     struct{ Test string }{Test: "foo"},
		})
		So(err, ShouldBeNil)
		So(tpl, ShouldNotBeNil)

		tpl, err = template.New(template.Options{
			Content: "{{.Test}}",
		})
		So(err, ShouldBeNil)
		So(tpl, ShouldNotBeNil)

		tpl, err = template.New(template.Options{
			Content: "{{.Test}}",
			Data:    struct{ Test string }{Test: "foo"},
		})
		So(err, ShouldBeNil)
		So(tpl, ShouldNotBeNil)
	})

	Convey("should fail to create a new template", t, func() {
		tpl, err := template.New(template.Options{
			FilePath: "./test.missing",
			Data:     struct{ Test string }{Test: "foo"},
		})
		So(err, ShouldBeError, errors.New("failed to create template due to open ./test.missing: no such file or directory"))
		So(tpl, ShouldBeNil)

		tpl, err = template.New(template.Options{
			Name:    "test",
			Content: "{{.Test",
		})
		So(err, ShouldBeError, errors.New("failed to parse template due to template: test:1: unclosed action"))
		So(tpl, ShouldBeNil)
	})
}

func TestExecute(t *testing.T) {
	Convey("should execute the template", t, func() {
		tpl, err := template.New(template.Options{})
		So(err, ShouldBeNil)
		So(tpl, ShouldNotBeNil)
		c, err := tpl.Execute(nil)
		So(err, ShouldBeNil)
		So(c, ShouldBeEmpty)

		tpl, err = template.New(template.Options{
			FilePath: "./test.txt",
		})
		So(err, ShouldBeNil)
		So(tpl, ShouldNotBeNil)
		c, err = tpl.Execute(nil)
		So(err, ShouldBeNil)
		So(c, ShouldEqual, "hello <no value>")
		c, err = tpl.Execute(struct{ Test string }{Test: "foo"})
		So(err, ShouldBeNil)
		So(c, ShouldEqual, "hello foo")

		tpl, err = template.New(template.Options{
			FilePath: "./test.txt",
			Data:     struct{ Test string }{Test: "foo"},
		})
		So(err, ShouldBeNil)
		So(tpl, ShouldNotBeNil)
		c, err = tpl.Execute(nil)
		So(err, ShouldBeNil)
		So(c, ShouldEqual, "hello foo")
		c, err = tpl.Execute(struct{ Test string }{Test: "bar"})
		So(err, ShouldBeNil)
		So(c, ShouldEqual, "hello bar")

		tpl, err = template.New(template.Options{
			Content: "{{.Test}}",
			Data:    struct{ Test string }{Test: "foo"},
		})
		So(err, ShouldBeNil)
		So(tpl, ShouldNotBeNil)
		c, err = tpl.Execute(nil)
		So(err, ShouldBeNil)
		So(c, ShouldEqual, "foo")
	})

	Convey("should fail to execute the template", t, func() {
		tpl, err := template.New(template.Options{
			Name:    "test",
			Content: "{{.Test}}",
		})
		So(err, ShouldBeNil)
		So(tpl, ShouldNotBeNil)
		c, err := tpl.Execute(struct{ test string }{test: "foo"})
		So(err, ShouldBeError, errors.New(`failed to execute template due to template: test:1:2: executing "test" at <.Test>: can't evaluate field Test in type struct { test string }`))
		So(c, ShouldEqual, "")
	})

	Convey("should execute the template by the given env variable", t, func() {
		tpl, err := template.New(template.Options{
			Content: `{{env}}`,
		})
		So(err, ShouldBeNil)
		So(tpl, ShouldNotBeNil)
		c, err := tpl.Execute(nil)
		So(err, ShouldBeNil)
		So(c, ShouldBeEmpty)

		tpl, err = template.New(template.Options{
			Content: `{{env "USER"}}`,
		})
		So(err, ShouldBeNil)
		So(tpl, ShouldNotBeNil)
		c, err = tpl.Execute(nil)
		So(err, ShouldBeNil)
		So(c, ShouldNotBeEmpty)

		tpl, err = template.New(template.Options{
			Content: `{{env "MISSING"}}`,
		})
		So(err, ShouldBeNil)
		So(tpl, ShouldNotBeNil)
		c, err = tpl.Execute(nil)
		So(err, ShouldBeNil)
		So(c, ShouldBeEmpty)

		tpl, err = template.New(template.Options{
			Content: `{{env "MISSING" "test"}}`,
		})
		So(err, ShouldBeNil)
		So(tpl, ShouldNotBeNil)
		c, err = tpl.Execute(nil)
		So(err, ShouldBeNil)
		So(c, ShouldEqual, "test")

		tpl, err = template.New(template.Options{
			Content: `{{env .Invalid}}`,
			Data:    struct{ Invalid []string }{Invalid: []string{"foo"}},
		})
		So(err, ShouldBeNil)
		So(tpl, ShouldNotBeNil)
		c, err = tpl.Execute(nil)
		So(err, ShouldBeNil)
		So(c, ShouldBeEmpty)
	})

	Convey("should execute the template by the given time format", t, func() {
		tpl, err := template.New(template.Options{
			Content: `{{time}}`,
		})
		So(err, ShouldBeNil)
		So(tpl, ShouldNotBeNil)
		c, err := tpl.Execute(nil)
		So(err, ShouldBeNil)
		So(c, ShouldContainSubstring, time.Now().Format("2006-01-02T15"))

		tpl, err = template.New(template.Options{
			Content: `{{time "format=2006-01-02T15"}}`,
		})
		So(err, ShouldBeNil)
		So(tpl, ShouldNotBeNil)
		c, err = tpl.Execute(nil)
		So(err, ShouldBeNil)
		So(c, ShouldEqual, time.Now().Format("2006-01-02T15"))

		tpl, err = template.New(template.Options{
			Content: `{{time "format=2006-01-02T15&add=2h"}}`,
		})
		So(err, ShouldBeNil)
		So(tpl, ShouldNotBeNil)
		c, err = tpl.Execute(nil)
		So(err, ShouldBeNil)
		So(c, ShouldEqual, time.Now().Add(time.Duration(time.Hour)*2).Format("2006-01-02T15"))

		tpl, err = template.New(template.Options{
			Content: `{{time "format=2006-01-02T15&add=-2h"}}`,
		})
		So(err, ShouldBeNil)
		So(tpl, ShouldNotBeNil)
		c, err = tpl.Execute(nil)
		So(err, ShouldBeNil)
		So(c, ShouldEqual, time.Now().Add(-(time.Duration(time.Hour) * 2)).Format("2006-01-02T15"))

		tpl, err = template.New(template.Options{
			Content: `{{time "format=2006-01-02T15&sub=2h"}}`,
		})
		So(err, ShouldBeNil)
		So(tpl, ShouldNotBeNil)
		c, err = tpl.Execute(nil)
		So(err, ShouldBeNil)
		So(c, ShouldEqual, time.Now().Add(-(time.Duration(time.Hour) * 2)).Format("2006-01-02T15"))
	})

	Convey("should execute the template by the given exec command", t, func() {
		tpl, err := template.New(template.Options{
			Content: `{{exec}}`,
		})
		So(err, ShouldBeNil)
		So(tpl, ShouldNotBeNil)
		c, err := tpl.Execute(nil)
		So(err, ShouldBeNil)
		So(c, ShouldBeEmpty)

		tpl, err = template.New(template.Options{
			Content: `{{exec "echo" "hello world"}}`,
		})
		So(err, ShouldBeNil)
		So(tpl, ShouldNotBeNil)
		c, err = tpl.Execute(nil)
		So(err, ShouldBeNil)
		So(c, ShouldEqual, "hello world")
	})
}
