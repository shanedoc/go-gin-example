package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Article struct {
	Model

	//tag信息 设置外键
	TagId int `json:"tag_id" gorm:"index`
	//嵌套结构体
	Tag Tag `json:"tag"`

	Title      string `json:"title"`
	Desc       string `json:"desc"`
	Content    string `json:"content"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
}

//sql更新操作时更新时间戳字段
func (art *Article) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedOn", time.Now().Unix())
	return nil
}

func (art *Article) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("ModifiedOn", time.Now().Unix())
	return nil
}

func GetArticles(pageNum, pageSize int, maps interface{}) (art []Article) {
	//preload:预加载器 执行两条sql
	//分别是SELECT * FROM blog_articles;
	//和SELECT * FROM blog_tag WHERE id IN (1,2,3,4);，
	//那么在查询出结构后，gorm内部处理对应的映射逻辑，
	//将其填充到Article的Tag中，会特别方便，并且避免了循环查询
	db.Preload("Tag").Where(maps).Offset(pageNum).Limit(pageSize).Find(&art)
	return
}

func GetArticleTotal(maps interface{}) (count int) {
	db.Model(&Article{}).Where(maps).Count(&count)
	return
}

func GetArticle(id int) (art Article) {
	db.Where("id=?", id).First(&art)
	//关联标签信息
	db.Model(&art).Related(&art.Tag)
	return
}

func AddArticle(data map[string]interface{}) bool {
	db.Create(&Article{
		TagId:     data["tag_id"].(int), //类型断言,判断结果值是否为某个类型
		Title:     data["title"].(string),
		Desc:      data["desc"].(string),
		Content:   data["content"].(string),
		CreatedBy: data["created_by"].(string),
		State:     data["state"].(int),
	})
	return true
}

func EditArticle(id int, data interface{}) bool {
	db.Model(&Article{}).Where("id=?", id).Updates(data)
	return true
}

func DeleteArticle(id int) bool {
	//TODO::删除不需要指针
	db.Where("id=?", id).Delete(Article{})
	return true
}

func ExistArticleById(id int) bool {
	var article Article
	db.Where("id=?", id).First(&article)
	if article.ID > 0 {
		return true
	}
	return false
}

func ExistArticleByTitle(title string) bool {
	var article Article
	db.Where("title=?", title).First(&article)
	if article.ID > 0 {
		return true
	}
	return false
}
