package userRepository

import (
	"errors"
	entities "hello-nurse/src/entities/user"
	"log"
)

func (dbase *controllerUser) DeleteNurse(params *entities.NurseDeleteParams) error {
	rows, err := dbase.DB.Exec("DELETE FROM products WHERE id = $1", params.UserId)

	if err != nil {
		log.Printf("Error deleting user: %s", err)
		return err
	}

	rowsAffected, err := rows.RowsAffected()
	if err != nil {
		log.Printf("Error checking row affected: %s", err)
		return err
	}

	if rowsAffected == 0 {
		log.Printf("Error no row affected: %s", err)
		return errors.New("No Row Affected")

	}

	return nil
}
