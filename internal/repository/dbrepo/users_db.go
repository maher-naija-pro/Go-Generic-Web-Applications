package dbrepo

import (
	
	"context"
	"web_server/internal/models"
	"time"
	"log"
	"errors"
	 "golang.org/x/crypto/bcrypt"

)

// search a user into the database
func (m *postgresDBRepo) AuthUser (email, testPaswword string) (int, error) {
   	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var id int
    var pass string
	 
    query := `select id, password
			from users where email = $1`

    row := m.DB.QueryRowContext(ctx, query,email)
    err := row.Scan(&id ,&pass)
	if err != nil  {
		return id, errors.New("user not found on database")
	}
	err = bcrypt.CompareHashAndPassword([]byte(pass),[]byte(testPaswword))
	if err == bcrypt.ErrMismatchedHashAndPassword  {
		log.Println("password not match:",err)
		return 0, errors.New("password not match on database")
	} 
    log.Println("password ok:",err)
	return id, nil
}


// search a user into the database
func (m *postgresDBRepo) GetUserByID (id int) (models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `select id,first_name,last_name,email, password,access_level 
			from users where id = $1`

	row := m.DB.QueryRowContext(ctx, query,id)
	var u  models.User

	err := row.Scan(
		&u.ID,
		&u.FirstName,
		&u.LastName,
		&u.Email,
		&u.Password,
		&u.AccessLevel,
		&u.Password,
	)
	if err != nil {
		return u , nil
	}
return u, nil
}

// inserts a user into the database
func (m *postgresDBRepo) InsertUser(res models.User) (int, error) {					
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(res.Password), 12)

	stmt := `insert into users (id,first_name,last_name,email, password,access_level) 
			values ($1, $2, $3, $4,$5,$6) returning id`
			
	_, err := m.DB.ExecContext(ctx, stmt,
        3,
		res.FirstName,
		res.LastName,
		res.Email,
		hashedPassword,
		3,
	)
	if err != nil {
		 log.Println("user not inserted:",err)
		return  0, err
	}
	 log.Println("user inserted")
	return 0 , nil
}