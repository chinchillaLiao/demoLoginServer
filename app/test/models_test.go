package test

import (
	databases "demoLoginServer/databases"
	mod "demoLoginServer/models"
	"os"
	"testing"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB_URL = os.Getenv("DB_URL")
var db, err = gorm.Open(postgres.Open(DB_URL), &gorm.Config{

	Logger: databases.DebugLogger,
})

// var db2 = db.Session(&gorm.Session{DryRun: true}).Debug()

func TestUser(t *testing.T) {
	// user := User{}
	// password := Password{}
	db = db.Session(&gorm.Session{DryRun: true}).Debug()
	type Company struct {
		ID   int
		Name string
	}
	type User struct {
		gorm.Model
		Name      string
		CompanyID int     // Foreign Key
		Company   Company // referenced Model
	}
	// user := User{}
	// company := Company{}
	// users := []User{}
	// companies := []Company{}
	u := mod.User{Email: "xxx"}
	p := mod.Password{}
	queryByUserEmail := mod.User{Email: "xxx"}

	db.Model(&p).Joins("User").Select("cipher_text").Where(&queryByUserEmail).Scan(&u.Password.CipherText)
	// db.Model(&queryByUserEmail).Joins("U INNER JOIN passwords P on users.user_id = P.user_id").Select("cipher_text").Where(queryByUserEmail).Scan(&u.Password.CipherText)
	// db.Model("User").Where(db.Where(&User{Name: "Gary"})).Association("Company").Find(&company)

	// db.Table("User").Where(db.Where(&User{Name: "Gary"})).InnerJoins("Company").Find(&companies)
	// db.Table("User").Where(db.Where(&User{Name: "Gary"})).Joins("Company").Find(&companies)
	// db.Table("User").Where(db.Where(&User{Name: "Gary"})).Joins("INNER JOIN Company").Find(&companies)
	// db.Table("User").Where(db.Where(&User{Name: "Gary"})).Joins("INNER JOIN companies").Find(&companies)
	// 	2023/06/15 09:47:50 /app/demoLoginServer/app/test/models_test.go:37

	// [0.011ms] [rows:0] SELECT "User"."id","User"."name" FROM "User" Company WHERE "User"."name" = 'Gary'

	// 2023/06/15 09:47:50 /app/demoLoginServer/app/test/models_test.go:38
	// [0.005ms] [rows:0] SELECT "User"."id","User"."name" FROM "User" Company WHERE "User"."name" = 'Gary'

	// 2023/06/15 09:47:50 /app/demoLoginServer/app/test/models_test.go:39
	// [0.004ms] [rows:0] SELECT "User"."id","User"."name" FROM "User" INNER JOIN Company WHERE "User"."name" = 'Gary'

	// 2023/06/15 09:47:50 /app/demoLoginServer/app/test/models_test.go:40
	// [0.004ms] [rows:0] SELECT "User"."id","User"."name" FROM "User" INNER JOIN companies WHERE "User"."name" = 'Gary'
	// db.InnerJoins("Company").Find(&users)
	// db.InnerJoins("User").Find(&companies)
	// db.InnerJoins("Company").Find(&users)
	// db.InnerJoins("User", db.Where(&User{Name: "Gary"})).Find(&companies)
	// db.InnerJoins("Company", db.Where(&Company{Name: "Edimax"})).Find(&users)
	// db.Joins("User").Find(&companies)
	// db.Joins("Company").Find(&users)
	// db.Joins("User", db.Where(&User{Name: "Gary"})).Find(&companies)
	// db.Joins("Company", db.Where(&Company{Name: "Edimax"})).Find(&users)

	// db.Model(&s).Where("users.email = ?", s.Email).Association("Password").Find(&mod.Password{}) //.Select("user_id", "cipher_text").InnerJoins("users", "passwords").Statement
	// stmt := db.Session(&gorm.Session{DryRun: true}).Debug().Find(&mod.User{})
	// t.Log(stmt.Statement.SQL.String())

	//db.InnerJoins("User").Find(&companies)
	// SELECT
	// "companies"."id"
	// ,"companies"."name"
	// FROM "companies" User -- what the ?

	// //db.InnerJoins("Company").Find(&users)
	// SELECT "users"."id"
	// ,"users"."created_at"
	// ,"users"."updated_at"
	// ,"users"."deleted_at"
	// ,"users"."name"
	// ,"users"."company_id"
	// ,"Company"."id" AS "Company__id"
	// ,"Company"."name" AS "Company__name"
	// FROM "users"
	// INNER JOIN "companies" "Company" ON "users"."company_id" = "Company"."id"
	// WHERE "users"."deleted_at" IS NULL

	// //db.InnerJoins("User", db.Where(&User{Name: "Gary"})).Find(&companies)
	// SELECT
	// "companies"."id"
	// ,"companies"."name"
	// FROM "companies" User -- what the ?

	// //db.InnerJoins("Company", db.Where(&Company{Name: "Edimax"})).Find(&users)
	// SELECT
	// "users"."id"
	// ,"users"."created_at"
	// ,"users"."updated_at"
	// ,"users"."deleted_at"
	// ,"users"."name"
	// ,"users"."company_id"
	// ,"Company"."id" AS "Company__id"
	// ,"Company"."name" AS "Company__name"
	// FROM "users"
	// INNER JOIN "companies" "Company"
	//   ON "users"."company_id" = "Company"."id" AND "Company"."name" = 'Edimax'
	// WHERE "users"."deleted_at" IS NULL

	// //db.Joins("User").Find(&companies)
	// SELECT
	// "companies"."id"
	// ,"companies"."name"
	// FROM "companies" User -- what the ?

	// //db.Joins("Company").Find(&users)
	// SELECT
	// "users"."id"
	// ,"users"."created_at"
	// ,"users"."updated_at"
	// ,"users"."deleted_at"
	// ,"users"."name"
	// ,"users"."company_id"
	// ,"Company"."id" AS "Company__id"
	// ,"Company"."name" AS "Company__name"
	// FROM "users"
	// LEFT JOIN "companies" "Company" -- 還會自動幫我變 Left Join
	//   ON "users"."company_id" = "Company"."id"
	// WHERE "users"."deleted_at" IS NULL

	// //db.Joins("User", db.Where(&User{Name: "Gary"})).Find(&companies)
	// SELECT
	// "companies"."id"
	// ,"companies"."name"
	// FROM "companies" User -- what the ?

	// //db.Joins("Company", db.Where(&Company{Name: "Edimax"})).Find(&users)
	// SELECT
	// "users"."id"
	// ,"users"."created_at"
	// ,"users"."updated_at"
	// ,"users"."deleted_at"
	// ,"users"."name"
	// ,"users"."company_id"
	// ,"Company"."id" AS "Company__id"
	// ,"Company"."name" AS "Company__name"
	// FROM "users"
	// LEFT JOIN "companies" "Company" -- 還會自動幫我變 Left Join
	//
	//	ON "users"."company_id" = "Company"."id" AND "Company"."name" = 'Edimax'
	//
	// WHERE "users"."deleted_at" IS NULL
}
