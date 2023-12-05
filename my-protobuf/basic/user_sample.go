package basic

import (
	"log"
	"my-protobuf/protogen/basic"
)

func BasicUser() {
	u := basic.User{
		Id:       99,
		Username: "johndoe",
		IsActive: true,
		Password: []byte("mypassword"),
	}

	log.Println(&u)
}
