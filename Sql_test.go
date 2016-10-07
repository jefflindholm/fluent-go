package fluentSQL

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var business = MakeBusiness()
var b = business.As("b")

func TestAlias(t *testing.T) {

	assert.Equal(t, b.Alias(), "b", "they should be equal")
	assert.Equal(t, business.Alias(), "businesses", "they should be equal")
}
func TestColumnAs(t *testing.T) {

}
