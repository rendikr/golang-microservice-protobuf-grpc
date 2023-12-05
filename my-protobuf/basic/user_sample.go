package basic

import (
	"log"
	"my-protobuf/protogen/basic"

	"google.golang.org/protobuf/encoding/protojson"
)

func BasicUser() {
	addr := basic.Address{
		Street:  "Jl. ABC",
		City:    "Jakarta",
		Country: "Indonesia",
		Coordinate: &basic.Address_Coordinate{
			Latitude:  192.11441784,
			Longitude: -128758198,
		},
	}

	u := basic.User{
		Id:       99,
		Username: "johndoe",
		IsActive: true,
		Password: []byte("mypassword"),
		Emails:   []string{"johndoe@example.com"},
		Gender:   basic.Gender_GENDER_MALE,
		Address:  &addr,
	}

	jsonBytes, _ := protojson.Marshal(&u)
	log.Println(&u)
	log.Println(string(jsonBytes))
}

func ProtoToJsonUser() {
	addr := basic.Address{
		Street:  "Jl. ABC",
		City:    "Jakarta",
		Country: "Indonesia",
	}

	u := basic.User{
		Id:       98,
		Username: "janedoe",
		IsActive: true,
		Password: []byte("mypassword"),
		Emails:   []string{"janedoe@example.com"},
		Gender:   basic.Gender_GENDER_FEMALE,
		Address:  &addr,
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
		"gender":"GENDER_MALE",
		"address": {
			"street": "JL. ABC",
			"city": "Jakarta",
			"country": "Indonesia"
		}
	}
	`

	var p basic.User

	err := protojson.Unmarshal([]byte(json), &p)

	if err != nil {
		log.Fatalln("Error:", err)
	}

	log.Println(&p)
}
