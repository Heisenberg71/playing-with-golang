package main

import (
	"fmt"
	"hello/examplepb"
	"log"
	"os"

	"google.golang.org/protobuf/proto"
)

func main() {
	// Create an object
	person := &examplepb.Person{
		Name: "Alice",
		Age:  25,
	}

	// Serialize
	data, err := proto.Marshal(person)
	if err != nil {
		log.Fatalf("Marshal error: %v", err)
	}

	fmt.Println("Serialized data: %s", data)
	os.WriteFile("person.bin", data, 0644)

	// Deserialize
	rawData, _ := os.ReadFile("person.bin")
	newPerson := &examplepb.Person{}
	err = proto.Unmarshal(rawData, newPerson)
	if err != nil {
		log.Fatalf("Unmarshal error: %v", err)
	}

	fmt.Printf("Name: %s, Age: %d\n", newPerson.Name, newPerson.Age)
}
