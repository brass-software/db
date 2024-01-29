package db

import (
	"fmt"

	"github.com/brass-software/files"
)

func GetUser(phone string) (*User, error) {
	return files.ReadJSON[User](fmt.Sprintf("/users/%s", phone))
}
