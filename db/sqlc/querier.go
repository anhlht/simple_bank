// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0

package db

import (
	"context"
)

type Querier interface {
	AddAccountsBalance(ctx context.Context, arg AddAccountsBalanceParams) (Account, error)
	CreateAccounts(ctx context.Context, arg CreateAccountsParams) (Account, error)
	CreateEntries(ctx context.Context, arg CreateEntriesParams) (Entry, error)
	CreateTransfers(ctx context.Context, arg CreateTransfersParams) (Transfer, error)
	CreateUsers(ctx context.Context, arg CreateUsersParams) (User, error)
	DeleteAccounts(ctx context.Context, id int64) error
	DeleteEntries(ctx context.Context, id int64) error
	DeleteTransfers(ctx context.Context, id int64) error
	GetAccounts(ctx context.Context, id int64) (Account, error)
	GetAccountsForUpdate(ctx context.Context, id int64) (Account, error)
	GetEntries(ctx context.Context, id int64) (Entry, error)
	GetTransfers(ctx context.Context, id int64) (Transfer, error)
	GetUsers(ctx context.Context, username string) (User, error)
	ListAccounts(ctx context.Context, arg ListAccountsParams) ([]Account, error)
	ListEntries(ctx context.Context, arg ListEntriesParams) ([]Entry, error)
	ListTransfers(ctx context.Context, arg ListTransfersParams) ([]Transfer, error)
	UpdateAccountsBalance(ctx context.Context, arg UpdateAccountsBalanceParams) (Account, error)
	UpdateEntriesAmount(ctx context.Context, arg UpdateEntriesAmountParams) (Entry, error)
	UpdateTransfersAmount(ctx context.Context, arg UpdateTransfersAmountParams) (Transfer, error)
}

var _ Querier = (*Queries)(nil)
