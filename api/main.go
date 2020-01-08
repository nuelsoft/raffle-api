package main

import (
	"encoding/json"
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"math/rand"
	"net/http"
	"os"
	"raffle-api/fxn"
	"raffle-api/schema"
	"strconv"
	"strings"
	"time"
)

const defaultPort = "9090"

var regC *mgo.Collection

func mkRegC(session *mgo.Session) (collection *mgo.Collection) {
	return session.DB("heroku_drh8mw8f").C("raffles")
}

func genRandom() (rands []int) {
	for i := 0; i < 100; i++ {
		r := rand.Intn(999)
		rands = append(rands, r)
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

	s, err := mgo.Dial("mongodb://heroku_drh8mw8f:8v8d10d1jhlo7crb7404psbtfg@ds161295.mlab.com:61295/heroku_drh8mw8f")
	//s, err := mgo.Dial("mongodb://127.0.0.1")
	if err != nil {
		log.Panicf("Error Occurred while Dialing Database: %s", err)
	} else {
		regC = mkRegC(s)
	}
}

func register(w http.ResponseWriter, r *http.Request) {
	email := r.Header.Get("Email")
	raffle := r.Header.Get("Raffle")

	iRaffle, _ := strconv.ParseInt(raffle, 10, 0)
	var ra *schema.RaffleEntry

	ra = &schema.RaffleEntry{
		Email:  email,
		Raffle: iRaffle,
		Date:   bson.Now().String(),
	}

	if err := regC.Insert(&ra); err != nil {
		log.Fatal("An error occurred Inserting to new Raffle DB\n")
	} else {
		fmt.Println("Placed Raffle Data. Awaiting Payment...")
		w.Header().Set("Content-Type", "application/json")
		res := &schema.Response{
			Msg: "success: raffle data saved",
		}

		json, _ := json.Marshal(res)
		if _, werr := w.Write(json); werr != nil {
			fmt.Println(werr.Error())
		}
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
				w.Header().Set("Content-Type", "application/json")

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
						con := fmt.Sprintf("<h1>Congratulations, <br> %s. You won!, you will recieve an email soon. Also check <a href =\"https://api-raffleit.herokuapp.com/winners\">here</a> to verify.</h1>", ra.Name)
						if _, werr := w.Write([]byte(con)); werr != nil {
							fmt.Println(werr.Error())
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
						con := fmt.Sprintf("<h1>Sorry <br> %s. That didn't go as planned! Please try again</h1>", ra.Name)
						if _, werr := w.Write([]byte(con)); werr != nil {
							fmt.Println(werr.Error())
						}
					}
				}
			} else {
				if ra != nil {
					con := fmt.Sprintf("<h1>You already won!, <br> %sPlease check <a href =\"https://api-raffleit.herokuapp.com/winners\">here</a> to verify.</h1>", ra.Email)
					if _, werr := w.Write([]byte(con)); werr != nil {
						fmt.Println(werr.Error())
					}
				}
			}
		} else {
			if ra != nil {
				if _, werr := w.Write([]byte("Payment ref has already been used")); werr != nil {
					fmt.Println(werr.Error())
				}
			}

		}
	} else {
		if _, werr := w.Write([]byte("Payment ref seems to be invalid")); werr != nil {
			fmt.Println(werr.Error())
		}
	}

}

func winners(w http.ResponseWriter, r *http.Request) {
	var ras []*schema.RaffleEntry
	if err := regC.Find(bson.M{"winner": true}).All(&ras); err != nil && err.Error() != "not found" {
		fmt.Println(err.Error())

	} else {
		//res := &schema.Response{Msg: "Check winner List", Winners: ras}
		//json, _ := json.Marshal(res)

		str := ""
		if len(ras) == 0 {
			str = "<h2>No Raffle Winners yet</h2>"
		} else {
			str = "<h2>Raffle Winners are:</h2>"

		}
		for i := 0; i < len(ras); i++ {
			str += fmt.Sprintf("<h3>%d. %s ::: %s ::: %s</h3>\n", i+1, ras[i].Name, ras[i].Email, ras[i].Phone)
		}

		if _, werr := w.Write([]byte(str)); werr != nil {
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

	http.HandleFunc("/api/register", register)
	http.HandleFunc("/api/draw", draw)
	http.HandleFunc("/winners", winners)

	log.Fatal(srv.ListenAndServe())
}
