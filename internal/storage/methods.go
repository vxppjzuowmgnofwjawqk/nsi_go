package storage

import (
	"fmt"
	"nsi_go/internal/scheme"
)

// -- GET ALL ++

func (s *storage) GetAllCountries() []scheme.NSI {
	var result []scheme.NSI
	s.getAllHandler("countries", &result)
	return result
}

func (s *storage) GetAllLanguages() []scheme.NSI {
	var result []scheme.NSI
	s.getAllHandler("languages", &result)
	return result
}

func (s *storage) GetAllStatuses() []scheme.NSI {
	var result []scheme.NSI
	s.getAllHandler("statuses", &result)
	return result
}

func (s *storage) getAllHandler(table string, target *[]scheme.NSI) {
	query := fmt.Sprintf("SELECT * FROM %s", table)
	rows, err := s.db.Query(query)
	if err != nil {
		fmt.Printf("an error occurred while querying the database:\n%v\n", err)
	}
	var element scheme.NSI
	for rows.Next() {
		if err := rows.Scan(&element.Id, &element.Name); err != nil {
			fmt.Printf("an error occurred while processing data:\n%v\n", err)
		}
		*target = append(*target, element)
	}
}

// -- GET BY ID ++

func (s *storage) GetCountryById(id string) scheme.NSI {
	var element scheme.NSI
	s.getByIdHandler("countries", id, &element)
	return element
}

func (s *storage) GetLanguageById(id string) scheme.NSI {
	var element scheme.NSI
	s.getByIdHandler("languages", id, &element)
	return element
}

func (s *storage) GetStatusById(id string) scheme.NSI {
	var element scheme.NSI
	s.getByIdHandler("statuses", id, &element)
	return element
}

func (s *storage) getByIdHandler(table, id string, target *scheme.NSI) {
	query := fmt.Sprintf("SELECT * FROM %s WHERE id=$1", table)
	rows, err := s.db.Query(query, id)
	if err != nil {
		fmt.Printf("an error occurred while querying the database:\n%v\n", err)
	}
	for rows.Next() {
		if err := rows.Scan(&target.Id, &target.Name); err != nil {
			fmt.Printf("an error occurred while processing data:\n%v\n", err)
		}
	}
}

// -- CREATE ++

func (s *storage) CreateCountry(name string) scheme.NSI {
	var element = scheme.NSI{
		Id:   genId(),
		Name: name,
	}
	s.createHandler("countries", element.Id, element.Name)
	return element
}

func (s *storage) CreateLanguage(name string) scheme.NSI {
	var element = scheme.NSI{
		Id:   genId(),
		Name: name,
	}
	s.createHandler("languages", element.Id, element.Name)
	return element
}

func (s *storage) CreateStatus(name string) scheme.NSI {
	var element = scheme.NSI{
		Id:   genId(),
		Name: name,
	}
	s.createHandler("statuses", element.Id, element.Name)
	return element
}

func (s *storage) createHandler(table, id, name string) {
	query := fmt.Sprintf("INSERT INTO %s VALUES ($1, $2)", table)
	_, err := s.db.Query(query, id, name)
	if err != nil {
		fmt.Printf("an error occurred while creating:\n%v\n", err)
	}
}

// -- UPDATE ++

func (s *storage) UpdateCountry(id, name string) {
	s.updateHandler("countries", id, name)
}

func (s *storage) UpdateLanguage(id, name string) {
	s.updateHandler("languages", id, name)
}

func (s *storage) UpdateStatus(id, name string) {
	s.updateHandler("statuses", id, name)
}

func (s *storage) updateHandler(table, id, name string) {
	query := fmt.Sprintf("UPDATE %s SET name=$1 WHERE id=$2", table)
	_, err := s.db.Query(query, name, id)
	if err != nil {
		fmt.Printf("an error occurred while updating the data:\n%v\n", err)
	}
}

// -- DELETE ++

func (s *storage) DeleteCountry(id string) {
	s.deleteHandler("countries", id)
}

func (s *storage) DeleteLanguage(id string) {
	s.deleteHandler("languages", id)
}

func (s *storage) DeleteStatus(id string) {
	s.deleteHandler("statuses", id)
}

func (s *storage) deleteHandler(table, id string) {
	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1", table)
	_, err := s.db.Query(query, id)
	if err != nil {
		fmt.Printf("an error occurred while deleting data:\n%v\n", err)
	}
}
