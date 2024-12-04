package models

import (
	"errors"

	"example.com/main/db"
	"example.com/main/utils"
)

type User struct {
	ID       int64  `json:"id"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (u *User) CreateUser() error{
	query := `INSERT INTO users(email, password) 
			  VALUES (?, ?)`
	stmt, err := db.DB.Prepare(query)
	if err != nil{
		return err 
	}
	defer stmt.Close()
	
	//Hash password
	hashedPassword, err := utils.HashPassword(u.Password)
	
	result, err := stmt.Exec(u.Email, hashedPassword)
	if err != nil{
		return err
	}
	userId, err := result.LastInsertId()
	if err != nil{
		return err
	}
	u.ID = userId
	return err
}

func (u *User) ValidateCredentials() error{
	query := `SELECT id, password FROM users WHERE email = ?`
	row := db.DB.QueryRow(query, u.Email)

	var retrievedPassword string
	err := row.Scan(&u.ID, &retrievedPassword)

	if err != nil{
		return errors.New("Credentials invalid")
	}

	passwordIsValid := utils.CheckPasswordHash(u.Password, retrievedPassword)

	if !passwordIsValid{
		return errors.New("Credentials invalid")
	}
	return nil
}