package datasources_test

import (
	"database/sql"
	"fmt"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/bxcodec/faker/v3"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"tiket.vip/src/domains"
	"tiket.vip/src/infrastructures/repositories/v1/datasources"
)

func OpenDBConnection(t *testing.T) (*sql.DB, *gorm.DB, sqlmock.Sqlmock) {
	dbMock, mock, errMock := sqlmock.New()
	assert.NoError(t, errMock)

	DB, errDb := gorm.Open(mysql.New(mysql.Config{
		Conn:                      dbMock,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{})
	assert.NoError(t, errDb)

	return dbMock, DB, mock
}

func TestGetEventSuccess(t *testing.T) {
	dbMock, DB, mock := OpenDBConnection(t)
	defer dbMock.Close()

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
func TestGetEventFailed(t *testing.T) {
	dbMock, DB, mock := OpenDBConnection(t)
	defer dbMock.Close()

	var mockEvents domains.Event
	errFaker := faker.FakeData(&mockEvents)
	assert.NoError(t, errFaker)

	queryEv := regexp.QuoteMeta("SELECT * FROM `event` WHERE id_event = ? ORDER BY `event`.`id_event` LIMIT 1")

	mock.ExpectQuery(queryEv).WithArgs("48d25390-1b71-11eb-9d1a-f94f590013ab6").WillReturnError(fmt.Errorf("error"))

	er := datasources.NewEventRepo(DB)

	_, errFetch := er.GetEvent("48d25390-1b71-11eb-9d1a-f94f590013ab6")
	assert.Equal(t, errFetch.Error(), fmt.Errorf("error").Error())
}

func TestGetEventFailedNotFound(t *testing.T) {
	dbMock, DB, mock := OpenDBConnection(t)
	defer dbMock.Close()

	var mockEvents domains.Event
	errFaker := faker.FakeData(&mockEvents)
	assert.NoError(t, errFaker)

	queryEv := regexp.QuoteMeta("SELECT * FROM `event` WHERE id_event = ? ORDER BY `event`.`id_event` LIMIT 1")

	mock.ExpectQuery(queryEv).WithArgs("48d25390-1b71-11eb-9d1a-f94f590013ab6").WillReturnError(fmt.Errorf("record not found"))

	er := datasources.NewEventRepo(DB)

	_, errFetch := er.GetEvent("48d25390-1b71-11eb-9d1a-f94f590013ab6")

	assert.Equal(t, errFetch.Error(), "event with id 48d25390-1b71-11eb-9d1a-f94f590013ab6 is not found")
}

func TestGetAllEventsSuccess(t *testing.T) {
	dbMock, DB, mock := OpenDBConnection(t)
	defer dbMock.Close()

	var mockEvent domains.Event
	errFaker := faker.FakeData(&mockEvent)
	assert.NoError(t, errFaker)

	rowsEv := sqlmock.NewRows([]string{"id_event", "id_location", "name", "desc", "start_date", "end_date", "created_at", "updated_at"}).
		AddRow(mockEvent.ID, mockEvent.IDLocation, mockEvent.Name, mockEvent.Desc, mockEvent.StartDate, mockEvent.EndDate, mockEvent.CreatedAt, mockEvent.UpdatedAt)

	rowsLoc := sqlmock.NewRows([]string{"id_location", "name", "address", "street", "city", "country", "zip", "latitude", "longitude", "created_at", "updated_at"}).
		AddRow(mockEvent.Location.ID, mockEvent.Location.Name, mockEvent.Location.Address, mockEvent.Location.Street, mockEvent.Location.City, mockEvent.Location.Country, mockEvent.Location.Zip, mockEvent.Location.Latitude, mockEvent.Location.Longitude, mockEvent.Location.CreatedAt, mockEvent.Location.UpdatedAt)

	queryEv := regexp.QuoteMeta("SELECT * FROM `event`")
	queryLoc := regexp.QuoteMeta("SELECT * FROM `location` WHERE `location`.`id_location` = ?")

	mock.ExpectQuery(queryEv).WillReturnRows(rowsEv)
	mock.ExpectQuery(queryLoc).WithArgs(mockEvent.IDLocation).WillReturnRows(rowsLoc)

	er := datasources.NewEventRepo(DB)
	var mockEvents []domains.Event
	mockEvents = append(mockEvents, mockEvent)

	res, errFetch := er.GetAllEvents()
	assert.Equal(t, mockEvents[0].ID, (*res)[0].ID)
	assert.NoError(t, errFetch)
}

func TestGetAllEventsFailed(t *testing.T) {
	dbMock, DB, mock := OpenDBConnection(t)
	defer dbMock.Close()

	var mockEvent domains.Event
	errFaker := faker.FakeData(&mockEvent)
	assert.NoError(t, errFaker)
	queryEv := regexp.QuoteMeta("SELECT * FROM `event`")
	mock.ExpectQuery(queryEv).WillReturnError(fmt.Errorf("error"))

	er := datasources.NewEventRepo(DB)
	var mockEvents []domains.Event
	mockEvents = append(mockEvents, mockEvent)

	_, errFetch := er.GetAllEvents()
	assert.Equal(t, errFetch.Error(), fmt.Errorf("error").Error())
}

func TestCreateEventFailedCount(t *testing.T) {
	dbMock, DB, mock := OpenDBConnection(t)
	defer dbMock.Close()
	var mockEvent domains.Event
	errFaker := faker.FakeData(&mockEvent)
	assert.NoError(t, errFaker)

	queryCn := regexp.QuoteMeta("SELECT count(1) FROM `event` WHERE id_event = ?")
	mock.ExpectQuery(queryCn).WillReturnError(fmt.Errorf("error"))
	// queryEv := regexp.QuoteMeta("INSERT INTO `event`")
	// mock.ExpectQuery(queryEv).WillReturnError(fmt.Errorf("error"))
	er := datasources.NewEventRepo(DB)
	_, errFetch := er.CreateEvent(mockEvent)
	assert.Equal(t, errFetch.Error(), fmt.Errorf("error").Error())
}

func TestCreateEventSuccess(t *testing.T) {
	id := uuid.New().String()

	dbMock, DB, mock := OpenDBConnection(t)
	defer dbMock.Close()
	var mockEvent domains.Event
	errFaker := faker.FakeData(&mockEvent)
	assert.NoError(t, errFaker)
	mockEvent.ID = id
	mockEvent.Location = &domains.Location{}
	mockEvent.StartDate = "2006-01-02 15:04:05"
	mockEvent.EndDate = "2006-01-02 15:04:05"

	rowsCn := sqlmock.NewRows([]string{"count"}).
		AddRow(0)

	queryCn := regexp.QuoteMeta("SELECT count(1) FROM `event`")
	mock.ExpectQuery(queryCn).WithArgs(mockEvent.ID).WillReturnRows(rowsCn)

	queryEv := regexp.QuoteMeta("INSERT INTO `event`")
	mock.ExpectExec(queryEv).WithArgs(mockEvent.ID, mockEvent.IDLocation, mockEvent.Name, mockEvent.Desc, mockEvent.StartDate, mockEvent.EndDate, mockEvent.CreatedAt, mockEvent.UpdatedAt).WillReturnResult(sqlmock.NewResult(1, 0))
	er := datasources.NewEventRepo(DB)
	res, errCreate := er.CreateEvent(mockEvent)
	assert.Equal(t, mockEvent.ID, res.ID)
	assert.NoError(t, errCreate)
}
func TestUpdateEventSuccess(t *testing.T) {
	id := uuid.New().String()

	dbMock, DB, mock := OpenDBConnection(t)
	defer dbMock.Close()
	var mockEvent domains.Event
	errFaker := faker.FakeData(&mockEvent)
	assert.NoError(t, errFaker)
	mockEvent.ID = id
	mockEvent.Location = &domains.Location{}
	mockEvent.StartDate = "2006-01-02 15:04:05"
	mockEvent.EndDate = "2006-01-02 15:04:05"

	queryEv := regexp.QuoteMeta("UPDATE `event` SET ")
	mock.ExpectExec(queryEv).WithArgs(mockEvent.ID, mockEvent.IDLocation, mockEvent.Name, mockEvent.Desc, mockEvent.StartDate, mockEvent.EndDate, mockEvent.CreatedAt, mockEvent.UpdatedAt, mockEvent.ID).WillReturnResult(sqlmock.NewResult(1, 1))
	er := datasources.NewEventRepo(DB)
	res, errCreate := er.UpdateEvent(mockEvent)
	assert.NoError(t, errCreate)
	assert.Equal(t, mockEvent.ID, res.ID)
}

func TestUpdateEventFailedNotFound(t *testing.T) {
	id := uuid.New().String()

	dbMock, DB, mock := OpenDBConnection(t)
	defer dbMock.Close()
	var mockEvent domains.Event
	errFaker := faker.FakeData(&mockEvent)
	assert.NoError(t, errFaker)
	mockEvent.ID = id
	mockEvent.Location = &domains.Location{}
	mockEvent.StartDate = "2006-01-02 15:04:05"
	mockEvent.EndDate = "2006-01-02 15:04:05"

	queryEv := regexp.QuoteMeta("UPDATE `event` SET ")
	mock.ExpectExec(queryEv).WithArgs(mockEvent.ID, mockEvent.IDLocation, mockEvent.Name, mockEvent.Desc, mockEvent.StartDate, mockEvent.EndDate, mockEvent.CreatedAt, mockEvent.UpdatedAt, mockEvent.ID).WillReturnResult(sqlmock.NewResult(0, 0))
	er := datasources.NewEventRepo(DB)
	_, errCreate := er.UpdateEvent(mockEvent)
	assert.Equal(t, fmt.Errorf("event with id %s is not found", mockEvent.ID).Error(), errCreate.Error())
}

func TestUpdateEventFailedError(t *testing.T) {
	id := uuid.New().String()

	dbMock, DB, mock := OpenDBConnection(t)
	defer dbMock.Close()
	var mockEvent domains.Event
	errFaker := faker.FakeData(&mockEvent)
	assert.NoError(t, errFaker)
	mockEvent.ID = id
	mockEvent.Location = &domains.Location{}
	mockEvent.StartDate = "2006-01-02 15:04:05"
	mockEvent.EndDate = "2006-01-02 15:04:05"

	queryEv := regexp.QuoteMeta("UPDATE `event` SET ")
	mock.ExpectExec(queryEv).WithArgs(mockEvent.ID, mockEvent.IDLocation, mockEvent.Name, mockEvent.Desc, mockEvent.StartDate, mockEvent.EndDate, mockEvent.CreatedAt, mockEvent.UpdatedAt, mockEvent.ID).WillReturnError(fmt.Errorf("error"))
	er := datasources.NewEventRepo(DB)
	_, errCreate := er.UpdateEvent(mockEvent)
	assert.Equal(t, fmt.Errorf("error").Error(), errCreate.Error())
}

func TestDeleteEventSuccess(t *testing.T) {
	id := uuid.New().String()

	dbMock, DB, mock := OpenDBConnection(t)
	defer dbMock.Close()
	var mockEvent domains.Event
	errFaker := faker.FakeData(&mockEvent)
	assert.NoError(t, errFaker)
	mockEvent.ID = id
	mockEvent.Location = &domains.Location{}
	mockEvent.StartDate = "2006-01-02 15:04:05"
	mockEvent.EndDate = "2006-01-02 15:04:05"

	queryEv := regexp.QuoteMeta("DELETE FROM `event`")
	mock.ExpectExec(queryEv).WithArgs(mockEvent.ID).WillReturnResult(sqlmock.NewResult(1, 1))
	er := datasources.NewEventRepo(DB)
	resId, errCreate := er.DeleteEvent(mockEvent.ID)

	assert.NoError(t, errCreate)
	assert.Equal(t, mockEvent.ID, *resId)
}

func TestDeleteEventFailed(t *testing.T) {
	id := uuid.New().String()

	dbMock, DB, mock := OpenDBConnection(t)
	defer dbMock.Close()
	var mockEvent domains.Event
	errFaker := faker.FakeData(&mockEvent)
	assert.NoError(t, errFaker)
	mockEvent.ID = id
	mockEvent.Location = &domains.Location{}
	mockEvent.StartDate = "2006-01-02 15:04:05"
	mockEvent.EndDate = "2006-01-02 15:04:05"

	queryEv := regexp.QuoteMeta("DELETE FROM `event`")
	mock.ExpectExec(queryEv).WithArgs(mockEvent.ID).WillReturnError(fmt.Errorf("error"))
	er := datasources.NewEventRepo(DB)
	_, errCreate := er.DeleteEvent(mockEvent.ID)

	assert.Equal(t, fmt.Errorf("error").Error(), errCreate.Error())
}

func TestGetAllEventsPaginateSuccess(t *testing.T) {
	dbMock, DB, mock := OpenDBConnection(t)
	defer dbMock.Close()

	var mockEvent domains.Event
	var mockEvents domains.Events
	var mockEventPagi domains.EventPagi

	errFaker := faker.FakeData(&mockEvent)
	assert.NoError(t, errFaker)

	mockEventPagi.EventName = mockEvent.Name
	mockEventPagi.EndDate = mockEvent.EndDate
	mockEventPagi.StartDate = mockEvent.StartDate
	mockEventPagi.LocAddress = mockEvent.Location.Address
	mockEventPagi.LocCountry = mockEvent.Location.Country
	mockEventPagi.LocName = mockEvent.Location.Name

	rowsCn := sqlmock.NewRows([]string{"count"}).
		AddRow(1)
	queryEvCn := regexp.QuoteMeta("SELECT count(1) FROM event event join")
	mock.ExpectQuery(queryEvCn).WillReturnRows(rowsCn)

	rowsEv := sqlmock.NewRows([]string{"event.id_event", "event.id_location", "event.name", "event.desc", "event.start_date", "event.end_date", "event.created_at", "event.updated_at", "location.id_location", "location.name", "location.address", "location.street", "location.city", "location.country", "location.zip", "location.latitude", "location.longitude", "location.created_at", "location.updated_at"}).
		AddRow(mockEvent.ID, mockEvent.IDLocation, mockEvent.Name, mockEvent.Desc, mockEvent.StartDate, mockEvent.EndDate, mockEvent.CreatedAt, mockEvent.UpdatedAt, mockEvent.Location.ID, mockEvent.Location.Name, mockEvent.Location.Address, mockEvent.Location.Street, mockEvent.Location.City, mockEvent.Location.Country, mockEvent.Location.Zip, mockEvent.Location.Latitude, mockEvent.Location.Longitude, mockEvent.Location.CreatedAt, mockEvent.Location.UpdatedAt)
	queryEv := regexp.QuoteMeta("SELECT * FROM event event join location location")
	mock.ExpectQuery(queryEv).WillReturnRows(rowsEv)

	mockEvents.Events = append(mockEvents.Events, mockEvent)
	mockEvents.Total = 1

	er := datasources.NewEventRepo(DB)
	res, errCreate := er.GetAllEventsPaginate(mockEventPagi)
	assert.NoError(t, errCreate)
	assert.Equal(t, mockEvents.Events[0].ID, (*res).Events[0].ID)
}

func TestGetAllEventsPaginateFailedCount(t *testing.T) {
	dbMock, DB, mock := OpenDBConnection(t)
	defer dbMock.Close()

	var mockEvent domains.Event
	var mockEvents domains.Events
	var mockEventPagi domains.EventPagi

	errFaker := faker.FakeData(&mockEvent)
	assert.NoError(t, errFaker)

	rowsCn := sqlmock.NewRows([]string{"count"}).
		AddRow(0)
	queryEvCn := regexp.QuoteMeta("SELECT count(1) FROM event event join")
	mock.ExpectQuery(queryEvCn).WillReturnRows(rowsCn)

	er := datasources.NewEventRepo(DB)
	res, errCreate := er.GetAllEventsPaginate(mockEventPagi)
	assert.NoError(t, errCreate)
	assert.Equal(t, mockEvents.Total, (*res).Total)
}

func TestGetAllEventsPaginateFailedCountError(t *testing.T) {
	dbMock, DB, mock := OpenDBConnection(t)
	defer dbMock.Close()

	var mockEvent domains.Event
	var mockEventPagi domains.EventPagi

	errFaker := faker.FakeData(&mockEvent)
	assert.NoError(t, errFaker)

	queryEvCn := regexp.QuoteMeta("SELECT count(1) FROM event event join")
	mock.ExpectQuery(queryEvCn).WillReturnError(fmt.Errorf("error"))

	er := datasources.NewEventRepo(DB)
	_, errCreate := er.GetAllEventsPaginate(mockEventPagi)

	assert.Equal(t, fmt.Errorf("error").Error(), errCreate.Error())

}
