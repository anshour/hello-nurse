package userRepository

import (
	"hello-nurse/src/constants"
	entities "hello-nurse/src/entities/user"
	"log"
)

func (i *controllerUser) CreateNurseAccess(params *entities.NurseRegisterAccess) error {

	res, err := i.DB.Exec(`UPDATE users SET password = $1, has_access = $2 WHERE id = $3`,
		params.Password, true, params.UserId,
	)

	if err != nil {
		log.Printf("Error giving access to nurse: %s", err)
		return err
	}

	rowsAffected, err := res.RowsAffected()
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
