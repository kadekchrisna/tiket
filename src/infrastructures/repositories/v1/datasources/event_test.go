package datasources_test

import (
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/bxcodec/faker/v3"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"tiket.vip/src/domains"
	"tiket.vip/src/infrastructures/repositories/v1/datasources"
)

func TestGetEvent(t *testing.T) {
	dbMock, mock, errMock := sqlmock.New()
	assert.NoError(t, errMock)

	defer dbMock.Close()

	// DB, _ := orm.New(dbMock)
	// assert.NoError(t, errDb)

	DB, errDb := gorm.Open(mysql.New(mysql.Config{
		Conn:                      dbMock,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{})
	assert.NoError(t, errDb)

	var mockEvents domains.Event
	errFaker := faker.FakeData(&mockEvents)
	assert.NoError(t, errFaker)

	rowsEv := sqlmock.NewRows([]string{"id_event", "id_location", "name", "desc", "start_date", "end_date", "created_at", "updated_at"}).
		AddRow(mockEvents.ID, mockEvents.IDLocation, mockEvents.Name, mockEvents.Desc, mockEvents.StartDate, mockEvents.EndDate, mockEvents.CreatedAt, mockEvents.UpdatedAt)

	rowsLoc := sqlmock.NewRows([]string{"id_location", "name", "address", "street", "city", "country", "zip", "latitude", "longitude", "created_at", "updated_at"}).
		AddRow(mockEvents.Location.ID, mockEvents.Location.Name, mockEvents.Location.Address, mockEvents.Location.Street, mockEvents.Location.City, mockEvents.Location.Country, mockEvents.Location.Zip, mockEvents.Location.Latitude, mockEvents.Location.Longitude, mockEvents.Location.CreatedAt, mockEvents.Location.UpdatedAt)

	queryEv := regexp.QuoteMeta("SELECT * FROM `event` WHERE id_event = ? ORDER BY `event`.`id_event` LIMIT 1")
	queryLoc := regexp.QuoteMeta("SELECT * FROM `location` WHERE `location`.`id_location` = ?")

	mock.ExpectQuery(queryEv).WithArgs("48d25390-1b71-11eb-9d1a-f94f590013b6").WillReturnRows(rowsEv)
	mock.ExpectQuery(queryLoc).WithArgs(mockEvents.IDLocation).WillReturnRows(rowsLoc)

	er := datasources.NewEventRepo(DB)

	res, errFetch := er.GetEvent("48d25390-1b71-11eb-9d1a-f94f590013b6")
	assert.Equal(t, mockEvents.ID, res.ID)
	assert.NoError(t, errFetch)

}
