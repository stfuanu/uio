package controller

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	// "errors"
	"uio/app/shared/session"
	"uio/app/shared/view"
	core "uio/core"

	"github.com/gorilla/context"
	"github.com/julienschmidt/httprouter"
)

// w.Write([]byte("Gorilla!\n"))

type Jsondata struct {
	Voter     string `json:"voter"`
	Candidate string `json:"candidate"`
	Contra    string `json:"contract"`
}

type BallData struct {
	Voterpvtkey string     `json:"voter"`
	Contra      string     `json:"contract"` // candidate to default hi rahega na
	Createdata  Jsonballot `json:"jsonballot"`
}

type Jsonballot struct {
	Name       string   `json:"naam"`
	Candidates []string `json:"candidates"`
	Start      string   `json:"start"`
	End        string   `json:"start"`
}

type ResponseToVoter struct {
	Status    int    `json:"status"`
	Message   string `json:"message"`
	Timestamp string `json:"timestamp"`
	Hash      string `json:"hash,omitempty"`
}

func APIGET(w http.ResponseWriter, r *http.Request) {
	// Display the view
	v := view.New(r)
	v.Name = "api/api"
	v.Render(w)
}

func LiveGet(w http.ResponseWriter, r *http.Request) {
	// Display the view
	v := view.New(r)
	v.Name = "result/live"
	v.Render(w)
}

func GetBlockchain(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	bb, err := json.MarshalIndent(core.Blockchain, "", "  ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	io.WriteString(w, string(bb))
}

func LiveStat(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	bb, err := json.MarshalIndent(core.RealLive, "", "  ")
	// fmt.Println(string(bb))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	io.WriteString(w, string(bb))
}

// GetBallots
func GetBallots(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	bb, err := json.MarshalIndent(core.Ballots, "", "  ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	io.WriteString(w, string(bb))
}

func Getaddrinfo(w http.ResponseWriter, r *http.Request) {

	var params httprouter.Params
	params = context.Get(r, "params").(httprouter.Params)
	Hashy := params.ByName("addr")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	// reqhash := mux.Vars(r)
	// fmt.Println(reqhash)

	// for _, block := range core.Blockchain {
	// 	if block.Hash == Hashy { 

	addrinfoo := core.GetAllInfoByAddr(Hashy)
	blk, err := json.MarshalIndent(addrinfoo, "", "  ")
	// fmt.Println(string(blk))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	io.WriteString(w, string(blk))

}

func GetElectionInfoWeb(w http.ResponseWriter, r *http.Request) {

	var params httprouter.Params
	params = context.Get(r, "params").(httprouter.Params)
	Hashy := params.ByName("btxhash")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	// reqhash := mux.Vars(r)
	// fmt.Println(reqhash)

	// for _, block := range core.Blockchain {
	// 	if block.Hash == Hashy {

	elkinfo := core.GetElectionInfo(Hashy)
	blk, err := json.MarshalIndent(elkinfo, "", "  ")
	// fmt.Println(string(blk))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	io.WriteString(w, string(blk))

	// }
	// }

}

func GetElectionInfoWeb_Fast(w http.ResponseWriter, r *http.Request) {

	var params httprouter.Params
	params = context.Get(r, "params").(httprouter.Params)
	Hashy := params.ByName("btxhash")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	// reqhash := mux.Vars(r)
	// fmt.Println(reqhash)
	var elkinfo core.ElectionInfo

	for _, ell := range core.AllElectionLive {
		if ell.ThisBallot.BTXhash == Hashy {
			elkinfo = ell
			break

		}
	}
	// 	if block.Hash == Hashy {

	// elkinfo := core.GetElectionInfo(Hashy)
	blk, err := json.MarshalIndent(elkinfo, "", "  ")
	// fmt.Println(string(blk))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	io.WriteString(w, string(blk))

	// }
	// }

}

func GetALLElectionInfoWeb(w http.ResponseWriter, r *http.Request) {

	// var params httprouter.Params
	// params = context.Get(r, "params").(httprouter.Params)
	// Hashy := params.ByName("btxhash")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	// reqhash := mux.Vars(r)
	// fmt.Println(reqhash)

	blk, err := json.MarshalIndent(core.AllElectionLive, "", "  ")
	// fmt.Println(string(blk))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	io.WriteString(w, string(blk))

}

func GetBlock(w http.ResponseWriter, r *http.Request) {

	var params httprouter.Params
	params = context.Get(r, "params").(httprouter.Params)
	Hashy := params.ByName("hash")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	// reqhash := mux.Vars(r)
	// fmt.Println(reqhash)

	for _, block := range core.Blockchain {
		if block.Hash == Hashy {
			blk, err := json.MarshalIndent(block, "", "  ")
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			io.WriteString(w, string(blk))

		}
	}

}

func NewVote(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	defer r.Body.Close()
	if r.Header.Get("Content-type") != "application/json" {
		w.WriteHeader(http.StatusUnsupportedMediaType)
		w.Write([]byte("415 - Unsupported Media Type. Only JSON files are allowed"))
		// return
	} else if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method not allowed."))
	} else {

		var newvote Jsondata

		// error :- web\api.go:75:24: cannot convert r.Body (type io.ReadCloser) to type []byte
		// note :- ye wala (io.reader) https://pkg.go.dev/encoding/json#NewDecoder

		abc := json.NewDecoder(r.Body) // outpts pointer

		errorAsreturn := abc.Decode(&newvote)
		// json.Unmarshal([]byte(r.Body), &newvote)
		// or maybe,,,, json.Unmarshal([]byte(r.Body.String()), &newvote)
		// err := decoder.Decode(&newvote);
		// http.StatusOK
		if errorAsreturn != nil {
			log.Fatal(errorAsreturn)
		}

		// if err != nil {
		// 	// panic(err)
		// 	fmt.Println(err)
		// 	// w.Write([]byte(err))
		//     return
		// }

		// New Block --->
		addedORnot, err, hashifany := core.Addnewblock("VOTE", newvote.Voter, newvote.Candidate, newvote.Contra)

		if addedORnot && err == nil {
			// fmt.Println("sukcess" , hashifany)
			io.WriteString(w, WhatHappened(http.StatusOK, "Success", hashifany))
			// w.Write([]byte("Method not allowed."))

		} else {
			// fmt.Println("NoHash" , hashifany)
			fullMessage := fmt.Sprintf("Failed : %v", err)
			io.WriteString(w, WhatHappened(http.StatusOK, fullMessage, hashifany))

		}

		// This will happen from funk.go
		// dat := WhatHappened(http.StatusOK, "Success")
		// bb, err := json.MarshalIndent(dat, "", "  ")
		// fmt.Println(string(bb))
	}

}

func WhatHappened(statuscode int, mess string, hashh string) string {
	// data := `{"status-sode": "%d", "message": "failed"}`
	// data := fmt.Sprintf(`{"status": "%d", "message": "%s"}`, statuscode, mess)

	data := &ResponseToVoter{
		Status:    statuscode,
		Message:   mess,
		Timestamp: core.Getlocaltime(),
		// Hash:      "newhash",
	}

	if hashh != "NAN" {
		data.Hash = hashh
	}

	// bb, err := json.MarshalIndent(data, "", "  ")
	// if err != nil {
	// 	fmt.Println(err)
	// 	// return
	// }
	// return *data
	bb, err := json.MarshalIndent(*data, "", "  ")
	// if err != nil {
	//     fmt.Println(err)
	// }
	core.Handle(err)
	// io.WriteString(w, string(bb))
	return string(bb)
}

func NewBallot(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	defer r.Body.Close()
	if r.Header.Get("Content-type") != "application/json" {
		w.WriteHeader(http.StatusUnsupportedMediaType)
		w.Write([]byte("415 - Unsupported Media Type. Only JSON files are allowed"))
		// return
	} else if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method not allowed."))
	} else {

		var NewBallot BallData
		errorAsreturn := json.NewDecoder(r.Body).Decode(&NewBallot)
		if errorAsreturn != nil {
			log.Fatal(errorAsreturn)
		}
		// jsonString, _ := json.Marshal(NewBallot)

		addedORnot, err, hashifany := core.Addnewblock(NewBallot.Voterpvtkey, "BALLOT", "SMART_CONTRACT", core.CreateBallot(NewBallot.Createdata.Name, NewBallot.Createdata.Candidates, NewBallot.Createdata.Start, NewBallot.Createdata.End))
		// core.CreateBallot("General-Elections-2021", []string{"1A3hszZSQ3X3uTKM4vBApsmdAzVb1JNesQ", "17hK9XqZr8K9mMV9BS4BPbEB7VAGYxkmVV"}, "1637016673", "1643216020")
		// core.Addnewblock("140edf6c44171ab7c93cb2df9da9cb56d253757c4b16badfde6cdfba86514b", "BALLOT", "", core.CreateBallot("General-Elections-2021", []string{"1A3hszZSQ3X3uTKM4vBApsmdAzVb1JNesQ", "17hK9XqZr8K9mMV9BS4BPbEB7VAGYxkmVV", "1P2wWctGp3YdaRXviVrjEc8Yy1Gm29e3zt", "1KPFGEbdDQJUTFG3JBth9rytNutSpk2WYH", "14GyPW5CZhz1PtMV9CgwCEBquXyPnr1pRK"}, "1637016673", "1643216020"))

		if addedORnot && err == nil {
			io.WriteString(w, WhatHappened(http.StatusOK, "Success", hashifany))

		} else {
			fullMessage := fmt.Sprintf("Failed : %v", err)
			io.WriteString(w, WhatHappened(http.StatusOK, fullMessage, hashifany))

		}
	}

}

func NewBallotweb(w http.ResponseWriter, r *http.Request) {
	// Get session
	sess := session.Instance(r)

	// Validate with required fields
	if validate, missingField := view.Validate(r, []string{"note"}); !validate {
		sess.AddFlash(view.Flash{"Field missing: " + missingField, view.FlashError})
		sess.Save(r, w)
		// NotepadCreateGET(w, r)
		return
	}
	// validate with database first
	// then gather all data to be supplied
	// figure out how to inlcude wallet with form data from user//

	// Get form values
	// content := r.FormValue("note")

	// Display the same page
	// NotepadCreateGET(w, r)

	// addedORnot, err, hashifany := core.Addnewblock("BALLOT", NewBallot.Voter, "SMART_CONTRACT", NewBallot.contra)

}
func NewVoteweb(w http.ResponseWriter, r *http.Request) {
	// Get session
	sess := session.Instance(r)

	// Validate with required fields
	if validate, missingField := view.Validate(r, []string{"poppvtkey", "popvoteorball", "popcontract", "popcandy"}); !validate {
		sess.AddFlash(view.Flash{"Field missing: " + missingField, view.FlashError})
		sess.Save(r, w)
		IndexGET(w, r)
		return
	}
	pvtkey := r.FormValue("poppvtkey")
	voter_wallet, err := core.FindWalletByPvtKey(pvtkey)
	if err != nil {
		sess.AddFlash(view.Flash{"Error while generating address from pvtkey : " + pvtkey, view.FlashError})
		sess.Save(r, w)
		IndexGET(w, r)
		return
	}
	// validate with database first
	// then gather all data to be supplied
	// figure out how to inlcude wallet with form data from user//

	// Get form values

	voteorball := r.FormValue("popvoteorball")
	contract := r.FormValue("popcontract")
	candyaddrr := r.FormValue("popcandy")

	addedORnot, err, hashifany := core.Addnewblock(pvtkey, voteorball, candyaddrr, contract)

	VOTER_ADDRESS := voter_wallet.Real.Address

	if err != nil || !addedORnot || hashifany == "NAN" {
		log.Println(err)
		sess.AddFlash(view.Flash{Message: err.Error(), Class: view.FlashError})
		sess.Values["addrinfo"] = core.GetAllInfoByAddr(VOTER_ADDRESS)
		sess.Save(r, w)
		http.Redirect(w, r, "/", http.StatusFound)
		IndexGET(w, r)
		return
	} else {
		jsonst, respodata := WhatHappened2(http.StatusOK, "Success", hashifany)
		sess.Values["whappjson"] = jsonst
		sess.Values["respodata"] = respodata
		sess.AddFlash(view.Flash{Message: jsonst, Class: view.FlashSuccess})
		sess.Values["addrinfo"] = core.GetAllInfoByAddr(VOTER_ADDRESS)
		sess.Save(r, w)
		http.Redirect(w, r, "/", http.StatusFound)
		IndexGET(w, r)

	}

	// Display the same page
	// NotepadCreateGET(w, r)

	// addedORnot, err, hashifany := core.Addnewblock("BALLOT", NewBallot.Voter, "SMART_CONTRACT", NewBallot.contra)

}

func WhatHappened2(statuscode int, mess string, hashh string) (string, ResponseToVoter) {

	data := &ResponseToVoter{
		Status:    statuscode,
		Message:   mess,
		Timestamp: core.Getlocaltime(),
		// Hash:      "newhash",
	}

	if hashh != "NAN" {
		data.Hash = hashh
	}
	bb, err := json.MarshalIndent(*data, "", "  ")

	core.Handle(err)
	return string(bb), *data
}

//external call , bahar se hi bsse58 bankr aana bhai

// func StartServer() {
// 	routervar := mux.NewRouter()

// 	routervar.HandleFunc("/api/votes.json", GetBlockchain).Methods("GET")
// 	routervar.HandleFunc("/api/vote/{hash}", GetBlock).Methods("GET")
// 	routervar.HandleFunc("/vote/newvtx", AppendNewBlock).Methods("POST")
// 	routervar.HandleFunc("/vote/newbtx", NewBallot).Methods("POST")

// 	log.Fatal(http.ListenAndServe(":80", routervar))
// }
