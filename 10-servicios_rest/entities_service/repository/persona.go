package repository

import (
	"database/sql"
	"log"
	"time"

	"github.com/mberliner/gobase/10-servicios_rest/entities_service/domain"
)

//TODO agregar los null
//y unique a Usuario en BD
type persona struct {
	ID              int
	Nombre          string
	Apellido        string
	FechaNacimiento sql.NullTime
}

type PersonaRepository interface {
	Persiste(p *domain.Persona) (*domain.Persona, error)
	Borra(id int) error
	BuscaTodo() ([]domain.Persona, error)
	BuscaPorID(id int) (*domain.Persona, error)
	Actualiza(p *domain.Persona) (*domain.Persona, error)
}

type personaRepository struct {
	db *sql.DB
}

func NewPersonaRepository(db *sql.DB) PersonaRepository {
	return &personaRepository{db}
}

func (pR personaRepository) Persiste(p *domain.Persona) (*domain.Persona, error) {
	//Inicio como nulo, si no lo es lo cambio
	fechaNull := sql.NullTime{
		Valid: false,
	}
	if p.FechaNacimiento != "" {
		fecha, err := time.Parse("02-01-2006", p.FechaNacimiento)
		if err != nil {
			log.Println("Error persiste Persona con fecha:", err)
			return nil, err
		}
		fechaNull = sql.NullTime{
			Time:  fecha,
			Valid: true,
		}
	}
	stmt, err := pR.db.Prepare("INSERT into persona(nombre, apellido, fecha_nacimiento) VALUES(?,?,?);")
	if err != nil {
		return nil, err
	}

	res, err := stmt.Exec(p.Nombre, p.Apellido, fechaNull)
	if err != nil {
		return nil, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	p.ID = int(id)

	return p, nil
}

func (pR personaRepository) Borra(id int) error {

	stmt, err := pR.db.Prepare("DELETE FROM persona WHERE id = ?;")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}

	return nil
}

func (pR personaRepository) BuscaTodo() ([]domain.Persona, error) {

	rows, err := pR.db.Query("SELECT id, nombre, apellido, fecha_nacimiento FROM persona;")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var rP []domain.Persona
	var p persona

	var fecha string
	for rows.Next() {
		err = rows.Scan(&p.ID, &p.Nombre, &p.Apellido, &p.FechaNacimiento)
		if err != nil {
			return nil, err
		}

		if p.FechaNacimiento.Valid == true {
			fecha = p.FechaNacimiento.Time.Format("02-01-2006")
		} else {
			fecha = ""
		}

		per := domain.Persona{ID: p.ID,
			Nombre:          p.Nombre,
			Apellido:        p.Apellido,
			FechaNacimiento: fecha,
		}

		rP = append(rP, per)
	}

	return rP, nil
}

func (pR personaRepository) BuscaPorID(id int) (*domain.Persona, error) {
	var p persona
	err := pR.db.QueryRow("SELECT id, nombre, apellido, fecha_nacimiento FROM persona WHERE id = ?;", id).
		Scan(&p.ID, &p.Nombre, &p.Apellido, &p.FechaNacimiento)
	if err != nil {
		return nil, err
	}

	var fecha string
	if p.FechaNacimiento.Valid == true {
		fecha = p.FechaNacimiento.Time.Format("02-01-2006")
	} else {
		fecha = ""
	}
	per := &domain.Persona{ID: p.ID,
		Nombre:          p.Nombre,
		Apellido:        p.Apellido,
		FechaNacimiento: fecha,
	}
	return per, nil
}

func (pR personaRepository) Actualiza(p *domain.Persona) (*domain.Persona, error) {

	//Inicio como nulo, si no lo es lo cambio
	fechaNull := sql.NullTime{
		Valid: false,
	}
	if p.FechaNacimiento != "" {
		fecha, err := time.Parse("02-01-2006", p.FechaNacimiento)
		if err != nil {
			log.Println("Error actualiza Persona con fecha:", err)
			return nil, err
		}
		fechaNull = sql.NullTime{
			Time:  fecha,
			Valid: true,
		}
	}

	stmt, err := pR.db.Prepare("Update persona SET nombre=?, apellido=?, fecha_nacimiento=? WHERE id=?;")
	if err != nil {
		return nil, err
	}

	_, err = stmt.Exec(p.Nombre, p.Apellido, fechaNull, p.ID)
	if err != nil {
		return nil, err
	}

	return p, nil
}
