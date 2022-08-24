package agents

import (
	"global.p2p.api/agents/adapters/messages"
	"global.p2p.api/agents/helpers"
	"global.p2p.api/agents/ports/http"
	"global.p2p.api/agents/repositories"
	"global.p2p.api/agents/services"
	"global.p2p.api/app"
	"global.p2p.api/app/notification"
)

type agent struct {
	app app.Gp2p
}

func New() *agent {
	return &agent{}
}

func (a agent) Init(app app.Gp2p) {
	a.app = app

	s := a.makeAgentService()
	ls := a.makeLoginService()
	vs := a.makeVerificationService()
	kyc := a.makeKycService()

	http.New(s, ls, vs, kyc).
		RegisterRoutes(app)
}

func (a agent) makeAgentService() *services.Agent {
	notifier := notification.New()
	wm := messages.NewWelcomeMessage(a.app.Env())

	dataStore := a.makeDataStore()
	service := services.NewAgentService(dataStore, notifier, wm)

	return service
}

func (a agent) makeDataStore() *repositories.AgentStore {
	db := a.app.Database()
	db = db.Table("agent")

	dataStore := repositories.NewAgentStore(db)
	return dataStore
}

func (a agent) makeLoginService() *services.Login {
	store := a.makeDataStore()
	return services.NewLoginService(store)
}

func (a agent) makeVerificationService() *services.AgentVerification {
	env := a.app.Env()
	otp := helpers.Otp(env)
	verificationMessage := messages.NewVerificationMessage(env, notification.New(), otp)
	dataStore := a.makeDataStore()
	return services.NewAgentVerificationService(dataStore, verificationMessage, otp)
}

func (a agent) makeKycService() *services.AgentKyc {
	r := a.makeDataStore()
	f := a.app.FileSystem()
	return services.NewAgentKyc(f, r)
}

