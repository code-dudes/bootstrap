package env

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnableEnvironments(t *testing.T) {

	assert := assert.New(t)

	DefineEnvironments()
	assert.False(Local.IsValid())

	assert.Nil(DefineLocal(Local))
	assert.Nil(DefineLocal(Local), "redifinition should not error")
	assert.NotNil(DefineDevelopment(Local), "redifinition in different category should error")
	assert.NotNil(DefineProduction(Local), "redifinition in different category should error")
	assert.True(Local.IsLocal())
	assert.False(Development.IsLocal())
	assert.False(Staging.IsLocal())
	assert.False(Production.IsLocal())

	assert.Nil(DefineDevelopment(Development))
	assert.Nil(DefineDevelopment(Development), "redifinition should not error")
	assert.NotNil(DefineLocal(Development), "redifinition in different category should error")
	assert.NotNil(DefineProduction(Development), "redifinition in different category should error")
	assert.True(Development.IsDevelopment())
	assert.False(Local.IsDevelopment())
	assert.False(Staging.IsDevelopment())
	assert.False(Production.IsDevelopment())

	assert.Nil(DefineProduction(Production))
	assert.Nil(DefineProduction(Production), "redifinition should not error")
	assert.NotNil(DefineLocal(Production), "redifinition in different category should error")
	assert.NotNil(DefineDevelopment(Production), "redifinition in different category should error")
	assert.True(Production.IsProduction())
	assert.False(Development.IsProduction())
	assert.False(Local.IsProduction())
	assert.False(Staging.IsProduction())

	// Custom Env
	var production1 Environment = "production1"
	assert.Nil(DefineProduction(production1))
	assert.Nil(DefineProduction(production1), "redifinition should not error")
	assert.NotNil(DefineLocal(production1), "redifinition in different category should error")
	assert.NotNil(DefineDevelopment(production1), "redifinition in different category should error")
	assert.True(production1.IsProduction())
}
