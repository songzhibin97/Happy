package sql

import (
	"Happy/dao/redis"
	"Happy/model/model"

	"go.uber.org/zap"
)

// 帖子相关数据库函数

// CreatePost:新建帖子
func CreatePost(postId, authorId, communityId int64, title, content string) error {
	sqlString := `INSERT INTO post(post_id,title,content,author_id,community_id) VALUE(?,?,?,?,?)`
	_, _, err := Exec(dbInstantiate, sqlString, postId, title, content, authorId, communityId)
	return err
}

// GetPostDetail:获取帖子详情
func GetPostDetail(postId int64) (*model.Post, error) {
	// todo 这里没有对帖子状态进行判断
	sqlString := `SELECT post_id,title,content,author_id,community_id,create_time FROM post WHERE post_id = ?`
	p := new(model.Post)
	err := SearchRow(dbInstantiate, sqlString, p, postId)
	if err != nil {
		return p, err
	}
	// todo 有可能会报错
	AnalyzeVoting(p)
	return p, nil
}

// GetApiPostDetail:获取详细信息(关于社区等信息进行拼接)
func GetApiPostDetail(postId int64) (*model.ApiPostDetail, error) {
	// 根据id找post详情
	p, err := GetPostDetail(postId)
	if err != nil {
		zap.L().Error("GetPostDetail", zap.Error(err), zap.Int("postId", int(postId)))
		return nil, err
	}
	// 根据uid找用户名
	u, err := GetUserName(p.AuthorID)
	if err != nil {
		zap.L().Error("GetUserName", zap.Error(err), zap.Int("uid", int(p.AuthorID)))
		return nil, err
	}
	// 根据CommunityId找社区相关信息
	c, err := CommunityDetail(p.CommunityID)
	if err != nil {
		zap.L().Error("CommunityDetail", zap.Error(err), zap.Int("cid", int(p.AuthorID)))
		return nil, err
	}
	return &model.ApiPostDetail{
		AuthorName:      u.Username,
		Post:            p,
		CommunityDetail: c,
	}, nil
}

// GetPostListToUid:根据用户ID获取帖子列表
func GetPostListToUid(uid int64, page, max int64) ([]*model.Post, error) {
	sqlString := `SELECT post_id,title,content,author_id,community_id,create_time FROM post WHERE author_id = ? limit ?,?`
	return GetPostList(uid, sqlString, page, max)
}

// GetPostListToCid:根据社区ID获取帖子列表
func GetPostListToCid(cid int64, page, max int64) ([]*model.Post, error) {
	sqlString := `SELECT post_id,title,content,author_id,community_id,create_time FROM post WHERE community_id = ? limit ?,?`
	return GetPostList(cid, sqlString, page, max)
}

// GetPostList:封装
func GetPostList(id int64, sqlString string, page, max int64) ([]*model.Post, error) {
	pList := new([]*model.Post)
	err := SearchAll(dbInstantiate, sqlString, pList, id, (page-1)*max, max)
	// 新增一步处理流程 如果有错误不进行处理 如果没错误增加投票信息
	if err != nil {
		return *pList, err
	}
	// todo 存在可优化区间 使用redis pipeline 减少 RTT(网络连接次数)
	for _, v := range *pList {
		AnalyzeVoting(v)
	}
	return *pList, err
}

// AnalyzeVoting:封装 获取投票数
func AnalyzeVoting(post *model.Post) *model.Post {
	if post.Votes == nil {
		post.Votes = new(model.Votes)
	}
	// 获取接口
	res, ok := redis.GetPostXY(post.ID)
	if !ok {
		post.Like = 0
		post.Unlike = 0
		return post
	}
	post.Like = res[0]
	post.Unlike = res[1]
	return post
}
