package client

import (
	"container/ring"
	"context"
	"fmt"
	pb "github.com/rhizomata/go-concurrency/grpc/proto"
	"google.golang.org/grpc"
	"log"
	"sync"
	"time"
)

type TestDriver struct {

}

func (driver *TestDriver) SendHelloTest(serverURLs []string, totalCount int, concurrent int) {
	fmt.Println("serverURLs=",serverURLs,"totalCount=", totalCount,",concurrent=",concurrent )
	
	aring := ring.New(len(serverURLs))
	for _, u := range serverURLs{
		aring.Next().Value = u
	}
	
	for row := 0; row< totalCount ; row++ {
		wg := sync.WaitGroup{}
		for col := 0; col < concurrent; col++ {
			name := fmt.Sprintf("%d-%d", row, col)
			url := aring.Next().Value.(string)
			wg.Add(1)
			go func(url string, name string) {
				defer wg.Done()
				sendHelloImpl(url,name)
			}(url, name)
		}
		wg.Wait()
	}
}

func sendHelloImpl(serverURL string, name string){
	conn, err := grpc.Dial(serverURL , grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)
	
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())
}