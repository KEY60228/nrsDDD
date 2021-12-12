package model

import "errors"

type modelNumber struct {
	productCode string
	branch      string
	lot         string
}

func New(productCode string, branch string, lot string) (*modelNumber, error) {
	if len(productCode) == 0 {
		return nil, errors.New("productCode's length is zero")
	}
	if len(branch) == 0 {
		return nil, errors.New("branch's length is zero")
	}
	if len(lot) == 0 {
		return nil, errors.New("lot's length is zero")
	}
	mn := &modelNumber{
		productCode: productCode,
		branch:      branch,
		lot:         lot,
	}
	return mn, nil
}
