package main

import (
	bookService "Or/Library/internal/application/book"
	loanService "Or/Library/internal/application/loan"
	userService "Or/Library/internal/application/user"
	"Or/Library/internal/domain/book"
	"Or/Library/internal/domain/loan"
	"Or/Library/internal/domain/user"
	bookDb "Or/Library/internal/infrastructure/db/postgres/book"
	loanDb "Or/Library/internal/infrastructure/db/postgres/loan"
	userDb "Or/Library/internal/infrastructure/db/postgres/user"
	"Or/Library/internal/interface/api/rest"

	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Database connection setup
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	// Migrate the schema
	err = db.AutoMigrate(&bookDb.BookEntity{}, &userDb.UserEntity{}, &loanDb.LoanEntity{})
	if err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	// Dependency Injection (using a simple container)
	container := &struct {
		BookRepository book.BookRepository
		UserRepository user.UserRepository
		LoanRepository loan.LoanRepository
		BookService    *bookService.BookService
		UserService    *userService.UserService
		LoanService    *loanService.LoanService
		BookController *rest.BookController
		UserController *rest.UserController
		LoanController *rest.LoanController
	}{}

	// Initialize repositories
	container.BookRepository = bookDb.NewBookRepository(db)
	//container.UserRepository = userDb.NewUserRepository(db)
	//container.LoanRepository = loanDb.NewLoanRepository(db)

	// Initialize services (injecting dependencies)
	container.BookService = bookService.NewBookService(container.BookRepository)
	container.UserService = userService.NewUserService(container.UserRepository, container.LoanRepository)
	container.LoanService = loanService.NewLoanService(container.LoanRepository, container.BookRepository, container.UserRepository)

	// Initialize controllers (injecting dependencies)
	container.BookController = rest.NewBookController(container.BookService)
	container.UserController = rest.NewUserController(container.UserService)
	container.LoanController = rest.NewLoanController(container.LoanService)

	// Routing setup (using injected controllers)
	r := mux.NewRouter()

	// Book routes
	r.HandleFunc("/books", container.BookController.CreateBook).Methods("POST")
	r.HandleFunc("/books/{id}", container.BookController.GetBook).Methods("GET")
	r.HandleFunc("/books/{id}", container.BookController.UpdateBook).Methods("PUT")
	r.HandleFunc("/books/{id}", container.BookController.DeleteBook).Methods("DELETE")
	r.HandleFunc("/books/search", container.BookController.SearchBooks).Methods("GET").Queries("query", "{query}")

	// User routes
	r.HandleFunc("/users", container.UserController.CreateUser).Methods("POST")
	r.HandleFunc("/users/{id}", container.UserController.GetUser).Methods("GET")
	r.HandleFunc("/users/{id}", container.UserController.UpdateUser).Methods("PUT")
	r.HandleFunc("/users/{id}", container.UserController.DeleteUser).Methods("DELETE")
	r.HandleFunc("/users/{id}/loans", container.UserController.GetUserLoans).Methods("GET")

	// Loan routes
	//r.HandleFunc("/loans", container.LoanController.CreateLoan).Methods("POST")
	//r.HandleFunc("/loans/{id}", container.LoanController.GetLoan).Methods("GET")
	//r.HandleFunc("/loans/{id}/return", container.LoanController.ReturnBook).Methods("POST")

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Server listening on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
