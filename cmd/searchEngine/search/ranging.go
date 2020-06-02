package search

import (
	"checkaem_server/cmd/searchEngine/entities"
	"checkaem_server/cmd/searchEngine/util"
	"log"
	"sort"
	"sync"
)

type idWeight struct {
	Id     string
	Weight float64
}

type idsSet struct {
	arr []*idWeight
}

func (set *idsSet) Insert(item *idWeight) {
	set.arr = append(set.arr, item)
	sort.SliceStable(set.arr, func(i, j int) bool {
		return set.arr[i].Weight < set.arr[j].Weight
	})
}

func (set *idsSet) IDs() (res []string) {
	for _, val := range set.arr {
		res = append(res, val.Id)
	}
	return
}

func newSet() *idsSet {
	return &idsSet{
		arr: []*idWeight{},
	}
}

func SearchRanking(query string) ([]string, error) {
	iDs, err := search(query)

	if err != nil {
		return nil, err
	}

	terms, err := util.PrepareQuery(query)

	if err != nil {
		return nil, err
	}

	bm25 := entities.NewBm25Builder()

	resChan := make(chan *idWeight)

	wg := &sync.WaitGroup{}

	for _, item := range iDs {
		wg.Add(1)
		go func(wg *sync.WaitGroup, c chan<- *idWeight, id string) {
			defer wg.Done()

			err = bm25.SetDoc(id)

			if err != nil {
				log.Println(err)
				return
			}

			calc, err := bm25.Calc(terms)

			if err != nil {
				log.Println(err)
				return
			}

			c <- &idWeight{
				Id:     id,
				Weight: calc,
			}

		}(wg, resChan, item)
	}

	go func(wg *sync.WaitGroup, c chan *idWeight) {
		wg.Wait()
		close(c)
	}(wg, resChan)

	resSet := newSet()

	for v := range resChan {
		resSet.Insert(v)
	}

	return resSet.IDs(), nil

}
