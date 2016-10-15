// This file is automatically generated by qtc from "readme.qtpl".
// See https://github.com/valyala/quicktemplate for details.

//line lib/templates/initreadme/qtpl/readme.qtpl:1
package qtpl

//line lib/templates/initreadme/qtpl/readme.qtpl:1
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line lib/templates/initreadme/qtpl/readme.qtpl:1
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line lib/templates/initreadme/qtpl/readme.qtpl:1
func StreamReadme(qw422016 *qt422016.Writer, name, template string) {
	//line lib/templates/initreadme/qtpl/readme.qtpl:1
	qw422016.N().S(`#`)
	//line lib/templates/initreadme/qtpl/readme.qtpl:1
	qw422016.E().S(name)
	//line lib/templates/initreadme/qtpl/readme.qtpl:1
	qw422016.N().S(` [![made with fap](https://img.shields.io/badge/made%20with-fap-brightgreen.svg)](https://github.com/kirillDanshin/fap)

Made with [fap](https://github.com/kirillDanshin/fap) using `)
	//line lib/templates/initreadme/qtpl/readme.qtpl:3
	qw422016.E().S(template)
	//line lib/templates/initreadme/qtpl/readme.qtpl:3
	qw422016.N().S(`
`)
//line lib/templates/initreadme/qtpl/readme.qtpl:4
}

//line lib/templates/initreadme/qtpl/readme.qtpl:4
func WriteReadme(qq422016 qtio422016.Writer, name, template string) {
	//line lib/templates/initreadme/qtpl/readme.qtpl:4
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line lib/templates/initreadme/qtpl/readme.qtpl:4
	StreamReadme(qw422016, name, template)
	//line lib/templates/initreadme/qtpl/readme.qtpl:4
	qt422016.ReleaseWriter(qw422016)
//line lib/templates/initreadme/qtpl/readme.qtpl:4
}

//line lib/templates/initreadme/qtpl/readme.qtpl:4
func Readme(name, template string) string {
	//line lib/templates/initreadme/qtpl/readme.qtpl:4
	qb422016 := qt422016.AcquireByteBuffer()
	//line lib/templates/initreadme/qtpl/readme.qtpl:4
	WriteReadme(qb422016, name, template)
	//line lib/templates/initreadme/qtpl/readme.qtpl:4
	qs422016 := string(qb422016.B)
	//line lib/templates/initreadme/qtpl/readme.qtpl:4
	qt422016.ReleaseByteBuffer(qb422016)
	//line lib/templates/initreadme/qtpl/readme.qtpl:4
	return qs422016
//line lib/templates/initreadme/qtpl/readme.qtpl:4
}