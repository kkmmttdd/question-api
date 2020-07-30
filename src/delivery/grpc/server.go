package main

import (
	"fmt"
	"github.com/kkmmttdd/question-api/src/delivery/grpc/question_pdb"
	"github.com/kkmmttdd/question-api/src/usecase"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

type server struct {}

func (*server) ListQuestion(ctx context.Context, req *question_pdb.QuestionListRequest) (*question_pdb.QuestionListResponse, error) {
	fmt.Printf("Received Sum RPC: %v\n", req)
	question_uc, _ := question_usecase.NewQuestionUseCase()
	questions, _ := question_uc.FetchAll()
	res := question_pdb.QuestionListResponse{
		Questions: []*question_pdb.Question{
			&question_pdb.Question{QuestionText: questions[0].QuestionText},
		},
	}
	return &res, nil
}

func main() {
	fmt.Println("Server Start")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	question_pdb.RegisterQuestionServiceServer(s, &server{})

	// Register reflection service on gRPC server.
	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
