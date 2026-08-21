package generic_test

import (
	"testing"

	"uuid"
)

func TestUUID(t *testing.T) {
	u := uuid.New()
	t.Log(u.String())
	t.Log(uuid.Max())

	v := uuid.New()
	t.Log(v.Compare(u))
	t.Log(u.Compare(u))
	t.Log(u.Compare(v))
}
