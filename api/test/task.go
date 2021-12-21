package test

import (
	"fmt"
	"ginTemplate/api"
	"ginTemplate/config"
	"strconv"

	"github.com/wonderivan/logger"
)

// ping测试链接
func Pong(c *api.Context) (int, string, interface{}) {
	fmt.Println(config.String("mode"))
	return 0, "success", api.H{"ping": "pong"}
}

// GET Param 解析参数
func DoTestParam(c *api.Context) (int, string, interface{}){
	//获取分页参数:PageIndex,PageSize
	page, pageSize, code, err := c.GetPager()
	if err != nil {
		return code, c.Error(err).Error(), nil
	}
	logger.Debug("page:", page)
	logger.Debug("pageSize:", pageSize)
	logger.Debug("code:", code)

	parConstruct := map[string]*api.ParamConstruct{
		"id":          {FieldName: "id", DefaultValue: "", CheckValue: nil, Need: false, Link: "and", Symbol: "="},
		"pId":         {FieldName: "pId", DefaultValue: "", CheckValue: nil, Need: false, Link: "and", Symbol: "="},
		"jId":      {FieldName: "jId", DefaultValue: "", CheckValue: nil, Need: true, Link: "and", Symbol: "="},
		"status":      {FieldName: "status", DefaultValue: "1", CheckValue: nil, Need: false, Link: "and", Symbol: "="},
		"name":        {FieldName: "name", DefaultValue: "", CheckValue: nil, Need: false, Link: "and", Symbol: "like"},
		"creator": {FieldName: "creator", DefaultValue: "", CheckValue: nil, Need: false, Link: "and", Symbol: "="},
		"orderStatus": {FieldName: "orderStatus", DefaultValue: "", CheckValue: nil, Need: false, Link: "and", Symbol: "="},
		"startTime": {FieldName: "statTime", DefaultValue: "", CheckValue: nil, Need: false, Link: "and", Symbol: ">="},
		"endTime": {FieldName: "statTime", DefaultValue: "", CheckValue: nil, Need: false, Link: "and", Symbol: "<="},
		"orderBy": {FieldName: "", DefaultValue: "statTime|desc", CheckValue: nil, Need: false},
	}
	c.Set("isNotAllCondition", false)
	strCondition, args, err := c.GetConditionByParam(parConstruct)
	if err != nil {
		return api.IllegalArgument, c.Error(err).Error(), nil
	}
	logger.Debug("strCondition:", strCondition)
	logger.Debug("args:", args)
	
	return 0, "success", strCondition
}

// GET Query 解析参数
func DoTestQuery(c *api.Context) (int, string, interface{}){
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		return api.IllegalArgument, c.Error(err).Error(), ""
	}
	logger.Debug("id:", id)
	return 0, "success", id
}



type RequestInfo struct{
	ID int `json:"id"`
	Type *int `json:"type" binding:"required"`
	Name string `json:"name" binding:"required"`
}

// 测试Post绑定参数
func DoTestPost(c *api.Context) (int, string, interface{}){
	reqData := RequestInfo{}
	if err := c.ShouldBind(&reqData); err != nil {
		return api.IllegalArgument, c.Error(err).Error(), nil
	}
	logger.Debug("reqData:",reqData);
	return 0, "success", reqData
}
