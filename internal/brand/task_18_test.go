package brand

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestQinghaiBrandTask18(t *testing.T) {
	s := NewService(NewRegistry(), time.Now)
	available, err := s.CheckAvailableStock(context.Background(), ProductListing{Stock: 20, Reserved: 5})
	require.NoError(t, err)
	require.Equal(t, 15, available)
}
