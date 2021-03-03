package services

import (
	"budget4home/src/models"
	"budget4home/src/repositories"
)

type ILabelService interface {
	Fetch() (res []models.Label, err error)
}

type LabelService struct {
	repo repositories.ILabelRepository
}

func NewLabelService(r repositories.ILabelRepository) ILabelService {
	return &LabelService{
		repo: r,
	}
}

func (this *LabelService) Fetch() (res []models.Label, err error) {
	res, err = this.repo.Fetch()
	if err != nil {
		return nil, err
	}
	return
}
