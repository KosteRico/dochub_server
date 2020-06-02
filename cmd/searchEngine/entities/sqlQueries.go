package entities

const (
	selectIndexInfoQuery    = "select data from index_info where id = $1 and term = $2;"
	selectDocumentMetaQuery = "select data from doc_meta where id = $1;"
	insertIndexInfoQuery    = "insert into index_info(id, term, data) VALUES ($1, $2, $3);"
	selectTermCountQuery    = "select data->'c' from index_info where id = $1 and term = $2;"
)
