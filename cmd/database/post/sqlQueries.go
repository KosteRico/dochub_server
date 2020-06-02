package post

const (
	AllColumns      = "id, title, description, date_created, date_updated, creator_username, bookmark_count"
	getQuery        = "select " + AllColumns + " from post where id = $1;"
	insertQuery     = "insert into post(id, description, creator_username, title)  values ($1, $2, $3, $4) returning " + AllColumns + ";"
	deleteQuery     = "delete from post where id = $1 returning " + AllColumns + ";"
	getCreatedQuery = "select " + AllColumns +
		` from "user" u join post p on
		u.username = p.creator_username where username = $1`
	getTagNamesQuery = `select tag_name
						from post
         					join tag_post tp on post.id = tp.post_id
						where post_id = $1;`
	addTagQuery         = "insert into tag_post(tag_name, post_id) VALUES ($1, $2);"
	updateTagCounter    = "update tag set counter = counter + 1 where name = $1;"
	deleteTagQuery      = "delete from tag_post where post_id = $1 returning tag_name;"
	getNamesByPostQuery = `select tag_name from tag_post where post_id = $1;`
	updatePostQuery     = `update post set description = $1, title = $2, date_updated = now() where id = $3 
						returning id, title, description, date_created, date_updated, creator_username, bookmark_count;`
	selectBookmarkQuery = `select username from bookmark where username = $1 and post_id = $2;`
	insertFileQuery     = `insert into file(post_id, data) values ($1, $2);`
	selectFileQuery     = `select data from file where post_id = $1;`
	selectNameQuery     = `select title from post where id = $1;`
)
