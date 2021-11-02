package main

import (
	"fmt"

	trippb "coolcar/proto/gen/go"

	"google.golang.org/protobuf/proto"
)

func main(){
	trip := trippb.Trip{
		Start: "abc",
		End: "def",
		DurationSec: 3600,
		FeeCent: 10000,
	}
	fmt.Println(&trip)
	b, err := proto.Marshal(&trip)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%X\n", b)

	var trip2 trippb.Trip
	err = proto.Unmarshal(b,&trip2)
	if err != nil {
		panic(err)
	}
	fmt.Println(&trip2)
}