package userRepository

import (
	"fmt"
	entities "hello-nurse/src/entities/user"
	"log"
	"reflect"
	"strconv"
	"strings"
)

func (i *controllerUser) List(filters *entities.UserListFilter) ([]*entities.UserListResponse, error) {
	baseQuery := "SELECT id, nip, name, created_at FROM users WHERE true"
	params := []interface{}{}

	n := (&entities.UserListFilter{})

	if !reflect.DeepEqual(filters, n) {
		conditions := []string{}

		if filters.Id != "" {
			conditions = append(conditions, "id = $"+strconv.Itoa(len(params)+1))
			params = append(params, filters.Id)
		}
		if filters.Name != "" {
			conditions = append(conditions, "name LIKE $"+strconv.Itoa(len(params)+1))
			params = append(params, "%"+filters.Name+"%")
		}
		if filters.Role != "" {
			conditions = append(conditions, "role = $"+strconv.Itoa(len(params)+1))
			params = append(params, filters.Role)
		}

		//TODO: FIX THIS
		if filters.Nip != 0 {
			conditions = append(conditions, fmt.Sprintf("CAST(nip AS VARCHAR LIKE $%d", len(params)+1))
			params = append(params, "%"+strconv.Itoa(filters.Nip)+"%")
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
	rows, err := i.DB.Query(baseQuery, params...)
	if err != nil {
		log.Printf("Error finding user: %s", err)
		return nil, err
	}
	defer rows.Close()

	var users = make([]*entities.UserListResponse, 0)
	for rows.Next() {
		user := new(entities.UserListResponse)
		if err := rows.Scan(
			&user.Id,
			&user.Nip,
			&user.Name,
			&user.CreatedAt,
		); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}
