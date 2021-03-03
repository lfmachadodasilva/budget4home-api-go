package label

type ILabelService interface {
	Fetch() (res []Label, err error)
}

type LabelService struct {
	repo ILabelRepository
}

func NewService(r ILabelRepository) ILabelService {
	return &LabelService{
		repo: r,
	}
}

func (this *LabelService) Fetch() (res []Label, err error) {
	res, err = this.repo.Fetch()
	if err != nil {
		return nil, err
	}
	return
}
