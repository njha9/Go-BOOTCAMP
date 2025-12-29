package models

type user struct {
	UserID   string
	Username string
	Email    string
	Password string
}

func NewUser(userID, username, email, password string) *user {
	return &user{
		UserID:   userID,
		Username: username,
		Email:    email,
		Password: password,
	}
}

func (u *user) UpdateEmail(newEmail string) {
	u.Email = newEmail
}

func (u *user) UpdatePassword(newPassword string) {
	u.Password = newPassword
}
func (u *user) Summary() string {
	return "User " + u.Username + " with email " + u.Email
}
