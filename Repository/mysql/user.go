package mysql

import "database/sql"

func (d DB) IsPhoneNumberUnique(phoneNumber string) (bool, error) {
	sql.Drivers()
}
