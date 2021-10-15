package goset

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSet(t *testing.T) {
	const str, i = "foo", 1
	strs := []string{"foo", "foga"}
	s := New()

	// add
	s.Add(str)
	s.Add(i)

	// exists
	require.True(t, s.Exists(str))
	require.False(t, s.Exists("aaa"))

	// strings
	require.Equal(t, s.Strings(), []string{str})

	// add string
	s.AddString(strs...)
	require.True(t, s.Exists("foga"))

	// sort strings
	require.Equal(t, s.SortStrings(), []string{strs[1], strs[0]})

	// delete
	s.Delete(i)
	require.False(t, s.Exists(i))

	// delete string
	s.DeleteString(strs...)
	require.False(t, s.Exists("foga"))

}

func TestNewSet(t *testing.T) {
	require.NotNil(t, NewSetWithStrings([]string{"a"}))
}
