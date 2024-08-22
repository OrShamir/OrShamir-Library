package user

type User struct {
	ID       string
	Name     string
	Email    string
	Role     string // "customer" or "employee"
	Password string // Consider hashing in a real application
}
