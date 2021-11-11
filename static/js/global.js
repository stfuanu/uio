$(function() {
    //$(document).foundation();
		
	// Hide any messages after a few seconds
    hideFlash();
});

// 14sNixMpyH2L1KQNaudxJm2CakhYgYEDtJ
// 140edf6c44171ab7c93cb2df9da9cb56d253757c4b16badfde6cdfba86514b

function hideFlash(rnum)
{    
    if (!rnum) rnum = '0';
    
    _.delay(function() {
        $('.alert-box-fixed' + rnum).fadeOut(300, function() {
            $(this).css({"visibility":"hidden",display:'block'}).slideUp();
            
            var that = this;
            
            _.delay(function() { that.remove(); }, 400);
        });
    }, 4000);
}

function showFlash(obj)
{
    $('#flash-container').html();
    $(obj).each(function(i, v) {
        var rnum = _.random(0, 100000);
		var message = '<div id="flash-message" class="alert-box-fixed'
		+ rnum + ' alert-box-fixed alert alert-dismissible '+v.cssclass+'">'
		+ '<button type="button" class="close" data-dismiss="alert" aria-label="Close"><span aria-hidden=true>&times;</span></button>'
		+ v.message + '</div>';
        $('#flash-container').prepend(message);
        hideFlash(rnum);
    });
}

function flashError(message) {
	var flash = [{Class: "alert-danger", Message: message}];
	showFlash(flash);
}

function flashSuccess(message) {
	var flash = [{Class: "alert-success", Message: message}];	
	showFlash(flash);
}

function flashNotice(message) {
	var flash = [{Class: "alert-info", Message: message}];
	showFlash(flash);
}

function flashWarning(message) {
	var flash = [{Class: "alert-warning", Message: message}];
	showFlash(flash);
}


function NewWallet(){
    document.getElementById("WhichButton").value = "NEW_WALLET";
    document.getElementById('walletform').submit();
}

function FindAddress(){
    document.getElementById("WhichButton").value = "FIND_ADDRESS";
    document.getElementById('walletform').submit();
}
function VoteNow(){

    if (document.getElementById("LockUnlock").innerText !== "Unlock & Edit") {
        alert("Verify & Lock Your vote , before Voting!")
        return
    }

    document.getElementById("popcontract").value = document.querySelector("#selectballot").value;
    document.getElementById("popcandy").value = document.querySelector("#ballcandidates").value;
    //  create form dynamically & submit it 
    // https://www.geeksforgeeks.org/how-to-create-a-form-dynamically-with-the-javascript/

    // alert("works")
    // reminder : validate all form input are there (later)
    document.getElementById('myForm').submit();
}

function closepopup() {
    document.getElementById("myForm").style.display = "none";
  }

function LockUnlock()
{
    if (document.getElementById("LockUnlock").innerText === "Verify & Lock") {


        if (document.getElementById("ballcandidates").value === '') {
            alert("Make sure , You have selected an Election & a Candidate to votw .")
            return
        } else if (/^1[a-zA-Z0-9]{30,40}/.test(document.getElementById("addr").value) === false ) {
            alert("Invalid Address!")
            return
        } else if (document.getElementById("ballcandidates").value === '') {
            alert("Private Key field , can't be empty")
            return
        }
        document.getElementById("LockUnlock").innerText = "Unlock & Edit"
        document.getElementById("LockUnlock").className = "btn btn-info";

        document.getElementById("ballcandidates").disabled = true;
        document.getElementById("addr").disabled = true;
        document.getElementById("pvtkey").disabled = true;
        document.getElementById("selectballot").disabled = true;

        document.getElementById("NewWalletbtn").style.pointerEvents="none";
        document.getElementById("NewWalletbtn").style.cursor="default";

        document.getElementById("FindAddressbtn").style.pointerEvents="none";
        document.getElementById("FindAddressbtn").style.cursor="default";
        alert("YOUR VOTE IS LOCKED , CLICK Vote !!!")

        return
    } else if (document.getElementById("LockUnlock").innerText === "Unlock & Edit") {
        document.getElementById("LockUnlock").innerText = "Verify & Lock"
        document.getElementById("LockUnlock").className = "btn btn-warning"; 
        
        document.getElementById("ballcandidates").disabled = false;
        document.getElementById("addr").disabled = false;
        document.getElementById("pvtkey").disabled = false;
        document.getElementById("selectballot").disabled = false;

        document.getElementById("NewWalletbtn").style.pointerEvents="auto";
        document.getElementById("NewWalletbtn").style.cursor="pointer";

        document.getElementById("FindAddressbtn").style.pointerEvents="auto";
        document.getElementById("FindAddressbtn").style.cursor="pointer";

        alert("Unlocked , You can Edit & resubmit !!!")
        return
    }
}

function display_ct7() {
var x = new Date()
var ampm = x.getHours( ) >= 12 ? ' PM' : ' AM';
hours = x.getHours( ) % 12;
hours = hours ? hours : 12;
hours=hours.toString().length==1? 0+hours.toString() : hours;

var minutes=x.getMinutes().toString()
minutes=minutes.length==1 ? 0+minutes : minutes;

var seconds=x.getSeconds().toString()
seconds=seconds.length==1 ? 0+seconds : seconds;

var month=(x.getMonth() +1).toString();
month=month.length==1 ? 0+month : month;

var dt=x.getDate().toString();
dt=dt.length==1 ? 0+dt : dt;

var x1=month + "/" + dt + "/" + x.getFullYear(); 
x1 = x1 + " - " +  hours + ":" +  minutes + ":" +  seconds + " " + ampm;
document.getElementById('ct7').innerHTML = x1;
display_c7();
 }
 function display_c7(){
var refresh=1000; // Refresh rate in milli seconds
mytime=setTimeout('display_ct7()',refresh)
}
display_c7()


function CopyClip() {
    /* Get the text field */

    var copyText = document.getElementById("pvtkey");
  
    copyText.select();
    document.execCommand('copy');
    
}

// function copyToClipboard(text) {

//     // exec command sirf textarea/input field pr work krta hai , becoz we can "select" only those tags ig ,
//     //  thats's y we create new input ofr copying p or h1 tag 

//     //  refer : https://www.youtube.com/watch?v=O269ctk5b5k
//     var keyy = document.querySelector("#pvtkey").value;

//     let input = document.crea
//     input.value = text;
//     input.focus();
//     input.select();
//     document.execCommand('copy');
//     input.parentNode.removeChild(input);
//   }

// ex:
// [{"contracthash":"ba56265934aa02b8ade28fd720119f627e8c51a08bec77f143cf4731b1b9697a","name":"BEST_COOK_2069",
// "candidates":["12tmRt6AADfQhfruF3RzFDdNhjiSEkwMvF","1HRK5H21wFguq5ecJF8FvQk28qYnPz1Qb9","1DngEcP2tCkxZNiAmm3Ar8VXXAAvAPfm8E"],
// "totalcandidates":3,"start":"1635032651","end":"1635043569"},
// 
// {"contracthash":"cf261d57c969a3d60c6ca8e2e31d0f812b160bf907f1b1d4052921984150ab1b","name":"BEST_SAMOSA_LOL",
// "candidates":["12tmRt6AADfQhfruF3RzFDdNhjiSEkwMvF","1HRK5H21wFguq5ecJF8FvQk28qYnPz1Qb9","1DngEcP2tCkxZNiAmm3Ar8VXXAAvAPfm8E"],
// "totalcandidates":3,"start":"1635032651","end":"1635043569"}]



// {{range .ballots}}
// 	{{if eq .ContractHash {{JS }} OR maybe "something" }}
// 	<span class="text-success">You have successfully logged in! : {{.ElectionName}}</span>
// 	{{end}}
// {{end}}


// <script>
// function showelectioninfo(){

// var d = document.getElementById("selectballot");
// var conhash = d.options[d.selectedIndex].text;
// alert(conhash);

// document.getElementById("balldata").value = conhash;

// </script>

// Test Example , jo console mai chla

// document.getElementById("selectballot");
// <select id=​"selectballot" class=​"form-select" aria-label=​"Default select example">​<option value disabled=​"disabled">​Open this select menu​</option>​<option value=​"ba56265934aa02b8ade28fd720119f627e8c51a08bec77f143cf4731b1b9697a">​BEST_COOK_2069​</option>​<option value=​"cf261d57c969a3d60c6ca8e2e31d0f812b160bf907f1b1d4052921984150ab1b">​BEST_SAMOSA_LOL​</option>​</select>​
// document.getElementById("selectballot").options[1].text;
// 'BEST_COOK_2069'
// document.getElementById("selectballot").options[document.getElementById("selectballot").selectedIndex].text;
// 'BEST_SAMOSA_LOL'
// document.getElementById("selectballot").options[document.getElementById("selectballot").selectedIndex].text;
// 'BEST_COOK_2069'