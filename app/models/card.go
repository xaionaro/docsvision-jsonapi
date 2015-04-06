package models

import (
	_ "github.com/revel/revel";
)

type Card struct {
	CardTypeID	uniqueidentifier;
	Alias		string;
	Version		int;
	SysVersion	int;
	LibraryID	uniqueidentifier;
	ControlInfo	string;
	Options		int;
	FetchMode	int;
	XMLSchema	string;
	XSDSchema	string;
	Icon		[]byte;
	SDID		uniqueidentifier;
	Timestamp	Time;
	TypeName	string;
}

type CardParams struct {
	
}

