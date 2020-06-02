package bookmark

const (
	selectPostsQuery = `select id, title, description, p.date_created, date_updated, creator_username, bookmark_count` +
		` from bookmark b
        	 join post p on b.post_id = p.id
		where b.username = $1 order by b.date_created desc;`
	insertQuery                 = `insert into bookmark(username, post_id) VALUES ($1, $2);`
	deleteQuery                 = `delete from bookmark where username = $1 and post_id = $2 returning username;`
	incrementBookmarkCountQuery = `update post set bookmark_count = bookmark_count + 1 where id = $1;`
	decrementBookmarkCountQuery = `update post set bookmark_count = bookmark_count - 1 where id = $1;`
)
