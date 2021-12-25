package blog

import (
	"ginTemplate/api"
	blog "ginTemplate/model/blog"
	user "ginTemplate/model/user"
	"ginTemplate/utils"
	"strconv"
	"time"

	"github.com/wonderivan/logger"
)

type RequestInfo struct{
	ID int `json:"id"`
	Address string `json:"address"  binding:"required"`
	Title string `json:"title"`
	BlogText string `json:"blogText" binding:"required"`
	UId int `json:"uId" binding:"required"`
	ImageUrl string `json:"imageUrl"`
}


func DoGetBlog(c *api.Context) (int, string, interface{}) {
	//获取分页参数:PageIndex,PageSize
	page, pageSize, code, err := c.GetPager()
	if err != nil {
		return code, c.Error(err).Error(), nil
	}
	logger.Debug("page:", page)
	logger.Debug("pageSize:", pageSize)
	logger.Debug("code:", code)

	parConstruct := map[string]*api.ParamConstruct{
		"id":      {FieldName: "id", DefaultValue: "", CheckValue: nil, Need: false, Link: "and", Symbol: "="},
		"address": {FieldName: "address", DefaultValue: "", CheckValue: nil, Need: false, Link: "and", Symbol: "="},
		"title": {FieldName: "title", DefaultValue: "", CheckValue: nil, Need: false, Link: "and", Symbol: "like"},
		"blogText":{FieldName: "blogText", DefaultValue: "", CheckValue: nil, Need: false, Link: "and", Symbol: "like"},
		"uId":   {FieldName: "uId", DefaultValue: "", CheckValue: nil, Need: false, Link: "and", Symbol: "="},
		"status":  {FieldName: "status", DefaultValue: "1", CheckValue: nil, Need: false, Link: "and", Symbol: "="},
		"creator": {FieldName: "creator", DefaultValue: "", CheckValue: nil, Need: false, Link: "and", Symbol: "="},
		"orderBy": {FieldName: "", DefaultValue: "id|desc", CheckValue: nil, Need: false},
	}
	c.Set("isNotAllCondition", false)
	strCondition, args, err := c.GetConditionByParam(parConstruct)
	if err != nil {
		return api.IllegalArgument, c.Error(err).Error(), nil
	}
	logger.Debug("strCondition:", strCondition)
	logger.Debug("args:", args)
	
	blogs, err := blog.FindBlog(strCondition, args...)
	if err != nil {
		return api.DatabaseError, c.Error(err).Error(), nil
	}


	// 封装Blogs
	res := make([]interface{}, 0)
	for _, v := range blogs {
		newMap := make(map[string]interface{})
		newMap["id"] = v.ID
		newMap["Address"] = v.Address
		newMap["title"] = v.Title
		newMap["blogText"] = v.BlogText
		newMap["uId"] = v.UID
		newMap["imageUrl"] = v.ImageURL
		newMap["status"] = v.Status
		newMap["creator"] = v.Creator
		newMap["createTime"] = v.CreateTime
		newMap["updater"] = v.Updater
		newMap["updateTime"] = v.UpdateTime

		// 获取点赞数
		blogLikeList, err := blog.FindBlogLike("1 = 1 AND status = 1 AND blogId = ?", v.ID)
		if err != nil {
			return api.DatabaseError, c.Error(err).Error(), nil
		}
		newMap["likeNum"] = len(blogLikeList)
		res = append(res, &newMap)
	}

	return 0, "success", res
}

func DoAddBlog(c *api.Context) (int, string, interface{}) {
	reqData := RequestInfo{}
	if err := c.ShouldBind(&reqData); err != nil {
		return api.IllegalArgument, c.Error(err).Error(), nil
	}
	logger.Debug("reqData:",reqData);

	if reqData.Address == "" {
		return api.IllegalArgument, c.Error("地址不能为空").Error(), nil
	}
	if reqData.UId == 0 {
		return api.IllegalArgument, c.Error("用户Id不能为空").Error(), nil
	}
	
	// 校验用户和id是否存在
	u, err := user.GetUser("1 = 1 AND status = 1 AND address = ? AND id = ?", reqData.Address, reqData.UId)
	if err != nil {
		return api.DatabaseError, c.Error(err).Error(), nil
	}
	if u == nil {
		return api.DatabaseError, c.Error("用户不存在").Error(), nil
	}

	id, err := blog.AddBlog(&blog.Blog{
		Address: reqData.Address,
		Title : reqData.Title,
		BlogText : reqData.BlogText,
		UID : reqData.UId,
		ImageURL : reqData.ImageUrl,
		Status: 1,
		Creator: reqData.Address,
		CreateTime: time.Now(),
		Updater: reqData.Address,
		UpdateTime: time.Now(),
	})
	if err != nil {
		return api.DatabaseError, c.Error(err).Error(), nil
	}
	return 0, "", api.H{"id": id, "blogUrl" : "http://wengyifan.com:8080/blog/getBlog?id=" + utils.String(id)}
}

func DoLike(c *api.Context) (int, string, interface{}) {
	uId, err := strconv.Atoi(c.Query("uId"))
	if err != nil {
		return api.IllegalArgument, c.Error(err).Error(), ""
	}
	blogId, err := strconv.Atoi(c.Query("blogId"))
	if err != nil {
		return api.IllegalArgument, c.Error(err).Error(), ""
	}
	like, err := strconv.Atoi(c.Query("like"))
	if err != nil {
		return api.IllegalArgument, c.Error(err).Error(), ""
	}
	if like < 0 || like > 1 {
		return api.IllegalArgument, c.Error("like参数不合法").Error(), ""
	}
	// 验证用户存在
	u, err := user.GetUser("1 = 1 AND status = 1 AND id = ?", uId)
	if err != nil {
		return api.DatabaseError, c.Error(err).Error(), nil
	}
	if u == nil {
		return api.DatabaseError, c.Error("用户不存在").Error(), nil
	}

	// 验证博客存在
	b, err := blog.GetBlog("1 = 1 AND status = 1 AND id = ?", blogId)
	if err != nil {
		return api.DatabaseError, c.Error(err).Error(), nil
	}
	if b == nil {
		return api.DatabaseError, c.Error("博客不存在").Error(), nil
	}

	// 检测这个点赞是否存在
	bl, err := blog.GetBlogLike("1 = 1 AND status = 1 AND blogId = ? AND uId = ?", blogId, uId);
	if err != nil {
		return api.DatabaseError, c.Error(err).Error(), nil
	}

	if bl != nil && like == 1{
		return api.DatabaseError, c.Error("请不要重复点赞").Error(), nil
	}
	if bl == nil && like == 0 {
		return api.DatabaseError, c.Error("请不要重复取消点赞").Error(), nil
	}
	if like == 0 && bl != nil{
		// update设置status = 0
		bl.Status = 0
		blog.UpdateBlogLike(bl)
	}
	if like == 1 && bl == nil {
		// 新插入一条数据
		blike := &blog.BlogLike{
			BlogID: blogId,
			UID: uId,
			Status: 1,
			Creator: "",
			CreateTime: time.Now(),
			Updater: "",
			UpdateTime: time.Now(),
		}
		_, err := blog.AddBlogLike(blike);
		if err != nil {
			return api.DatabaseError, c.Error(err).Error(), ""
		}
	}

	return 0, "success", nil
}