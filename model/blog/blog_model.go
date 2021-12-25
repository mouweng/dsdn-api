// Code generated by go-dbmodel-gen. DO NOT EDIT.
// versions:
// 	go-dbmodel-gen v1.0.2
// nolint
package model

import (
    "time"
    "fmt"

	"ginTemplate/model/common"
)

/** JSGEN({type: "model", paged: true})
CREATE TABLE tbBlog (
    id int(10) NOT NULL AUTO_INCREMENT  COMMENT 'ID',
    Address varchar(512) NOT NULL COMMENT '地址',
    title varchar(1024) NOT NULL DEFAULT '' COMMENT '标题',
    blogText longtext NOT NULL COMMENT '博文内容',
    uId int(10) NOT NULL DEFAULT 0 COMMENT 'userId',
    imageUrl varchar(1024) NOT NULL DEFAULT '' COMMENT '图片地址',
    status int(10) NOT NULL DEFAULT 1 COMMENT '记录状态：1表示可用，0表示不可用',
    creator varchar(64) NOT NULL DEFAULT '' COMMENT '创建者',
    createTime datetime NOT NULL COMMENT '创建时间@now',
    updater varchar(64) NOT NULL DEFAULT '' COMMENT '最后更新人',
    updateTime datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (id)
);
JSGEN **/

var _ = time.Now

// BlogConnection Blog连接类型
type BlogConnection func() common.DBConnect

// DefaultBlogConnection DefaultBlog默认连接
var DefaultBlogConnection BlogConnection

// Blog Blog值类型
type Blog struct {
  ID int `json:"id"`
  Address string `json:"Address"`
  Title string `json:"title"`
  BlogText string `json:"blogText"`
  UID int `json:"uId"`
  ImageURL string `json:"imageUrl"`
  Status int `json:"status"`
  Creator string `json:"creator"`
  CreateTime time.Time `json:"createTime"`
  Updater string `json:"updater"`
  UpdateTime time.Time `json:"updateTime"`
}

// Add 插入Blog
func (c BlogConnection) Add(model *Blog) (int64, error) {
    sqlStr := "INSERT INTO `tbBlog` (`Address`, `title`, `blogText`, `uId`, `imageUrl`, `status`, `creator`, `createTime`, `updater`) VALUES(?, ?, ?, ?, ?, ?, ?, now(), ?)"
    result, err := c().Exec(sqlStr, model.Address, model.Title, model.BlogText, model.UID, model.ImageURL, model.Status, model.Creator, model.Updater)
    if err != nil {
        return 0, err
    } 
    
    return result.LastInsertId()
}

// AddBlog 插入Blog
func AddBlog(model *Blog) (int64, error) {
    return DefaultBlogConnection.Add(model)
}

// Find 查询Blog
func (c BlogConnection) Find(condition string, args ...interface{}) ([]*Blog, error) {
    sqlStr := "SELECT `id`, `Address`, `title`, `blogText`, `uId`, `imageUrl`, `status`, `creator`, `createTime`, `updater`, `updateTime` FROM `tbBlog`"
    if len(condition) > 0 {
        sqlStr = sqlStr + " WHERE " + condition
    }
    results := make([]*Blog, 0)

	stmt, err := c().Prepare(sqlStr)
	if err != nil {
		return results, err
	}
	defer stmt.Close()
	rows, err := stmt.Query(args...)
    if err != nil {
        return results, err
    } 
    
        defer rows.Close()
        for rows.Next() {
            model := Blog{}
            values := []interface{}{
              new(interface{}),
              new(interface{}),
              new(interface{}),
              new(interface{}),
              new(interface{}),
              new(interface{}),
              new(interface{}),
              new(interface{}),
              new(interface{}),
              new(interface{}),
              new(interface{}),
            }
            rows.Scan(values...)
            if *(values[0].(*interface{})) != nil {
                tmp := int((*(values[0].(*interface{}))).(int64))
                model.ID = tmp
            }
            if *(values[1].(*interface{})) != nil {
                tmp := string((*(values[1].(*interface{}))).([]uint8))
                model.Address = tmp
            }
            if *(values[2].(*interface{})) != nil {
                tmp := string((*(values[2].(*interface{}))).([]uint8))
                model.Title = tmp
            }
            if *(values[3].(*interface{})) != nil {
                tmp := string((*(values[3].(*interface{}))).([]uint8))
                model.BlogText = tmp
            }
            if *(values[4].(*interface{})) != nil {
                tmp := int((*(values[4].(*interface{}))).(int64))
                model.UID = tmp
            }
            if *(values[5].(*interface{})) != nil {
                tmp := string((*(values[5].(*interface{}))).([]uint8))
                model.ImageURL = tmp
            }
            if *(values[6].(*interface{})) != nil {
                tmp := int((*(values[6].(*interface{}))).(int64))
                model.Status = tmp
            }
            if *(values[7].(*interface{})) != nil {
                tmp := string((*(values[7].(*interface{}))).([]uint8))
                model.Creator = tmp
            }
            if *(values[8].(*interface{})) != nil {
                tmp := (*(values[8].(*interface{}))).(time.Time)
                model.CreateTime = tmp
            }
            if *(values[9].(*interface{})) != nil {
                tmp := string((*(values[9].(*interface{}))).([]uint8))
                model.Updater = tmp
            }
            if *(values[10].(*interface{})) != nil {
                tmp := (*(values[10].(*interface{}))).(time.Time)
                model.UpdateTime = tmp
            }
            results = append(results, &model)
        }
    return results, nil
}

// FindBlog 查询Blog
func FindBlog(condition string, args ...interface{}) ([]*Blog, error) {
    return DefaultBlogConnection.Find(condition, args...)
}
// PagedQuery 分页查询Blog
func (c BlogConnection) PagedQuery(condition string, pageSize uint, page uint, args ...interface{}) (totalCount uint, rows []*Blog, err error) {
	sqlStr := "SELECT COUNT(1) as cnt FROM `tbBlog`"
	if len(condition) > 0 {
		sqlStr = sqlStr + " WHERE " + condition
	}

	cr := c().QueryRow(sqlStr, args...)

	err = cr.Scan(&totalCount)
	if err != nil {
		return 0, nil, err
	}
	if page > 0 {
		page = page - 1
	}
	offset := page * pageSize
	if totalCount <= offset {
		return totalCount, []*Blog{}, nil
	}

	if len(condition) == 0 {
		condition = fmt.Sprintf("1=1")
	}
	condition = condition + fmt.Sprintf(" LIMIT %d OFFSET %d", pageSize, offset)
	rows, err = c.Find(condition, args...)
	return
}

// BlogPagedQuery 分页查询Blog
func BlogPagedQuery(condition string, pageSize uint, page uint, args ...interface{}) (totalCount uint, rows []*Blog, err error) {
	return DefaultBlogConnection.PagedQuery(condition, pageSize, page, args...)
}

// Get 获取Blog
func (c BlogConnection) Get(condition string, args ...interface{}) (*Blog, error) {
    results, err := c.Find(condition, args...)

    if err != nil {
        return nil, err
    } 
    
    if len(results) > 0 {
        return results[0], nil
    } 
        
    return nil, nil
}


// GetBlog 获取Blog
func GetBlog(condition string, args ...interface{}) (*Blog, error) {
    return DefaultBlogConnection.Get(condition, args...)
}

// Update 更新Blog
func (c BlogConnection) Update(model *Blog) (int64, error) {
    sqlStr := "UPDATE `tbBlog` SET `Address` = ?, `title` = ?, `blogText` = ?, `uId` = ?, `imageUrl` = ?, `status` = ?, `creator` = ?, `updater` = ? WHERE `id` = ?"
    result, err := c().Exec(sqlStr, model.Address, model.Title, model.BlogText, model.UID, model.ImageURL, model.Status, model.Creator, model.Updater, model.ID)
    if err != nil {
        return 0, err
    }
    return result.RowsAffected()
}

// UpdateBlog 更新Blog
func UpdateBlog(model *Blog) (int64, error) {
    return DefaultBlogConnection.Update(model)
}