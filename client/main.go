package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	pb "github.com/NewTanachot/learn-go-grpc/proto"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type config struct {
	Host *string
	Port *string
}

const (
	defaultId string = "9bc62ee1-2bf9-4cc7-b81d-71b3140815c0"
)

func printStructJSON(input interface{}) {
	val, _ := json.MarshalIndent(input, "", "  ")
	fmt.Println(string(val))
}

func main() {
	// Load env
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("error, can't load dotenv with an error: %v\n", err)
	}
	cfg := config{
		Host: flag.String("host", os.Getenv("HOST"), "The server host"),
		Port: flag.String("port", os.Getenv("PORT"), "The server port"),
	}
	url := fmt.Sprintf("%s:%s", *cfg.Host, *cfg.Port)

	// product id
	// productId := flag.String("product_id", defaultId, "Product id")

	flag.Parse()

	conn, err := grpc.Dial(url, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("error, failed to connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewTransferClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	r, err := client.GetProduct(ctx, &pb.Order{
		Id: "1305e1b4-bb31-4a18-9f06-261750d92beb",
	})
	if err != nil {
		log.Fatalf("could not send data with an: %v", err)
	}

	printStructJSON(r)
}
