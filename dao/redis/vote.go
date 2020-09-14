/******
** @创建时间 : 2020/9/13 10:33
** @作者 : SongZhiBin
******/
package redis

import (
	"Happy/model/model"
	"Happy/pkg/snowflake"
	"github.com/go-redis/redis"
	"math"
	"strconv"
	"time"
)

/*
关于投票功能的一些辅助函数
*/

// CreatePost:创建文章后同时在Redis中生成对应的对列 用于记录 t,x,y等信息
// Ps:创建文章默认作者给自己一票赞 作者不允许取消或者点踩自己
func CreatePost(postId int64, uid int64) {
	var StandardTime = snowflake.GetStartTime().Unix()
	t := time.Now().Unix() - StandardTime
	objs := map[string]interface{}{
		KeyPostHashT: t,
		KeyPostHashX: 1,
		KeyPostHashY: 0,
		KeyPostAuth:  uid,
	}
	// 创建hash
	HashMAddSplice(rdb, KeyPostHash, strconv.FormatInt(postId, 10), objs)
	// 创建zSet
	AddScore(postId, CalculateScore(float64(t), 1, 0))
}

// PostIdExist:判断是否存在
func PostIdExist(postID int64) bool {
	return ExistsKey(rdb, KeyPostHash, strconv.FormatInt(postID, 10))
}

// PostIdExistReturnG:判断帖子ID是否有效并返回对应HSet结构
func PostIdExistReturnG(postID int64) (*model.GVoted, bool) {
	pid := strconv.FormatInt(postID, 10)
	// 判断是否存在
	if !ExistsKey(rdb, KeyPostHash, pid) {
		return nil, false
	}
	// 存在后,获取对应数据
	ires := HashMGet(rdb, KeyPostHash, pid, KeyPostHashT, KeyPostHashX, KeyPostHashY)
	// 如果ires == nil 代表批量查找失败
	// 如果ires != 3 代表查找值有误
	if ires == nil || len(ires) != 3 {
		return nil, false
	}
	res := new(model.GVoted)
	var ok bool
	// 断言
	t, ok := AToi(ires[0])
	res.T = int64(t)
	if !ok {
		return nil, false
	}
	res.X, ok = AToi(ires[1])
	if !ok {
		return nil, false
	}
	res.Y, ok = AToi(ires[2])
	if !ok {
		return nil, false
	}
	return res, true
}

// JudgeAuth:判断投票是否是作者本人
func JudgeAuth(postId, uid int64) bool {
	auth := HashGet(rdb, KeyPostHash, strconv.FormatInt(postId, 10), KeyPostAuth)
	if auth == nil {
		return false
	}
	v, ok := auth.(string)
	if !ok {
		return false
	}
	iv, err := strconv.Atoi(v)
	if err != nil {
		return false
	}
	return int64(iv) == uid
}

// GetPostXY:获取当前文章点赞数以及踩数
func GetPostXY(postId int64) ([2]int, bool) {
	pid := strconv.FormatInt(postId, 10)
	// 1.先判断postId文章是否存在redis
	if HashIsExists(rdb, KeyPostHash, pid, KeyPostHashX) &&
		HashIsExists(rdb, KeyPostHash, pid, KeyPostHashY) {
		// 2.如果存在redis中则取出
		ires := HashMGet(rdb, KeyPostHash, pid, KeyPostHashX, KeyPostHashY)
		// 如果ires == nil 代表批量查找失败
		// 如果ires != 2 代表查找的值有误
		if ires == nil || len(ires) != 2 {
			return [2]int{0, 0}, false
		}
		x, ok := AToi(ires[0])
		if !ok {
			return [2]int{0, 0}, false
		}
		y, ok := AToi(ires[1])
		if !ok {
			return [2]int{0, 0}, false
		}

		return [2]int{x, y}, true
	}
	// todo 3.如果不存在redis从mysql中寻找
	return [2]int{0, 0}, false
}

// ChangeX:更新点赞数
func ChangeX(postId int64, i int64) {
	HashChange(rdb, KeyPostHash, strconv.FormatInt(postId, 10), KeyPostHashX, i)
}

// ChangeY:更新点踩数
func ChangeY(postId int64, i int64) {
	HashChange(rdb, KeyPostHash, strconv.FormatInt(postId, 10), KeyPostHashY, i)
}

// ChangeScore:更新分数
func ChangeScore(postId int64) bool {
	obj, ok := PostIdExistReturnG(postId)
	if !ok {
		return false
	}
	ZSetChangeV(rdb, KeyPostScore, CalculateScore(float64(obj.T), float64(obj.X), float64(obj.Y)),
		strconv.FormatInt(postId, 10))
	return true
}

// GetUserVote:获取用户是否存在
func GetUserVote(postId, uid int64) (int, bool) {
	return GetZSetV(rdb, KeyPostVotePF, strconv.FormatInt(postId, 10), strconv.FormatInt(uid, 10))
}

// AddScore:添加分数/修改 ZSet
func AddScore(postId int64, score float64) {
	ZSetAdd(rdb, KeyPostScore, redis.Z{
		Score:  score,
		Member: strconv.FormatInt(postId, 10),
	})
}

// AddUserVote:添加用户提交 ZSet
func AddUserVote(postId int64, uid int64, ic int64) {
	ZSetAddSplice(rdb, KeyPostVotePF, strconv.FormatInt(postId, 10), redis.Z{
		Score:  float64(ic),
		Member: strconv.FormatInt(uid, 10),
	})
}

// CalculateScore:计算分数
func CalculateScore(t, x, y float64) float64 {
	sX := x - y // 赞成票 - 反对票
	sY := 0
	if sX > 0 {
		sY = 1
	} else if sX < 0 {
		sY = -1
	}
	sZ := 1
	if x != 0 {
		sZ = int(math.Abs(x))
	}
	return (math.Log10(float64(sZ))) + ((float64(sY) * t) / float64(4500))
}

// 工具函数
// AToi:interface转化为string后转化为int64
func AToi(i interface{}) (int, bool) {
	si, ok := i.(string)
	if !ok {
		return 0, false
	}
	is, err := strconv.Atoi(si)
	if err != nil {
		return 0, false
	}
	return is, true
}
