package models


import (
	_ "github.com/revel/revel";
)

type Section struct {
	SectionTypeID	uniqueidentifier;
	Alias		string;
	ParentSectionId	uniqueidentifier;
	CardTypeId	uniqueidentifier;
	SecurityType	uint8;
	UserDependent	bool;
	NestLevel	uint8;
	Type		uint8;
	Flags		uint16;
	IsDynamic	bool;
}

type SectionParams struct {
	CardTypeId	uniqueidentifier;
}

