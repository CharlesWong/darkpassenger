package sqlite

import (
	"github.com/charleswong/darkpassenger/model"
	"log"
	"testing"
	"time"
)

func TestAccount(t *testing.T) {
	dataFile := "test_data.dat"
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	err := InitDB(dataFile)
	if err != nil {
		t.Fatal(err)
	}

	defer CloseDB()

	user := &model.User{
		Id:               123,
		Email:            "user1@dp.com",
		PwdHash:          "qwertyuiop123",
		MaxConn:          10,
		IsEnabled:        true,
		CreatedTimestamp: time.Now().Unix(),
		ExpiredTimestamp: time.Now().Add(time.Hour).Unix(),
	}

	err = AddUser(user)
	if err != nil {
		t.Fatal(err)
	}

	user2, err := GetUser(user.Id)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(user2)

	session := &model.UserSession{
		Id:    123,
		Token: "token",
		IP:    "127.0.0.1",
	}
	err = AddSession(session)
	if err != nil {
		t.Fatal(err)
	}
	err = AddSessionHistory(session)
	if err != nil {
		t.Fatal(err)
	}

	sessions, err := GetSessionsByUserId(user.Id)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(sessions)

	if len(sessions) != 1 {
		t.Error("sessions is incorrect: ", sessions)
	} else {
		err = DelSession(sessions[0].Token)
		if err != nil {
			t.Fatal(err)
		}
	}

	sessions, err = GetSessionsByUserId(user.Id)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(sessions)

	userCredit := &model.UserCredit{
		Id:             123,
		Credit:         24 * 30,
		TopUpTimestamp: time.Now().Unix(),
		AdminToken:     "qwertyuiop123",
	}
	err = AddCreditHistory(userCredit)
	if err != nil {
		t.Fatal(err)
	}
}
