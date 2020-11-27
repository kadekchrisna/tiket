package routers

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	usecase "tiket.vip/src/applications/usecases/v1"
	"tiket.vip/src/domains"
	elastics "tiket.vip/src/infrastructures/elastic"
	"tiket.vip/src/infrastructures/repositories/v1/datasources"
	"tiket.vip/src/interfaces/controllers/v1"
)

type Handler struct {
	Event    domains.EventController
	Location domains.LocController
}

func NewHandler(db *gorm.DB, ec elastics.ESClientInterface) *Handler {

	return &Handler{
		Event:    controllers.NewEventController(usecase.NewEventUseCase(datasources.NewEventRepo(db), datasources.NewLocRepo(db))),
		Location: controllers.NewLocController(usecase.NewLocUseCase(datasources.NewLocRepo(db), datasources.NewItemRepo(ec))),
	}
}

func (h *Handler) Register(v1 *echo.Group) {
	ev := v1.Group("/event")
	ev.GET("/all", h.Event.GetAllEvents)
	ev.GET("/all/paginate", h.Event.GetAllEventsPaginate)
	ev.GET("/get_info/:id", h.Event.GetEvent)
	ev.POST("/create", h.Event.CreateEvent)
	ev.PUT("/update", h.Event.UpdateEvent)
	ev.DELETE("/delete/:id", h.Event.DeleteEvent)

	loc := v1.Group("/location")
	loc.POST("/create", h.Location.CreateLocation)
	loc.POST("/search", h.Location.SearchItem)
}
