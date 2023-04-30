package request

type DoorCreateUpdate struct {
	Id        string `validate:"required,uuid4"`
	Status    bool   `validate:"required,oneof=true false"`
	CreatedAt string `validate:"required,timestamp"`
	UpdatedAt string `validate:"required,timestamp"`
}
