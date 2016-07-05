package account

import (
	db "github.com/charleswong/darkpassenger/db/sqlite"
	"github.com/charleswong/darkpassenger/model"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"testing"
	"time"
)

func TestAccountService(t *testing.T) {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	err := InitConfig("dp_test.conf")
	if err != nil {
		t.Fatal(err)
	}
	err = db.InitDB(GetConfig().DataFile)
	if err != nil {
		t.Fatal(err)
	}

	defer db.CloseDB()

	go StartAccountService()

	c := NewAccountClient(config.ListenAddr)
	c.Start()

	user := &model.User{
		Id:               123,
		Email:            "user1@dp.com",
		PwdHash:          "qwertyuiop123",
		MaxConn:          10,
		IsEnabled:        true,
		CreatedTimestamp: time.Now().Unix(),
		ExpiredTimestamp: time.Now().Add(time.Hour).Unix(),
	}

	_, err = c.Client.SignUp(context.Background(), user)
	if err != nil {
		t.Error(err)
	}
	session := &model.UserSession{
		IP: "127.0.0.1",
	}
	user.Sessions = append(user.Sessions, session)
	r, err := c.Client.LogIn(context.Background(), user)
	if err != nil {
		t.Error(err)
	}
	t.Log(r)

	userCredit := &model.UserCredit{
		Id:             123,
		Credit:         24 * 30,
		TopUpTimestamp: time.Now().Unix(),
		AdminToken:     "qwertyuiop123",
	}
	_, err = c.Client.TopUp(context.Background(), userCredit)
	if err != nil {
		t.Error(err)
	}
	_, err = c.Client.LogOut(context.Background(), user)
	if err != nil {
		t.Error(err)
	}
}

type AccountClient struct {
	Conn       *grpc.ClientConn
	Client     model.AccountServiceClient
	ServerAddr string
}

func NewAccountClient(serverAddr string) *AccountClient {
	return &AccountClient{
		ServerAddr: serverAddr,
	}
}

func (client *AccountClient) Start() {
	// if client.Conn == nil || client.State == grpc.Idle  || client.State == grpc.Shutdown
	var err error
	client.Conn, err = grpc.Dial(client.ServerAddr, grpc.WithInsecure())
	if err != nil {
		log.Println("did not connect: %v", err)
	}
	client.Client = model.NewAccountServiceClient(client.Conn)
}

func (client *AccountClient) Stop() {
	defer client.Conn.Close()
}
