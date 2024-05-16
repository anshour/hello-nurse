package userRepository

import (
	"errors"
	entities "hello-nurse/src/entities/user"
	"log"
)

func (dbase *controllerUser) EditNurse(params *entities.NurseEditParams) error {
	rows, err := dbase.DB.Exec("UPDATE users SET nip = $1, name = $2 WHERE id = $3", params.Nip, params.Name, params.UserId)

	if err != nil {
		log.Printf("Error editing nurse: %s", err)
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
