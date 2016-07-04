package sqlite

import (
	"database/sql"
	"fmt"
	"github.com/charleswong/darkpassenger/model"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

var (
	db *sql.DB
)

func InitDB(file string) (err error) {
	log.Println("Open db file ", file)
	db, err = sql.Open("sqlite3", file)
	if err != nil {
		log.Println("Created db file")
		err = createDBFile(file)
		if err != nil {
			return err
		}
	}
	if db == nil {
		panic("db nil")
	}
	return nil
}

func createDBFile(file string) (err error) {
	// create table if not exists
	sqlTable := `
CREATE TABLE IF NOT EXISTS user(
	Id INTEGER NOT NULL PRIMARY KEY,
	Email TEXT,
	PwdHash TEXT,
	MaxConn INTEGER,
	CreatedTimestamp INTEGER,
	ExpiredTimestamp INTEGER,
	IsEnabled INTEGER
);

CREATE TABLE IF NOT EXISTS session(
	Token TEXT NOT NULL PRIMARY KEY,
	Id INTEGER,
	LoginTimestamp INTEGER,
	Traffic INTEGER,
	IP TEXT
);

CREATE TABLE IF NOT EXISTS session_history(
	Token TEXT NOT NULL PRIMARY KEY,
	Id INTEGER,
	LoginTimestamp INTEGER,
	Traffic INTEGER,
	IP TEXT,
	Expired INTEGER
);

CREATE TABLE IF NOT EXISTS user_credit_history(
	Id INTEGER NOT NULL PRIMARY KEY,
	Credit INTEGER,
	TopUpTimestamp INTEGER
);
`
	_, err = db.Exec(sqlTable)
	if err != nil {
		panic(err)
	}
	return nil
}

func AddUser(user *model.User) (err error) {
	sqlStmt := `
		INSERT OR REPLACE INTO user(
			Id,
			Email,
			PwdHash,
			MaxConn,
			CreatedTimestamp,
			ExpiredTimestamp,
			IsEnabled
		) values(?, ?, ?, ?, ?, ?, ?)
		`

	stmt, err := db.Prepare(sqlStmt)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		user.Id,
		user.Email,
		user.PwdHash,
		user.MaxConn,
		user.CreatedTimestamp,
		user.ExpiredTimestamp,
		user.IsEnabled,
	)

	if err != nil {
		return err
	}
	return nil
}

func UpdateUser(user *model.User) (err error) {
	return AddUser(user)
}

func GetUser(id int64) (*model.User, error) {
	sqlStmt := fmt.Sprintf(`
		SELECT * FROM user WHERE user.Id = %d
		`, id)

	rows, err := db.Query(sqlStmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	user := &model.User{}
	for rows.Next() {
		err = rows.Scan(
			&user.Id,
			&user.Email,
			&user.PwdHash,
			&user.MaxConn,
			&user.CreatedTimestamp,
			&user.ExpiredTimestamp,
			&user.IsEnabled,
		)
		if err != nil {
			return nil, err
		}
		return user, nil
	}
	return nil, nil
}

func GetSessionsByUserId(id int64) ([]*model.UserSession, error) {
	sqlStmt := fmt.Sprintf(`
		SELECT * FROM session WHERE session.Id = %d
		`, id)

	rows, err := db.Query(sqlStmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	sessions := make([]*model.UserSession, 0)
	for rows.Next() {
		session := &model.UserSession{}
		err = rows.Scan(
			&session.Token,
			&session.Id,
			&session.LoginTimestamp,
			&session.Traffic,
			&session.IP,
		)
		if err != nil {
			return sessions, err
		}
		sessions = append(sessions, session)
	}
	return sessions, nil
}

func AddSession(session *model.UserSession) (err error) {
	return addSessionImpl(session, "session")
}

func AddSessionHistory(session *model.UserSession) (err error) {
	return addSessionImpl(session, "session_history")
}

func addSessionImpl(session *model.UserSession, table string) (err error) {
	sqlStmt := fmt.Sprintf(`
		INSERT OR REPLACE INTO %s(
			Token,
			Id,
			LoginTimestamp,
			Traffic,
			IP
		) values(?, ?, ?, ?, ?)
		`, "session")

	stmt, err := db.Prepare(sqlStmt)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		session.Token,
		session.Id,
		session.LoginTimestamp,
		session.Traffic,
		session.IP)

	if err != nil {
		return err
	}
	return nil
}

func DelSession(token string) (err error) {
	sqlStmt := fmt.Sprintf(`
		DELETE session where token = %s`, token)

	stmt, err := db.Prepare(sqlStmt)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec()

	if err != nil {
		return err
	}
	return nil
}

func AddCreditHistory(userCredit *model.UserCredit) (err error) {
	sqlStmt := `
		INSERT OR REPLACE INTO user_credit_history(
			Id,
			Credit,
			TopUpTimestamp
		) values(?, ?, ?)
		`

	stmt, err := db.Prepare(sqlStmt)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		userCredit.Id,
		userCredit.Credit,
		userCredit.TopUpTimestamp,
	)
	if err != nil {
		return err
	}
	return nil
}
