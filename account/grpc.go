package account

import (
	"github.com/deepglint/deep-data/model"
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
	return nil, SignUp(req)
}

func (s *AccountService) LogIn(ctx context.Context, req *model.User) (*model.NullMessage, error) {
	if len(req.Sessions) != 1 {
		return nil.errors.New("No session provided.")
	}
	user, session := LogIn(req, req.Sessions[0])
	user.Sessions = []*UserSession{session}
	return nil, user
}

func (s *AccountService) LogOut(ctx context.Context, req *model.User) (*model.NullMessage, error) {
	if len(req.Sessions) != 1 {
		return nil.errors.New("No session provided.")
	}
	return nil, LogOut(req, req.Sessions[0])
}

func (s *AccountService) TopUp(ctx context.Context, req *model.UserCredit) (*model.NullMessage, error) {
	return nil, TopUp(req)
}

func (s *AccountService) Update(ctx context.Context, req *model.User) (*model.NullMessage, error) {
	return nil, Update(req)
}

func (s *AccountService) Enable(ctx context.Context, req *model.User) (*model.NullMessage, error) {
	return nil, Enable(req)
}

func (s *AccountService) Disable(ctx context.Context, req *model.User) (*model.NullMessage, error) {
	return nil, Disable(req)
}

func StartAccountService(addr string) {
	// Don't enter this servcie twice since it will listen the same port twice.
	select {
	case accountServiceLock <- 1:
		defer func() {
			<-accountServiceLock
		}()
	default:
		return
	}

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Println("failed to listen: %v", err)
		return
	}
	s := grpc.NewServer()
	model.RegisterAccountServiceServer(s, &AccountService{})
	log.Println("Account Service started.")
	s.Serve(lis)
}
