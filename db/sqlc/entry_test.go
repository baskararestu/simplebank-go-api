package db

import (
	"context"
	"testing"
	"time"

	"github.com/baskararestu/simplebank-go-api/util"
	"github.com/stretchr/testify/require"
)

func createRandomEntry(t *testing.T) Entry {
	account1 := createRandomAccount(t)
	arg := CreateEntriesParams{
		Amount:    util.RandomAmountEntries(),
		CreatedAt: time.Now().Local(),
		AccountID: account1.ID,
	}
	entry, err := testQueries.CreateEntries(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account1)
	require.NotEmpty(t, entry)

	require.NotZero(t, entry.ID)
	require.NotZero(t, entry.CreatedAt)

	require.Equal(t, arg.AccountID, entry.AccountID)
	require.Equal(t, arg.Amount, entry.Amount)
	require.WithinDuration(t, arg.CreatedAt, entry.CreatedAt, time.Second)

	return entry
}

func TestCreateEntry(t *testing.T) {
	createRandomEntry(t)
}

func TestGetEnteries(t *testing.T) {
	entry1 := createRandomEntry(t)
	entry2, err := testQueries.GetEntries(context.Background(), entry1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, entry2)

	require.NotZero(t, entry2.ID)
	require.NotZero(t, entry2.CreatedAt)

	require.Equal(t, entry1.ID, entry2.ID)
	require.Equal(t, entry1.AccountID, entry2.AccountID)
	require.Equal(t, entry1.Amount, entry2.Amount)
	require.WithinDuration(t, entry1.CreatedAt, entry2.CreatedAt, time.Second)
}

func TestListEnteries(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomEntry(t)
	}

	arg := ListEntriesParams{
		Limit:  5,
		Offset: 5,
	}

	entries, err := testQueries.ListEntries(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, entries, 5)

	for _, entery := range entries {
		require.NotZero(t, entery.ID)
		require.NotZero(t, entery.CreatedAt)
	}
}
