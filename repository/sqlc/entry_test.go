package repository

import (
	"context"
	"testing"

	"github.com/jim124/simple-bank/util"
	"github.com/stretchr/testify/require"
)

func createRandomEntry(t *testing.T) (Account, Entry) {
	account := createRandomAccount(t)
	arg := CreateEntryParams{
		AccountID: account.ID,
		Amount:    util.RandomMoney(),
	}
	entry, err := testQueries.CreateEntry(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, entry)
	require.Equal(t, account.ID, entry.AccountID)
	require.Equal(t, arg.Amount, entry.Amount)
	return account, entry
}

func TestCreateEntry(t *testing.T) {
	createRandomEntry(t)
}

func TestGetEntry(t *testing.T) {
	_, entry := createRandomEntry(t)
	entry2, err := testQueries.GetEntry(context.Background(), entry.ID)
	require.NoError(t, err)
	require.NotEmpty(t, entry2)
	require.Equal(t, entry.ID, entry2.ID)
	require.Equal(t, entry.AccountID, entry2.AccountID)
	require.Equal(t, entry.Amount, entry2.Amount)
}
func TestListEntries(t *testing.T) {
	var accountId int64
	for range 10 {
		account, _ := createRandomEntry(t)
		accountId = account.ID
	}
	arg := ListEntriesParams{
		AccountID: accountId,
		Limit:     5,
		Offset:    5,
	}
	_, err := testQueries.ListEntries(context.Background(), arg)
	require.NoError(t, err)

}
