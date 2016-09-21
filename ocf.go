package com_novoo_epub3

import (
	"encoding/xml"
	"os"
	"path"
)

type container struct{
	XMLName xml.Name	`xml:"container"`
	Version string		`xml:"version,attr"`
	Xmlns 	string 		`xml:"xmlns:attr"`
}

type rights struct{}
type encryption struct{}
type metadata struct{}
type signatures struct{}
type manifest struct{}

func Init(){
	tempPath := path.Join("projects","user","thingsfallapart")
	os.Mkdir(tempPath, os.ModePerm)
}