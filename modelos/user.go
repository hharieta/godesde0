package modelos

import (
	"time"
)

type User struct {
	Id        int
	Name      string
	CreatedAt time.Time
	Status    bool
}

func (this *User) AddUser(id int, name string, createdAt time.Time, status bool) {

	this.Id = id
	this.Name = name
	this.CreatedAt = createdAt
	this.Status = status
}
