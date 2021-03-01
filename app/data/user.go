package data

type User struct {
	ID       int64  `db:"id" json:"id"`
	Username string `db:"username" json:"username"`
	Name     string `db:"name,omitempty" json:"name,omitempty"`
	Surname  string `db:"surname,omitempty" json:"surname,omitempty"`
	Email    string `db:"email" json:"email"`
	Password string `db:"password" json:"password"`
	// todo for better times :)
	Salt     string `db:"salt,omitempty" json:"salt,omitempty"`
}

func (u User) ToMap() map[string]interface{} {
	result := map[string]interface{}{
		"username": u.Username,
		"name":     u.Name,
		"surname":  u.Surname,
		"email":    u.Email,
		"password": u.Password,
		"salt":     u.Salt,
	}

	return result
}
