package user

import (
	"crypto/md5"
	"errors"
	"fmt"
	db "github.com/charleswong/darkpassenger/db/sqlite"
	"github.com/charleswong/darkpassenger/model"
	"sync"
	"time"
)

var (
	userLock = &sync.Mutex{}
)

func SignUp(user *model.User) error {
	return db.AddUser(user)
}

func Login(user *model.User, session *model.UserSession) (*model.User, *model.UserSession, error) {
	if user == nil {
		return nil, nil, errors.New("Invalid user to login.")
	}
	stored, err := db.GetUser(user.Id)
	if err != nil {
		return nil, nil, err
	}
	if !stored.IsEnabled || stored.MaxConn <= 0 {
		return nil, nil, errors.New("User account is disabled.")
	}
	if stored.ExpiredTimestamp <= time.Now() {
		return nil, nil, errors.New("User account is expired.")
	}
	if session.IP == "" {
		return nil, nil, errors.New("Invalid IP.")
	}
	if stored != nil && stored.PwdHash == user.PwdHash {
		userLock.Lock()
		defer userLock.Unlock()
		// Check session limits.
		sessions := db.GetSessionsByUserId(user.Id)
		for len(sessions) >= stored.MaxConn {
			db.DelSession(sessions[0].Token)
			sessions = sessions[1:]
		}
		// Create a new session.
		session.Id = stored.Id
		session.LoginTimestamp = time.Now()
		session.Traffic = 0
		session.Token = newSessionToken(stored.Id)
		db.AddSession(session)
		db.AddSessionHistory(session)
		return user, session, nil
	} else {
		return nil, nil, errors.New("Username or password is incorrect. ", err)
	}
}

func Logout(user *model.User, session *model.UserSession) error {
	db.DelSession(session.Token)
	return nil
}

func TopUp(credit *model.UserCredit) error {
	if !verifyAdminToken(credit.AdminToken) {
		return errors.New("Invalid operation.")
	}
	user, err := db.GetUser(credit.Id)
	if err != nil {
		return err
	}
	db.AddCreditHistory(credit)
	user.ExpiredTimestamp = time.Now().Unix() + 3600*credit.Credit
	return Update(user)
}

func Update(user *model.User) error {
	if !verifyAdminToken(user.AdminToken) {
		return errors.New("Invalid operation.")
	}
	return db.UpdateUser(user)
}

func Enable(user *model.User) error {
	if !verifyAdminToken(user.AdminToken) {
		return errors.New("Invalid operation.")
	}
	stored, err := db.GetUser(user.Id)
	if err != nil {
		return err
	}
	stored.IsEnabled = true
	return db.UpdateUser(user)
}

func Disable(user *model.User) error {
	if !verifyAdminToken(user.AdminToken) {
		return errors.New("Invalid operation.")
	}
	stored, err := db.GetUser(user.Id)
	if err != nil {
		return err
	}
	stored.IsEnabled = false
	return db.UpdateUser(user)
}

func newSessionToken(id int64) string {
	identity := fmt.Sprintf("%d-%d", time.Now().UnixNano())
	return fmt.Sprintf("%X", md5.Sum([]byte(identity)))
}

func verifyAdminToken(token string) bool {
	return token == config.AdminToken
}
