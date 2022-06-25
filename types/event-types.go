package types

import (
	"github.com/google/uuid"
	"time"
)

type OrganizationCreatedEvent struct {
	OrganizationID uuid.UUID `json:"organizationID" validate:"required"`
	AdminEmail     string    `json:"adminID" validate:"required"`
}

type ApplicationShortlisted struct {
	ApplicationID uuid.UUID `json:"applicationID" validate:"required"`
}

type AvailabilityCreated struct {
	AvailabilityID  uuid.UUID `json:"availabilityID" validate:"required"`
	CandidateID     uuid.UUID `json:"candidateID" validate:"required"`
	RoleID          uuid.UUID `json:"roleID" validate:"required"`
	InterviewStepID uuid.UUID `json:"interviewStepID" validate:"required"`
	StartTime       time.Time `json:"startTime" validate:"required"`
}
