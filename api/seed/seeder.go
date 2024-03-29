package seed

import (
	"log"

	"github.com/jinzhu/gorm"
	"github.com/sundjatak/fullstack/api/models"
)

var users = []models.User{
	models.User{
		Nickname: "Kourouma Leo",
		Email:    "lkourouma@telkea.com",
		Password: "password",
	},
	models.User{
		Nickname: "DeCarpenterie Leo",
		Email:    "leo.kourouma@gmail.com",
		Password: "password",
	},
}
var posts = []models.Post{

	models.Post{
		Title:   "Title 1",
		Content: "Hello world 1",
	},
	models.Post{
		Title:   "Title 2",
		Content: "Hello world 2",
	},
}

func Load(db *gorm.DB) {

	err := db.Debug().DropTableIfExists(&models.Post{}, &models.User{}).Error

	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}

	err = db.Debug().AutoMigrate(&models.User{}, &models.Post{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	for i, _ := range users {
		err = db.Debug().Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}
		posts[i].AuthorID = users[i].ID

		err = db.Debug().Model(&models.Post{}).Create(&posts[i]).Error
		if err != nil {
			log.Fatalf("cannot seed posts table: %v", err)
		}
	}
}
