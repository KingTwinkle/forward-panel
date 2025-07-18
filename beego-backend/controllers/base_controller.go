package controllers

import (
	"encoding/json"
	"github.com/beego/beego/v2/server/web"
	"net/http"
)

type BaseController struct {
	web.Controller
}

func (c *BaseController) ServeJSON() {
	c.Ctx.Output.Header("Content-Type", "application/json; charset=utf-8")

	var err error

	data := c.Data["json"]

	encoder := json.NewEncoder(c.Ctx.Output.Context.ResponseWriter)
	encoder.SetEscapeHTML(false)

	encoder.SetIndent("", "  ")
	err = encoder.Encode(data)

	if err != nil {
		http.Error(c.Ctx.Output.Context.ResponseWriter, err.Error(), http.StatusInternalServerError)
		return
	}
}
