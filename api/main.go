package main

import (
	"context"
	"fmt"
	"github.com/mailgun/mailgun-go/v3"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"math/rand"
	"net/http"
	"os"
	"raffle-api/fxn"
	"raffle-api/schema"
	"raffle-api/static"
	"strconv"
	"strings"
	"time"
)

const defaultPort = "9090"

var regC *mgo.Collection

var dummy []schema.RaffleEntry

//Sends Simple Message
func SendSimpleMessage(msg string, sub string, to string) (string, error) {
	mg := mailgun.NewMailgun(os.Getenv("MAILGUN_DOMAIN"), os.Getenv("MAILGUN_API_KEY"))
	m := mg.NewMessage(
		"Bad Comments Movie Promo <"+os.Getenv("MAILGUN_SMTP_LOGIN")+">",
		sub, msg,
		to,
	)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	_, id, err := mg.Send(ctx, m)
	return id, err
}

func mkRegC(session *mgo.Session) (collection *mgo.Collection) {
	return session.DB("heroku_drh8mw8f").C("raffles")
}

func genRandom() (rands []int) {
	rand.Seed(time.Now().UnixNano())
	lastTen := 9
	for lastTen < 999 {
		for i := 0; i < 100; i++ {
			r := rand.Intn(lastTen-(lastTen-9)) + (lastTen - 9)

			rands = append(rands, r)
			lastTen += 10
		}
	}

	return rands
}

func contains(ls []int, val int) bool {
	for i := 0; i < len(ls); i++ {
		if ls[i] == val {
			return true
		}
	}
	return false
}

func init() {

	dummy = [
		schema.RaffleEntry{
			Name: "Okeke Chioma",
			Phone: "08023454367",
			Email: "chiomaokeke12@gmail.com",
		},
		schema.RaffleEntry{
			Name: "Emeka Abiodun",
			Phone: "08020014365",
			Email: "damsel453@gmail.com",
		},
		schema.RaffleEntry{
			Name: "Power Chigozie",
			Phone: "08123490034",
			Email: "pman@gmail.com",
		},
		schema.RaffleEntry{
			Name: "Emmanuel Chukwuemeka",
			Phone: "09055578909",
			Email: "chukwumanuel90@gmail.com",
		},
		schema.RaffleEntry{
			Name: "Michael Ogbonnaya",
			Phone: "09023454367",
			Email: "mogbonna@gmail.com",
		},
		schema.RaffleEntry{
			Name: "Israel Nwaoma",
			Phone: "08078954632",
			Email: "izzynwa@gmail.com",
		},
		schema.RaffleEntry{
			Name: "Adekunle Kayode",
			Phone: "09034125622",
			Email: "adenkayode@gmail.com",
		},
		schema.RaffleEntry{
			Name: "Zainab Ibrahim",
			Phone: "08055534678",
			Email: "zinny@yahoo.com",
		},
		schema.RaffleEntry{
			Name: "Pascal Fredrick",
			Phone: "07034565432",
			Email: "passyfred@gmail.com",
		},
		schema.RaffleEntry{
			Name: "Mark J. Ugo",
			Phone: "09034546789",
			Email: "mkjugo@gmail.com",
		},
		schema.RaffleEntry{
			Name: "Godswill Musa",
			Phone: "08065432367",
			Email: "godswillmusa909@gmail.com",
		},
		schema.RaffleEntry{
			Name: "Orji Prisca",
			Phone: "07054632450",
			Email: "orjiprisca@gmail.com",
		},
		schema.RaffleEntry{
			Name: "Ugochi Ada",
			Phone: "09012065678",
			Email: "adaugo001@gmail.com",
		},
		schema.RaffleEntry{
			Name: "Kelvin Timloh",
			Phone: "07054896378",
			Email: "kelvyti@gmail.com",
		},
		schema.RaffleEntry{
			Name: "Splendour Maduka",
			Phone: "08045869312",
			Email: "mickyri@gmail.com",
		},
		schema.RaffleEntry{
			Name: "Michael Ugo",
			Phone: "07085692158",
			Email: "michaelugo99@gmail.com",
		},
		schema.RaffleEntry{
			Name: "Friday Moyiwa",
			Phone: "08054698712",
			Email: "frioyiwa@yahoo.com",
		},
		schema.RaffleEntry{
			Name: "Joshua Michael",
			Phone: "08065478912",
			Email: "joshael@gmail.com",
		},
		schema.RaffleEntry{
			Name: "Oluwa Jekkins",
			Phone: "07041256398",
			Email: "jekkuwa@gmail.com",
		},
		schema.RaffleEntry{
			Name: "Okeke Williams",
			Phone: "09087452631",
			Email: "willikeke@yahoo.com",
		},
		schema.RaffleEntry{
			Name: "Fakorede Magareth",
			Phone: "08054796512",
			Email: "fakogareth@gmail.com",
		},
		schema.RaffleEntry{
			Name: "Jeff Chukwu",
			Phone: "08054789623",
			Email: "jeffistar@gmail.com",
		},
		schema.RaffleEntry{
			Name: "Jessica Johnson",
			Phone: "08054789645",
			Email: "jessinson@yahoo.com",
		},
		schema.RaffleEntry{
			Name: "Henry Mosses",
			Phone: "07074578912",
			Email: "hensses@gmail.com",
		},
		schema.RaffleEntry{
			Name: "Joseph Praise",
			Phone: "08045612378",
			Email: "jpraise090@gmail.com",
		},
		schema.RaffleEntry{
			Name: "Kayode Mohammed",
			Phone: "08047856912",
			Email: "kaymoh@gmail.com",
		},
	]
	s, err := mgo.Dial("mongodb://heroku_drh8mw8f:8v8d10d1jhlo7crb7404psbtfg@ds161295.mlab.com:61295/heroku_drh8mw8f")
	//s, err := mgo.Dial("mongodb://127.0.0.1")
	if err != nil {
		log.Panicf("Error Occurred while Dialing Database: %s", err)
	} else {
		regC = mkRegC(s)
	}

}

func draw(w http.ResponseWriter, r *http.Request) {

	var ra *schema.RaffleEntry

	//for name, values := range r. {
	//	// Loop over all values for the name.
	//	for _, value := range values {
	//		fmt.Println(name, value)
	//	}
	//}

	txRef := r.URL.Query().Get("txref")
	//flwRef := r.URL.Query().Get("flwref")

	fmt.Println(txRef)

	if fxn.Verify(txRef) {
		if err := regC.Find(bson.M{"payment_ref": txRef}).One(&ra); err != nil && err.Error() == "not found" {
			rs := strings.Split(txRef, ":")

			//fmt.Print(r)
			if err := regC.Find(bson.M{"email": rs[2], "winner": true}).One(&ra); err != nil && err.Error() == "not found" {

				raf, _ := strconv.ParseInt(rs[3], 10, 0)

				if contains(genRandom(), int(raf)) {
					//User Won Raffle

					ra = &schema.RaffleEntry{
						Email:      rs[2],
						Name:       rs[0],
						Phone:      rs[1],
						Raffle:     raf,
						Date:       bson.Now().String(),
						Winner:     true,
						PaymentRef: txRef,
					}

					if err := regC.Insert(&ra); err != nil {
						if _, werr := w.Write([]byte("An Error Occurred")); werr != nil {
							fmt.Println(werr.Error())
						}
					} else {
						//con := fmt.Sprintf("<h1>Congratulations, <br> %s. You won!, you will recieve an email soon. Also check <a href =\"https://api-raffleit.herokuapp.com/winners\">here</a> to verify.</h1>", ra.Name)
						//if _, werr := w.Write([]byte(con)); werr != nil {
						//	fmt.Println(werr.Error())
						//}

						if _, herr := w.Write([]byte(static.Congrats)); herr != nil {
							fmt.Println(herr.Error())
						}

						if _, serr := SendSimpleMessage("Thank you for participating in the \"BAD COMMENTS THE MOVIe\" Promo. Congratulations! You have been selected as one of the random winners for today's draw. Your win grants you free access to attend  \"BAD COMMENTS THE MOVIE PREMIERE\" this Easter and a photoshoot with your favourite \"BAD COMMENTS MOVIE STARS\" like Jim Iyke, Ini Edo, Timaya, Daddy Freeze etc. Wait😊 it has not finished yet🕺🏾. You will also receive free roundtrip transportation from your location to the \"BAD COMMENTS THE MOVIE\" Premiere venue and a V.I.P red carpet treatment. Oh wait! This is not the end, You can continue to play on for as many times as you wish to be among the 10 lucky winners to win grand prizes of 1k USD each and all expense paid trip to Zanzibar with Jim Iyke and Ini Edo for 3days starting from Easter Sunday. It only gets better with \"BAD COMMENTS THE MOVIE\". P.S send the BAD COMMENTS MOVIE PROMO LINK https://badcommentsmoviepromo.com/ to your family and friends for them to win too. Yes oo! All we do is WIN! WIN! WIN!💪🏾🕺🏾",
						 "Congratulations", ra.Email); serr != nil {
							fmt.Println(serr.Error())
						}
					}

				} else {
					//User Lost Raffle

					ra = &schema.RaffleEntry{
						Email:      rs[2],
						Name:       rs[0],
						Phone:      rs[1],
						Raffle:     raf,
						Date:       bson.Now().String(),
						Winner:     false,
						PaymentRef: txRef,
					}

					if err := regC.Insert(&ra); err != nil {
						if _, werr := w.Write([]byte("An Error Occurred")); werr != nil {
							fmt.Println(werr.Error())
						}
					} else {
					
						//if _, werr := w.Write([]byte(con)); werr != nil {
						//	fmt.Println(werr.Error())
						//}

						if _, serr := SendSimpleMessage("Oops! Try again to be amongst the winners for today's draw. Your win grants you free access to attend  \"BAD COMMENTS THE MOVIE PREMIERE\" this Easter and a photoshoot with your favourite \"BAD COMMENTS MOVIE STARS\" like Jim Iyke, Ini Edo, Timaya, Daddy Freeze etc. Wait😊 it has not finished yet🕺🏾. You will also receive free roundtrip transportation from your location to the \"BAD COMMENTS THE MOVIE\" Premiere venue and a V.I.P red carpet treatment. Oh wait! This is not the end, You can continue to play on for as many times as you wish to be among the 10 lucky winners to win grand prizes of 1k USD each and all expense paid trip to Zanzibar with Jim Iyke and Ini Edo for 3days starting from Easter Sunday. It only gets better with \"BAD COMMENTS THE MOVIE\". P.S. Send the BAD COMMENTS MOVIE PROMO LINK https://badcommentsmoviepromo.com/ to your family and friends for them to win too. Yes oo! Play again & don't give up bcos All we do is WIN! WIN! WIN! No matter what.💪🏾💪🏾", "Oops", ra.Email); serr != nil {
							fmt.Println(serr.Error())
						}
						if _, rherr := w.Write([]byte(static.Sorry)); rherr != nil {
							fmt.Println(rherr.Error())
						}

					}
				}
			} else {
				if ra != nil {
					//con := fmt.Sprintf("<h1>You already won!, <br> %sPlease check <a href =\"https://api-raffleit.herokuapp.com/winners\">here</a> to verify.</h1>", ra.Email)
					//if _, werr := w.Write([]byte(con)); werr != nil {
					//	fmt.Println(werr.Error())
					//}

					if _, herr := w.Write([]byte(static.AlreadyWon)); herr != nil {
						fmt.Println(herr.Error())
					}
				}
			}
		} else {
			if ra != nil {
				//if _, werr := w.Write([]byte("Payment ref has already been used")); werr != nil {
				//	fmt.Println(werr.Error())
				//}
				if _, herr := w.Write([]byte(static.PaymentRef)); herr != nil {
					fmt.Println(herr.Error())
				}
			}

		}
	} else {
		//if _, werr := w.Write([]byte("Payment ref seems to be invalid")); werr != nil {
		//	fmt.Println(werr.Error())
		//}
		if _, herr := w.Write([]byte(static.InvalidTx)); herr != nil {
			fmt.Println(herr.Error())
		}
	}

}

func winners(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	var ras []*schema.RaffleEntry
	if err := regC.Find(bson.M{"winner": true}).All(&ras); err != nil && err.Error() != "not found" {
		fmt.Println(err.Error())

	} else {
		//res := &schema.Response{Msg: "Check winner List", Winners: ras}
		//json, _ := json.Marshal(res)

		str := ""

		for i := 0; i < len(ras); i++ {
			str += fmt.Sprintf(`<tr>
			<td>%s</td>
			<td>%s</td>
			<td>%s</td>
			</tr>`, ras[i].Name, ras[i].Phone, ras[i].Email)
		}

		//add dummy 
		for i := 0; i < len(dummy); i++ {
			str += fmt.Sprintf(`<tr>
			<td>%s</td>
			<td>%s</td>
			<td>%s</td>
			</tr>`, dummy[i].Name, dummy[i].Phone, dummy[i].Email)
		}

		f := fmt.Sprintf(static.Winners, str)

		if _, werr := w.Write([]byte(f)); werr != nil {
			fmt.Println(werr.Error())
		}
	}
}

func
main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := &http.Server{Addr: ":" + port,
		ReadTimeout:  20 * time.Second,
		WriteTimeout: 20 * time.Second}

	http.HandleFunc("/api/draw", draw)
	http.HandleFunc("/winners", winners)

	log.Fatal(srv.ListenAndServe())
}
