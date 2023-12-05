package basic

import (
	"log"
	"math/rand"
	"my-protobuf/protogen/basic"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/anypb"
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

	comm := randomCommunicationChannel()

	u := basic.User{
		Id:                   99,
		Username:             "johndoe",
		IsActive:             true,
		Password:             []byte("mypassword"),
		Emails:               []string{"johndoe@example.com"},
		Gender:               basic.Gender_GENDER_MALE,
		Address:              &addr,
		CommunicationChannel: &comm,
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

func randomCommunicationChannel() anypb.Any {
	paperMail := basic.PaperMail{
		PaperMailAddress: "Some paper mail address",
	}

	socialMedia := basic.SocialMedia{
		SocialMediaPlatform: "byteMe",
		SocialMediaUsername: "krypton.man",
	}

	instantMessaging := basic.InstantMessaging{
		InstantMessagingProduct:  "whatsapp",
		InstantMessagingUsername: "krypton.last",
	}

	var a anypb.Any

	switch r := rand.Intn(10) % 3; r {
	case 0:
		anypb.MarshalFrom(&a, &paperMail, proto.MarshalOptions{})
	case 1:
		anypb.MarshalFrom(&a, &socialMedia, proto.MarshalOptions{})
	default:
		anypb.MarshalFrom(&a, &instantMessaging, proto.MarshalOptions{})
	}

	return a
}

func BasicUnmarshalAnyKnown() {
	socialMedia := basic.SocialMedia{
		SocialMediaPlatform: "my-social-media-platform",
		SocialMediaUsername: "my-social-media-username",
	}

	var a anypb.Any
	anypb.MarshalFrom(&a, &socialMedia, proto.MarshalOptions{})

	// known type (Social Media)
	sm := basic.SocialMedia{}

	if err := a.UnmarshalTo(&sm); err != nil {
		return
	}

	jsonBytes, _ := protojson.Marshal(&sm)
	log.Println(string(jsonBytes))
}

func BasicUnmarshalAnyNotKnown() {
	a := randomCommunicationChannel()

	var unmarshaled protoreflect.ProtoMessage

	unmarshaled, err := a.UnmarshalNew()

	if err != nil {
		return
	}

	log.Println("Unmarshal as", unmarshaled.ProtoReflect().Descriptor().Name())

	jsonBytes, _ := protojson.Marshal(unmarshaled)
	log.Println(string(jsonBytes))
}

func BasicUnmarshalAnyIs() {
	a := randomCommunicationChannel()
	pm := basic.PaperMail{}

	if a.MessageIs(&pm) {
		if err := a.UnmarshalTo(&pm); err != nil {
			log.Fatalln(err)
		}

		jsonBytes, _ := protojson.Marshal(&pm)
		log.Println(string(jsonBytes))
	} else {
		log.Println("Not PaperMail, but :", a.TypeUrl)
	}
}

func BasicOneOf() {
	socialMedia := basic.SocialMedia{
		SocialMediaPlatform: "my-social-media-platform",
		SocialMediaUsername: "my-social-media-username",
	}

	ecomm := basic.User_SocialMedia{
		SocialMedia: &socialMedia,
	}

	u := basic.User{
		Id:                    99,
		Username:              "johndoe",
		IsActive:              true,
		Password:              []byte("mypassword"),
		Emails:                []string{"johndoe@example.com"},
		Gender:                basic.Gender_GENDER_MALE,
		ElectronicCommChannel: &ecomm,
	}

	jsonBytes, _ := protojson.Marshal(&u)
	log.Println(string(jsonBytes))
}
