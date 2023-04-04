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
	core.CreateAllElectionInfo()
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

	// Test :
	// _, _, _ = core.Addnewblock("140edf6c44171ab7c93cb2df9da9cb56d253757c4b16badfde6cdfba86514b", "BALLOT", "", core.CreateBallot("General-Elections-2021", []string{"1A3hszZSQ3X3uTKM4vBApsmdAzVb1JNesQ", "17hK9XqZr8K9mMV9BS4BPbEB7VAGYxkmVV", "1P2wWctGp3YdaRXviVrjEc8Yy1Gm29e3zt", "1KPFGEbdDQJUTFG3JBth9rytNutSpk2WYH", "14GyPW5CZhz1PtMV9CgwCEBquXyPnr1pRK"}, "1637016673", "1643216020"))
	// _, _, _ = core.Addnewblock("140edf6c44171ab7c93cb2df9da9cb56d253757c4b16badfde6cdfba86514b", "BALLOT", "", core.CreateBallot("General-Elections-2023", []string{"1A3hszZSQ3X3uTKM4vBApsmdAzVb1JNesQ", "17hK9XqZr8K9mMV9BS4BPbEB7VAGYxkmVV", "1P2wWctGp3YdaRXviVrjEc8Yy1Gm29e3zt", "1KPFGEbdDQJUTFG3JBth9rytNutSpk2WYH", "14GyPW5CZhz1PtMV9CgwCEBquXyPnr1pRK"}, "1637016673", "1611681673"))

	//1. Address : 1A3hszZSQ3X3uTKM4vBApsmdAzVb1JNesQ , 29652bac6f13a7c2f77a77d9e0c6bd185b45aaab4541034c81281f8f9080f8fa
	//  2.17hK9XqZr8K9mMV9BS4BPbEB7VAGYxkmVV , 78b6d844c33fc49cdb432923d66d87605405419f02a64f9d5a950f59f7706ac0
	// 3. 1P2wWctGp3YdaRXviVrjEc8Yy1Gm29e3zt , a769cfbf683aa7804c5c9f6d5c53160fb3320e86e1a8aa1ae73c1f2aec9cf40b 
	// 4. 1KPFGEbdDQJUTFG3JBth9rytNutSpk2WYH , 8824cc98ffdafc26b86a21d00a33153501d33926e90be16271cfad7c1e292f97
	// 5 . 14GyPW5CZhz1PtMV9CgwCEBquXyPnr1pRK , 457e578b0bc0f6910955f74c0a1aa189ca8a4bca558728179ddde91a7cfac477

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
