package basic

import (
	"log"
	"my-protobuf/protogen/basic"

	"google.golang.org/protobuf/encoding/protojson"
)

func BasicUser() {
	u := basic.User{
		Id:       99,
		Username: "johndoe",
		IsActive: true,
		Password: []byte("mypassword"),
		Emails:   []string{"johndoe@example.com"},
		Gender:   basic.Gender_GENDER_MALE,
	}

	jsonBytes, _ := protojson.Marshal(&u)
	log.Println(&u)
	log.Println(string(jsonBytes))
}

func ProtoToJsonUser() {
	u := basic.User{
		Id:       98,
		Username: "janedoe",
		IsActive: true,
		Password: []byte("mypassword"),
		Emails:   []string{"janedoe@example.com"},
		Gender:   basic.Gender_GENDER_FEMALE,
	}

	jsonBytes, _ := protojson.Marshal(&u)
	log.Println(string(jsonBytes))
}

func JsonToProtoUser() {
	json := `
	{
		"id":97,
		"username":"peter",
		"is_active":true,
		"password":"YmF0bWFucGFzc3dvcnQ=",
		"emails":[
			"peter@example.com"
		],
		"gender":"GENDER_MALE"
	}
	`

	var p basic.User

	err := protojson.Unmarshal([]byte(json), &p)

	if err != nil {
		log.Fatalln("Error:", err)
	}

	log.Println(&p)
}
