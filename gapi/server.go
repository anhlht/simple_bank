package gapi

import (
	db "bank_account/db/sqlc"
	"bank_account/pb"
	"bank_account/token"
	"bank_account/util"
	"fmt"
)

// Server serves HTTP request for banking service.
type Server struct {
	pb.UnimplementedBankAcountServer
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
}

func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMakter(config.TokenSymmectricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}
	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}

	return server, nil
}
