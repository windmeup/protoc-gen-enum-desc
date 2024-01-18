package example

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestClient_ClientEvent_Description(t *testing.T) {
	require.Equal(t, "client start", Client_CLIENT_EVENT_START.Description())
}
