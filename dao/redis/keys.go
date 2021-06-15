package redis

const (
	KeyPrefix              = "bluebell:"
	KeyPostTimeZset        = "post:time"   // 帖子及发帖时间
	KeyPostScoreZset       = "post:score"  // 帖子及分数
	KeyPostVotedZsetPrefix = "post:voted:" //记录用户及投票类型，参数是post_id
)

func getRedisKey(key string) string {
	return KeyPrefix + key
}
