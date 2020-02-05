package partners

import (
	"testing"
)

type ServicePartnerMock struct{}

func (spm *ServicePartnerMock) SavePartner(p *Partner) error {
	if p == nil {
		return nilPartnerError
	}
	return nil
}
func (spm *ServicePartnerMock) GetPartner(id string) (*Partner, error) {
	if len(id) <= 0 {
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
	id      string
	partner *Partner
	error   error
}

var getTestCases = map[string]getPartnerTestCase{
	"Success GetPartner": {
		"1",
		&Partner{ID: "1"},
		nil,
	},
	"Error GetPartner": {
		"",
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

type searchPartnerTestCase struct {
	length int
	error  error
}

var searchTestCases = map[string]searchPartnerTestCase{
	"Success GetPartner": {
		3,
		nil,
	},
	"Error GetPartner": {
		0,
		invalidCoordError,
	},
}

func TestServicePartner_SearchPartner(t *testing.T) {
	svc := NewService(new(ServicePartnerMock))
	for caseName, tCase := range searchTestCases {
		coords := make(Coordinates, tCase.length)
		caseResult, err := svc.SearchPartners(coords)
		if err != tCase.error {
			t.Errorf("case: %s\n expected: %+e\n got: %+e\n", caseName, tCase.error, err)
		}
		if len(caseResult) != tCase.length {
			t.Errorf("case: %s\n expected length: %d\n got: %+v\n", caseName, tCase.length, caseResult)
		}
	}
}
