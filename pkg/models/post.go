package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	ID       uint32 `gorm:"primary_key;auto_increment"`
	Title    string `gorm:"not null"`
	Content  string `gorm:"not null"`
	Author   User
	AuthorID uint32 `gorm:"not null"`
}

type PostRepository struct {
	db *gorm.DB
}

func (pr *PostRepository) FindAll() (*[]Post, error) {
	posts := []Post{}

	if err := pr.db.Debug().
		Model(&[]Post{}).
		Limit(100).
		Preload("Author").
		Find(&posts).
		Error; err != nil {
		return &[]Post{}, err
	}

	return &posts, nil
}

func (pr *PostRepository) FindOne(id uint32) (*Post, error) {
	post := Post{}

	var err error
	if err = pr.db.Debug().
		Model(&Post{}).
		Where("id = ?", id).
		Preload("Author").
		Take(&post).
		Error; err != nil {
		return &Post{}, err
	}

	return &post, nil
}

func (pr *PostRepository) DeleteOne(id uint32) (int64, error) {
	db := pr.db.Debug().
		Model(&Post{}).
		Where("id = ?", id).
		Take(&Post{}).
		Delete(&Post{})

	if db.Error != nil {
		return 0, db.Error
	}

	return db.RowsAffected, nil
}

func (pr *PostRepository) Save(data *Post) error {
	// use db transaction
	trx := pr.db.Begin()

	if err := trx.Debug().
		Create(&data).
		Error; err != nil {
		trx.Rollback()
		return err
	}

	if err := trx.Debug().
		Where(data.ID).
		Preload("Author").
		First(&data).
		Error; err != nil {
		trx.Rollback()
		return err
	}

	return trx.Commit().Error
}

func (pr *PostRepository) UpdateOne(id uint32, data *Post) error {
	// use db transaction
	trx := pr.db.Begin()

	if err := trx.Debug().
		Model(data).
		Updates(&data).
		Error; err != nil {
		trx.Rollback()
		return err
	}
	if err := trx.Debug().
		Where(data.ID).
		Preload("Author").
		First(&data).
		Error; err != nil {
		trx.Rollback()
		return err
	}

	return trx.Commit().Error
}
