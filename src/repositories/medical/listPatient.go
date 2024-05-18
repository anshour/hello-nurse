package medicalRepository

import (
	entities "hello-nurse/src/entities/medical"
	"log"
	"reflect"
	"strconv"
	"strings"
)

func (dbase *controllerMedical) ListPatient(filters *entities.PatientListFilter) ([]*entities.PatientListResult, error) {
	baseQuery := "SELECT identity_number, name, phone_number, birth_date, gender, created_at FROM patients WHERE true"
	params := []interface{}{}
	n := (&entities.PatientListFilter{})

	if !reflect.DeepEqual(filters, n) {
		conditions := []string{}

		if filters.IdentityNumber != 0 {
			conditions = append(conditions, "identity_number = $"+strconv.Itoa(len(params)+1))
			params = append(params, filters.IdentityNumber)
		}
		if filters.Name != "" {
			conditions = append(conditions, "name LIKE $"+strconv.Itoa(len(params)+1))
			params = append(params, "%"+filters.Name+"%")
		}
		if filters.PhoneNumber != "" {
			conditions = append(conditions, "phone_number LIKE $"+strconv.Itoa(len(params)+1))
			params = append(params, "%"+filters.PhoneNumber+"%")
		}

		if len(conditions) > 0 {
			baseQuery += " AND "
		}
		baseQuery += strings.Join(conditions, " AND ")

	}

	if filters.Limit == 0 {
		filters.Limit = 5
	}

	if filters.CreatedAt == "" {
		filters.CreatedAt = "DESC"
	}
	if filters.CreatedAt == "asc" {
		filters.CreatedAt = "ASC"
	}

	baseQuery += " ORDER BY created_at " + filters.CreatedAt

	baseQuery += " LIMIT $" + strconv.Itoa(len(params)+1)
	params = append(params, filters.Limit)

	if filters.Offset == 0 {
		filters.Offset = 0
	} else {
		baseQuery += " OFFSET $" + strconv.Itoa(len(params)+1)
		params = append(params, filters.Offset)
	}

	println("baseQuery: ", baseQuery)
	rows, err := dbase.DB.Query(baseQuery, params...)
	if err != nil {
		log.Printf("Error finding patients: %s", err)
		return nil, err
	}
	defer rows.Close()

	var patients = make([]*entities.PatientListResult, 0)
	for rows.Next() {
		patient := new(entities.PatientListResult)
		if err := rows.Scan(
			&patient.IdentityNumber,
			&patient.Name,
			&patient.PhoneNumber,
			&patient.BirthDate,
			&patient.Gender,
			&patient.CreatedAt,
		); err != nil {
			return nil, err
		}
		patients = append(patients, patient)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return patients, nil
}
