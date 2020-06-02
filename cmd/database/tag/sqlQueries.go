package tag

const (
	getQuery = "select name, counter, subscribers_counter from tag where name = $1;"

	getAllQuery = "select name, counter, subscribers_counter from tag;"

	getPostsQuery = `select p.id,
						   title,
						   description,
						   p.date_created,
						   date_updated,
						   creator_username,
						   bookmark_count
					from tag
							 join tag_post tp on tag.name = tp.tag_name
							 join post p on tp.post_id = p.id
					where name = $1;`

	insertQuery = `insert into tag(name) values ($1) returning name, counter;`
)
