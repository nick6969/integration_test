package mysql

import "time"

type Customer struct {
	ID        uint       
	Username  string     
	Password  string     `json:"-"`
	CreatedAt time.Time  
	UpdatedAt time.Time  
	DeletedAt *time.Time
}

func (d *Database) FindUserWithUsername(name string) (user Customer, err error) {
	err = d.Where(&Customer{Username: name}).First(&user).Error
	return
}
