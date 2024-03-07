package table

import "github.com/jinzhu/gorm"

type ArticleImage struct {
	gorm.Model
	Aid  uint
	Path string
}
