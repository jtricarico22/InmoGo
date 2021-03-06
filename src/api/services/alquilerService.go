package services

import (
	"InmoGo/src/api/models"
	"InmoGo/src/api/repositories"
)

type AlquilerService struct {
	repository *repositories.AlquilerRepository
}

func NewAlquilerService(repository *repositories.AlquilerRepository) *AlquilerService {
	return &AlquilerService{repository: repository}
}

func (i *AlquilerService) Get(ID int) (*models.Alquiler, error) {
	return i.repository.Get(ID)
}

func (i *AlquilerService) Save(propietario *models.Alquiler) *models.Alquiler {
	i.repository.Save(propietario)
	return nil
}

func (i *AlquilerService) GetAllByInmueble(propietarioID int) ([]*models.Alquiler, error) {
	return i.repository.GetAllByInmueble(propietarioID)
}
