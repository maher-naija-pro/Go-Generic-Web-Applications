package dbrepo


import (
	"context"
	"web_server/pkg/models"
	"time"
)

func (m *postgresDBRepo) AllUsers() bool {
	return true
}


// InsertReservation inserts a user into the database
func (m *postgresDBRepo) InsertUser(res models.User) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	

	stmt := `insert into users (id,username, password) 
			values ($1, $2, $3) returning id`

	_, err := m.DB.ExecContext(ctx, stmt,
        2,
		res.Username,
		res.Password,
	)

	

	if err != nil {
		return 0, err
	}

	return 0 ,nil
}