package repository_interfaces

import "github.com/RandySteven/Idolized/entities/models"

type UserRepository interface {
	Saver[models.User]
}
