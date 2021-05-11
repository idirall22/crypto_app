package redismem_test

import (
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/require"
)

func TestSetGet(t *testing.T) {
	key := gofakeit.Animal()
	value := int(gofakeit.Int32())

	err := memTest.SetLoginAttemps(key, value, time.Second*3)
	require.NoError(t, err)

	val, err := memTest.GetLoginAttemps(key)
	require.NoError(t, err)
	require.Equal(t, value, val)
}
