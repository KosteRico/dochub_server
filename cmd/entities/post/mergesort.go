package post

import "sort"

func Sort(posts []*Post) {
	sort.SliceStable(posts, func(i, j int) bool {
		return posts[i].DateCreated.After(posts[j].DateCreated)
	})
}
