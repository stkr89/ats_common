package types

import "github.com/google/uuid"

type OrganizationCreatedEvent struct {
	OrganizationID uuid.UUID `json:"organizationID" validate:"required"`
	AdminEmail     string    `json:"adminID" validate:"required"`
}

type ApplicationShortlisted struct {
	ApplicationID uuid.UUID `json:"applicationID" validate:"required"`
}
