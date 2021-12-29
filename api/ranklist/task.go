package ranklist

import (
	"ginTemplate/api"
	blog "ginTemplate/model/blog"
	user "ginTemplate/model/user"
	"sort"
	"strconv"
)


type RequestInfo struct{
	ID int `json:"id"`
	Address string `json:"address"  binding:"required"`
	NickName string `json:"nickName"`
	Email string `json:"email"`
	ImageUrl string `json:"imageUrl"`
	Introduction string `json:"introduction"`
	Num int `json:"num"`
}


func DoLikeList(c *api.Context) (int, string, interface{}) {
	num, err := strconv.Atoi(c.Query("num"))
	if err != nil {
		return api.IllegalArgument, c.Error(err).Error(), ""
	}
	// 检测这个点赞是否存在
	list, err := blog.GetBlogLikeList(num)
	if err != nil {
		return api.DatabaseError, c.Error(err).Error(), ""
	}
	// 排序
	res := make([]RequestInfo, 0)
	// 关联用户
	for _, l := range list {
		u, err := user.GetUser("1 = 1 AND status = 1 AND id = ?", l.UID);
		if err != nil {
			return api.DatabaseError, c.Error(err).Error(), ""
		}
		req := RequestInfo {
			ID : u.ID,
			Address : u.Address,
			NickName : u.NickName,
			Email : u.Email,
			ImageUrl : u.ImageURL,
			Introduction: u.Introduction,
			Num: l.Num,
		}
		res = append(res, req)
	}
	sort.Sort(List(res))

	return 0, "success", res
}


type List []RequestInfo
// Swap
func (p List) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
// Len
func (p List) Len() int           { return len(p) }
// Less
func (p List) Less(i, j int) bool { return p[i].Num > p[j].Num }
