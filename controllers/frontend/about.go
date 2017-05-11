package frontend

type AboutController struct {
	BaseController
}

func (this *AboutController) Index()  {
	this.init()
	this.Data["title"] = "关于 Sunaloe"
	this.TplName = "frontend/about/index.html"
	this.Render()
}
