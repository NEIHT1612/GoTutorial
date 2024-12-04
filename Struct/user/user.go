package user

import "time"

type User struct {
	FirstName string
	LastName  string
	BirthDate string
	CreateAt  time.Time
}

type Admin struct {
	Email string
	Password string
	User User
}

func NewUser(FirstName, LastName, BirthDate string) User{
	return User{
		FirstName: FirstName,
		LastName: LastName,
		BirthDate: BirthDate,
		CreateAt: time.Now(),
	}
}

func NewAdmin(Email, Password string) Admin{
	return Admin{
		Email: Email,
		Password: Password,
		User: User{
			FirstName: "ADMIN",
			LastName: "ADMIN",
			BirthDate: "16/12/2003",
			CreateAt: time.Now(),
		},
	}
}