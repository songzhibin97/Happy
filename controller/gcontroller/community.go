/******
** @创建时间 : 2020/8/26 20:21
** @作者 : SongZhiBin
******/
package gcontroller

import (
	sqls "Happy/dao/sql"
	"Happy/model/gmodel"
	"Happy/model/model"
	pb "Happy/model/pmodel/community"
	"context"
	"database/sql"
	"encoding/json"
	"go.uber.org/zap"
)

// CommunityServer:定义CommunityServer
type CommunityServer struct{}

// CommunityList:获取社区列表
func (c *CommunityServer) CommunityList(ctx context.Context, request *pb.CommunityListRequest) (*pb.Response, error) {
	// 获取社区列表
	res, err := sqls.GetCommunityList()
	if err != nil {
		if err == sql.ErrNoRows {
			return (*pb.Response)(gmodel.ResponseError(model.CodeGetListEmpty)), nil
		}
		return (*pb.Response)(gmodel.ResponseError(model.CodeGetListError)), nil
	}
	resString, err := json.Marshal(res)
	if err != nil {
		zap.L().Error("CommunityList Marshal Error", zap.Error(err))
		return (*pb.Response)(gmodel.ResponseError(model.CodeServerBusy)), nil
	}
	return (*pb.Response)(gmodel.ResponseSuccess(map[string]string{"communityList": string(resString)})), nil
}

// CommunityDetail:根据社区id获取社区详情
func (c *CommunityServer) CommunityDetail(ctx context.Context, request *pb.CommunityDetailRequest) (*pb.Response, error) {
	// 参数校验
	r, err := _verification(request)
	if err != nil {
		return (*pb.Response)(r), nil
	}

	// 获取社区详情
	res, err := sqls.CommunityDetail(request.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			return (*pb.Response)(gmodel.ResponseError(model.CodeGetListEmpty)), nil
		}
		return (*pb.Response)(gmodel.ResponseError(model.CodeGetListError)), nil
	}
	resString, err := json.Marshal(res)
	if err != nil {
		zap.L().Error("CommunityDetail Marshal Error", zap.Error(err))
		return (*pb.Response)(gmodel.ResponseError(model.CodeServerBusy)), nil
	}
	return (*pb.Response)(gmodel.ResponseSuccess(map[string]string{"CommunityDetail": string(resString)})), nil
}
