package datatbases

import (
	mod "demoLoginServer/models"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
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

	newLogger := logger.New(
		// golog,
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      false,       // Don't include params in the SQL log
			Colorful:                  false,       // Disable color

		},
	)
	DB_URL := os.Getenv("DB_URL")
	db, err = gorm.Open(postgres.Open(DB_URL), &gorm.Config{
		Logger: newLogger,
	})

	if err == nil {
		db = db.Debug()
		var user mod.User
		var password mod.Password
		db.AutoMigrate(&user, &password)

	}
	return db, err
}
