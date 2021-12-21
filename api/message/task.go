package message

import (
	"database/sql"
	"strconv"
	"time"

	"ginTemplate/api"
	"ginTemplate/config"
	center "ginTemplate/model/center"
	"ginTemplate/model/common"
)

// DoGetCenter
func DoGetCenter(c *api.Context) (int, string, interface{}) {
	centerList, err := center.FindCenter("1 = 1")
	if err != nil {
		return api.IllegalArgument, c.Error(err).Error(), nil
	}
	return 0, "success", centerList
}

// centerRequest
type centerRequest struct {
	ID              int    `json:"id"`
	CenterName            string `json:"centerName" binding:"required"`
	CenterCode          string `json:"centerCode" binding:"required"`
}

// DoGetCenter
func DoAddCenter(c *api.Context) (int, string, interface{}) {
	reqData := centerRequest{}
	if err := c.ShouldBind(&reqData); err != nil {
		return api.IllegalArgument, c.Error(err).Error(), nil
	}


	id, err := center.AddCenter(&center.Center{
		CenterName: reqData.CenterName,
		CenterCode: reqData.CenterCode,
		Status: 1,
		Creator: "creator",
		CreateTime: time.Now(),
		Updater: "updater",
		UpdateTime: time.Now(),
	})
	if err != nil {
		return api.DatabaseError, c.Error(err).Error(), nil
	}

	return 0, "success", api.H{"id": id}
}

// DoDelCenter
func DoDelCenter(c *api.Context) (int, string, interface{}) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		return api.IllegalArgument, c.Error(err).Error(), ""
	}

	// 使用事务保证一致性
	var tx *sql.Tx
	tx, err = config.GetDBConnect(config.DbName)().Begin()
	if err != nil {
		return api.DatabaseError, c.Error(err).Error(), nil
	}
	txCenterExecutor := center.CenterConnection(common.Tx2DBConnect(tx));
	defer tx.Rollback()


	ct,err := txCenterExecutor.Get(" 1 = 1 AND id = ?", id);
	if err != nil {
		return api.DatabaseError, c.Error(err).Error(), nil
	}
	if ct == nil {
		return api.DatabaseError, c.Error("数据不存在").Error(), nil
	}
	ct.Status = 0;
	res, err := txCenterExecutor.Update(ct);
	if err != nil {
		return api.DatabaseError, c.Error(err).Error(), nil
	}

	// 结束事务
	err = tx.Commit()
	if err != nil {
		return api.DatabaseError, c.Error(err).Error(), nil
	}

	return 0, "success", api.H{"id": res}
}