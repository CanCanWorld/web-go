package table

import "github.com/jinzhu/gorm"

type Article struct {
	gorm.Model
	Uid        uint
	Mid        uint
	Title      string
	Content    string
	LikeNum    int
	CommentNum int
	CollectNum int
}
