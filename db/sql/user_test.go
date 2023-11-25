package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/sjxiang/simplebank/util"
)

func createRandomUser(t *testing.T) {
	hashedPassword, err := util.HashPassword(util.RandomString(6))
	require.NoError(t, err)

	arg := CreateUserParams{
		Username:        util.RandomOwner(),
		HashedPassword:  hashedPassword,
		FullName:        util.RandomOwner(),
		Email:           util.RandomEmail(),
		IsEmailVerified: true,
		Role:            "guest",
	}

	id, err := testData.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotZero(t, id)
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}