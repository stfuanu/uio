package controller

import (
	"net/http"

	"uio/app/shared/session"
	"uio/app/shared/view"
	"uio/core"
)

// NotepadCreatePOST handles the note creation form submission
func GenerateWallet(w http.ResponseWriter, r *http.Request) {
	// Get session

	NEW_WALLET := &core.Wallet{}
	sess := session.Instance(r)
	pvtkey := r.FormValue("pvtkey")
	buttoName := r.FormValue("change123")
	// fmt.Println(buttoName)
	// this doesnot work

	// ; document.getElementById('walletform').submit();
	if buttoName == "NEW_WALLET" {
		// fmt.Println("new_wall")
		NEW_WALLET = core.MakeWallet()
		sess.AddFlash(view.Flash{"New Wallet Created! : " + NEW_WALLET.Real.Address, view.FlashSuccess})
		sess.Values["wallet"] = NEW_WALLET
		sess.Values["addrinfo"] = core.GetAllInfoByAddr(NEW_WALLET.Real.Address)
		sess.Save(r, w)

		// Display the same page
		//
		IndexGET(w, r)
		// http.Redirect(w, r, "/", http.StatusFound)

	} else if buttoName == "FIND_ADDRESS" {

		if pvtkey == "" {
			sess.AddFlash(view.Flash{Message: "Private Key Needed , for finding Address!", Class: view.FlashError})
			sess.Save(r, w)
			http.Redirect(w, r, "/", http.StatusFound)
			return
		}
		var err error
		NEW_WALLET, err = core.FindWalletByPvtKey(pvtkey)
		if err != nil {
			sess.AddFlash(view.Flash{Message: "Error : Couldn't reproduce PrivateKey .", Class: view.FlashError})
			sess.Save(r, w)
			// IndexGET(w, r)
			http.Redirect(w, r, "/", http.StatusFound)
			return
		}
		sess.Values["wallet"] = NEW_WALLET
		sess.Values["addrinfo"] = core.GetAllInfoByAddr(NEW_WALLET.Real.Address)
		sess.Save(r, w)

		// Display the same page
		//
		IndexGET(w, r)
		// http.Redirect(w, r, "/", http.StatusFound)

	}

}
