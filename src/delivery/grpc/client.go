package main

import (
	"fmt"
	"github.com/kkmmttdd/question-api/src/delivery/grpc/question_pdb"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
)

func main() {
	fmt.Println("Calculator Client")
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer cc.Close()

	c := question_pdb.NewQuestionServiceClient(cc)
	connectUnary(c)
}

func connectUnary(c question_pdb.QuestionServiceClient) {
	fmt.Println("Starting to do a Sum Unary RPC...")
	req := &question_pdb.QuestionListRequest{
		QuestionVersion: 1,
	}
	res, err := c.ListQuestion(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling Sum RPC: %v", err)
	}
	log.Printf("Response from Sum: %v", res.Questions[0].QuestionText)
}
