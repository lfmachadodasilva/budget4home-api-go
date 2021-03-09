package group

type IGroupService interface {
	Fetch() (res []Group, err error)
}

type GroupService struct {
	repo IGroupRepository
}

func NewService(r IGroupRepository) IGroupService {
	return &GroupService{
		repo: r,
	}
}

func (this *GroupService) Fetch() (res []Group, err error) {
	res, err = this.repo.Fetch()
	if err != nil {
		return nil, err
	}
	return
}
