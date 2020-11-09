package routers

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	usecase "tiket.vip/src/applications/usecases/v1"
	"tiket.vip/src/domains"
	"tiket.vip/src/infrastructures/repositories/v1/datasources"
	"tiket.vip/src/interfaces/controllers/v1"
)

type Handler struct {
	Event domains.EventController
}

func NewHandler(db *gorm.DB) *Handler {
	return &Handler{
		Event: controllers.NewEventController(usecase.NewEventUseCase(datasources.NewRepo(db))),
	}
}

func (h *Handler) Register(v1 *echo.Group) {
	ev := v1.Group("/event")
	ev.GET("/all", h.Event.GetAllEvents)
	ev.GET("/get_info/:id", h.Event.GetEvent)
	ev.POST("/create", h.Event.CreateEvent)
	ev.PUT("/update", h.Event.UpdateEvent)
	ev.DELETE("/delete/:id", h.Event.DeleteEvent)
}
