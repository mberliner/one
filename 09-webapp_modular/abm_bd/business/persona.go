package business

import (
	"database/sql"
	"github.com/mberliner/gobase/09-webapp_modular/abm_bd/model"
	"github.com/mberliner/gobase/09-webapp_modular/abm_bd/repository"
	"log"
	"strconv"
	"time"
)

func CreaPersona(nom string, ape string, fechaNacimiento string) model.Personas {

	//Inicio como nulo, si no lo es lo cambio
	fechaNull := sql.NullTime{
		Valid: false,
	}
	if fechaNacimiento != "" {
		fecha, err := time.Parse("02-01-2006", fechaNacimiento)
		if err != nil {
			log.Println("Error persiste Persona con fecha:", err)
			mP := model.Personas{}
			mP.Error = err
			return mP
		}
		fechaNull = sql.NullTime{
			Time:  fecha,
			Valid: true,
		}
	}

	p := repository.Persona{Nombre: nom, Apellido: ape, FechaNacimiento: fechaNull}
	p, err := repository.PR.Persiste(p)
	if err != nil {
		log.Println("Error persiste Pesona:", err)
		mP := model.Personas{}
		mP.Error = err
		return mP
	}

	var fecha string
	if p.FechaNacimiento.Valid == true {
		fecha = p.FechaNacimiento.Time.Format("02-01-2006")
	} else {
		fecha = ""
	}
	per := model.Persona{ID: p.ID,
		Nombre:          p.Nombre,
		Apellido:        p.Apellido,
		FechaNacimiento: fecha,
	}

	personas := []model.Persona{per}

	mP := model.Personas{
		PersonasM: personas,
		Error:     nil,
		Mensaje:   "Persona Creada Ok",
	}
	return mP
}

func BuscaTodo() model.Personas {

	ps, err := repository.PR.BuscaTodo()
	if err != nil {
		log.Println("Error buscaTodo:", err)
		mP := model.Personas{}
		mP.Error = err
		return mP
	}
	personas := []model.Persona{}
	var fecha string
	for _, p := range ps {

		if p.FechaNacimiento.Valid == true {
			fecha = p.FechaNacimiento.Time.Format("02-01-2006")
		} else {
			fecha = ""
		}

		per := model.Persona{ID: p.ID,
			Nombre:          p.Nombre,
			Apellido:        p.Apellido,
			FechaNacimiento: fecha,
		}
		personas = append(personas, per)
	}
	mP := model.Personas{
		PersonasM: personas,
		Error:     nil,
		Mensaje:   "Carga ok",
	}

	return mP
}

func BorraPersona(id int) model.Personas {

	err := repository.PR.Borra(id)
	if err != nil {
		log.Println("Error borraPersona:", err)
		mP := model.Personas{}
		mP.Error = err
		return mP
	}
	personas := []model.Persona{}
	per := model.Persona{ID: id,
		Nombre:   "",
		Apellido: "",
	}

	personas = append(personas, per)

	mP := model.Personas{
		PersonasM: personas,
		Error:     nil,
		Mensaje:   "Borrado ok",
	}

	return mP
}

func BuscaPersona(id int) model.Personas {

	p, err := repository.PR.BuscaPorId(id)
	log.Println("Error en editarPersona1-------:", id, p, err)
	if err != nil {
		log.Println("Error buscaPersona:", err)
		mP := model.Personas{}
		mP.Error = err
		return mP
	}
	personas := []model.Persona{}

	var fecha string
	if p.FechaNacimiento.Valid == true {
		fecha = p.FechaNacimiento.Time.Format("02-01-2006")
	} else {
		fecha = ""
	}
	per := model.Persona{ID: p.ID,
		Nombre:          p.Nombre,
		Apellido:        p.Apellido,
		FechaNacimiento: fecha,
	}
	personas = append(personas, per)

	mP := model.Personas{
		PersonasM: personas,
		Error:     nil,
		Mensaje:   "Busca ok",
	}
	log.Println("Error en editarPersona2-------:", mP)
	return mP
}

func ActualizaPersona(id string, nom string, ape string, fechaNacimiento string) model.Personas {

	//Inicio como nulo, si no lo es lo cambio
	fechaNull := sql.NullTime{
		Valid: false,
	}
	if fechaNacimiento != "" {
		fecha, err := time.Parse("02-01-2006", fechaNacimiento)
		if err != nil {
			log.Println("Error actualiza Persona con fecha:", err)
			mP := model.Personas{}
			mP.Error = err
			return mP
		}
		fechaNull = sql.NullTime{
			Time:  fecha,
			Valid: true,
		}
	}
	idd, _ := strconv.Atoi(id)
	p := repository.Persona{Nombre: nom, Apellido: ape, FechaNacimiento: fechaNull, ID: idd}
	p, err := repository.PR.Actualiza(p)
	log.Println("Error actualiza Persona con fecha:", p, err)
	if err != nil {
		log.Println("Error actualiza Pesona:", err)
		mP := model.Personas{}
		mP.Error = err
		return mP
	}

	var fecha string
	if p.FechaNacimiento.Valid == true {
		fecha = p.FechaNacimiento.Time.Format("02-01-2006")
	} else {
		fecha = ""
	}
	per := model.Persona{ID: p.ID,
		Nombre:          p.Nombre,
		Apellido:        p.Apellido,
		FechaNacimiento: fecha,
	}

	personas := []model.Persona{per}

	mP := model.Personas{
		PersonasM: personas,
		Error:     nil,
		Mensaje:   "Persona Actualizada Ok",
	}
	return mP
}
