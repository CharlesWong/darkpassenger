package account

import (
	"errors"
	"github.com/charleswong/darkpassenger/model"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"net"
)

var (
	accountServiceLock = make(chan int, 1)
)

type AccountService struct {
}

func (s *AccountService) SignUp(ctx context.Context, req *model.User) (*model.NullMessage, error) {
	return &model.NullMessage{}, SignUp(req)
}

func (s *AccountService) LogIn(ctx context.Context, req *model.User) (*model.User, error) {
	if len(req.Sessions) != 1 {
		return &model.User{}, errors.New("No session provided.")
	}
	user, session, err := LogIn(req, req.Sessions[0])
	if err != nil {
		return &model.User{}, err
	}
	user.Sessions = []*model.UserSession{session}
	return user, nil
}

func (s *AccountService) LogOut(ctx context.Context, req *model.User) (*model.NullMessage, error) {
	if len(req.Sessions) != 1 {
		return &model.NullMessage{}, errors.New("No session provided.")
	}
	return &model.NullMessage{}, LogOut(req, req.Sessions[0])
}

func (s *AccountService) TopUp(ctx context.Context, req *model.UserCredit) (*model.NullMessage, error) {
	return &model.NullMessage{}, TopUp(req)
}

func (s *AccountService) Update(ctx context.Context, req *model.User) (*model.NullMessage, error) {
	return &model.NullMessage{}, Update(req)
}

func (s *AccountService) Enable(ctx context.Context, req *model.User) (*model.NullMessage, error) {
	return &model.NullMessage{}, Enable(req)
}

func (s *AccountService) Disable(ctx context.Context, req *model.User) (*model.NullMessage, error) {
	return &model.NullMessage{}, Disable(req)
}

func StartAccountService() {
	if config.ListenAddr == "" {
		return
	}
	// Don't enter this servcie twice since it will listen the same port twice.
	select {
	case accountServiceLock <- 1:
		defer func() {
			<-accountServiceLock
		}()
	default:
		return
	}

	lis, err := net.Listen("tcp", config.ListenAddr)
	if err != nil {
		log.Println("failed to listen: %v", err)
		return
	}
	s := grpc.NewServer()
	model.RegisterAccountServiceServer(s, &AccountService{})
	log.Println("Account Service started.")
	s.Serve(lis)
}
