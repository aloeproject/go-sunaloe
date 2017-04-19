package admin


type IndexController struct {
	BaseController
}

func (c *IndexController) Index() {
	c.init()
	c.TplName = "backend/index/index.html"
	c.Render()
	//c.TplName = "login/login.html"
}
