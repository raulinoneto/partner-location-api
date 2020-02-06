package partners

import (
	"reflect"
	"testing"
)

type repoPartnerMock struct{}

func (rpm *repoPartnerMock) SavePartner(p *Partner) error {
	if p == nil {
		return nilPartnerError
	}
	return nil
}
func (rpm *repoPartnerMock) GetPartner(id string) (*Partner, error) {
	if len(id) <= 0 {
		return nil, invalidIdError
	}
	return &Partner{ID: id}, nil
}
func (rpm *repoPartnerMock) SearchPartners(point *Point) ([]Partner, error) {
	if point == nil {
		return nil, invalidPointError
	}
	partners := make([]Partner, int(point.Latitude))
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
	svc := NewService(new(repoPartnerMock))
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
	svc := NewService(new(repoPartnerMock))
	for caseName, tCase := range getTestCases {
		caseResult, err := svc.GetPartner(tCase.id)
		if err != tCase.error {
			t.Errorf("case: %s\n expected: %+e\n got: %+e\n", caseName, tCase.error, err)
		}
		if caseResult != nil && !reflect.DeepEqual(*caseResult, *tCase.partner) {
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
		invalidPointError,
	},
}

func TestServicePartner_SearchPartner(t *testing.T) {
	svc := NewService(new(repoPartnerMock))
	for caseName, tCase := range searchTestCases {
		var point *Point
		if tCase.length > 0 {
			point = &Point{Latitude: float64(tCase.length)}
		}
		caseResult, err := svc.SearchPartners(point)
		if err != tCase.error {
			t.Errorf("case: %s\n expected: %+e\n got: %+e\n", caseName, tCase.error, err)
		}
		if caseResult != nil && len(caseResult.Pdvs) != tCase.length {
			t.Errorf("case: %s\n expected length: %d\n got: %+v\n", caseName, tCase.length, caseResult)
		}
	}
}
