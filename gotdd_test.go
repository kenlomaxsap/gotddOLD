package gotdd

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

// The Example function collects Green, up-to-date, use-cases/examples, as specified with the PO

func Example() {

	// I would like to specify some lunch talks and then..
	lt1 := LunchTalk{Title: "TDD and Go", Speaker: "Ken Lomax"}
	lt2 := LunchTalk{Title: "Kyma Drones", Speaker: "JE"}

	fmt.Printf("lt1: %+v\n", lt1)
	fmt.Printf("lt2: %+v\n", lt2)

	// .. store them
	r := Register{}
	r.AddLunchTalk(lt1)
	r.AddLunchTalk(lt2)

	// .. list them
	lts3 := r.GetLunchTalks()
	fmt.Printf("lts3: %+v\n", lts3)

	// .. add reviews to them
	r.AddReview(0, Review{Comment: "Big pile of poo"})
	lts4 := r.GetLunchTalks()
	fmt.Printf("lts4: %+v\n", lts4)

	// .. adjust those reviews
	r.AdjustReview(0, 0, Review{Comment: "Just amazing!!!"})
	lts5 := r.GetLunchTalks()
	fmt.Printf("lts5: %+v\n", lts5)

	// More requirements to come after that is implemented..
	// Can we have eror handling - to look after invalid input gracefully..
	e6 := r.AddLunchTalk(LunchTalk{Title: "", Speaker: ""})
	fmt.Printf("e6: %+v\n", e6)

	e7 := r.AddReview(999, Review{Comment: "Big pile of poo"})
	fmt.Printf("e7: %+v\n", e7)

	e8 := r.AdjustReview(999, 999, Review{Comment: "Big pile of poo"})
	fmt.Printf("e8: %+v\n", e8)

	e9 := r.AdjustReview(999, 999, Review{Comment: ""})
	fmt.Printf("e9: %+v\n", e9)

	e10 := r.AddReview(999, Review{Comment: ""})
	fmt.Printf("e10: %+v\n", e10)

	// Some measure of performance please..
	// go test -bench=.
	// gobenchui -last 10 github.com/kenlomaxhybris/gotdd

	// It needs to be threads safe - sp multiple people can call this at once
	// go test -v -race

	// More to come..

	//Output:
	// lt1: {Id:0 Title:TDD and Go Speaker:Ken Lomax Reviews:[]}
	// lt2: {Id:0 Title:Kyma Drones Speaker:JE Reviews:[]}
	// lts3: [{Id:0 Title:TDD and Go Speaker:Ken Lomax Reviews:[]} {Id:1 Title:Kyma Drones Speaker:JE Reviews:[]}]
	// lts4: [{Id:0 Title:TDD and Go Speaker:Ken Lomax Reviews:[{Id:0 Comment:Big pile of poo}]} {Id:1 Title:Kyma Drones Speaker:JE Reviews:[]}]
	// lts5: [{Id:0 Title:TDD and Go Speaker:Ken Lomax Reviews:[{Id:0 Comment:Just amazing!!!}]} {Id:1 Title:Kyma Drones Speaker:JE Reviews:[]}]
	// e6: Missing Data
	// e7: Out of bounds
	// e8: Out of bounds
	// e9: Missing Data
	// e10: Missing Data

}

func TestMultipleCalls(t *testing.T) {
	n := 100
	r := Register{}
	for i := 0; i < n; i++ {
		go r.AddLunchTalk(LunchTalk{Title: "Some", Speaker: "Thing"})
	}
	for i := 0; i < n; i++ {
		go r.AddReview(0, Review{Comment: "Just amazing!!!"})
	}
	for i := 0; i < n; i++ {
		go r.AdjustReview(0, 0, Review{Comment: "Just amazing!!!"})
	}
	time.Sleep(2 * time.Second)

}

func Benchmark(b *testing.B) {
	r := Register{}
	for i := 0; i < b.N; i++ {
		r.AddLunchTalk(LunchTalk{Title: "Some", Speaker: "Thing"})
		r.AddReview(rand.Intn(100), Review{Comment: "Just amazing!!!"})
		r.AdjustReview(rand.Intn(100), rand.Intn(100), Review{Comment: "Just amazing!!!"})
	}
}
