package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "github.com/go-usermgmt-grpc/usermgmt"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func main(){
	conn , err := grpc.Dial(address,grpc.WithInsecure() , grpc.WithBlock())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	c := pb.NewUserManagementClient(conn)
	ctx , cancel := context.WithTimeout(context.Background() , time.Second)
	defer cancel()
	var newUsers = make(map[string]int32)
	newUsers["Alice"] = 43
	newUsers["Bob"] = 34
	newUsers["John"] = 12
	for name , age := range newUsers{
		r , err := c.CreateNewUser(ctx , &pb.NewUser{Name: name , Age: age})
		if err != nil {
			log.Fatal(err)
		}
		log.Printf(`user details : 
		name : %s
		age : %v
		id : %v
		` , r.GetName() , r.GetAge() , r.GetId() )
	}
	params := &pb.GetUsersParams{}
	r , err := c.GetUsers(ctx , params)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("user list")
	fmt.Printf("the users %v\n" , r.GetUsers())
}