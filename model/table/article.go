package table

import "github.com/jinzhu/gorm"

type Article struct {
	gorm.Model
	Uid        uint
	Username   string
	Mid        uint
	Title      string
	Content    string
	LikeNum    int
	CommentNum int
	CollectNum int
}
