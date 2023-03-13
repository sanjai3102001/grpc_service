package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "github.com/tech-with-moss/go-servicemgmt-grpc/servicemgmt"
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
	c := pb.NewserviceManagementClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	// var service_id = "sanjai"
	// var new_services1 = make(map[string]string)
	// new_services1["service_id"] = "sanjai"
	var new_services = make(map[int]string)
	new_services[3] = "sanjai@gmail.com"
	new_services[2] = "sanjai@gmail.com"
	new_services[1] = "sanjai@gmail.com"
	for service_id, email := range new_services {
		r, err := c.CreateNewservice(ctx, &pb.Newservice{serviceId: int32(service_id), Email: string(email)})
		if err != nil {
			log.Fatalf("could not create service: %v", err)
		}
		log.Printf(`service Details:
service_id: %d
email: %s
IsActive: %d`, r.GetserviceId(), r.GetEmail(), r.GetIsActive())

	}
	params := &pb.GetservicesParams{}
	r, err := c.Getservices(ctx, params)
	if err != nil {
		log.Fatalf("could not create service: %v", err)
	}
	log.Print("\nservice LIST: \n")
	fmt.Printf("r.Getservices(): %v\n", r.Getservices())
}
