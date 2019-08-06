// Package gotdd - a program to provide and list reviews of recent lunchtalks
package gotdd

import (
	"errors"
	"os"
	"strconv"
	"sync"

	"github.com/sonyarouje/simdb/db"
)

type Review struct {
	Id      string
	Comment string
}

type LunchTalk struct {
	Id      string
	Title   string
	Speaker string
	Reviews []Review
}

type Register struct {
	lts   []LunchTalk
	mutex sync.Mutex
	db    *db.Driver
}

func (lt LunchTalk) ID() (jsonField string, value interface{}) {
	value = lt.Id
	jsonField = "Id"
	return
}

func (r *Register) AddLunchTalk(lt LunchTalk) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	if len(lt.Speaker) == 0 || len(lt.Title) == 0 {
		return errors.New("Missing Data")
	}
	lt.Id = strconv.Itoa(len(r.lts))
	r.lts = append(r.lts, lt)
	r.db.Insert(lt)
	return nil
}

func (r *Register) GetLunchTalks() []LunchTalk {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	var lunchtalks []LunchTalk
	r.db.Open(LunchTalk{}).Get().AsEntity(&lunchtalks)
	return lunchtalks
}

func (r *Register) AddReview(i int, rev Review) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	if len(rev.Comment) == 0 {
		return errors.New("Missing Data")
	}
	if i > len(r.lts)-1 {
		return errors.New("Out of bounds")
	}
	rev.Id = strconv.Itoa(len(r.lts[i].Reviews))
	r.lts[i].Reviews = append(r.lts[i].Reviews, rev)
	r.db.Update(r.lts[i])
	return nil
}

func (r *Register) AdjustReview(i int, j int, rev Review) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	if len(rev.Comment) == 0 {
		return errors.New("Missing Data")
	}

	if i > len(r.lts)-1 || j > len(r.lts[i].Reviews)-1 {
		return errors.New("Out of bounds")
	}

	rev.Id = strconv.Itoa(j)
	r.lts[i].Reviews[j] = rev
	r.db.Update(r.lts[i])
	return nil
}

func (r *Register) Clear() {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	r.lts = []LunchTalk{}
	os.Remove("./data/LunchTalk")
}
