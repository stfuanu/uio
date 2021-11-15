package core

import (
	"crypto/elliptic"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	// "regexp"
	"crypto/ecdsa"
	"crypto/rand"
	"math/big"
	"time"

	// "github.com/mr-tron/base58"
	"log"
	// api "pro/web"
)

var difficulty int = 5

type Block struct {
	Index     int    `json:"index"`
	Timestamp string `json:"timestamp"`
	Votes     []Vote `json:"votes"`
	Nonce     int    `json:"nonce"`
	PrevHash  string `json:"prevhash"`
	Hash      string `json:"hash"`
}

type Vote struct { // add timestamp of transaction
	TXID      string    `json:"txhash"`
	Timestamp string    `json:"timestamp"`
	Voter     VoterInfo `json:"voter"`
	Candidate string    `json:"candidate"`
	Status    bool      `json:"status"`
	Contract  string    `json:"contract"`
}

type Ballot struct {
	BTXhash         string   `json:"btxhash"`
	ContractHash    string   `json:"contracthash"`
	ElectionName    string   `json:"name"`
	Candidates      []string `json:"candidates"`
	TotalCandidates int      `json:"totalcandidates"`
	StartTimeStamp  string   `json:"start"`
	EndTimeStamp    string   `json:"end"`
}

// GODLEVELL STUFF ON ERORS : https://dave.cheney.net/2016/04/27/dont-just-check-errors-handle-them-gracefully

type VoterInfo struct {
	Address   string `json:"address"`
	PublicKey string `json:"pubkey"`
	Signature string `json:"signature"`
}

func CalculateTXN_HASH(vtx Vote) string {
	record := vtx.Voter.Address + vtx.Candidate + strconv.FormatBool(vtx.Status) + vtx.Contract + vtx.Timestamp
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

type AjaxLive struct {
	Total_Blocks         int    `json:"total_blocks"`
	Total_txn            int    `json:"total_txn"`
	Total_btx            int    `json:"total_btx"`
	Total_vtx            int    `json:"total_vtx"`
	Total_ballots        int    `json:"total_ballots"`
	Latest_Block         int    `json:"total_bloc"`
	Latest_BlockHash     string `json:"latest_blocHash"`
	Latest_TxnHash       string `json:"latest_txnHash"`
	LatestBlockTimeStamp string `json:"latest_block_stamp"`
	LatestVoteTimeStamp  string `json:"latest_vote_stamp"`
}

type ElectionInfo struct {
	ThisBallot      Ballot         `json:"ballot"`
	Candidate_votes map[string]int `json:candyvotes`
	Total_Votes     int            `json:"tvote"`
}

var AllElectionLive []ElectionInfo

func GetElectionInfo(btxhash string) ElectionInfo {
	var ball Ballot
	// var candyvotes map[string]int
	candyvotes := make(map[string]int)
	var tvot int
	// error : https://yourbasic.org/golang/gotcha-assignment-entry-nil-map/

	for _, Ballott := range Ballots {
		if Ballott.BTXhash == btxhash {
			btxhash = Ballott.BTXhash
			ball = Ballott
			break
		}
	}
	for _, candy := range ball.Candidates {
		candyvotes[candy] = 0
	}

	for _, block := range Blockchain {
		for _, vtx := range block.Votes {
			if vtx.Contract == btxhash {
				tvot++
				for candyaddress, _ := range candyvotes {
					if vtx.Candidate == candyaddress {
						candyvotes[candyaddress] = candyvotes[candyaddress] + 1

					}
				}
			}
		}
	}

	elinfo := ElectionInfo{ball, candyvotes, tvot}

	return elinfo
}

var RealLive AjaxLive

func GetLiveStat() {

	// latest wale creation ke time update ho jayenge
	// Total wale one time loop se
	livestat := &AjaxLive{
		Total_Blocks:  len(Blockchain),
		Total_ballots: len(Ballots),
	}
	livestat.SetVoteStat()

	RealLive = *livestat
	// fmt.Println(RealLive)

	// return *livestat

}

func CreateAllElectionInfo() {
	var btxhash_o string

	for _, ballot := range Ballots {
		btxhash_o = ballot.BTXhash

		elkinfo := GetElectionInfo(btxhash_o)
		AllElectionLive = append(AllElectionLive, elkinfo)
	}
}

func (stat *AjaxLive) SetVoteStat() {

	for _, block := range Blockchain {
		// stat.Total_Blocks++
		for _, vtx := range block.Votes {
			stat.Total_txn++
			if vtx.Candidate == "SMART_CONTRACT" {
				stat.Total_btx++
				// stat.Total_ballots++

			} else {
				stat.Total_vtx++
			}
		}
	}
	Block := Blockchain[stat.Total_Blocks-1]
	vtx := Block.Votes[len(Block.Votes)-1]

	stat.Latest_TxnHash = vtx.TXID
	stat.LatestVoteTimeStamp = vtx.Timestamp
	stat.LatestBlockTimeStamp = Block.Timestamp
	stat.Latest_BlockHash = Block.Hash
	stat.Latest_Block = Block.Index

}

var Blockchain []Block
var Ballots []Ballot

func CalContractHash(bb Ballot) string {

	var sss strings.Builder
	for _, s := range bb.Candidates {
		// fmt.Println(s)
		sss.WriteString(s)
	}

	record := bb.ElectionName + bb.StartTimeStamp + bb.EndTimeStamp + sss.String() + strconv.Itoa(bb.TotalCandidates)
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

func CreateBallot(name string, candies []string, start string, end string) string {
	// []string{"12tmRt6AADfQhfruF3RzFDdNhjiSEkwMvF", "1DngEcP2tCkxZNiAmm3Ar8VXXAAvAPfm8E", "1HRK5H21wFguq5ecJF8FvQk28qYnPz1Qb9"}

	NewBallot := Ballot{
		ElectionName:    name,
		Candidates:      candies,
		StartTimeStamp:  start,
		EndTimeStamp:    end,
		TotalCandidates: len(candies),
	}
	NewBallot.ContractHash = CalContractHash(NewBallot)
	// return NewBallot

	jsonString, _ := json.Marshal(NewBallot)
	enco := string(base58Encode(jsonString))
	return enco

	// //37EuM2iWfq5VzDyvBVPnaAsdD8steUBYqNiLfYjcJJ1rVoYWaoVBqWn89qZAvmC7Up34GNcmYQY5soD9vXyscz4WvH3qxaqnFrF92tmd3L8S9JKgrCJT1nRdK3MgmqvEaSwLAJdUZAMBmJpnBDavLKjLg1rEnoifaHqFhrWsXDMyDGVexQaoU9K3J

	// deco := base58Decode([]byte(enco))
	// fmt.Println(string(deco))
}

func Getlocaltime() string {
	loc, _ := time.LoadLocation("Asia/Kolkata")
	now := time.Now().In(loc)
	// return now.Format("2006-01-02 15:04:05.000000")
	return now.String()
}

// func CoinbaseVote() *Vote {

// 	cb := &Vote{
// 		Voter:     "me",
// 		Candidate: "anu"
// 	}

// 	return cb
// }

func LoadChains() {
	Blockchain = BlockchainFileToArray("votes.json")
	Ballots = BallotFileToArray("ballots.json")
}

func Genesisblock() {

	if len(Blockchain) != 0 {
		// fmt.Println("first block")
		return
	}
	bloc := Block{}

	// t := Getlocaltime()
	// fmt.Println(t)

	vinfo := VoterInfo{"VOTER_ADDRESS", "VOTER_PUBLIC_KEY", "SIGNATURE_PLUS_R"}

	cb := Vote{
		Timestamp: Getlocaltime(),
		Voter:     vinfo,
		Candidate: "CANDIDATE_ADDRESS",
		Status:    true,
		Contract:  "CONTRACT_HASH",
	}
	cb.TXID = CalculateTXN_HASH(cb)

	bloc = Block{
		Index:     0,
		Timestamp: Getlocaltime(),
		Votes:     []Vote{cb},
		// Nonce:     0,
		PrevHash: "0",
		// Hash:      CalHash(bloc),
	}

	bloc.Nonce, bloc.Hash = MineBlock(bloc)

	// bloc = Block{
	// 	Nonce: noncee,
	// 	Hash: hashh,
	// }

	Blockchain = append(Blockchain, bloc)
}

func HashVali(hash string) bool {
	// r, _ := regexp.Compile("^0{+strconv.Itoa(difficulty)+}")
	// r, _ := regexp.Compile("^xxx")
	// return r.MatchString(hash)

	prefix := strings.Repeat("0", difficulty)
	// prefix := "00"
	return strings.HasPrefix(hash, prefix)
}

func MineBlock(nblk Block) (int, string) {
	better_hash := nblk.Hash
	nblk.Nonce = 0

	for {

		better_hash = CalHash(nblk)
		// fmt.Println(nblk.Nonce, better_hash)
		if !HashVali(better_hash) {
			nblk.Nonce++
			continue
		} else {
			break
		}
	}
	time.Sleep(2 * time.Second)
	return nblk.Nonce, better_hash

}

func (vtx *Vote) Sign(privKey ecdsa.PrivateKey, txhash string) {

	r, s, err := ecdsa.Sign(rand.Reader, &privKey, []byte(txhash))
	if err != nil {
		panic(err)
	}
	signatureplusr := append(r.Bytes(), s.Bytes()...)
	// fmt.Printf("signature: %x\n", sig) OR including r ?? idk
	// we need BOTH , because verify ke time r aur sig dono params mai hain!!
	// ex. : ecdsa.Verify(&PVT.PublicKey, hash[:], r, sig)

	vtx.Voter.Signature = fmt.Sprintf("%x", signatureplusr)

}

func Verify(vtx Vote) bool {
	// txCopy := tx.TrimmedCopy()
	// curve := elliptic.P256()

	// for inID, vin := range tx.Vin {

	// prevTx := prevTXs[hex.EncodeToString(vin.Txid)]

	// refVoutPubKeyHash := prevTx.Vout[vin.Vout].PubKeyHash

	// // check that the spend coin is owned by vin.PubKey
	// if !bytes.Equal(PublicKeyHash(vin.PubKey), refVoutPubKeyHash) {
	// 	return false
	// }

	// txCopy.Vin[inID].Signature = nil
	// // txCopy.Vin[inID].PubKey = prevTx.Vout[vin.Vout].PubKeyHash
	// txCopy.Vin[inID].PubKey = refVoutPubKeyHash
	// txCopy.ID = txCopy.Hash() // txID hash mil gayi

	// txCopy.Vin[inID].PubKey = nil

	// idhar r aur sig nikal lena hai BIGsignature se
	r := big.Int{}
	s := big.Int{}

	signatureplusrINbYTES := HexDecode(vtx.Voter.Signature)
	// fmt.Println(vtx.Voter.Signature , fmt.Sprintf("%s",signatureplusrINbYTES) )
	sigLen := len(signatureplusrINbYTES)
	r.SetBytes([]byte(signatureplusrINbYTES)[:(sigLen / 2)])
	s.SetBytes(signatureplusrINbYTES[(sigLen / 2):])

	// idhar (x,y) ki zaroorat hai &publickeyStruct banane mai
	// jo publickey-byte version se nikala
	x := big.Int{}
	y := big.Int{}
	PubkeyINbYtes := HexDecode(vtx.Voter.PublicKey)
	keyLen := len(PubkeyINbYtes)
	x.SetBytes(PubkeyINbYtes[:(keyLen / 2)])
	y.SetBytes(PubkeyINbYtes[(keyLen / 2):])

	rawPubKey := ecdsa.PublicKey{elliptic.P256(), &x, &y}

	if ecdsa.Verify(&rawPubKey, []byte(vtx.TXID), &r, &s) == false {
		return false
	}
	// }

	return true
}

func NewTXN(w Wallet, candidateID string, contract string) Vote {

	// fmt.Println("here txn", contract, candidateID)

	// fortestonly
	// w := MakeWallet()
	// vinfo := VoterInfo{voterID,w.PublicKey,"signature"}
	voterID := fmt.Sprintf("%s", w.Address())

	vinfo := VoterInfo{
		Address:   voterID,
		PublicKey: fmt.Sprintf("%x", w.PublicKey),
	}

	vtx := Vote{
		Timestamp: Getlocaltime(),
		Voter:     vinfo,
		Candidate: candidateID,
		Status:    false,
		Contract:  contract,
	}
	vtx.TXID = CalculateTXN_HASH(vtx)

	vtx.Sign(w.PrivateKey, vtx.TXID) // sets signature to sign+r !! we need r in verify

	return vtx

}

func NewBTX(w Wallet, Base58_Contract string) Vote { // voterid & candy(defualt) dono idhar hi ban rahe toh kya hi args pass krna

	// NewBallot.TotalCandidates = len(NewBallot.Candidates) // do this when making ballot & it's enco

	// refer jsonballot.go in /trash

	// fortestonly
	// fmt.Println("here btx", Base58_Contract)
	// w := MakeWallet()
	// vinfo := VoterInfo{voterID,w.PublicKey,"signature"}
	owneraddr := fmt.Sprintf("%s", w.Address())

	vinfo := VoterInfo{
		Address:   owneraddr,
		PublicKey: fmt.Sprintf("%x", w.PublicKey),
	}

	vtx := Vote{
		Timestamp: Getlocaltime(),
		Voter:     vinfo,
		Candidate: "SMART_CONTRACT",
		Status:    false,
		Contract:  Base58_Contract,
		// Contract:  contract, //ballot idhar hi save hoga
	}

	// jsonString, _ := json.Marshal(NewBallot) // ye sab api side ho jayega re
	vtx.TXID = CalculateTXN_HASH(vtx)

	vtx.Sign(w.PrivateKey, vtx.TXID) // sets signature to sign+r !! we need r in verify
	// fmt.Println(vtx)

	return vtx

}

// func CandidateBlock(voterID string, candidateID string ,Prevblock Block ) (Block ) {

// 	// var blk_new Block

// 	tx , Verify_Signature_Status := NewTXN(voterID , candidateID) // new txn
// 	candyBloc := Block{}

// 	if !Verify_Signature_Status {
// 		return
// 	}

// 	// vinfo := VoterInfo{voterID,"pubkey","signature"}

// 	// // t := Getlocaltime()
// 	// tx := Vote{
// 	// 	Voter:     vinfo,
// 	// 	Candidate: candidateID,
// 	// 	Value:		1 ,

// 	// }

// 	// tx.TXID = CalculateTXN_HASH(tx)

// 	candyBloc = Block{
// 		Index:     Prevblock.Index + 1,
// 		Timestamp: Getlocaltime(),
// 		Votes:      []Vote{tx},
// 		// Nonce:     0,
// 		PrevHash:  Prevblock.Hash,
// 		// Hash:      CalHash(bloc),
// 	}

// 	return candyBloc
// }

// smart contracts are safe & immutable becoz ,
// I make a test.go file ,
// write go code/functions/struct/variables in it ,
// Hash tf out it , put that 256 hash inside VTX blockchain , make sure it contributes to block's hash .
// Now if someones changes anything in my code , hash changes & he's effd

// make smart.go , it has variables with values ,
// example : total candidates : 5 , candidate 1 = samosaJi etc.. , hash it then after creating RULES
// this smart.go

// ex: voter chooses that he wants to vote in Lok sabha election
// we then pick loksabha struct/func/smart_contra , (already populated or we will fill it , hash it(values of variables) & it will have a 256hash to be included in txn)
// loksabha struct can have starting date & ending date OR specific dates for specific voter (PHASE wise voting)

// func Vote_Status(address string, contract string) bool {

// 	// if !ValidateAddress(address) {
// 	// 	return false , errors.New("NOTFOUND")
// 	// }

// 	for _, block := range Blockchain {
// 		for _, vtx := range block.Votes {
// 			if vtx.Voter.Address == address && vtx.Contract == contract {
// 				return true
// 			}
// 		}
// 	}
// 	return false
// }

// IDEA : define address's of all candidates in a array ,
// forloop check , candidate address exists in this Candy_array
// concat all address's from array , Hash it , make it contract_Address , include it in VTX struct , so it affects block + include it in block hash too

// (VOTE_COUNT , VOTE_STATUS)
// func GetContractDataFromTXHASH(contracthash string) string {

// 	for _, block := range Blockchain {
// 		for _, vtx := range block.Votes {
// 			if vtx.TXID == contracthash && vtx.Candidate == "SMART_CONTRACT" {
// 				return vtx.Contract
// 			}
// 		}
// 	}
// 	return "NAN"
// }

func Vote_StatusANDContractData(address string, contracthash string) (bool, string) {

	// if this param contacthash is 256 wala then it's send by voter
	// if it's long encoded then it's of Encoded Data

	var votestat bool
	var ContractDATA string
	ContractDATA = "NAN"

	for _, block := range Blockchain {
		for _, vtx := range block.Votes {
			if vtx.TXID == contracthash && vtx.Candidate == "SMART_CONTRACT" { // verify this contract exists & get it's Data
				ContractDATA = vtx.Contract
			} else if vtx.Voter.Address == address && vtx.Contract == contracthash && vtx.Candidate != "SMART_CONTRACT" {
				votestat = true
			}
		}
	}
	return votestat, ContractDATA
}

func ContractDataExists(encodata string) (string, bool) {

	var databool bool
	var btxhash string

	for _, block := range Blockchain {
		for _, vtx := range block.Votes {
			if vtx.Contract == encodata && vtx.Candidate == "SMART_CONTRACT" {
				databool = true
				btxhash = vtx.TXID
			}
		}
	}
	return btxhash, databool
}

// for all addresses in a contract
func CountVotes(addr string, contract string) int {
	vote_count := 0

	for _, block := range Blockchain {
		for _, vtx := range block.Votes {
			if vtx.Candidate == addr && vtx.Contract == contract {
				vote_count++
			}
		}
	}

	return vote_count

}

func Addnewblock(voter_pvtkey string, VoteORBallot string, CANDY_ADDRESS string, CONTRACT_HASH string) (bool, error, string) {
	var err error
	var voter_wallet *Wallet

	voter_wallet, err = FindWalletByPvtKey(voter_pvtkey)
	if err != nil {
		return false, err, "NAN"
	}

	VOTER_ADDRESS := voter_wallet.Real.Address

	// ADD Contract ballots txns to vote.json , make more if/elses for contract timestamp/candidate exists stuff .

	var tx Vote
	var new_ballot Ballot
	var verify bool

	if VoteORBallot == "VOTE" {

		if !ValidateAddress(VOTER_ADDRESS) {
			// return false , errors.New("Voter Address")
			return false, fmt.Errorf("!Invalid Voter Address : %s ", VOTER_ADDRESS), "NAN"

		} else if !ValidateAddress(CANDY_ADDRESS) {
			return false, fmt.Errorf("!Invalid Candidate Address : %s ", CANDY_ADDRESS), "NAN"
		}

		Vote_Status, EncodedContract := Vote_StatusANDContractData(VOTER_ADDRESS, CONTRACT_HASH)

		if Vote_Status {
			return false, fmt.Errorf("Error : Already Voted as %s --- OR --- Same Contract Hash/Data as %s  , Plz check ", VOTER_ADDRESS, CONTRACT_HASH), "NAN"
		} else if EncodedContract == "NAN" {
			return false, fmt.Errorf("Error : Contract doesn't Exist : %s ", VOTER_ADDRESS, CONTRACT_HASH), "NAN"
		}
		CONTRACT_DATA := EncodedContract

		BallotStruct, verify := CheckCon(CONTRACT_DATA)
		if !verify {
			return false, fmt.Errorf("Contract Json format error : %s ", CONTRACT_DATA), "NAN"
		} else if !CandyExists(BallotStruct.Candidates, CANDY_ADDRESS) {
			return false, fmt.Errorf("Error : Candidate - %s Doesn't Exist in %v from Contract %s ", BallotStruct.Candidates, CANDY_ADDRESS, CONTRACT_HASH), "NAN"
		}

		// check if contract is available & candidates are there in encoded
		// GET ballot type using checkcon

		tx = NewTXN(*voter_wallet, CANDY_ADDRESS, CONTRACT_HASH)

		if !Verify(tx) {
			return false, fmt.Errorf("Signature Verification Error : false "), "NAN"
		}

		// VERIFIED !!
		// tx.Status = true
		// SEND it to all other nodes & add it your own local blockchain

	} else if VoteORBallot == "BALLOT" {
		CANDY_ADDRESS = "SMART_CONTRACT"

		if !ValidateAddress(VOTER_ADDRESS) {
			// return false , errors.New("Voter Address")
			return false, fmt.Errorf("!Invalid Voter Address : %s ", VOTER_ADDRESS), "NAN"
		} // CANDY ADDRESS ALREADY HARDCODED HAI , so no need for checks there

		btxhash, boolexist := ContractDataExists(CONTRACT_HASH)

		if boolexist { // same contractDATAENCO exists already
			fmt.Println("Already exists contra", btxhash)
			return false, fmt.Errorf("Error : Contract Data already exists at : %s , \nPassedContract : %s ", btxhash, CONTRACT_HASH), "NAN"
		}

		new_ballot, verify = CheckCon(CONTRACT_HASH)
		if !verify {
			return false, fmt.Errorf("Contract Json format error : %s ", CONTRACT_HASH), "NAN"
		}

		tx = NewBTX(*voter_wallet, CONTRACT_HASH)
		new_ballot.BTXhash = tx.TXID
		if !Verify(tx) {
			return false, fmt.Errorf("Signature Verification Error : false "), "NAN"
		}

	} else {
		return false, fmt.Errorf("TXN TYPE NOT MENTIONED "), "NAN"
	}

	// VERIFIED !!
	tx.Status = true

	// if !ValidateAddress(VOTER_ADDRESS) && VOTER_ADDRESS == "SMART_CONTRACT" {
	// 	// return false , errors.New("Voter Address")
	// 	return false, fmt.Errorf("!Invalid Voter Address : %s ", VOTER_ADDRESS), "NAN"

	// } else if !ValidateAddress(CANDY_ADDRESS) {
	// 	return false, fmt.Errorf("!Invalid Candidate Address : %s ", CANDY_ADDRESS), "NAN"
	// }

	// if !ContractSafe(CONTRACT_HASH) {
	// 	return false, fmt.Errorf("Contract Doesn't Exists or Match , check /api/Ballots.json : %s ", CONTRACT_HASH), "NAN"
	// }
	// iske func main , ballots pe nahi blocchain/votes par loop maarna hai !
	// ya fir baarbaar blockchain/votes par loop maarne se aache ek baar mai hi saari info nikal lo

	// -------------------------------------------------------------------
	// if Vote_Status(VOTER_ADDRESS, CONTRACT_HASH) {
	// 	return false, fmt.Errorf("Error : Already Voted : %s ", VOTER_ADDRESS), "NAN"
	// }
	// encod := GetContractDataFromTXHASH(CONTRACT_HASH)
	// if encod != "NAN" {
	// 	CONTRACT_HASH = encod
	// } else {
	// 	return false, fmt.Errorf("Smart Contract Not found : %s ", CONTRACT_HASH), "NAN"
	// }
	// ----------------------------------------------------------------

	// Merged Votestatus & if true then contractData in One (One Loop only)

	// Vote_Status, EncodedContract := Vote_StatusANDContractData(VOTER_ADDRESS, CONTRACT_HASH)

	// if Vote_Status {
	// 	return false, fmt.Errorf("Error : Already Voted as %s --- OR --- Same Contract Hash/Data as %s  , Plz check ", VOTER_ADDRESS, CONTRACT_HASH), "NAN"
	// } else if EncodedContract != "NAN" {
	// 	CONTRACT_HASH = EncodedContract
	// }

	// var tx Vote

	// Now CONTRACT_HASH in it's ENCODED_DATA FORM -------------------->>>

	// if CANDY_ADDRESS == "SMART_CONTRACT" && EncodedContract == "NAN" { // check here later ki bas58 contract json ban bhi rha ya nahi !
	// 	// and also verify that it properly casts into Ballot Struct mold , THAT WILL verify ki key/value etc.. sab sahi hai

	// 	// 0). Same contract already exists in blockchain

	// 	// 1). Decoded bas58 is indeed json format & has all those stuff wchi we need & len(candidates bhi set hai)
	// 	_, verify := CheckCon(CONTRACT_HASH)
	// 	if !verify {
	// 		return false, fmt.Errorf("Contract Json format error : %s ", CONTRACT_HASH), "NAN"
	// 	}

	// 	// later will check regarding timestamp decodedballot checkcon se mil jayega

	// 	tx = NewBTX(CONTRACT_HASH) // baad hai voter bhi pass krna hai , abhi toh idhar hi wallet hai

	// } else {

	// 	// check if contract is available & candidates are there in encoded

	// 	tx = NewTXN(VOTER_ADDRESS, CANDY_ADDRESS, CONTRACT_HASH)
	// }

	// if !Verify(tx) {
	// 	return false, fmt.Errorf("Signature Verification Error : false "), "NAN"
	// }

	// VERIFIED !!
	// tx.Status = true
	// SEND it to all other nodes & add it your own local blockchain

	// fmt.Println(tx)

	Prevblock := Blockchain[len(Blockchain)-1]

	// blk_new := CandidateBlock(voterID , candidateID , Prevblock)
	// blk_new := Block{}
	// var blk_new Block

	blk_new := Block{
		Index:     Prevblock.Index + 1,
		Timestamp: Getlocaltime(),
		Votes:     []Vote{tx}, // append vtx
		// Nonce:     0,
		PrevHash: Prevblock.Hash,
		// Hash:      CalHash(bloc),
	}

	// t := time.Now()

	// blk_new.Index = Prevblock.Index + 1
	// blk_new.Timestamp = Getlocaltime()
	// cb := &Vote{
	// 	Voter:     voterID,
	// 	Candidate: candidateID,
	// }

	// blk_new = Block{
	// 	Votes:      []*Vote{cb},
	// }
	// blk_new.Votes.Voter = voterID
	// blk_new.Votes.Candidate = candidateID
	// blk_new.PrevHash = Prevblock.Hash
	blk_new.Nonce, blk_new.Hash = MineBlock(blk_new)

	// fmt.Println(blk_new.Votes)
	// fmt.Println(blk_new.Index,blk_new)

	// adjust settings before next mining
	// adjustdiff()

	// blk_new.Nonce =
	// blk_new.Hash = CalHash(blk_new)

	// verify
	if !Valid(blk_new, Prevblock) {
		// fmt.Println("invalid??")
		return false, fmt.Errorf("Block is Invalid!"), "NAN"
	}
	// now append it (maybe idk)
	Blockchain = append(Blockchain, blk_new)
	file, _ := json.MarshalIndent(Blockchain, "", " ")

	err = ioutil.WriteFile("votes.json", file, 0644)
	if err != nil {
		log.Println(err)
	}
	// time.Sleep(2 * time.Second)
	if VoteORBallot == "BALLOT" {
		Ballots = append(Ballots, new_ballot)
		file, _ := json.MarshalIndent(Ballots, "", " ")

		err := ioutil.WriteFile("ballots.json", file, 0644)
		if err != nil {
			log.Println(err)
		}
	}

	latest_Hash := Blockchain[len(Blockchain)-1].Hash
	fmt.Println("New Block Added ", blk_new.Index, blk_new.Hash, UnixTime(blk_new.Timestamp))

	fmt.Println("Vote_Count : ", blk_new.Votes[0].Candidate, CountVotes(blk_new.Votes[0].Candidate, blk_new.Votes[0].Contract))
	// Top()

	if latest_Hash == blk_new.Hash {
		blk_new.UpdateBlockLive()
		tx.UpdateVoteLive(VoteORBallot) // later would have to change maybe
		tx.UpdateAllEllinfo(VoteORBallot)
		return true, nil, latest_Hash
	} else {
		return false, fmt.Errorf("Error : Block Not Added to Blockchain "), "NAN"
	}

}

func (block *Block) UpdateBlockLive() {
	RealLive.Latest_BlockHash = block.Hash
	RealLive.Latest_Block = block.Index
	RealLive.LatestBlockTimeStamp = block.Timestamp
	RealLive.Total_Blocks = len(Blockchain)

}
func (tx *Vote) UpdateVoteLive(voteorball string) {
	RealLive.Latest_TxnHash = tx.TXID
	RealLive.Total_txn++
	RealLive.LatestVoteTimeStamp = tx.Timestamp

	if voteorball == "BALLOT" {
		RealLive.Total_btx++
		// latest btxhash maybe next idk

	} else if voteorball == "VOTE" {
		RealLive.Total_vtx++
	}
}

func (tx *Vote) UpdateAllEllinfo(voteorball string) {

	if voteorball == "BALLOT" { // create new electioninfo if it's ballot type (i.e , it's new txn and new ballot)
		ball, verify := CheckCon(tx.Contract)
		if !verify {
			fmt.Println("func: UpdateAllEllinfo , func:CheckCon")
		}

		var tvot int
		candyvotes := make(map[string]int)
		for _, candy := range ball.Candidates {
			candyvotes[candy] = 0
		}

		newelk := ElectionInfo{ball, candyvotes, tvot}

		AllElectionLive = append(AllElectionLive, newelk)

	} else if voteorball == "VOTE" { // if vote type Then we just have to update candidateVotes and total votes in this contract/btxhash
		for _, elk := range AllElectionLive {
			if elk.ThisBallot.BTXhash == tx.Contract {
				elk.Candidate_votes[tx.Candidate] = elk.Candidate_votes[tx.Candidate] + 1 // add 1 to vote count of candidate
				elk.Total_Votes++
			}
		}
	}

}

func CandyExists(candidates []string, cand string) bool {

	for _, candies := range candidates {
		if candies == cand {
			return true
		}
	}

	return false
}

func CheckCon(base58enco string) (Ballot, bool) {
	var nbl Ballot
	jsonbyte := base58Decode([]byte(base58enco))
	err := json.Unmarshal(jsonbyte, &nbl)
	if err != nil {
		fmt.Println("error:", err)
		return nbl, false
	}
	nbl.TotalCandidates = len(nbl.Candidates)
	return nbl, true
}

func CalHash(thisblock Block) string {

	// Hash of Transaction will be calculated differently : MERKLE TREE (LATER) :: REMOVE VOTER/CANDIDATE_ID for now
	record := strconv.Itoa(thisblock.Index) + thisblock.Timestamp + thisblock.PrevHash + strconv.Itoa(thisblock.Nonce)
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

func Valid(newblk, prevblk Block) bool {
	if prevblk.Index+1 != newblk.Index {
		// fmt.Println("1",prevblk.Index+1,newblk.Index)
		return false
	} else if newblk.PrevHash != prevblk.Hash {
		// print("2")
		return false
	} else if CalHash(newblk) != newblk.Hash {
		// print("3")
		return false
	}
	return true
}

func print(s string) {
	fmt.Println(s)
}
func Handle(err error) {
	if err != nil {
		log.Panic(err)
	}
}

func PrintblockchainStdout() {
	bb, err := json.MarshalIndent(Blockchain, "", "  ")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(bb))
}

// func adjustdiff() {

// 	length := len(Blockchain)
// 	lasttime := UnixTime(Blockchain[0].Timestamp)
// 	thistime := 420
// 	// fmt.Println(lasttime)
// 	tt := 0

// 	// currently loops whole blockchain , later will cut it down to last 10-20

// 	for _, block := range Blockchain[1:] {

// 		thistime = UnixTime(block.Timestamp)
// 		tt = tt + (thistime - lasttime)
// 		lasttime = thistime
// 		// fmt.Println(thistime , lasttime ,tt)
// 	}
// 	// fmt.Println(lasttime,thistime)

// 	var avgt float64 = float64(tt) / float64(length)

// 	if avgt > 7 {
// 		difficulty = difficulty - 1
// 	} else if avgt < 1 {
// 		// fmt.Println("here")
// 		difficulty++
// 	}
// 	// fmt.Println("AvgTime : ",avgt, difficulty)
// 	fmt.Printf("AvgTime : %.2fs , Difficulty : %d   ", avgt, difficulty)

// }

// // const longForm = ""
func UnixTime(tim string) int {
	loc, _ := time.LoadLocation("Asia/Kolkata")
	// 2021-10-08 19:55:56.7413308 +0530 IST
	lay := "2006-01-02 15:04:05 MST"
	// t, err := time.Parse(lay, "2021-10-08 20:11:06 +0530 IST")
	t, err := time.ParseInLocation(lay, strings.Replace(tim, "+0530", "", -1), loc)

	// t, err := time.Parse(lay, tim)
	if err != nil {
		fmt.Println(err)
	}
	return int(t.UTC().Unix())
}
