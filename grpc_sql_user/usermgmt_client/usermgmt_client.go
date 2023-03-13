package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "github.com/tech-with-moss/go-usermgmt-grpc/usermgmt"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func main() {

	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewUserManagementClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	// var user_id = "sanjai"
	// var new_users1 = make(map[string]string)
	// new_users1["user_id"] = "sanjai"
	var new_users = make(map[int]string)
	new_users[3] = "sanjai@gmail.com"
	new_users[2] = "sanjai@gmail.com"
	new_users[1] = "sanjai@gmail.com"
	for user_id, email := range new_users {
		r, err := c.CreateNewUser(ctx, &pb.NewUser{UserId: int32(user_id), Email: string(email)})
		if err != nil {
			log.Fatalf("could not create user: %v", err)
		}
		log.Printf(`User Details:
user_id: %d
email: %s
IsActive: %d`, r.GetUserId(), r.GetEmail(), r.GetIsActive())

	}
	params := &pb.GetUsersParams{}
	r, err := c.GetUsers(ctx, params)
	if err != nil {
		log.Fatalf("could not create user: %v", err)
	}
	log.Print("\nUSER LIST: \n")
	fmt.Printf("r.GetUsers(): %v\n", r.GetUsers())
}
