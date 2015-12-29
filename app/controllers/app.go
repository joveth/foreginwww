package controllers

import (
	"ForeginWWW/app/utils"
	"github.com/revel/revel"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	return c.Render()
}
func (c App) Mima() revel.Result {
	content, _ := utils.GetHTMLContent("https://www.findmima.com/")
	return c.RenderHtml(content)
}
