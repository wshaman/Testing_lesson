package user

import (
	"gtest_example/app/internal/db"
	"gtest_example/app/internal/models"
	"gtest_example/app/utils/naming"

	"github.com/pkg/errors"
)

var ErrEmailInvalid = errors.New("invalid email format")

func Register(d *db.DB, name string) (*models.User, error) {
	eml, err := naming.NameToEmail(name)
	if err != nil {
		return nil, errors.Wrap(err, "failed to register user")
	}
	m := &models.User{
		Email: eml,
		Name:  name,
	}
	if err = models.UserSave(d, m); err != nil {
		return nil, errors.Wrap(err, "failed to register user")
	}
	return m, nil
}

func getNextNameAvailable(d *db.DB, eml string) (string, error) {
	namePart, err := naming.NameFromEmail(eml)
	if err != nil {
		return "", errors.Wrap(err, "failed to parse email")
	}
	users, err := models.UserListEmailLike(d, namePart)
	if err != nil {
		return "", errors.Wrap(err, "failed to find users")
	}
	emails := make([]string, 0, len(users))
	for _, u := range users {
		name, err := naming.NameFromEmail(u.Email)
		if err != nil {
			return "", errors.Wrap(err, "failed to parse email")
		}
		emails = append(emails, name)
	}
	newName, err := naming.NameWithNumberIncrement(namePart, emails)
	if err != nil {
		return "", errors.Wrap(err, "failed to generate new email")
	}
	return newName, nil
}
