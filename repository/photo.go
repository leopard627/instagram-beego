package repository

import (
	"github.com/astaxie/beego/orm"
	"github.com/instagram-beego/models"
	"strconv"
)

type PhotoRepository struct{}

func (this *PhotoRepository) GetAll() ([]*models.Photo, error) {
	var photos []*models.Photo
	o := orm.NewOrm()
	_, err := o.
		QueryTable(&models.Photo{}).
		RelatedSel().
		All(&photos)

	for _, photo := range photos {
		photoIdStr := strconv.Itoa(photo.Id)

		o.
			QueryTable(&models.Comment{}).
			Filter("Photo__Id", photoIdStr).
			RelatedSel("User").
			OrderBy("-id").
			All(&photo.Comments)
	}

	return photos, err
}

func (this *PhotoRepository) GetByUserId(userId int) ([]*models.Photo, error) {
	var photos []*models.Photo
	o := orm.NewOrm()
	ps := o.QueryTable(&models.Photo{})
	_, err := ps.Filter("User", strconv.Itoa(userId)).RelatedSel().All(&photos)
	return photos, err
}
