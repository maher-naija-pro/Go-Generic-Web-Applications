package dbrepo


import (
	
	"context"
	"web_server/internal/models"
	"time"
	 "golang.org/x/crypto/bcrypt"

)

// search a user into the database
func (m *postgresDBRepo) AuthUser (email, testPaswword string) (int,string, error) {
   	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var id int  

   var hashpassword string

   query := `select (id, password)
			from users where email = $1`

   row := m.DB.QueryRowContext(ctx, query,email)
   err := row.Scan(&id,&hashpassword)
	if err != nil  {
		return id, "", err
	}
	err = bcrypt.CompareHashAndPassword([]byte(hashpassword),[]byte(testPaswword))
	if err == bcrypt.ErrMismatchedHashAndPassword  {
		return 0,"",  err
	} else 
	{
      return 0, "", err
	}

	return id, hashpassword, nil
}


// search a user into the database
func (m *postgresDBRepo) GetUserByID (id int) (models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `select (id,first_name,last_name,email, password,access_level) 
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
	return 0 , nil
}