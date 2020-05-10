package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"
	"time"

	"google.golang.org/grpc"

	helloworld "github.com/jomoespe/grpc/helloworld"
)

const (
	address            = "localhost:50051"
	defaultName        = "world"
	defaultInvokations = 10
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(
		address,
		grpc.WithInsecure(),
		grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := helloworld.NewGreeterClient(conn)

	// Contact the server and print out its response.
	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	invokations := defaultInvokations
	if len(os.Args) > 2 {
		if invokations, err = strconv.Atoi(os.Args[2]); err != nil {
			invokations = defaultInvokations
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Invoke service in gorutines
	var wg sync.WaitGroup
	for i := 0; i < invokations; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			r, err := c.SayHello(ctx, &helloworld.HelloRequest{Name: name + " # " + strconv.Itoa(i)})
			if err != nil {
				log.Fatalf("could not greet: %v", err)
			}
			fmt.Printf("Greeting: %s\n", r.GetMessage())
		}(i)
	}
	wg.Wait()
}
