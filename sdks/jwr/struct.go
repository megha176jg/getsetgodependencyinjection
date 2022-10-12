package jwr

import (
	"log"
	"strings"
	"time"
)

type UserProfile struct {
	FirstName  string `json:"firstName"`
	MiddleName string `json:"middleName"`
	LastName   string `json:"lastName"`
	DOB        int64  `json:"dateOfBirth"`
	Pin        string `json:"pin"`
	Address    string `json:"address"`
	Address2   string `json:"address2"`
	City       string `json:"city"`
	State      string `json:"state"`
	Gender     string `json:"gender"`
}

type JWRSDKConfig struct {
	BaseURL    string
	Token      string
	APITimeout int
}

func (u UserProfile) GetFullName() string {
	if strings.ToLower(u.FirstName) == "fnu" {
		u.FirstName = ""
	}
	if strings.ToLower(u.LastName) == "lnu" {
		u.LastName = ""
	}
	if u.MiddleName == "" {
		return strings.TrimSpace(u.FirstName + " " + u.LastName)
	}
	return strings.TrimSpace(u.FirstName + " " + u.MiddleName + " " + u.LastName)
}

func (u UserProfile) GetDOB(log *log.Logger, traceID string) string {
	t := time.Unix(0, int64(u.DOB)*int64(time.Millisecond))
	loc, _ := time.LoadLocation("Asia/Kolkata")
	t = t.In(loc)
	s := t.Format("02/01/2006")
	log.Printf("GETDOB %s : profileDOB: %d (%s) : output DOB: %s\n", traceID, u.DOB, t.Format(time.RFC3339), s)
	return s
}
