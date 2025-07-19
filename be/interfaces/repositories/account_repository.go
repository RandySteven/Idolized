package repository_interfaces

import "github.com/RandySteven/Idolized/entities/models"

type AccountRepository interface {
	Saver[models.Account]
}
