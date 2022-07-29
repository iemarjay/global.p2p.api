package agents

import (
	"global.p2p.api/agents/adapters/messages"
	"global.p2p.api/agents/data"
	"global.p2p.api/agents/ports/http"
	"global.p2p.api/agents/services"
	"global.p2p.api/gp2p"
	"global.p2p.api/gp2p/notification"
)

type agent struct {
	app gp2p.Gp2p
}

func New() *agent {
	return &agent{}
}

func (a agent) Init(app gp2p.Gp2p) {
	a.app = app
	service := a.makeAgentService()
	httpHandlers := http.New(service)
	httpHandlers.Register(app.Router())
}

func (a agent) makeAgentService() *services.AgentService {
	notifier := notification.New()
	wm := messages.NewWelcomeMessage(a.app.Env())

	db := a.app.Database()
	db = db.Table("agent")

	dataStore := data.NewAgentStore(db)
	service := services.NewAgentService(dataStore, notifier, wm)

	return service
}

