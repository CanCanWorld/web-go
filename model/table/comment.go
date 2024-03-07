package table

import "github.com/jinzhu/gorm"

type Comment struct {
	gorm.Model
	Aid        uint
	Uid        uint
	Content    uint
	LikeNum    int
	CommentNum int
}
