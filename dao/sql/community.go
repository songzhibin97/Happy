/******
** @创建时间 : 2020/8/26 20:50
** @作者 : SongZhiBin
******/
package sql

import (
	"Happy/model/model"
	"go.uber.org/zap"
)

// GetCommunityList:获取社区列表
func GetCommunityList() ([]*model.CommunityList, error) {
	sqlString := `SELECT community_id,community_name FROM community`
	res := make([]*model.CommunityList, 0)
	err := SearchAll(dbInstantiate, sqlString, &res)
	if err != nil {
		zap.L().Error("GetCommunityList Error", zap.Error(err))
		return res, err
	}
	return res, nil
}

// CommunityDetail:获取社区详情
func CommunityDetail(id int64) (*model.CommunityDetail, error) {
	sqlString := `SELECT community_id, community_name, introduction, create_time FROM community WHERE community_id = ?`
	res := new(model.CommunityDetail)
	err := SearchRow(dbInstantiate, sqlString, res, id)
	if err != nil {
		zap.L().Error("GetCommunityList Error", zap.Error(err))
		return res, err
	}
	return res, nil
}
