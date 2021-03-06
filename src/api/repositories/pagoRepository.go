package repositories

import (
	"InmoGo/src/api/models"
	"InmoGo/src/api/utils"
	"fmt"
	"gorm.io/gorm"
)

type PagoRepository struct {
	db *gorm.DB
}

func NewPagoRepository(db *gorm.DB) *PagoRepository {
	return &PagoRepository{db: db}
}

func (p *PagoRepository) Save(pago *models.Pago) {
	p.db.Save(pago)
}

func (p *PagoRepository) Get(ID int) (*models.Pago, error) {
	var pago *models.Pago
	p.db.First(&pago, ID)

	if pago.ID != 0 {
		return pago, nil
	}
	return nil, utils.InmoError{
		Code:    404,
		Message: fmt.Sprintf("not found pago with id: %v", ID),
	}
}

func (p *PagoRepository) GetAll(alquilerID int) []*models.Pago {
	var pagos []*models.Pago
	p.db.Where("alquiler_id = ?", alquilerID).Find(&pagos)

	return pagos
}
