package model

import (
	"ginTemplate/config"
	"ginTemplate/model/common"
)

func init() {
	DefaultBlogLikeConnection = BlogLikeConnection(common.GetDB2DBConnect(config.GetDBConnect(config.DbName)))
}

type LikeList struct {
	UID int `json:"uId"`
	Num int `json:"num"`
}

// FindLabel 查询Label
func GetBlogLikeList(n int) ([]LikeList, error) {
	return DefaultBlogLikeConnection.GetBlogLikeList(n)
}

// Find 查询Label
func (c BlogLikeConnection) GetBlogLikeList(n int) ([]LikeList, error) {
	sqlStr := "select b.uid,count(*) as num from tbBlog as b left join tbBlogLike as l on b.id = l.blogId where b.status = 1 and l.status = 1 group by b.uid order by num desc limit ?"

	results := make([]LikeList, 0)
	stmt, err := c().Prepare(sqlStr)
	if err != nil {
		return results, err
	}
	defer stmt.Close()
	rows, err := stmt.Query(n)
	if err != nil {
		return results, err
	}
	defer rows.Close()
	for rows.Next() {
		model := LikeList{}
		values := []interface{}{
			new(interface{}),
			new(interface{}),
		}
		rows.Scan(values...)
		if *(values[0].(*interface{})) != nil {
			tmp := int((*(values[0].(*interface{}))).(int64))
			model.UID = tmp
		}
		if *(values[1].(*interface{})) != nil {
			tmp := int((*(values[0].(*interface{}))).(int64))
			model.Num = tmp
		}
		results = append(results, model)
	}
	return results, nil
}