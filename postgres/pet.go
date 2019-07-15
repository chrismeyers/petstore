package postgres

import (
	"github.com/chrismeyers/petstore"
	"github.com/jmoiron/sqlx"
)

type PetService struct {
	DB *sqlx.DB
}

func (p *PetService) Get(id int) (*petstore.Pet, error) {
	var pet petstore.Pet
	err := p.DB.Get(&pet, "select * from pets where id = $1", id)

	if err != nil {
		return nil, err
	}

	return &pet, nil
}

func (p *PetService) GetAll() ([]*petstore.Pet, error) {
	pets := []*petstore.Pet{}
	err := p.DB.Select(&pets, "select * from pets")

	if err != nil {
		return nil, err
	}

	return pets, nil
}

func (p *PetService) Create(pet petstore.Pet) (int, error) {
	var id int
	err := p.DB.Get(&id, "insert into pets (name) values ($1) returning id", pet.Name)

	if err != nil {
		return 0, err
	}

	return id, nil
}

func (p *PetService) Delete(id int) error {
	_, err := p.DB.Exec("delete from pets where id = $1", id)
	return err
}
