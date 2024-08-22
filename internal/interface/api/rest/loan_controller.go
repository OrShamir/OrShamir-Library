package rest

import (
	"Or/Library/internal/application/loan"
	"net/http"
)

type LoanController struct {
	loanService *loan.LoanService
}

func NewLoanController(loanService *loan.LoanService) *LoanController {
	return &LoanController{loanService}
}

func (c *LoanController) CreateLoan(w http.ResponseWriter /* r *http.Request */) {
	/*	var loanDTO loan.LoanDTO
		err := json.NewDecoder(r.Body).Decode(&loanDTO)
		if err != nil {
			http.Error(w, "invalid request body", http.StatusBadRequest)
			return
		}

		l := // mapToLoanDomainFromDTO(&loanDTO)
		err = c..CreateLoan(r.Context(), l)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	*/
	w.WriteHeader(http.StatusCreated)
}

// func (c *LoanController) GetLoan(w http.ResponseWriter, r *http.Request) { ...}

// func (c *LoanController) UpdateLoan(w http.ResponseWriter, r *http.Request) { ...}

// func (c *LoanController) DeleteLoan(w http.ResponseWriter, r *http.Request) { ...}

// func (c *LoanController) SearchLoan(w http.ResponseWriter, r *http.Request) { ...}

// func (c *LoanController) ReturnLoan(w http.ResponseWriter, r *http.Request) { ...}

// func (c *LoanController) ExtendLoan(w http.ResponseWriter, r *http.Request) { ...}

// func (c *LoanController) GetLoansByUser(w http.ResponseWriter, r *http.Request) { ...}

// func (c *LoanController) GetLoansBy(w http.ResponseWriter, r *http.Request) { ...}
