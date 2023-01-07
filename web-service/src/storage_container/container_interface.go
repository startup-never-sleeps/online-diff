package storage_container

import (
	guuid "github.com/google/uuid"
	utils "web-service/src/utils"
)

type ResStatus string

const (
	Success ResStatus = "Success"
	Error   ResStatus = "Error"
	Pending ResStatus = "Pending"
)

type ClientStorageInterface interface {
	GetResValue(id guuid.UUID) (*utils.Pair, error)
	SavePendingClient(id guuid.UUID, res string) error
	ClientExists(id guuid.UUID) bool
	SaveErrorClient(id guuid.UUID, res string) error
	SaveSuccessClient(id guuid.UUID, res string) error
	GetRemoveStaleClients(back_period int64) ([]string, error)
}
