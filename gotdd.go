// Package gotdd - a program to provide and list reviews of recent lunchtalks
package gotdd

import (
	"errors"
	"sync"
)

type Review struct {
	Id      int
	Comment string
}

type LunchTalk struct {
	Id      int
	Title   string
	Speaker string
	Reviews []Review
}

type Register struct {
	lts   []LunchTalk
	mutex sync.Mutex
}

func (r *Register) AddLunchTalk(lt LunchTalk) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	if len(lt.Speaker) == 0 || len(lt.Title) == 0 {
		return errors.New("Missing Data")
	}
	lt.Id = len(r.lts)
	r.lts = append(r.lts, lt)
	return nil
}

func (r *Register) GetLunchTalks() []LunchTalk {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	return r.lts
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
	rev.Id = len(r.lts[i].Reviews)
	r.lts[i].Reviews = append(r.lts[i].Reviews, rev)
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

	rev.Id = j
	r.lts[i].Reviews[j] = rev
	return nil
}
