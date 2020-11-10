package datasources_test

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/bxcodec/faker/v3"
	"github.com/stretchr/testify/assert"
	"tiket.vip/src/domains"
	"tiket.vip/src/infrastructures/repositories/v1/datasources"
)

func TestGetLocationSuccss(t *testing.T) {
	dbMock, DB, mock := OpenDBConnection(t)
	defer dbMock.Close()

	var mockLoc domains.Location
	errFaker := faker.FakeData(&mockLoc)
	assert.NoError(t, errFaker)

	rowsLoc := sqlmock.NewRows([]string{"id_location", "name", "address", "street", "city", "country", "zip", "latitude", "longitude", "created_at", "updated_at"}).
		AddRow(mockLoc.ID, mockLoc.Name, mockLoc.Address, mockLoc.Street, mockLoc.City, mockLoc.Country, mockLoc.Zip, mockLoc.Latitude, mockLoc.Longitude, mockLoc.CreatedAt, mockLoc.UpdatedAt)

	queryLoc := regexp.QuoteMeta("SELECT * FROM `location` WHERE id_location =")
	mock.ExpectQuery(queryLoc).WithArgs(mockLoc.ID).WillReturnRows(rowsLoc)

	lr := datasources.NewLocRepo(DB)
	res, errLoc := lr.GetLocation(mockLoc.ID)
	assert.NoError(t, errLoc)
	assert.Equal(t, mockLoc.ID, res.ID)

}

func TestGetLocationNotFound(t *testing.T) {
	dbMock, DB, mock := OpenDBConnection(t)
	defer dbMock.Close()

	var mockLoc domains.Location
	errFaker := faker.FakeData(&mockLoc)
	assert.NoError(t, errFaker)

	queryLoc := regexp.QuoteMeta("SELECT * FROM `location` WHERE id_location =")
	mock.ExpectQuery(queryLoc).WithArgs(mockLoc.ID).WillReturnError(fmt.Errorf("record not found"))

	lr := datasources.NewLocRepo(DB)
	_, errLoc := lr.GetLocation(mockLoc.ID)

	assert.Equal(t, fmt.Errorf("location with id %s is not found", mockLoc.ID).Error(), errLoc.Error())
}
func TestGetLocationError(t *testing.T) {
	dbMock, DB, mock := OpenDBConnection(t)
	defer dbMock.Close()

	var mockLoc domains.Location
	errFaker := faker.FakeData(&mockLoc)
	assert.NoError(t, errFaker)

	queryLoc := regexp.QuoteMeta("SELECT * FROM `location` WHERE id_location =")
	mock.ExpectQuery(queryLoc).WithArgs(mockLoc.ID).WillReturnError(fmt.Errorf("error"))

	lr := datasources.NewLocRepo(DB)
	_, errLoc := lr.GetLocation(mockLoc.ID)

	assert.Equal(t, fmt.Errorf("error").Error(), errLoc.Error())
}
