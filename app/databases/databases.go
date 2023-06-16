package datatbases

import (
	mod "demoLoginServer/models"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() (db *gorm.DB, err error) {
	golog := log.New(os.Stdout, "\r\n", log.LstdFlags) // io writer
	f, err := os.Create("testlogfile.log")
	if err != nil {
		golog.Fatalf("error opening file: %v", err)
	} else {
		golog.SetOutput(f)
		defer f.Close()
		golog.Println("Log begins")
	}

	DB_URL := os.Getenv("DB_URL")
	db, err = gorm.Open(postgres.Open(DB_URL), &gorm.Config{
		Logger: EnvLogger(),
	})

	// AutoMigrate
	if err == nil {
		db = db.Debug()
		var user mod.User
		var password mod.Password
		db.AutoMigrate(&user, &password)

	}
	return db, err
}
