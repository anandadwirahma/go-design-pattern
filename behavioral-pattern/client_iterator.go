package main

import (
	"fmt"
	"go-design-pattern/behavioral-pattern/iterator"
)

func CallIteratorMethod() {

	user1 := &iterator.User{
		Name: "Steve Job",
		Age:  25,
	}

	user2 := &iterator.User{
		Name: "Larry Page",
		Age:  25,
	}

	userCollection := iterator.UserCollection{Users: []*iterator.User{user1, user2}}
	iterator := userCollection.CreateIterator()

	for iterator.HasNext() {
		user := iterator.GetNext()
		fmt.Println(user)
	}

}
