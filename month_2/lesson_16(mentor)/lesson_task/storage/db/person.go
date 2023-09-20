package db

import (
	"context"
	"fmt"
	"playground/cpp-bootcamp/models"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type personRepo struct {
	db *pgxpool.Pool
}

func NewPerson(db *pgxpool.Pool) *personRepo {
	return &personRepo{
		db: db,
	}
}
func (p *personRepo) Create(newPerson models.CreatePerson) (string, error) {
	fmt.Println("person create")
	id := uuid.NewString()

	query := `
	INSERT INTO 
		persons(id,name,job,age) 
	VALUES($1,$2,$3,$4)`
	_, err := p.db.Exec(context.Background(), query,
		id,
		newPerson.Name,
		newPerson.Job,
		newPerson.Age,
	)
	if err != nil {
		fmt.Println("error:", err.Error())
		return "", err
	}
	return id, nil
}

func (p *personRepo) Update(newPerson models.Person) (string, error) {
	query := `
	UPDATE persons
	SET name=$2,job=$3,age=$4
	WHERE id=$1`
	resp, err := p.db.Exec(context.Background(), query,
		newPerson.Id,
		newPerson.Name,
		newPerson.Job,
		newPerson.Age,
	)
	if err != nil {
		return "", err
	}
	if resp.RowsAffected() == 0 {
		return "", pgx.ErrNoRows
	}
	return "OK", nil
}

func (p *personRepo) GetAll(req models.GetAllRequest) ([]models.Person, error) {

	return []models.Person{}, nil
}

func (p *personRepo) Get(req models.RequestByID) (models.Person, error) {
	query := `
	SELECT * FROM persons WHERE id = $1`

	row := p.db.QueryRow(context.Background(), query, req.ID)

	var person models.Person
	err := row.Scan(
		&person.Id,
		&person.Name,
		&person.Job,
		&person.Age,
	)

	if err != nil {
		return models.Person{}, err
	}
	return person, nil
}

func (p *personRepo) Delete(req models.RequestByID) (string, error) {
	query := `
	DELETE FROM persons WHERE id = $1`

	result, err := p.db.Exec(context.Background(), query, req.ID)
	if err != nil {
		return "", err
	}

	rowsAffected := result.RowsAffected()

	if rowsAffected == 0 {
		return "not found", nil
	}

	return "deleted", nil
}
