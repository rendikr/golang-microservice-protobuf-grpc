package basic

import (
	"log"
	"my-protobuf/protogen/basic"

	"google.golang.org/protobuf/encoding/protojson"
)

func BasicUserGroup() {
	john := basic.User{
		Id:       99,
		Username: "johndoe",
		IsActive: true,
		Password: []byte("mypassword"),
		Emails:   []string{"johndoe@example.com"},
		Gender:   basic.Gender_GENDER_MALE,
		Address: &basic.Address{
			Street:  "JL. ABC",
			City:    "Jakarta",
			Country: "Indonesia",
			Coordinate: &basic.Address_Coordinate{
				Latitude:  192.11441784,
				Longitude: -128758198,
			},
		},
	}

	jane := basic.User{
		Id:       98,
		Username: "janedoe",
		IsActive: true,
		Password: []byte("mypassword"),
		Emails:   []string{"janedoe@example.com"},
		Gender:   basic.Gender_GENDER_FEMALE,
		Address: &basic.Address{
			Street:  "JL. DEF",
			City:    "Jakarta",
			Country: "Indonesia",
			Coordinate: &basic.Address_Coordinate{
				Latitude:  192.11441784,
				Longitude: -128758198,
			},
		},
	}

	persons := basic.UserGroup{
		GroupId:     90,
		GroupName:   "Family",
		Users:       []*basic.User{&john, &jane},
		Description: "Happy Family",
	}

	jsonBytes, _ := protojson.Marshal(&persons)
	log.Println(&persons)
	log.Println(string(jsonBytes))
}
