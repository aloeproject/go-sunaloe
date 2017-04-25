package admin

type ManageController struct {
	BaseController
}

func (this *ManageController) Realtime (){
	this.init()
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["Scripts"] = "backend/manage/script.html"
	this.TplName = "backend/manage/realtime.html"
	this.Render()
}
