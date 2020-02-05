package partners

type PartnerRepository interface {
	SavePartner(p *Partner) error
	GetPartner(id int) (*Partner, error)
	SearchPartners(coords Coordinates) ([]Partner, error)
}

type ServicePartner struct {
	repo PartnerRepository
}

func NewService(repo PartnerRepository) *ServicePartner {
	return &ServicePartner{repo}
}

func (ps *ServicePartner) CreatePartner(_ *Partner) (*Partner, error) {
	return nil, nil
}

func (ps *ServicePartner) GetPartner(_ int) (*Partner, error) {
	return nil, nil
}
