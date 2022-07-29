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
	loginService := a.makeLogicService()
	httpHandlers := http.New(service, loginService)
	httpHandlers.Register(app.Router())
}

func (a agent) makeAgentService() *services.AgentService {
	notifier := notification.New()
	wm := messages.NewWelcomeMessage(a.app.Env())

	dataStore := a.makeDataStore()
	service := services.NewAgentService(dataStore, notifier, wm)

	return service
}

func (a agent) makeDataStore() *data.AgentStore {
	db := a.app.Database()
	db = db.Table("agent")

	dataStore := data.NewAgentStore(db)
	return dataStore
}

func (a agent) makeLogicService() *services.LoginService {
	store := a.makeDataStore()
	return services.NewLoginService(store)
}

