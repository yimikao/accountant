package database

import (
	"encoding/json"
	"errors"
	"os"
	"strings"
)

type User struct {
	Username     string        `json:"username"`
	Balance      float64       `json:"balance"`
	Transactions []Transaction `json:"transactions"`
}

type Transaction struct {
	Amount  float64 `json:"amount"`
	Type    string  `json:"type"`
	Details string  `json:"details"`
}

func getUsers() ([]User, error) {
	data, err := os.ReadFile("database/db.json")
	var users []User

	if err == nil {
		if err := json.Unmarshal(data, &users); err != nil {
			return users, err
		}
		return users, nil
	}
	return users, err
}

func updateDB(users []User) {
	data, err := json.Marshal(users)
	if err != nil {
		panic(err)
	}
	err = os.WriteFile("database/db.json", data, 0644)
	if err != nil {
		panic(err)
	}
}

func CreateUser(username string) {
	var user User
	user.Username = username
	user.Balance = 0
	user.Transactions = []Transaction{}

	users, err := getUsers()
	if err != nil {
		panic(err)
	}
	users = append(users, user)
	updateDB(users)
}

func FindUser(username string) (*User, error) {
	users, err := getUsers()
	for _, u := range users {
		if strings.EqualFold(u.Username, username) {
			return &u, nil
		}
		return nil, errors.New("user not found")
	}
	return nil, err
}

func AllUsers() ([]User, error) {
	users, err := getUsers()
	if err != nil {
		return users, err
	}

	return users, nil
}

func (u *User) UpdateAmount(amount float64, details string) {
	tx := Transaction{
		Type:    ttype,
		Details: details,
		Amount:  amount,
	}
	u.Transactions = append(u.Transactions, tx)

	newuser := User{
		Username:     u.Username,
		Transactions: u.Transactions,
	}

	newuser.Balance = u.Balance + amount

	users, _ := getUsers()
	for i, usr := range users {
		if usr.Username == u.Username {
			users[i] = newuser
		}
	}
	updateDB(users)
}

// func RemoveUser(username string) {

// }
