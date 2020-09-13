/******
** @创建时间 : 2020/9/13 10:33
** @作者 : SongZhiBin
******/
package redis

import (
	"Happy/pkg/snowflake"
	"github.com/go-redis/redis"
	"math"
	"strconv"
	"time"
)

/*
投票功能:

*/

var StandardTime = snowflake.GetStartTime().Unix()

// CreatePost:创建文章后同时在Redis中生成对应的对列 用于记录 t,x,y等信息
func CreatePost(postId int64) {
	t := time.Now().Unix() - StandardTime
	objs := map[string]interface{}{
		KeyPostHashT: t,
		KeyPostHashX: 0,
		KeyPostHashY: 0,
	}
	// 创建hash
	HashMAddSplice(rdb, KeyPostHash, strconv.FormatInt(postId, 10), objs)
	// 创建zSet
	AddScore(postId, CalculateScore(t, 0, 0))
}

// ChangeX:更新点赞数
func ChangeX(postId int64, i int64) {
	HashChange(rdb, KeyPostHash, strconv.FormatInt(postId, 10), KeyPostHashX, i)
}

// ChangeY:更新点踩数
func ChangeY(postId int64, i int64) {
	HashChange(rdb, KeyPostHash, strconv.FormatInt(postId, 10), KeyPostHashY, i)
}

// AddScore:添加分数/修改 ZSet
func AddScore(postId int64, score int64) {
	ZSetAdd(rdb, KeyPostScore, redis.Z{
		Score:  float64(score),
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
func CalculateScore(t, x, y int64) int64 {
	sX := x - y // 赞成票 - 反对票
	sY := 0
	if sX > 0 {
		sY = 1
	} else if sX < 0 {
		sY = -1
	}
	sZ := 1
	if x != 0 {
		sZ = int(math.Abs(float64(x)))
	}
	return int64(math.Log10(float64(sZ))) + ((int64(sY) * t) / 4500)
}
