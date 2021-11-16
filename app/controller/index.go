package controller

import (
	"encoding/json"
	"net/http"

	"uio/app/shared/session"
	"uio/app/shared/view"
	"uio/core"

	"github.com/josephspurrier/csrfbanana"
)

// IndexGET displays the home page
func IndexGET(w http.ResponseWriter, r *http.Request) {
	// Get session
	sess := session.Instance(r)
	// error this : caused by: gob: type core.RealWallet has no exported fields
	// In go, fields and variables that start with an Uppercase letter are "Exported", and are visible to other packages.
	// Fields that start with a lowercase letter are "unexported", and are only visible inside their own package.
	// https://stackoverflow.com/questions/40256161/exported-and-unexported-fields-in-go-language/40256229
	// fmt.Println(sess, sess.Values["id"], sess.Values["wallet"])
	// fmt.Println(sess.Values["token"])

	if sess.Values["id"] != nil {
		// Display the view
		v := view.New(r)
		v.Name = "index/auth2"
		v.Vars["first_name"] = sess.Values["first_name"]
		v.Vars["token"] = csrfbanana.Token(w, r, sess)
		v.Vars["wallet"] = sess.Values["wallet"]
		// v.Vars["addrinfo"] = sess.Values["addrinfo"]
		v.Vars["ballots"] = core.Ballots
		jsonString, _ := json.Marshal(core.Ballots)
		v.Vars["ballotjson"] = string(jsonString)
		// fmt.Println(v.Vars["ballotjson"])
		v.Render(w)
	} else {
		// Display the view
		v := view.New(r)
		v.Name = "index/anon"
		v.Render(w)
		// fmt.Println("kkwk")
		return
	}
}
