package gbt36104

import (
	"github.com/stretchr/testify/assert"
	"testing"
)
import "gorm.io/gorm"
import "gorm.io/driver/sqlite"

func TestGormGen(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	assert.NoError(t, err)
	assert.NoError(t, db.AutoMigrate(Organization{}))
}
