package db

import (
	"context"
	"testing"
	"time"

	"github.com/renatoviolin/simplebank/util"

	"github.com/stretchr/testify/require"
)

func createRandomEntry(t *testing.T) Entry {
	acc := createRandomAccount(t)

	arg := CreateEntryParams{
		AccountID: acc.ID,
		Amount:    util.RandomMoney(),
	}

	entry, err := testQueries.CreateEntry(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, entry)

	require.Equal(t, arg.Amount, entry.Amount)
	require.Equal(t, arg.AccountID, acc.ID)

	require.NotZero(t, entry.ID)
	require.NotZero(t, entry.CreatedAt)
	return entry
}

func TestCreateEntry(t *testing.T) {
	createRandomEntry(t)
}

func TestGetEntry(t *testing.T) {
	ent1 := createRandomEntry(t)
	ent2, err := testQueries.GetEntry(context.Background(), ent1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, ent2)

	require.Equal(t, ent1.ID, ent2.ID)
	require.Equal(t, ent1.Amount, ent2.Amount)
	require.WithinDuration(t, ent1.CreatedAt, ent2.CreatedAt, time.Second)
}

func TestListEntry(t *testing.T) {
	acc := createRandomAccount(t)
	for i := 0; i < 10; i++ {
		arg := CreateEntryParams{
			AccountID: acc.ID,
			Amount:    util.RandomMoney(),
		}
		entry, err := testQueries.CreateEntry(context.Background(), arg)
		require.NoError(t, err)
		require.NotEmpty(t, entry)
	}

	arg := ListEntriesParams{
		AccountID: acc.ID,
		Limit:     5,
		Offset:    5,
	}

	accounts, err := testQueries.ListEntries(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, accounts, 5)

	for _, acc := range accounts {
		require.NotEmpty(t, acc)
	}
}
