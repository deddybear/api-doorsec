package repository

type MqqtRepositoryInterface interface {
}

type MqqtImpl struct {
}

func NewMqqtRepository() MqqtRepositoryInterface {
	return &MqqtImpl{}
}
