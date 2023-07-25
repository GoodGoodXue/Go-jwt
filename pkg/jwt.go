package pkg

import (
	"Practice/Go-Projects/jwt/models"
	"time"

	"github.com/golang-jwt/jwt"
)

type MyClaims struct {
	UserId uint
	jwt.StandardClaims
}

var MySercet = []byte("waldjf1la")

// 生成token
func CreateToken(user models.MyUser) (string, error) {

	// 声明部分
	claims := &MyClaims{
		UserId: user.ID,

		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24 * 30).Unix(),

			IssuedAt: time.Now().Unix(),

			Issuer: "daquan",

			Subject: "UserID",

			Audience: "UserID",
		},
	}

	// 加密声明 // 加私钥生成token
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(MySercet)
	if err != nil {
		return "Failed to token", err
	}

	return token, nil

}

// 解析token
func ParseToken(StringToken string) (*jwt.Token, *MyClaims, error) {

	// 解析出声明模型内容
	claims := &MyClaims{}

	token, _ := jwt.ParseWithClaims(StringToken, claims, func(t *jwt.Token) (interface{}, error) {
		return MySercet, nil
	})

	// 返回参数
	return token, claims, nil
}
