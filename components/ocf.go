//This file contain structs which will be mapped against their xml template counter-parts
//transformed and the written out to a destination folder

package components

import (
	"encoding/xml"
	"fmt"
)

//debug
var (
	p = fmt.Println
)

//single rootfile element
type rootfile struct {
	fullPath string		`xml:"full-path,attr"`
	mediaType string	`xml:"media-type,attr"`
}

//rootfiles container element for single roofile
type rootfiles struct {
	rootfile []rootfile
}

//container element
type container struct {
	XMLName xml.Name	`xml:"container"`
	Version string		`xml:"version,attr"`
	Xmlns 	string 		`xml:"xmlns,attr"`
	rootfiles rootfiles
}

type rights struct {}
type encryption struct {
	XMLName xml.Name	`xml:"encryption`

}
type metadata struct {}
type signatures struct {}
type manifest struct {}

type OCF struct{
	container container
	encryption encryption
	rights rights
	signatures signatures
	manifest manifest
}

type OCFHelper struct{}

func (ctx OCF) Initialize() {}
func (ctx OCF) writer() {}

