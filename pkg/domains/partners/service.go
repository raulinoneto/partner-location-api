package partners

import (
	"errors"
	"math"
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

func (ps *ServicePartner) SearchPartners(point *Point) (*Partner, error) {
	var nearest float64
	partnerResponse := new(Partner)
	partners, err := ps.repo.SearchPartners(point)
	if err != nil {
		return nil, err
	}
	for i, partner := range partners {
		distance := getDistance(point, partner)
		if i == 0 || distance < nearest {
			partnerResponse = &partner
			nearest = distance
			continue
		}
	}
	return partnerResponse, nil
}

func getDistance(point *Point, partner Partner) float64 {
	if len(partner.Address.Coordinates) < 2 {
		panic("Malformed coordinates")
	}
	rl1 := float64(math.Pi * point.Latitude / 180)
	rl2 := float64(math.Pi * partner.Address.Coordinates[0] / 180)

	theta := float64(point.Longitude - partner.Address.Coordinates[1])
	radtheta := float64(math.Pi * theta / 180)

	d := math.Sin(rl1)*math.Sin(rl2) + math.Cos(rl1)*math.Cos(rl2)*math.Cos(radtheta)
	d = math.Acos(d)
	d = d * 180 / math.Pi
	d = d * 60 * 1.1515
	d = d * 1.609344

	return d
}
