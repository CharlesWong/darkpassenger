// package sqlite

// import (
// 	"database/sql"
// 	_ "github.com/mattn/go-sqlite3"
// )

// func InitDB(filepath string) *sql.DB {
// 	db, err := sql.Open("sqlite3", filepath)
// 	if err != nil {
// 		panic(err)
// 	}
// 	if db == nil {
// 		panic("db nil")
// 	}
// 	return db
// }

// func CreateTable(db *sql.DB) {
// 	// create table if not exists
// 	sqlTable := `
// 	CREATE TABLE IF NOT EXISTS items(
// 		Id TEXT NOT NULL PRIMARY KEY,
// 		Name TEXT,
// 		Phone TEXT,
// 		InsertedDatetime DATETIME
// 	);
// 	`

// 	_, err := db.Exec(sqlTable)
// 	if err != nil {
// 		panic(err)
// 	}
// }

// func StoreItem(db *sql.DB, items []TestItem) {
// 	sql_additem := `
// 	INSERT OR REPLACE INTO items(
// 		Id,
// 		Name,
// 		Phone,
// 		InsertedDatetime
// 	) values(?, ?, ?, CURRENT_TIMESTAMP)
// 	`

// 	stmt, err := db.Prepare(sql_additem)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer stmt.Close()

// 	for _, item := range items {
// 		_, err2 := stmt.Exec(item.Id, item.Name, item.Phone)
// 		if err2 != nil {
// 			panic(err2)
// 		}
// 	}
// }

// func ReadItem(db *sql.DB) []TestItem {
// 	sql_readall := `
// 	SELECT Id, Name, Phone FROM items
// 	ORDER BY datetime(InsertedDatetime) DESC
// 	`

// 	rows, err := db.Query(sql_readall)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer rows.Close()

// 	var result []TestItem
// 	for rows.Next() {
// 		item := TestItem{}
// 		err2 := rows.Scan(&item.Id, &item.Name, &item.Phone)
// 		if err2 != nil {
// 			panic(err2)
// 		}
// 		result = append(result, item)
// 	}
// 	return result
// }
// func Add(user *model.User) error {
// 	return nil
// }

// func Login(user *model.User) (*model.UserSession, error) {
// 	return nil, nil
// }

// func Logout(user *model.User) error {
// 	return nil
// }

// func Update(user *model.User) error {
// 	return nil
// }

// func Enable(user *model.User) error {
// 	return nil
// }

// func Disable(user *model.User) error {
// 	return nil
// }
