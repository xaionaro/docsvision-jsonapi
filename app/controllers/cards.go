package controllers;

import (
	"github.com/revel/revel";
	"github.com/xaionaro/docsvision-jsonapi/app/models";
	"github.com/xaionaro/docsvision-jsonapi/app";
)

type Cards struct {
	*revel.Controller;
}

func (c Cards) Find(p *models.CardParams) (cards []models.Card) {
	_, err := app.Dbm.Select(&cards, "SELECT * FROM [dbo].[dvsys_carddefs]");
	if (err != nil) {
		panic(err);
	}

	//fmt.Printf("test: %v\n", cards);
	return cards;
}

func (c Cards) Index() (cards []models.Card) {
	return c.Find(nil);
}

func (c Cards) IndexHtml() revel.Result {
	return c.Render();
}

func (c Cards) IndexJson() revel.Result {
	return c.RenderJson(c.Index());
}
