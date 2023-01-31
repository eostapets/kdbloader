package main

import (
	kdb "github.com/sv/kdbgo"
	"log"
)

type tickerval struct {
	volum    int32
	exchange string
}
type ticker map[string]tickerval

func main() {
	m := ticker{
		"meta":   {4000, "nasdaq"},
		"nvidia": {9000, "dow"},
	}
	con, err := kdb.DialKDB("localhost", 5001, "")
	if err != nil {
		log.Panicf("Error connect: %s\n", err)
	}
	for k, v := range m {
		log.Printf("%s - %d %s\n", k, v.volum, v.exchange)
		tick := &kdb.K{kdb.KS, kdb.NONE, []string{k}}
		vol := &kdb.K{kdb.KI, kdb.NONE, []int32{v.volum}}
		ex := &kdb.K{kdb.KS, kdb.NONE, []string{v.exchange}}
		//tab := &kdb.K{kdb.XT, kdb.NONE,
		//kdb.Table{[]string{"ts", "source", "sym", "ask", "mid", "bid", "status"},
		//[]*kdb.K{ts, source, sym, ask, mid, bid, st}}}                    // insert tab sync
		//                    _, err = con.Call("insert", &kdb.K{-kdb.KS, kdb.NONE, "prices"}, tab)
		tab := &kdb.K{kdb.XT, kdb.NONE,
			kdb.Table{[]string{"sym", "mcap", "ex"},
				[]*kdb.K{tick, vol, ex}}}
		_, err = con.Call("insert", &kdb.K{-kdb.KS, kdb.NONE, "trade"}, tab)
		if err != nil {
			log.Panicf("Insert Query failed: %v", err)
			return
		}
	}
}
