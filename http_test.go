package down

import (
	"testing"

	"github.com/gsamokovarov/assert"
)

func Test_successfulHTTP(t *testing.T) {
	t.Parallel()

	assert.
		True(t, successfulHTTP(200)).
		True(t, successfulHTTP(204)).
		False(t, successfulHTTP(400))
}
