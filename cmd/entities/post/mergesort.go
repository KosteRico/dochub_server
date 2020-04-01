package post

func merge(a []*Post, b []*Post, f func(a, b *Post) bool) []*Post {

	var r = make([]*Post, len(a)+len(b))
	var i = 0
	var j = 0

	for i < len(a) && j < len(b) {

		//<=
		if f(a[i], b[j]) {
			r[i+j] = a[i]
			i++
		} else {
			r[i+j] = b[j]
			j++
		}

	}

	for i < len(a) {
		r[i+j] = a[i]
		i++
	}
	for j < len(b) {
		r[i+j] = b[j]
		j++
	}

	return r
}

func Mergesort(items []*Post, f func(a, b *Post) bool) []*Post {

	if len(items) < 2 {
		return items

	}

	var middle = len(items) / 2
	var a = Mergesort(items[:middle], f)
	var b = Mergesort(items[middle:], f)
	return merge(a, b, f)
}
