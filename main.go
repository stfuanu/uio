package main

import (
	"encoding/json"
	"log"
	"path/filepath"

	"uio/core"

	"uio/app/route"
	"uio/app/shared/database"
	"uio/app/shared/email"
	"uio/app/shared/jsonconfig"
	"uio/app/shared/recaptcha"
	"uio/app/shared/server"
	"uio/app/shared/session"
	"uio/app/shared/view"
	"uio/app/shared/view/plugin"
)

func init() {
	// Verbose logging with file name and line number
	log.SetFlags(log.Lshortfile)
}

var (
	addedORnot bool
	err        error
	hashifany  string
)

func main() {

	// fmt.Println(app2.STFU())
	core.LoadChains()
	core.GetLiveStat()
	// rand.Seed(time.Now().UnixNano())
	// var randnum int
	// core.TestBallot("ELECTION_1")
	// core.TestBallot("ELECTION_2")
	// core.TestBallot("ELECTION_3")

	// core.Genesisblock()

	// Load the configuration file
	jsonconfig.Load(filepath.Join("config", "config.json"), config)

	// Configure the session cookie store
	session.Configure(config.Session)

	// Connect to database
	// database.Connect(config.Database)
	database.Connect()

	// Configure the Google reCAPTCHA prior to loading view plugins
	recaptcha.Configure(config.Recaptcha)

	// Setup the views
	view.Configure(config.View)
	view.LoadTemplates(config.Template.Root, config.Template.Children)
	view.LoadPlugins(
		plugin.TagHelper(config.View),
		plugin.NoEscape(),
		plugin.PrettyTime(),
		recaptcha.Plugin())

	// _, _, _ = core.Addnewblock("BALLOT", "19yaXUBokMBzdqFex5qpZPwvV3CStnRVff", "", core.CreateBallot("BEST_SAMOSA_LOL", []string{"12tmRt6AADfQhfruF3RzFDdNhjiSEkwMvF", "1HRK5H21wFguq5ecJF8FvQk28qYnPz1Qb9", "1DngEcP2tCkxZNiAmm3Ar8VXXAAvAPfm8E"}))

	// Candidates := []string{"12tmRt6AADfQhfruF3RzFDdNhjiSEkwMvF", "1HRK5H21wFguq5ecJF8FvQk28qYnPz1Qb9", "1DngEcP2tCkxZNiAmm3Ar8VXXAAvAPfm8E"}

	// for i := 1; i < 3; i++ {
	// 	randnum = rand.Intn(2-0+1) + 0
	// 	addedORnot, err, hashifany = core.Addnewblock("VOTE", "19yaXUBokMBzdqFex5qpZPwvV3CStnRVff", Candidates[randnum], "4191a0fc9cabc39454827d41ac33137ca92d8710e5930a9653bf16afc44a7e0f")
	// 	fmt.Println(addedORnot, err, hashifany)
	// }

	// core.PrintblockchainStdout()

	// Start the listener
	server.Run(route.LoadHTTP(), route.LoadHTTPS(), config.Server)
}

// config the settings variable
var config = &configuration{}

// configuration contains the application settings
type configuration struct {
	Database  database.Info   `json:"Database"`
	Email     email.SMTPInfo  `json:"Email"`
	Recaptcha recaptcha.Info  `json:"Recaptcha"`
	Server    server.Server   `json:"Server"`
	Session   session.Session `json:"Session"`
	Template  view.Template   `json:"Template"`
	View      view.View       `json:"View"`
}

// ParseJSON unmarshals bytes to structs
func (c *configuration) ParseJSON(b []byte) error {
	return json.Unmarshal(b, &c)
}
