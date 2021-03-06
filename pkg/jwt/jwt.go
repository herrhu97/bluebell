package jwt

import (
	"errors"
	"time"

	"github.com/spf13/viper"

	"github.com/dgrijalva/jwt-go"
)

// MyClaims 自定义声明结构体并内嵌jwt.StandardClaims
// jwt包自带的jwt.StandardClaims只包含了官方字段
// 我们这里需要额外记录一个username字段，所以要自定义结构体
// 如果想要保存更多信息，都可以添加到这个结构体中
type MyClaims struct {
	Username string `json:"username"`
	UserID   int64  `json:"user_id"`
	jwt.StandardClaims
}

var MySecret = []byte("夏天的风")

// GenToken 生成JWT
func GenToken(username string, userID int64) (string, error) {
	// 创建一个我们自己的声明
	c := MyClaims{
		username, // 自定义字段
		userID,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(viper.GetDuration("auth.jwt_expire") * time.Hour).Unix(), // 过期时间
			Issuer:    "bluebell",                                                              // 签发人
		},
	}
	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	return token.SignedString(MySecret)
}

// ParseToken 解析JWT
func ParseToken(tokenString string) (*MyClaims, error) {
	var mc = new(MyClaims)
	// 解析token
	token, err := jwt.ParseWithClaims(tokenString, mc, func(token *jwt.Token) (i interface{}, err error) {
		return MySecret, nil
	})
	if err != nil {
		return nil, err
	}
	if token.Valid { // 校验token
		return mc, nil
	}
	return nil, errors.New("invalid token")
}
