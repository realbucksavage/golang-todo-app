package controllers

import (
	"github.com/revel/revel"

	"gitlab.com/golang-commonmark/markdown"

	"io/ioutil"
)

type Application struct {
	*revel.Controller
}

func (c Application) Index() revel.Result {
	md := markdown.New(markdown.XHTMLOutput(true))
	buf, err := ioutil.ReadFile("README.md")

	if err != nil {
		buf = []byte("Cannot open README")
	}

	return c.RenderHTML(md.RenderToString(buf))
}
