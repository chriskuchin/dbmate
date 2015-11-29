package driver_test

import (
	"github.com/adrianmacneil/dbmate/driver"
	"github.com/adrianmacneil/dbmate/driver/postgres"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGet_Postgres(t *testing.T) {
	drv, err := driver.Get("postgres")
	require.Nil(t, err)
	_, ok := drv.(postgres.Driver)
	require.Equal(t, true, ok)
}

func TestGet_Error(t *testing.T) {
	drv, err := driver.Get("foo")
	require.Equal(t, "Unknown driver: foo", err.Error())
	require.Nil(t, drv)
}
