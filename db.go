/**
	DB Interface
*/

package main

import (
	"fmt"
	"log"
	"crypto/sha1"
	"database/sql"
	_ "github.com/lib/pq"
)

type Flag_Record struct {
	Id int `json:"id"`
	Hash string `json:"hash"`
	Stub string `json:"stub"`
	Name string `json:"name"`
}

type User_Record struct {
	Id int `json:"id"`
	Hash string `json:"hash"`
	Stub string `json:"stub"`
	Name string `json:"name"`
}

// Maybe make a Global DB Conection Here?

func db_init() *sql.DB {

	db, err := sql.Open("postgres", "user=ffs dbname=ffs sslmode=disable")
	//db, err := sql.Open("postgres", Config.Database.link)
	//db, err := sql.Open("postgres", "user=ffs dbname=ffs sslmode=disable")
	if err != nil {
		panic("Failed to connect to the database")
	}

	return db
}



/**
	Find the Flag or Returns Null
*/
func DB_find_flag(stub string) Flag_Record {

	log.Println("DB_find_flag()")

	db := db_init()

	hash0 := sha1.Sum([]byte(stub))
	hashS := fmt.Sprintf("% x", hash0)

	var F Flag_Record;

	sql := "SELECT id, hash, name FROM f WHERE (hash = $1 OR stub = $2)"
	err := db.QueryRow(sql, hashS, stub).Scan(&F.Id, &F.Stub, &F.Name)
	if err != nil {
		log.Println(err)
		panic("Failed to query the Database")
	}

	if 0 == F.Id {

		// Insert
		sql = "INSERT INTO f (hash, name, stub) VALUES ($1, $2, $3) RETURNING id"
		err := db.QueryRow(sql, hashS, stub, stub).Scan(&F.Id)
		if err != nil {
			log.Println(err)
			panic("Failed to create the new Flag")
		}

		F.Hash = hashS
		F.Name = stub
		F.Stub = stub

	}

	return F

}

func DB_flag_list() *sql.Rows {

	log.Println("DB_flag_list()")

	db := db_init()

	sql := "SELECT id, stub, name, hash FROM f ORDER BY stub"

	rows, err := db.Query(sql)
	if err != nil {
		log.Print(err)
		panic("Failed to fetch flag_list")
	}

	return rows
}

func DB_user_list() *sql.Rows {

	log.Println("DB_user_list()")

	db := db_init()

	sql := "SELECT id, hash, name FROM u ORDER BY name"

	rows, err := db.Query(sql)
	if err != nil {
		log.Print(err)
		panic("Failed to fetch user_list")
	}

	return rows
}

/**
	Find the User or Returns Null
*/
func DB_user_pick(hash string) User_Record {

	log.Println("DB_find_user(%s)", hash)

	db := db_init()

	var u User_Record

	sql := "SELECT id, stub, name FROM u WHERE hash = $1"
	err := db.QueryRow(sql, hash).Scan(&u.Id, &u.Stub, &u.Name)
	if err != nil {
		log.Print(err)
		panic("Failed to DB_user_select")
	}

	if (0 == u.Id) {

		u.Name = hash;
		u.Stub = hash;
		u.Hash = hash;

		// Insert
		sql = "INSERT INTO u (name, stub, hash) VALUES ($1, $2, $3) RETURNING id"
		err := db.QueryRow(sql, u.Name, u.Stub, u.Hash).Scan(&u.Id)
		if err != nil {
			panic("Failed to create the new User")
		}

	}

	return u
}

func DB_user_flag_list(user int) *sql.Rows {

	db := db_init()

	sql := "SELECT f.stub, f.name, fu.v AS data FROM f JOIN fu ON f.id = fu.f_id WHERE fu.u_id = $1"

	res, err := db.Query(sql, user)
	if err != nil {
		log.Print(err)
		panic("Failed to fetch flag_list")
	}

	return res
}


func DB_user_flag(user string, flag string) User_Record {

	db := db_init()

	sql := "SELECT * FROM u JOIN uf ON u.id = uf.user_id JOIN f ON uf.flag_id = f.id WHERE u.stub = $1 AND f.stub = $2"
	row := db.QueryRow(sql, user, flag)
	log.Println(row)

	//age := 21
	//rows, err := db.Query("SELECT name FROM users WHERE age = $1", age)

	u := User_Record{
		Id: 2,
		Stub: user,
		Name: user,
	}

	return u
}

