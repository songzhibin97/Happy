package redis

import (
	"fmt"
	"testing"
)

func TestIsZSet(t *testing.T) {
	rdb = CRedis(
		"127.0.0.1",
		6379,
		"",
		0)
	if rdb == nil {
		panic("Redis init is nil ")
	}
	fmt.Println(IsZSet(rdb, "", "zzz", "z1z2"))
}

func TestZSetChangeV(t *testing.T) {
	rdb = CRedis(
		"127.0.0.1",
		6379,
		"",
		0)
	if rdb == nil {
		panic("Redis init is nil ")
	}
	fmt.Println(ZSetChangeVSplit(rdb, "", "zzz", 1, "z1z1"))
}

func TestZSetRemove(t *testing.T) {
	rdb = CRedis(
		"127.0.0.1",
		6379,
		"",
		0)
	if rdb == nil {
		panic("Redis init is nil ")
	}
	ZSetRemoveSplit(rdb, "", "zzz", "z1z1")
}
