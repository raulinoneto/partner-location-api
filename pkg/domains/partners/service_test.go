package partners

import (
	"errors"
	"testing"
)

var nilPartnerError = errors.New("nil partner, cannot save")
var invalidIdError = errors.New("invalid partner id")
var invalidCoordError = errors.New("invalid coordinates")

type ServicePartnerMock struct{}

func (spm *ServicePartnerMock) SavePartner(p *Partner) error {
	if p == nil {
		return nilPartnerError
	}
	return nil
}
func (spm *ServicePartnerMock) GetPartner(id int) (*Partner, error) {
	if id <= 0 {
		return nil, invalidIdError
	}
	return &Partner{ID: id}, nil
}
func (spm *ServicePartnerMock) SearchPartners(coords Coordinates) ([]Partner, error) {
	if coords == nil {
		return nil, invalidCoordError
	}
	partners := make([]Partner, len(coords))
	return partners, nil
}

type createPartnerTestCase struct {
	error   error
	payload *Partner
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
		caseResult, err := svc.CreatePartner(tCase.payload)
		if err != tCase.error {
			t.Errorf("case: %s\n expected: %+e\n got: %+e\n", caseName, tCase.error, err)
		}
		if caseResult != tCase.payload {
			t.Errorf("case: %s\n expected: %+v\n got: %+v\n", caseName, tCase.payload, caseResult)
		}
	}
}

type getPartnerTestCase struct {
	id      int
	partner *Partner
	error   error
}

var getTestCases = map[string]getPartnerTestCase{
	"Success GetPartner": {
		1,
		&Partner{ID: 1},
		nil,
	},
	"Error GetPartner": {
		0,
		nil,
		invalidIdError,
	},
}

func TestServicePartner_GetPartner(t *testing.T) {
	svc := NewService(new(ServicePartnerMock))
	for caseName, tCase := range getTestCases {
		caseResult, err := svc.GetPartner(tCase.id)
		if err != tCase.error {
			t.Errorf("case: %s\n expected: %+e\n got: %+e\n", caseName, tCase.error, err)
		}
		if caseResult != tCase.partner {
			t.Errorf("case: %s\n expected: %+v\n got: %+v\n", caseName, tCase.partner, caseResult)
		}
	}
}
