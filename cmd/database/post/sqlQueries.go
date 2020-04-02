package post

const (
	AllColumns      = "id, title, description, date_created, date_updated, creator_username, like_count, bookmark_count"
	getQuery        = "select " + AllColumns + " from post where id = $1;"
	insertQuery     = "insert into post(description, creator_username, title)  values ($1, $2, $3) returning " + AllColumns + ";"
	deleteQuery     = "delete from post where id = $1 returning " + AllColumns + ";"
	getCreatedQuery = "select " + AllColumns +
		` from "user" u join post p on
		u.username = p.creator_username where username = $1`
	getTagNamesQuery = `select tag_name
						from post
         					join tag_post tp on post.id = tp.post_id
						where post_id = $1;`
	addTagQuery         = "insert into tag_post(tag_name, post_id) VALUES ($1, $2);"
	deleteTagQuery      = "delete from tag_post where post_id = $1 returning tag_name;"
	getNamesByPostQuery = `select tag_name from tag_post where post_id = $1;`
	updatePostQuery     = `update post set description = $1, title = $2, date_updated = now() where id = $3 
						returning id, title, description, date_created, date_updated, creator_username, like_count, bookmark_count;`
)
