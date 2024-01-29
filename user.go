package db

import (
	"fmt"

	"github.com/brass-software/files"
	"github.com/mikerybka/util"
)

type User struct {
	Phone         string
	LoginCodes    map[string]bool
	SessionTokens map[string]bool
}

func (u *User) ID() string {
	return u.Phone
}

func (u *User) Save() error {
	return files.WriteJSON(fmt.Sprintf("/users/%s", u.ID()), u)
}

func (u *User) CreateLoginCode() (string, error) {
	code := util.RandomCode(6)
	u.LoginCodes[code] = true
	return code, u.Save()
}

func (u *User) DeleteLoginCode(code string) error {
	u.LoginCodes[code] = false
	return u.Save()
}

func (u *User) Login(code string) (string, error) {
	if !u.LoginCodes[code] {
		return "", fmt.Errorf("bad login code")
	}
	delete(u.LoginCodes, code)
	token := util.RandomToken(32)
	u.SessionTokens[token] = true
	err := u.Save()
	if err != nil {
		return "", err
	}
	return token, nil
}
