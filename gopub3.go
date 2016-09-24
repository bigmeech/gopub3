package gopub3

import components "github.com/bigmeech/gopub3/components"


type EpubV3 struct {
	author Author;
}

func (ctx EpubV3) Create(title string, authorName string, isbn int32) {
	components.OCF{}
	author := Author{
		title: title,
		authorName: authorName,
		isbn : isbn,
	}
}