package repository

import "github.com/adibhauzan/go-unit-test/entity"

type CategoryRepository interface {
	FindById(id string) *entity.Category
}



