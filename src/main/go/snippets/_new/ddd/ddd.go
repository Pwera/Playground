package ddd

import (
	"fmt"
	"github.com/pwera/Playground/src/main/go/snippets/_new/ddd/application"
	"github.com/pwera/Playground/src/main/go/snippets/_new/ddd/controller"
	"github.com/pwera/Playground/src/main/go/snippets/_new/ddd/domain"
	"github.com/pwera/Playground/src/main/go/snippets/_new/ddd/persistence/db"
	"github.com/pwera/Playground/src/main/go/snippets/_new/ddd/persistence/memory"
	"net/http"
)

func ExamppleUserRepository() {
	userRepo := memory.NewUserRepository()
	issueRepo := db.NewIssueRepository()
	userService := application.UserService{UserRepository: userRepo}
	issueService := application.IssueService{IssueRepository: issueRepo}
	userController := controller.UserController{UserService: userService}
	issueController := controller.IssueController{IssueService: issueService}
	prepareUsers(userService)
	prepareIssues(issueService)
	mux := http.NewServeMux()
	mux.HandleFunc("/api/users", userController.List)
	mux.HandleFunc("/api/issues", issueController.List)

	http.ListenAndServe(":8090", mux)

}

func prepareIssues(issueService application.IssueService) {
	issue := domain.Issue{
		Title:    "Title",
		Priority: domain.PriorityLow,
		OwnerId:  1,
		ProjectId: 1,
		Description:"??",
	}
	err := issueService.Create(&issue)
	if err != nil {
		fmt.Println(err)
	}
}

func prepareUsers(userService application.UserService) {
	for i := 0; i < 10; i += 1 {
		userService.Create(&domain.User{Name: fmt.Sprintf("User_%d", i)})
	}
}
