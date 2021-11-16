package route

import (
	"net/http"

	"uio/app/controller"
	"uio/app/route/middleware/acl"
	hr "uio/app/route/middleware/httprouterwrapper"
	"uio/app/route/middleware/logrequest"
	"uio/app/route/middleware/pprofhandler"
	"uio/app/shared/session"

	"github.com/gorilla/context"
	"github.com/josephspurrier/csrfbanana"
	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
)

// Load returns the routes and middleware
func Load() http.Handler {
	return middleware(routes())
}

// LoadHTTPS returns the HTTP routes and middleware
func LoadHTTPS() http.Handler {
	return middleware(routes())
}

// LoadHTTP returns the HTTPS routes and middleware
func LoadHTTP() http.Handler {
	return middleware(routes())

	// Uncomment this and comment out the line above to always redirect to HTTPS
	//return http.HandlerFunc(redirectToHTTPS)
}

// Optional method to make it easy to redirect from HTTP to HTTPS
func redirectToHTTPS(w http.ResponseWriter, req *http.Request) {
	http.Redirect(w, req, "https://"+req.Host, http.StatusMovedPermanently)
}

func routes() *httprouter.Router {
	r := httprouter.New()

	// Set 404 handler
	r.NotFound = alice.
		New().
		ThenFunc(controller.Error404)

	// Serve static files, no directory browsing
	r.GET("/static/*filepath", hr.Handler(alice.
		New().
		ThenFunc(controller.Static)))

	// Home page
	r.GET("/", hr.Handler(alice.
		New().
		ThenFunc(controller.IndexGET)))

	// Login
	r.GET("/login", hr.Handler(alice.
		New(acl.DisallowAuth).
		ThenFunc(controller.LoginGET)))
	r.POST("/login", hr.Handler(alice.
		New(acl.DisallowAuth).
		ThenFunc(controller.LoginPOST)))
	r.GET("/logout", hr.Handler(alice.
		New().
		ThenFunc(controller.LogoutGET)))

	// Register
	r.GET("/register", hr.Handler(alice.
		New(acl.DisallowAuth).
		ThenFunc(controller.RegisterGET)))
	r.POST("/register", hr.Handler(alice.
		New(acl.DisallowAuth).
		ThenFunc(controller.RegisterPOST)))

	// About
	r.GET("/about", hr.Handler(alice.
		New().
		ThenFunc(controller.AboutGET)))

	// API
	r.GET("/api", hr.Handler(alice.
		New().
		ThenFunc(controller.APIGET)))

	r.GET("/live", hr.Handler(alice.
		New().
		ThenFunc(controller.LiveGet)))

	r.GET("/api/live", hr.Handler(alice.
		New().
		ThenFunc(controller.LiveStat)))

	r.GET("/api/votes.json", hr.Handler(alice.
		New().
		ThenFunc(controller.GetBlockchain)))
	r.GET("/api/ballots.json", hr.Handler(alice.
		New().
		ThenFunc(controller.GetBallots)))

	r.GET("/api/vote/:hash", hr.Handler(alice.
		New().
		ThenFunc(controller.GetBlock)))
	r.GET("/api/addrinfo/:addr", hr.Handler(alice.
		New().
		ThenFunc(controller.Getaddrinfo)))

	r.GET("/api/electioninfo/:btxhash", hr.Handler(alice.
		New().
		// ThenFunc(controller.GetElectionInfoWeb)))
		ThenFunc(controller.GetElectionInfoWeb_Fast)))
	r.GET("/api/electioninfo", hr.Handler(alice.
		New().
		ThenFunc(controller.GetALLElectionInfoWeb)))

	r.POST("/vote/newvtx", hr.Handler(alice.
		New().
		ThenFunc(controller.NewVoteweb)))

	r.POST("/vote/newbtx", hr.Handler(alice.
		New().
		ThenFunc(controller.NewBallotweb)))

	r.POST("/api/vote/newvtx", hr.Handler(alice.
		New(acl.DisallowAuth).
		ThenFunc(controller.NewVote)))

	r.POST("/api/vote/newbtx", hr.Handler(alice.
		New(acl.DisallowAuth).
		ThenFunc(controller.NewBallot)))

	r.POST("/generate", hr.Handler(alice.
		New(acl.DisallowAnon).
		ThenFunc(controller.GenerateWallet)))

	// Enable Pprof
	r.GET("/debug/pprof/*pprof", hr.Handler(alice.
		New(acl.DisallowAnon).
		ThenFunc(pprofhandler.Handler)))

	return r
}

func middleware(h http.Handler) http.Handler {
	// Prevents CSRF and Double Submits
	cs := csrfbanana.New(h, session.Store, session.Name)
	cs.FailureHandler(http.HandlerFunc(controller.InvalidToken))
	cs.ClearAfterUsage(true)
	cs.ExcludeRegexPaths([]string{"/static(.*)"})
	csrfbanana.TokenLength = 32
	csrfbanana.TokenName = "token"
	csrfbanana.SingleToken = false
	h = cs

	// Log every request
	h = logrequest.Handler(h)

	// Clear handler for Gorilla Context
	h = context.ClearHandler(h)

	return h
}
