package storage

import "playground/cpp-bootcamp/models"

type StorageI interface {
	Person() PersonsI
}
type PersonsI interface {
	Create(models.CreatePerson) (string, error)
	Update(models.Person) (string, error)
	Get(req models.RequestByID) (models.Person, error)
	GetAll(models.GetAllRequest) ([]models.Person, error)
	Delete(req models.RequestByID) (string, error)
}
