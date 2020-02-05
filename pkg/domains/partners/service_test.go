package partners

import (
	"errors"
	"testing"
)

var nilPartnerError = errors.New("nil partner, cannot save")

type ServicePartnerMock struct{}

func (spm *ServicePartnerMock) SavePartner(p *Partner) error {
	if p == nil {
		return nilPartnerError
	}
	return nil
}
func (spm *ServicePartnerMock) GetPartner(id int) (*Partner, error) {
	if id <= 0 {
		return nil, errors.New("invalid partner id")
	}
	return &Partner{ID: id}, nil
}
func (spm *ServicePartnerMock) SearchPartners(coords Coordinates) ([]Partner, error) {
	if coords == nil {
		return nil, errors.New("invalid coordinates")
	}
	return nil, nil
}

type createPartnerTestCase struct {
	expected error
	payload  *Partner
}

var createTestCases = map[string]createPartnerTestCase{
	"Success CreatePartner": {
		nil,
		&Partner{},
	},
	"Error CreatePartner": {
		nilPartnerError,
		nil,
	},
}

func TestServicePartner_CreatePartner(t *testing.T) {
	svc := NewService(new(ServicePartnerMock))
	for caseName, tCase := range createTestCases {
		caseResult := svc.CreatePartner(tCase.payload)
		if caseResult != tCase.expected {
			t.Errorf("case: %s\n expected: %+e\n got: %+e\n", caseName, tCase.expected, caseResult)
		}
	}
}
