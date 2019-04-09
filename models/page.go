package models

//定义分页信息
type Page struct {
	PrevPage	int //前一页页码
	NextPage	int //后一页页码
	Total		int //总记录数
	CurrentPage int //当前页码
	Limit 		int //每页显示条数
}

//SetPrevAndNextPage：设置分页前一页和后一页页码
func (p *Page) SetPrevAndNextPage()  {
	if p.CurrentPage > 1 {
		p.PrevPage = p.CurrentPage - 1
	}
	if (p.Total -1)/p.Limit >= p.CurrentPage{
		p.NextPage = p.CurrentPage + 1
	}
}

//SetPage：设置分页信息
func (p *Page)SetPage(totalCount,currentPage, limit int)  {
	p.Total = totalCount
	p.CurrentPage = currentPage
	p.Limit = limit
	p.SetPrevAndNextPage()
}
