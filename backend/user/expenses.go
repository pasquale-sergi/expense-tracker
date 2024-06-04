package user

import (
	"database/sql"
	"fmt"
	"time"
)

type ExpenseRequest struct {
	UserID      int       `json:"userid"`
	CategoryID  int       `json:"categoryid"`
	Amount      float64   `json:"amount"`
	Description string    `json:"description"`
	Date        time.Time `json:"date"`
}

type Expense struct {
	ExpenseID   int
	UserID      int
	Amount      float64
	CategoryID  int
	Description string
	Date        time.Time
	CreatedAt   time.Time
}

func AddExpense(db *sql.DB) error {
	query := `INSERT INTO expenses(userid, categoryid, amount,description, date, createdat) VALUES ($1,$2,$3,$4,$5,$6)`
	var req ExpenseRequest
	_, err := db.Exec(query, req.UserID, req.CategoryID, req.Amount, req.Description, req.Date)
	return err

}

func ShowExpensesOfTheMonth(db *sql.DB, userID int) ([]Expense, error) {
	query := `SELECT * FROM expenses WHERE userid = $1 AND date <= $2 AND date >= $3`
	currentDate := time.Now()

	firstDayOfMonth := time.Date(currentDate.Year(), currentDate.Month(), 1, 0, 0, 0, 0, time.UTC)
	lastDayOfMonth := firstDayOfMonth.AddDate(0, 1, -1)

	rows, err := db.Query(query, userID, lastDayOfMonth, firstDayOfMonth)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var expenses []Expense
	for rows.Next() {
		var expense Expense

		if err := rows.Scan(&expense.ExpenseID, &expense.UserID, &expense.CategoryID, &expense.Amount, &expense.Description, &expense.Date, &expense.CreatedAt); err != nil {
			return nil, fmt.Errorf("error scanning row: %v", err)
		}
		expenses = append(expenses, expense)
	}

	return expenses, nil

}

func ShowExpensesHistory(db *sql.DB, userID int) ([]Expense, error) {
	query := `SELECT * FROM expenses WHERE userid = $1`

	rows, err := db.Query(query, userID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var expenses []Expense
	for rows.Next() {
		var expense Expense
		if err := rows.Scan(&expense.ExpenseID, &expense.UserID, &expense.CategoryID, &expense.Amount, &expense.Description, &expense.Date, &expense.CreatedAt); err != nil {
			fmt.Printf("error scanning row: %v ", err)
		}
		expenses = append(expenses, expense)
	}
	return expenses, nil
}

func GetExpensesByUserID(db *sql.DB, userID int) ([]Expense, error) {
	query := "SELECT expenseid, amount, categoryid, description, date FROM expenses WHERE userID = $1"
	rows, err := db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var expenses []Expense
	for rows.Next() {
		var expense Expense
		if err := rows.Scan(&expense.ExpenseID, &expense.Amount, &expense.CategoryID, &expense.Description, &expense.Date); err != nil {
			return nil, err
		}
		expenses = append(expenses, expense)
	}
	return expenses, nil
}
