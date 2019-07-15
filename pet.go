package petstore

type Pet struct {
	ID   string `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
}

type PetService interface {
	Get(id int) (*Pet, error)
	GetAll() ([]*Pet, error)
	Create(p Pet) (int, error)
	Delete(id int) error
}
