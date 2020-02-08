package partners

import (
	"errors"
)

var nilPartnerError = errors.New("nil partner, cannot save")
var invalidIdError = errors.New("invalid partner id")
var invalidPointError = errors.New("invalid coordinates")

type PartnerRepository interface {
	SavePartner(partner *Partner) (*Partner, error)
	GetPartner(id string) (*Partner, error)
	SearchPartners(point *Point) ([]Partner, error)
}

type ServicePartner struct {
	repo PartnerRepository
}

func NewService(repo PartnerRepository) *ServicePartner {
	return &ServicePartner{repo}
}

func (ps *ServicePartner) CreatePartner(p *Partner) (*Partner, error) {
	if p == nil {
		return nil, nilPartnerError
	}
	p, err := ps.repo.SavePartner(p)
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (ps *ServicePartner) GetPartner(id string) (*Partner, error) {
	if len(id) == 0 {
		return nil, invalidIdError
	}
	partner, err := ps.repo.GetPartner(id)
	if err != nil {
		return nil, err
	}
	return partner, nil
}

func (ps *ServicePartner) SearchPartners(point *Point) (*Pdvs, error) {
	partners, err := ps.repo.SearchPartners(point)
	if err != nil {
		return nil, err
	}
	return &Pdvs{partners}, nil
}
