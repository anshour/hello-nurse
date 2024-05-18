package userRepository

import (
	"hello-nurse/src/constants"
	entities "hello-nurse/src/entities/user"
	"log"
)

func (dbase *controllerUser) DeleteNurse(params *entities.NurseDeleteParams) error {
	rows, err := dbase.DB.Exec("DELETE FROM users WHERE id = $1 AND role = $2", params.UserId, constants.ROLE_NURSE)

	if err != nil {
		log.Printf("Error deleting user: %s", err)
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
