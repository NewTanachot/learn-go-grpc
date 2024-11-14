package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	pb "mikelopster_grpc/mikelopter_grpc/proto"
)

const (
	// Port ที่ server
	port = ":50051"
)

// Struct server ใช้ในการ implement gRPC services
type server struct {
	// Embedding UnimplementedGreeterServer เพื่อให้มั่นใจว่า struct นี้ implement ทุก method ของ Greeter service
	pb.UnimplementedGreeterServer
}

// Implement SayHello method ที่จะตอบกลับคำทักทาย
// รับ Context และ HelloRequest เป็น input และส่งกลับ HelloReply หรือ error
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	// Log ชื่อที่ได้รับจาก request
	log.Printf("Received: %v", in.GetName())
	// ตอบกลับข้อความทักทาย
	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func main() {
	// สร้าง TCP listener ที่ฟังการเชื่อมต่อบน port ที่กำหนด
	lis, err := net.Listen("tcp", port)
	if err != nil {
		// ถ้าเกิดข้อผิดพลาดขณะสร้าง listener, ให้ log error แล้วออกจากโปรแกรม
		log.Fatalf("failed to listen: %v", err)
	}

	// สร้าง gRPC server instance ใหม่
	grpcServer := grpc.NewServer()

	// ลงทะเบียน Greeter service กับ server
	pb.RegisterGreeterServer(grpcServer, &server{})

	// Log ว่า server กำลังฟังการเชื่อมต่อบน port ที่กำหนด
	log.Printf("Server is listening on port %v", port)

	// เริ่มฟังการเชื่อมต่อและให้บริการ gRPC request
	if err := grpcServer.Serve(lis); err != nil {
		// ถ้าเกิดข้อผิดพลาดขณะให้บริการ, ให้ log error แล้วออกจากโปรแกรม
		log.Fatalf("failed to serve: %v", err)
	}
}
