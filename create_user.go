package db

func CreateUser(phone string) (*User, error) {
	u := &User{
		Phone: phone,
	}
	err := u.Save()
	if err != nil {
		return nil, err
	}
	return u, nil
}
