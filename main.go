package main

import (
	"bytes"
	"fmt"
	"golang-protobuf/model"
	"os"
	"strings"

	"github.com/golang/protobuf/jsonpb"
	"google.golang.org/protobuf/runtime/protoiface"
)

func main() {
	user1 := &model.User{
		Id:       "u001",
		Name:     "Ante Rebic",
		Password: "f0rZ4 m1l4N",
		Gender:   model.UserGender_MALE,
	}

	userList := &model.UserList{
		List: []*model.User{
			user1,
		},
	}

	garage1 := &model.Garage{
		Id:   "g001",
		Name: "Kalimdor",
		Coordinate: &model.GarageCoordinate{
			Latitude:  23.2212847,
			Longitude: 53.22033123,
		},
	}

	garageList := &model.GarageList{
		List: []*model.Garage{
			garage1,
		},
	}

	// garageListByUser := &model.GarageListByUser{
	// 	List: map[string]*model.GarageList{
	// 		user1.Id: garageList,
	// 	},
	// }

	// ==== Original
	fmt.Printf("# ==== Original \n		%#v \n", user1)

	// ==== As string
	fmt.Printf("# ==== As String \n		%v \n", user1.String())

	// ==== as json string
	var buf bytes.Buffer
	err1 := protoToJSONString(&buf, garageList)
	if err1 != nil {
		fmt.Println(err1.Error())
		os.Exit(0)
	}

	jsonString := buf.String()
	fmt.Printf("# ==== As JSON String\n		%v \n", jsonString)

	// ==== json string to proto object using jsonpb.Unmarshaler
	buf2 := strings.NewReader(jsonString)
	protoObject := new(model.GarageList)

	err2 := (&jsonpb.Unmarshaler{}).Unmarshal(buf2, protoObject)
	if err2 != nil {
		fmt.Println(err2.Error())
		os.Exit(0)
	}

	fmt.Printf("# ==== Json string to proto as string\n		%v \n", protoObject.String())

	// ==== json string to proto object using jsonpb.UnmarshalString
	var buf3 bytes.Buffer
	err3 := protoToJSONString(&buf3, userList)
	if err3 != nil {
		fmt.Println(err3.Error())
		os.Exit(0)
	}

	jsonString3 := buf3.String()
	fmt.Printf("==== from a json string\n		%v \n", jsonString3)

	protoObject2 := new(model.UserList)
	err4 := jsonpb.UnmarshalString(jsonString3, protoObject2)
	if err4 != nil {
		fmt.Println(err4.Error())
		os.Exit(0)
	}

	fmt.Printf("==== to a proto object\n		%v \n", protoObject2.String())
}

func protoToJSONString(buf *bytes.Buffer, protoObject protoiface.MessageV1) error {
	err := (&jsonpb.Marshaler{}).Marshal(buf, protoObject)
	if err != nil {
		return err
	}

	return nil
}
