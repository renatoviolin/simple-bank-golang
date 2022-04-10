package db

import (
	"context"
	"testing"
	"time"

	"github.com/renatoviolin/simplebank/util"

	"github.com/stretchr/testify/require"
)

func createRandomTransfer(t *testing.T) Transfer {
	acc_from := createRandomAccount(t)
	acc_to := createRandomAccount(t)

	arg := CreateTransferParams{
		FromAccountID: acc_from.ID,
		ToAccountID:   acc_to.ID,
		Amount:        util.RandomMoney(),
	}

	transfer, err := testQueries.CreateTransfer(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, transfer)

	require.Equal(t, arg.Amount, transfer.Amount)
	require.Equal(t, arg.FromAccountID, acc_from.ID)
	require.Equal(t, arg.ToAccountID, acc_to.ID)

	require.NotZero(t, transfer.ID)
	require.NotZero(t, transfer.CreatedAt)
	return transfer
}

func TestCreateTransfer(t *testing.T) {
	createRandomTransfer(t)
}

func TestGetTransfer(t *testing.T) {
	t1 := createRandomTransfer(t)
	t2, err := testQueries.GetTransfer(context.Background(), t1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, t2)

	require.Equal(t, t1.ID, t2.ID)
	require.Equal(t, t1.Amount, t2.Amount)
	require.Equal(t, t1.FromAccountID, t2.FromAccountID)
	require.Equal(t, t1.ToAccountID, t2.ToAccountID)
	require.WithinDuration(t, t1.CreatedAt, t2.CreatedAt, time.Second)
}

func TestListTransfer(t *testing.T) {
	acc_from := createRandomAccount(t)
	acc_to := createRandomAccount(t)
	for i := 0; i < 10; i++ {
		arg := CreateTransferParams{
			FromAccountID: acc_from.ID,
			ToAccountID:   acc_to.ID,
			Amount:        util.RandomMoney(),
		}
		entry, err := testQueries.CreateTransfer(context.Background(), arg)
		require.NoError(t, err)
		require.NotEmpty(t, entry)
	}

	arg := ListTransfersParams{
		FromAccountID: acc_from.ID,
		ToAccountID:   acc_to.ID,
		Limit:         5,
		Offset:        5,
	}

	transfers, err := testQueries.ListTransfers(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, transfers, 5)

	for _, acc := range transfers {
		require.NotEmpty(t, acc)
	}
}
