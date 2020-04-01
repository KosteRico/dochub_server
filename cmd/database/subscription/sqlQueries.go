package subscription

const (
	getPostsQuery = `select p.id,
						   title,
						   description,
						   p.date_created,
						   date_updated,
						   creator_username,
						   like_count,
						   bookmark_count
					from subscription
							 join tag t on subscription.tag_name = t.name
							 join tag_post tp on t.name = tp.tag_name
							 join post p on tp.post_id = p.id
					where username = $1;`

	getTagNamesQuery = `select name
						from subscription
								 join tag t on subscription.tag_name = t.name
						where username = $1;`
	insertQuery = "insert into subscription(username, tag_name) values ($1, $2);"
	deleteQuery = "delete from subscription where username = $1 and tag_name = $2;"
)
