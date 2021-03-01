package data

type User struct {
	ID       int64  `db:"id" json:"id"`
	Username string `db:"username" json:"username"`
	Name     string `db:"name,omitempty" json:"name,omitempty"`
	Surname  string `db:"surname,omitempty" json:"surname,omitempty"`
	Email    string `db:"email" json:"email"`
	ImageURL string `db:"image_url" json:"image_url"`
	Password string `db:"password" json:"password"`
	Salt     string `db:"salt,omitempty" json:"salt,omitempty"`
}

func (u User) ToMap() map[string]interface{} {
	result := map[string]interface{}{
		"username":  u.Username,
		"name":      u.Name,
		"surname":   u.Surname,
		"email":     u.Email,
		"image_url": u.ImageURL,
		"password":  u.Password,
		"salt":      u.Salt,
	}

	return result
}

func (u User) ToReturn() map[string]interface{} {
	result := map[string]interface{}{
		"username":  u.Username,
		"name":      u.Name,
		"surname":   u.Surname,
		"email":     u.Email,
		"image_url": u.ImageURL,
	}

	return result
}
