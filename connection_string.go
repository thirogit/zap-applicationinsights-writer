package zapappinsigths

import (
	"errors"
	"strings"
)

type ConnectionString struct {
	InstrumentationKey string
	IngestionEndpoint  string
	LiveEndpoint       string
	ApplicationId      string
}

const (
	KeyWord_InstrumentationKey = "InstrumentationKey"
	KeyWord_IngestionEndpoint  = "IngestionEndpoint"
	KeyWord_LiveEndpoint       = "LiveEndpoint"
	KeyWord_ApplicationId      = "ApplicationId"
)

func NewConnectionString(connStr string) (*ConnectionString, error) {
	kcsb := ConnectionString{}
	if isEmpty(connStr) {
		return nil, errors.New("connection string cannot be empty")
	}
	connStrArr := strings.Split(connStr, ";")

	for _, kvp := range connStrArr {
		if isEmpty(strings.Trim(kvp, " ")) {
			continue
		}
		kvparr := strings.Split(kvp, "=")
		val := strings.Trim(kvparr[1], " ")
		if isEmpty(val) {
			continue
		}
		if err := assignValue(&kcsb, kvparr[0], val); err != nil {
			return nil, err
		}
	}

	return &kcsb, nil
}

func isEmpty(str string) bool {
	return strings.TrimSpace(str) == ""
}

func assignValue(kcsb *ConnectionString, key string, value string) error {

	switch key {
	case KeyWord_InstrumentationKey:
		kcsb.InstrumentationKey = value
	case KeyWord_IngestionEndpoint:
		kcsb.IngestionEndpoint = value
	case KeyWord_LiveEndpoint:
		kcsb.LiveEndpoint = value
	case KeyWord_ApplicationId:
		kcsb.ApplicationId = value
	}

	return nil
}
