package author

import (
	"fmt"
	"time"
)

type Author struct {
	Name      string
	Email     string
	Timestamp time.Time
}

func New(name string, email string, timestamp time.Time) *Author {
	return &Author{
		Name:      name,
		Email:     email,
		Timestamp: timestamp,
	}
}

func (a *Author) GetData() string {
	formatterTime := a.Timestamp.Format(time.UnixDate)
	return fmt.Sprintf("%s <%s> %s", a.Name, a.Email, formatterTime)
}

func (a *Author) AssignOid(oid string) {}
