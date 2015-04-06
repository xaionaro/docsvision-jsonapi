package controllers;


import (
	"github.com/revel/revel";
	"github.com/xaionaro/docsvision-jsonapi/app/models";
	"github.com/xaionaro/docsvision-jsonapi/app";
)

type Sections struct {
	*revel.Controller;
}

func (c Sections) Find(p *models.SectionParams) (sections []models.Section) {
	_, err := app.Dbm.Select(&sections, "SELECT * FROM [dbo].[dvsys_sectiondefs]");
	if (err != nil) {
		panic(err);
	}

	//fmt.Printf("test: %v\n", sections);
	return sections;
}

func (c Sections) Index() (sections []models.Section) {
	return c.Find(nil);
}

func (c Sections) IndexHtml() revel.Result {
	return c.Render();
}

func (c Sections) IndexJson() revel.Result {
	return c.RenderJson(c.Index());
}


