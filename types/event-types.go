package types

import "github.com/google/uuid"

type OrganizationCreatedEvent struct {
	OrganizationID uuid.UUID `json:"organizationID" validate:"required"`
	AdminEmail     string    `json:"adminEmail" validate:"required"`
}
