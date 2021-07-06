package naming_test

import (
	"gtest_example/app/utils/naming/testdata"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/stretchr/testify/assert"

	"gtest_example/app/utils/naming"
)

func TestNameToEmail_valid(t *testing.T) {
	for _, item := range testdata.CrewData {
		out, _ := naming.NameToEmail(item.InName)
		assert.Equal(t, item.OutEmail, out,
			"Name to email conversion failed: '%s'", item.InName)
	}
}

func TestNameToEmail_ErrNameIsEmpty(t *testing.T) {
	_, err := naming.NameToEmail("")
	assert.ErrorIs(t, err, naming.ErrNameIsEmpty,
		"must throw an empty name error")
}

func Test_AddDomain(t *testing.T) {
	in := "dummy"
	out := "dummy@acme.com"
	eml, err := naming.AddDomain(in)
	require.NoError(t, err)
	assert.Equal(t, out, eml)
}

func Test_AddDomain_Err(t *testing.T) {
	in := "dummy@hi!@hoy"
	_, err := naming.AddDomain(in)
	assert.Error(t, err)
}
