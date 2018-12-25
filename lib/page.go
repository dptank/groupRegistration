package lib
type Page struct {
	pageNum     int
	pageSize   int
	totalPage  int
	totalCount int
	nextPage  bool
	list       interface{}
}
func PageUtil(count int, pageNum int, pageSize int, list interface{},nextPage bool) Page {
	tp := count / pageSize
	if count % pageSize > 0 {
		tp = count / pageSize + 1
	}
	return Page{pageNum: pageNum, pageSize: pageSize, totalPage: tp, totalCount: count, nextPage:nextPage , list: list}
}