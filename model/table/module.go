package table

import "github.com/jinzhu/gorm"

type Module struct {
	gorm.Model
	Title string
}
