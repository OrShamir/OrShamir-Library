package user

import (
	"Or/Library/internal/domain/loan"
	"Or/Library/internal/domain/user"
	"context"
	"errors"
)

type UserService struct {
	userRepository user.UserRepository
	loanRepository loan.LoanRepository
}

func NewUserService(userRepository user.UserRepository, loanRepository loan.LoanRepository) *UserService {
	return &UserService{userRepository, loanRepository}
}

func (s *UserService) CreateUser(ctx context.Context, user *user.User) error {
	// ... validations (e.g., name, email, password strength, unique email)

	return s.userRepository.Create(ctx, user)
}

func (s *UserService) GetUser(ctx context.Context, id string) (*user.User, error) {
	return s.userRepository.GetByID(ctx, id)
}

func (s *UserService) UpdateUser(ctx context.Context, user *user.User) error {
	// ... validations

	return s.userRepository.Update(ctx, user)
}

func (s *UserService) DeleteUser(ctx context.Context, id string) error {
	// Check if the user has any active loans
	loans, err := s.loanRepository.GetByUser(ctx, id)
	if err != nil {
		return err
	}
	if len(loans) > 0 {
		return errors.New("cannot delete user with active loans")
	}

	return s.userRepository.Delete(ctx, id)
}

func (s *UserService) GetAllUsers(ctx context.Context) ([]*user.User, error) {
	return s.userRepository.GetAll(ctx)
}

func (s *UserService) GetUserLoans(ctx context.Context, userID string) ([]loan.Loan, error) {
	loans, err := s.loanRepository.GetByUser(ctx, userID)
	if err != nil {
		return nil, err
	}

	return loans, nil
}
