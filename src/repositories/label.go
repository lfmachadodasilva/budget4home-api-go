package repositories

import (
	"budget4home/src/models"
	"fmt"
	"math/rand"
)

type ILabelRepository interface {
	Fetch() (res []models.Label, err error)
}

type LabelRepository struct {
}

func NewLabelRepository() ILabelRepository {
	return &LabelRepository{}
}

func (this *LabelRepository) Fetch() (res []models.Label, err error) {
	id := rand.Intn(100)
	return []models.Label{
		{
			Id:   id,
			Name: "Label " + fmt.Sprint(id),
		},
	}, nil
}
