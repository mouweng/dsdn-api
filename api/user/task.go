package user

import (
	"fmt"
	"ginTemplate/api"
	user "ginTemplate/model/user"
	"strconv"
	"time"

	"github.com/wonderivan/logger"
)

type RequestInfo struct{
	ID int `json:"id"`
	Address string `json:"address"  binding:"required"`
	NickName string `json:"nickName"`
	Email string `json:"email"`
	ImageUrl string `json:"imageUrl"`
	Introduction string `json:"introduction"`
}

// 获取用户
func DoGetUser(c *api.Context) (int, string, interface{}) {
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
		"address": {FieldName: "address", DefaultValue: "-1", CheckValue: nil, Need: false, Link: "and", Symbol: "="},
		"nickName":{FieldName: "nickName", DefaultValue: "", CheckValue: nil, Need: false, Link: "and", Symbol: "like"},
		"email":   {FieldName: "email", DefaultValue: "", CheckValue: nil, Need: false, Link: "and", Symbol: "="},
		"introduction":{FieldName: "introduction", DefaultValue: "", CheckValue: nil, Need: false, Link: "and", Symbol: "="},
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
	fmt.Println(strCondition)
	fmt.Println(args...)
	
	user, err := user.GetUser(strCondition, args...)
	if err != nil {
		return api.IllegalArgument, c.Error(err).Error(), nil
	}
	if (user == nil) {
		return api.DatabaseError, c.Error("用户不存在").Error(), nil
	}
	return 0, "success", user
}


// 注册用户
func DoRegisterUser(c *api.Context) (int, string, interface{}) {
	reqData := RequestInfo{}
	if err := c.ShouldBind(&reqData); err != nil {
		return api.IllegalArgument, c.Error(err).Error(), nil
	}
	logger.Debug("reqData:",reqData);

	if reqData.Address == "" {
		return api.IllegalArgument, c.Error("地址不能为空").Error(), nil
	}

	u, err := user.GetUser("1 = 1 AND status = 1 AND address = ?", reqData.Address)
	if err != nil {
		return api.DatabaseError, c.Error(err).Error(), nil
	}
	if u != nil {
		return api.DatabaseError, c.Error("用户已存在").Error(), nil
	}

	id, err := user.AddUser(&user.User{
		Address: reqData.Address,
		NickName: reqData.NickName,
		Email: reqData.Email,
		ImageURL: reqData.ImageUrl,
		Introduction: reqData.Introduction,
		Status: 1,
		Creator: reqData.Address,
		CreateTime: time.Now(),
		Updater: reqData.Address,
		UpdateTime: time.Now(),
	})
	if err != nil {
		return api.DatabaseError, c.Error(err).Error(), nil
	}
	return 0, "", api.H{"id": id}
}


// 注册用户
func DoFollowUser(c *api.Context) (int, string, interface{}) {
	uId1, err := strconv.Atoi(c.Query("uId1"))
	if err != nil {
		return api.IllegalArgument, c.Error(err).Error(), ""
	}
	uId2, err := strconv.Atoi(c.Query("uId2"))
	if err != nil {
		return api.IllegalArgument, c.Error(err).Error(), ""
	}
	// 校验u1和u2是否存在
	user1, err := user.GetUser("1 = 1 AND status = 1 AND id = ?", uId1)
	if err != nil {
		return api.DatabaseError, c.Error(err).Error(), nil
	}
	if user1 == nil {
		return api.DatabaseError, c.Error("用户不存在").Error(), nil
	}
	// 校验u1和u2是否存在
	user2, err := user.GetUser("1 = 1 AND status = 1 AND id = ?", uId2)
	if err != nil {
		return api.DatabaseError, c.Error(err).Error(), nil
	}
	if user2 == nil {
		return api.DatabaseError, c.Error("用户不存在").Error(), nil
	}

	// 关注表， 和like的逻辑类似！


	return 0, "", nil
}
