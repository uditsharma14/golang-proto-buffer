package main

import (
	"fmt"
	"reflect"

	"google.golang.org/protobuf/proto"

	pb "protobufexample.com/m/proto"
)

func doSimple() *pb.Simple {
	return &pb.Simple{
		Id:         32,
		IsSimple:   true,
		Name:       "Udit Sharma",
		SampleList: []int32{1, 2, 3, 4},
	}
}

func doComplex() *pb.Complex {
	return &pb.Complex{
		OneDummy: &pb.DummyType{Id: 42, Name: "Udit Sharma"},
		MultipleDummies: []*pb.DummyType{
			{Id: 42, Name: "Udit Sharma"},
			{Id: 43, Name: "Jatin Saxena"},
			{Id: 44, Name: "Bijal shah"},
		},
	}
}

func doEnum() *pb.Enumeration {
	return &pb.Enumeration{
		EyeColor: pb.EyeColor_EYE_COLOR_BLUE,
	}
}

func oneOf(message interface{}) {
	switch x := message.(type) {
	case *pb.Result_Id:
		fmt.Println(message.(*pb.Result_Id).Id)
	case *pb.Result_Message:
		fmt.Println(message.(*pb.Result_Message).Message)

	default:
		fmt.Errorf("message has unexpected type : %v", x)

	}
}

func doMap() *pb.MapExample {
	return &pb.MapExample{
		Ids: map[string]*pb.IdWrapper{
			"myid":  {Id: 42},
			"myid2": {Id: 42},
			"myid3": {Id: 42},
			"myid4": {Id: 42},
		},
	}
}

func doFile(p proto.Message) {
	path := "simple.text.file"
	writeToFile(path, p)
	message := &pb.Simple{}
	readFromFile(path, message)
	fmt.Println(message)
}

func doToJSON(p proto.Message) string {
	jsonString := toJSON(p)
	fmt.Println(jsonString)
	return jsonString
}

func dofromJSON(jsonString string, t reflect.Type) proto.Message {
	message := reflect.New(t).Interface().(proto.Message)
	fromJSON(jsonString, message)
	return message
}

func main() {
	//fmt.Println(doSimple())
	//fmt.Println(doComplex())
	//ft.Println(doEnum())
	//doFile(doSimple())
	//oneOf(&pb.Result_Id{Id: 42})
	//oneOf(&pb.Result_Message{Message: "Udit Sharma"})
	//fmt.Println(doMap())

	jsonString := doToJSON(doSimple())
	message := dofromJSON(jsonString, reflect.TypeOf(pb.Simple{}))
	fmt.Println(message)

	jsonString2 := doToJSON(doComplex())
	message2 := dofromJSON(jsonString2, reflect.TypeOf(pb.Complex{}))
	fmt.Println(message2)

	fmt.Println(dofromJSON(`{
		"id":  32,
		"isSimple":  true,
		"name":  "Goga Pandit",
		"sampleList":  [
		  1,
		  2,
		  3,
		  4
		]
	  }`, reflect.TypeOf(pb.Simple{})))
}
