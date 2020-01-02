package main

import (
	"fmt"
	"github.com/pwera/Playground/src/main/go/snippets/_new/ddd/application"
	"github.com/pwera/Playground/src/main/go/snippets/_new/ddd/controller"
	"github.com/pwera/Playground/src/main/go/snippets/_new/ddd/domain"
	"github.com/pwera/Playground/src/main/go/snippets/_new/ddd/persistence/memory"
	"github.com/pwera/Playground/src/main/go/snippets/_new/ddd/protocol/protocol"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
	"time"
)

type server struct{}

func (s *server) SubmitNewUser(context context.Context, ur *protocol.NewUserRequest) (*protocol.NewUserResponse, error) {
	fmt.Println("OK")
	return &protocol.NewUserResponse{Status: len(ur.Email)>0}, nil
}
func main() {
	userRepo := memory.NewUserRepository()
	userService := application.UserService{
		UserRepository: userRepo,
	}
	userController := controller.UserController{
		UserService: userService,
	}
	for i := 0; i < 10; i += 1 {
		userService.Create(&domain.User{Name: fmt.Sprintf("User_%d", i)})
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	mux.HandleFunc("/", userController.List)

	opts := []grpc.ServerOption{}
	grpcServer := grpc.NewServer(opts...)
	s := server{}
	protocol.RegisterUserServer(grpcServer, &s)
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 10000))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	go func() {
		grpcServer.Serve(lis)
	}()
	server := &http.Server{
		Addr:           ":8091",
		Handler:        mux,
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   10 * time.Second,
		IdleTimeout:    120 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Fatal(server.ListenAndServe())
}
