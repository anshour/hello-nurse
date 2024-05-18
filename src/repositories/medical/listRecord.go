package medicalRepository

import (
	entities "hello-nurse/src/entities/medical"
	"log"
	"reflect"
	"strconv"
	"strings"
)

func (dbase *controllerMedical) ListRecord(filters *entities.RecordListFilter) ([]*entities.RecordListResult, error) {
	baseQuery := `SELECT 
		p.identity_number, p.phone_number, p.name, p.birth_date, p.gender, p.identity_image,
		r.symptoms, r.medications, r.created_at,
		u.name, u.nip, u.id
		FROM medical_records r
		JOIN users u ON r.user_id = u.id
		JOIN patients p ON r.identity_number = p.identity_number
		WHERE true `

	params := []interface{}{}
	n := (&entities.PatientListFilter{})

	if !reflect.DeepEqual(filters, n) {
		conditions := []string{}

		if filters.IdentityNumber != 0 {
			conditions = append(conditions, "r.identity_number = $"+strconv.Itoa(len(params)+1))
			params = append(params, filters.IdentityNumber)
		}

		if filters.UserId != "" {
			conditions = append(conditions, "u.id = $"+strconv.Itoa(len(params)+1))
			params = append(params, filters.IdentityNumber)
		}

		if filters.Nip != "" {
			conditions = append(conditions, "u.nip = $"+strconv.Itoa(len(params)+1))
			params = append(params, filters.Nip)
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

	rows, err := dbase.DB.Query(baseQuery, params...)
	if err != nil {
		log.Printf("Error finding recordss: %s", err)
		return nil, err
	}
	defer rows.Close()

	var records = make([]*entities.RecordListResult, 0)
	for rows.Next() {
		record := new(entities.RecordListResult)
		if err := rows.Scan(
			&record.IdentityDetail.IdentityNumber,
			&record.IdentityDetail.PhoneNumber,
			&record.IdentityDetail.Name,
			&record.IdentityDetail.BirthDate,
			&record.IdentityDetail.Gender,
			&record.IdentityDetail.IdentityImage,
			&record.Symptoms,
			&record.Medications,
			&record.CreatedAt,
			&record.CreatedBy.Name,
			&record.CreatedBy.Nip,
			&record.CreatedBy.UserId,
		); err != nil {
			return nil, err
		}
		records = append(records, record)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return records, nil
}
