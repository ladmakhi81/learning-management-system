package userentity

import "time"

type Teacher struct {
	Bio               string
	ResumeFile        string
	SalaryAmount      float64
	Email             string
	NationalID        string
	IsVerified        bool
	VerifiedDate      *time.Time
	VerifiedByID      *uint
	IsProfileComplete bool
}
