package main

import (
	datatbases "demoLoginServer/databases"
	mod "demoLoginServer/models"
	"fmt"
)

// User has and belongs to many languages, `user_languages` is the join table

func main() {

	db, _ := datatbases.Init()
	// (&mod.User{Name: "Gary", Password: mod.Password{Plaintext: "1234"}}).Create(db)
	// (&mod.User{Name: "Gary2", Password: mod.Password{Plaintext: "1234"}}).Create(db)
	user1 := mod.User{
		Email: "gary@example.com",
		Password: mod.Password{
			Plaintext: "1234",
		},
	}

	err1 := user1.Create(db)
	if err1 != nil {
		fmt.Println("user1", err1)
	} else {
		fmt.Println("user1 created")
	}

	var user2 = user1
	err2 := user2.Create(db)
	if err2 != nil {
		fmt.Println("user2", err2)
	} else {
		fmt.Println("user2 created")
	}
	success1, _ := user1.Login(db)
	fmt.Println("user1 login:", success1)

	var user3 = user1
	user3.Password.Plaintext = "abcd"
	success3, _ := user3.Login(db)
	fmt.Println("user3 login:", success3)

}
