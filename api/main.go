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
	"raffle-api/schema"
	"strconv"
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
	email := r.Header.Get("customer_email")
	var ra *schema.RaffleEntry

	for name, values := range r.Header {
		// Loop over all values for the name.
		for _, value := range values {
			fmt.Println(name, value)
		}
	}
	if err := regC.Find(bson.M{"email": email}).One(&ra); err != nil {
		fmt.Printf("An error occurred finding raffler %s\n", err.Error())
	} else {
		w.Header().Set("Content-Type", "application/json")

		if contains(genRandom(), int(ra.Raffle)) {
			res := &schema.Response{Msg: "WOW! You won"}

			if uerr := regC.Update(bson.M{"email": email}, bson.M{"$set": bson.M{
				"winner": true,
			}}); uerr != nil {
				fmt.Println("An Error occurred registering User as winner")
			} else {
				json, _ := json.Marshal(res)
				if _, werr := w.Write(json); werr != nil {
					fmt.Println(werr.Error())
				}
			}

		} else {
			res := &schema.Response{Msg: "Sorry, Your draw didn't win. Try again."}

			json, _ := json.Marshal(res)
			if _, werr := w.Write(json); werr != nil {
				fmt.Println(werr.Error())
			}

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
		var str string
		if len(ras) == 0 {
			str = "<h1>No Raffle Winners yet</h1>"
		} else {
			str = "<h1>Raffle Winners are:</h1>"

		}
		for i := 0; i < len(ras); i++ {
			str = string(i) + str + fmt.Sprintf("<h3>%s</h3>", ras[i])
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
