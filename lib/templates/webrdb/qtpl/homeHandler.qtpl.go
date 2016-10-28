// This file is automatically generated by qtc from "homeHandler.qtpl".
// See https://github.com/valyala/quicktemplate for details.

//line lib/templates/webrdb/qtpl/homeHandler.qtpl:1
package qtpl

//line lib/templates/webrdb/qtpl/homeHandler.qtpl:1
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line lib/templates/webrdb/qtpl/homeHandler.qtpl:1
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line lib/templates/webrdb/qtpl/homeHandler.qtpl:1
func StreamHomeHandler(qw422016 *qt422016.Writer, packageName string) {
	//line lib/templates/webrdb/qtpl/homeHandler.qtpl:1
	if packageName == "main" {
		//line lib/templates/webrdb/qtpl/homeHandler.qtpl:1
		qw422016.N().S(`package main`)
		//line lib/templates/webrdb/qtpl/homeHandler.qtpl:1
	} else {
		//line lib/templates/webrdb/qtpl/homeHandler.qtpl:1
		qw422016.N().S(`package `)
		//line lib/templates/webrdb/qtpl/homeHandler.qtpl:1
		qw422016.E().S(packageName)
		//line lib/templates/webrdb/qtpl/homeHandler.qtpl:1
	}
	//line lib/templates/webrdb/qtpl/homeHandler.qtpl:1
	qw422016.N().S(`

import (
	"fmt"
	"github.com/valyala/fasthttp"
	r "gopkg.in/dancannon/gorethink.v2"
)

var (
	homeGreetsTpl = []byte("Hi, ")
	homeGreetsNameParam = "name"
	homeGreetsDefaultName = []byte("Guest")
	homeCounterTpl = []byte("This page was viewed: ")
	homeCounterEnd = []byte(" times")
)

func homeHandler(ctx *fasthttp.RequestCtx) {
	ctx.Write(homeGreetsTpl)
	if name := ctx.UserValue(homeGreetsNameParam); name != nil && len(name.(string)) > 0 {
		ctx.WriteString(name.(string))
	} else {
		ctx.Write(homeGreetsDefaultName)
	}
	filter := map[string]interface{}{
	    "uri": string(ctx.URI().Path()),
	}
	var (
		tmp []interface{}
		count = float64(0)
	)
 	f, _ := r.Table("stats").Filter(filter).Run(rdb)
	f.All(&tmp)
	f.Close()
	if len(tmp) > 0 {
		r.Table("stats").Filter(filter).Update(map[string]interface{}{
			"views": r.Row.Field("views").Add(1).Default(0),
		}).RunWrite(rdb)
		v, _ := r.Table("stats").Filter(filter).Run(rdb)
		c := map[string]interface{}{}
		v.Peek(&c)
		count = c["views"].(float64)
		v.Close()
	} else {
		filter["views"] = 1
		r.Table("stats").Insert(filter).RunWrite(rdb)
		count = 1
	}
	fmt.Fprintln(ctx)
	ctx.Write(homeCounterTpl)
	fmt.Fprint(ctx, count)
	ctx.Write(homeCounterEnd)
}
`)
//line lib/templates/webrdb/qtpl/homeHandler.qtpl:53
}

//line lib/templates/webrdb/qtpl/homeHandler.qtpl:53
func WriteHomeHandler(qq422016 qtio422016.Writer, packageName string) {
	//line lib/templates/webrdb/qtpl/homeHandler.qtpl:53
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line lib/templates/webrdb/qtpl/homeHandler.qtpl:53
	StreamHomeHandler(qw422016, packageName)
	//line lib/templates/webrdb/qtpl/homeHandler.qtpl:53
	qt422016.ReleaseWriter(qw422016)
//line lib/templates/webrdb/qtpl/homeHandler.qtpl:53
}

//line lib/templates/webrdb/qtpl/homeHandler.qtpl:53
func HomeHandler(packageName string) string {
	//line lib/templates/webrdb/qtpl/homeHandler.qtpl:53
	qb422016 := qt422016.AcquireByteBuffer()
	//line lib/templates/webrdb/qtpl/homeHandler.qtpl:53
	WriteHomeHandler(qb422016, packageName)
	//line lib/templates/webrdb/qtpl/homeHandler.qtpl:53
	qs422016 := string(qb422016.B)
	//line lib/templates/webrdb/qtpl/homeHandler.qtpl:53
	qt422016.ReleaseByteBuffer(qb422016)
	//line lib/templates/webrdb/qtpl/homeHandler.qtpl:53
	return qs422016
//line lib/templates/webrdb/qtpl/homeHandler.qtpl:53
}
