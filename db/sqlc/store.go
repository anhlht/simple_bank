package db

import (
	"context"
	"database/sql"
	"fmt"
)

// Store provide all functions to excute db quries and transactions
type Store struct {
	*Queries
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db:      db,
		Queries: New(db),
	}
}

func (store *Store) exexTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return err
	}
	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb %v", err, rbErr)
		}
		return err
	}
	return tx.Commit()
}

// TransferTxParams contains the input parammeters of the transfer transaction
// type TransferTxParams struct {
// 	FromAccountID int64 `json:"from_account_id"`
// 	ToAccountID   int64 `json:"to_account_id"`
// 	Amount        int64 `json:"amount"`
// }

// TransferTxResult is the result of the transfer transaction
type TransferTxResult struct {
	Transfer    Transfer `json:"transfer"`
	FromAccount Account  `json:"from_account"`
	ToAccount   Account  `json:"to_account"`
	FromEntry   Entry    `json:"from_entry"`
	ToEntry     Entry    `json:"to_entry"`
}

// TransferTx performs a money transfer from one account to the other.
// It creates a transfer record, add account entries, and update accounts's balance within a single database transaction

func (store *Store) TransferTx(ctx context.Context, args CreateTransfersParams) (TransferTxResult, error) {
	var result TransferTxResult
	err := store.exexTx(ctx, func(q *Queries) error {
		var err error

		result.Transfer, err = q.CreateTransfers(ctx, args)
		if err != nil {
			return err
		}

		result.FromEntry, err = q.CreateEntries(ctx, CreateEntriesParams{
			AccountID: args.FromAccountID,
			Amount:    -args.Amount,
		})

		if err != nil {
			return err
		}

		result.ToEntry, err = q.CreateEntries(ctx, CreateEntriesParams{
			AccountID: args.ToAccountID,
			Amount:    args.Amount,
		})

		if err != nil {
			return err
		}

		if args.FromAccountID < args.ToAccountID {
			result.FromAccount, err = q.AddAccountsBalance(ctx, AddAccountsBalanceParams{
				ID:     args.FromAccountID,
				Amount: -args.Amount,
			})

			result.ToAccount, err = q.AddAccountsBalance(ctx, AddAccountsBalanceParams{
				ID:     args.ToAccountID,
				Amount: args.Amount,
			})
		} else {
			result.ToAccount, err = q.AddAccountsBalance(ctx, AddAccountsBalanceParams{
				ID:     args.ToAccountID,
				Amount: args.Amount,
			})

			result.FromAccount, err = q.AddAccountsBalance(ctx, AddAccountsBalanceParams{
				ID:     args.FromAccountID,
				Amount: -args.Amount,
			})
		}

		return nil
	})
	return result, err
}
