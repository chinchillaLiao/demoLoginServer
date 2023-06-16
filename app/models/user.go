package models

import (
	"fmt"
	"time"

	"github.com/btcsuite/btcutil/base58"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	UserID    string `gorm:"size:22;primaryKey"`
	Email     string `gorm:"size:100;unique"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Password  Password `gorm:"foreignKey:UserID;references:UserID"`
}

func (u *User) Create(db *gorm.DB) error {
	u.Password.Encrypt()
	uu := uuid.New()
	u.UserID = base58.Encode(uu[:])
	// .Session(&gorm.Session{DryRun: true})
	return db.Create(u).Error
}

func (u *User) Login(db *gorm.DB) (success bool, err error) {
	// 我覺得GORM在查詢這塊，並沒有辦法完全做到ORM，例如欄位名稱必須硬寫資料庫內的欄位名稱，那不如直接寫SQL吧。
	tx := db.Table("users U").Joins("inner join passwords P on U.user_id = P.user_id").Where("U.email = ?", u.Email).Select("cipher_text").Scan(&u.Password.CipherText)
	// tx := db.Raw("SELECT cipher_text FROM users U, passwords P WHERE U.user_id = P.user_id and U.email = ?", u.Email).Scan(&u.Password.CipherText)
	fmt.Println(u)
	if tx.Error == nil {
		success = u.Password.Match()
	}
	// success bool default: false
	return success, tx.Error
}
