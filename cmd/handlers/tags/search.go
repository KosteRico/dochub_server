package tags

import (
	tag2 "checkaem_server/cmd/database/tag"
	"checkaem_server/cmd/entities/tag"
	"github.com/lithammer/fuzzysearch/fuzzy"
)

type tagsWithDistance struct {
	tag      *tag.Tag
	distance int
}

func insertionSort(arr []tagsWithDistance) []tagsWithDistance {
	n := len(arr)

	for i := 1; i <= n-1; i++ {

		j := i
		for j > 0 {

			if arr[j].distance > arr[j-1].distance {
				arr[j], arr[j-1] = arr[j-1], arr[j]
			}
			j -= 1
		}
	}

	return arr
}

func SearchTagsFunc(query string, limit int) (res []*tag.Tag, err error) {
	tags, err := tag2.GetAll()

	if err != nil {
		return
	}

	var withDistances []tagsWithDistance

	for _, t := range tags {
		dist := fuzzy.RankMatchNormalized(query, t.Name)

		withDistances = append(withDistances, tagsWithDistance{
			tag:      t,
			distance: dist,
		})

		withDistances = insertionSort(withDistances)
	}

	for i := 0; i < len(withDistances); i++ {
		if withDistances[i].distance >= 0 {
			res = append(res, withDistances[i].tag)
		}

		if i == limit-1 {
			break
		}
	}

	return
}
