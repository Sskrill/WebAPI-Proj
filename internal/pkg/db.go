package webAPIUsers

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
)

type CRUD interface {
	Insert(u User) error
	Get(id int) (User, error)
	Delete(id int) error
	Update(id int, u User) error
	GetAll() []User
}
type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type DataBase struct {
	DB *sql.DB
}

func NewDB() *DataBase {

	connStr := fmt.Sprintf("host=localhost port=5432 user=username password=password dbname=postgres sslmode=disable")
	db, err := sql.Open("postgres", connStr)
	if err != nil {

		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Println("Failed to connect to the database.")
		log.Fatal(err)
	}
	DtBs := DataBase{DB: db}
	log.Println("Succes")
	return &DtBs
}
func (d *DataBase) Get(id int) (User, error) {
	var count int
	user := User{}
	err := d.DB.QueryRow("SELECT COUNT(*) FROM users WHERE id=$1", id).Scan(&count)
	if err != nil || count == 0 {
		return user, errors.New("Not found user")
	}

	rows, err := d.DB.Query("SELECT id,name,age FROM users WHERE id=$1", id)
	if err != nil {
		log.Println("Not found user")
		return user, errors.New("Not found user")
	}

	for rows.Next() {
		err := rows.Scan(&user.Id, &user.Name, &user.Age)
		if err != nil {
			log.Println("error Scan")
			log.Fatal(err)
		}

	}
	return user, nil
}
func (d *DataBase) GetAll() []User {
	rows, err := d.DB.Query("SELECT id,name,age FROM users")
	if err != nil {
		log.Println("Not found users")
		log.Fatal(err)
	}
	user := User{}
	users := make([]User, 0)

	for rows.Next() {
		err := rows.Scan(&user.Id, &user.Name, &user.Age)
		if err != nil {
			log.Println("error Scan")
			log.Fatal(err)
		}
		users = append(users, user)
	}
	return users
}
func (d *DataBase) Insert(u User) error {

	_, err := d.DB.Exec("INSERT INTO users (id,name,age) VALUES ($1,$2,$3)", u.Id, u.Name, u.Age)
	if err != nil {
		log.Println("error Insert user")
		return err
	}
	return nil
}

func (d *DataBase) Update(id int, u User) error {
	var count int

	err := d.DB.QueryRow("SELECT COUNT(*) FROM users WHERE id=$1", id).Scan(&count)
	if err != nil || count == 0 {
		return errors.New("Not found user")
	}

	result, err := d.DB.Exec("UPDATE users SET name=$1,age=$2 WHERE id=$3 ", u.Name, u.Age, id)
	resultRows, _ := result.RowsAffected()
	if err != nil || resultRows == 0 {
		log.Println("error Update")
		return err
	}
	return nil
}

func (d *DataBase) Delete(id int) error {
	var count int

	err := d.DB.QueryRow("SELECT COUNT(*) FROM users WHERE id=$1", id).Scan(&count)
	if err != nil || count == 0 {
		return errors.New("Not found user")
	}
	result, err := d.DB.Exec("DELETE FROM users WHERE id=$1", id)
	resultRows, _ := result.RowsAffected()
	if err != nil || resultRows == 0 {
		log.Println("error Delete")
		return err
	}
	return nil
}
