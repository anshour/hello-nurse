package userController

import (
	"fmt"
	"hello-nurse/src/http/models/user"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

func (i *V1User) UserList(c echo.Context) (err error) {
	baseQuery := "SELECT id, nip, name, created_at FROM users WHERE true"
	var params []interface{}
	var conditions, orders []string
	var limitQuery, offsetQuery, orderByQuery string

	if userId := c.QueryParam("userId"); userId != "" {
		conditions = append(conditions, fmt.Sprintf("userId = $%d", len(params)+1))
		params = append(params, userId)
	}

	if name := c.QueryParam("name"); name != "" {
		conditions = append(conditions, fmt.Sprintf("name = $%d", len(params)+1))
		params = append(params, name)
	}

	//TODO: FIX THIS
	// if nip := c.QueryParam("nip"); nip != "" {
	// 	conditions = append(conditions, fmt.Sprintf("nip LIKE $%d", len(params)+1))
	// 	params = append(params, "%"+nip+"%")
	// }

	if role := c.QueryParam("role"); role != "" {
		conditions = append(conditions, fmt.Sprintf("role = $%d", len(params)+1))
		params = append(params, role)
	}

	if limit := c.QueryParam("limit"); limit != "" {
		limitQuery = fmt.Sprintf("LIMIT $%d", len(params)+1)
		params = append(params, limit)
	} else {
		limitQuery = fmt.Sprintf("LIMIT $%d", len(params)+1)
		params = append(params, 20)
	}

	if offset := c.QueryParam("offset"); offset != "" {
		offsetQuery = fmt.Sprintf("OFFSET $%d", len(params)+1)
		params = append(params, offset)
	}

	if createdAt := c.QueryParam("createdAt"); createdAt != "" {
		if createdAt == "desc" {
			createdAt = "DESC"
		} else {
			createdAt = "ASC"
		}
		orders = append(orders, fmt.Sprintf("created_at %s", createdAt))
	} else {
		orders = append(orders, "created_at DESC")
	}

	if len(orders) > 0 {
		orderByQuery = "ORDER BY " + strings.Join(orders, ", ")
	}

	if len(conditions) > 0 {
		baseQuery += " AND " + strings.Join(conditions, " AND ")
	}

	if orderByQuery != "" {
		baseQuery += " " + orderByQuery
	}

	if offsetQuery != "" {
		baseQuery += " " + offsetQuery
	}

	if limitQuery != "" {
		baseQuery += " " + limitQuery
	}

	rows, err := i.DB.Query(baseQuery, params...)

	if err != nil {
		return c.JSON(http.StatusConflict, ErrorResponse{Message: err.Error()})
	}
	defer rows.Close()

	var users = make([]user.UserList, 0)
	for rows.Next() {
		var user user.UserList
		if err := rows.Scan(
			&user.UserId,
			&user.Nip,
			&user.Name,
			&user.CreatedAt,
		); err != nil {
			return c.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return c.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, SuccessResponse{
		Message: "User list successfully",
		Data:    users,
	})

}
