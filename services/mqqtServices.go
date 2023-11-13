package services

import "context"

type MqqtService interface {
	GetStatusDoor(ctx context.Context)
}
