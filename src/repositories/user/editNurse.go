package userRepository

import (
	"hello-nurse/src/constants"
	entities "hello-nurse/src/entities/user"
	"log"

	"github.com/lib/pq"
)

func (dbase *controllerUser) EditNurse(params *entities.NurseEditParams) error {
	rows, err := dbase.DB.Exec("UPDATE users SET nip = $1, name = $2 WHERE id = $3 AND role = $4",
		params.Nip,
		params.Name,
		params.UserId,
		constants.ROLE_NURSE)

	if err != nil {
		if err, ok := err.(*pq.Error); ok {
			log.Printf("pg error %s", err.Error())
			if err.Code == constants.UniqueViolationExistData {
				return constants.ErrConflict
			}
		}

		log.Printf("Error editing nurse: %s", err)
		return err
	}

	rowsAffected, err := rows.RowsAffected()
	if err != nil {
		log.Printf("Error checking row affected: %s", err)
		return constants.ErrInternalServer
	}

	if rowsAffected == 0 {
		log.Printf("Error no row affected: %s", err)
		return constants.ErrNotFound

	}

	return nil
}
