package main

import (
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// User has and belongs to many languages, `user_languages` is the join table

func (u *User) Birth() uint {

}

func main() {

	golog := log.New(os.Stdout, "\r\n", log.LstdFlags) // io writer
	f, err := os.OpenFile("testlogfile", os.O_RDWR, 0666)
	if err != nil {
		golog.Fatalf("error opening file: %v", err)
	} else {
		golog.SetOutput(f)
		defer f.Close()
		golog.Println("Log begins")
	}

	newLogger := logger.New(
		golog,

		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      false,       // Don't include params in the SQL log
			Colorful:                  true,        // Disable color

		},
	)
	dsn := "host=172.17.0.3 user=cgg password=1234_abc dbname=fengshui port=5432 sslmode=disable TimeZone=Asia/Taipei"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
		// NamingStrategy: MyNamingStrategy{},
	})

	// db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err == nil {
		db = db.Debug()

		// db.Session(&gorm.Session{DryRun: true}).AutoMigrate(&cards, &user)
		// db.AutoMigrate(&cards, &user)
		db.Session(&gorm.Session{DryRun: true}).Create(&user)

	}

}