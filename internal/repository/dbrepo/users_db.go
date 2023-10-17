package dbrepo


import (
	
	"context"
	"web_server/internal/models"
	"time"

)



// InsertReservation inserts a user into the database
func (m *postgresDBRepo) InsertUser(res models.User) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	

	stmt := `insert into users (id,first_name,last_name,email, password,access_level) 
			values ($1, $2, $3, $4,$5,$6) returning id`

	_, err := m.DB.ExecContext(ctx, stmt,
        1,
		res.FirstName,
		res.LastName,
		res.Email,
		res.Password,
		1,
	)

	if err != nil {
		return 0, err
	}

	return 0 ,nil
}