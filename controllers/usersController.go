package controller

import (
	"Practice/Go-Projects/jwt/initializes"
	"Practice/Go-Projects/jwt/models"
	"Practice/Go-Projects/jwt/pkg"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(c *gin.Context) {
	// 接受参数
	var body struct {
		Name     string
		Email    string
		Password string
	}
	err := c.Bind(&body)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to query",
		})
		return
	}

	// 判断邮箱是否已注册
	var user models.MyUser
	result := initializes.DB.First(&user, "email = ?", body.Email)

	// 存储操作的错误信息，有值为nil，未找到含错误信息
	if result.Error == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Email is signuped",
		})
		return
	}

	// 密码进行加密
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)

	if err != nil {
		return
	}

	// 将用户信息添加到数据库
	NowUsers := models.MyUser{
		Name:     body.Name,
		Email:    body.Email,
		Password: string(hashedPassword),
	}

	result = initializes.DB.Create(&NowUsers)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create user",
		})
		return
	}

	// 返回用户信息，注册成功
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "success to signuped",
	})

}

func Login(c *gin.Context) {
	// 接受参数
	var body struct {
		Name     string
		Email    string
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to query",
		})
		return
	}

	// 判断邮箱未注册
	var user models.MyUser
	result := initializes.DB.First(&user, "email = ?", body.Email)
	// fmt.Println(user.Email)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Email is not signuped",
		})
		return
	}

	// 查询数据库，找出对应的已加密的密码
	// 解析密码，进行数据库对比，是否正确

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":  http.StatusBadRequest,
			"error": "Invalid email or password",
		})
		return
	}

	// 登陆发放token
	token, _ := pkg.CreateToken(user)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create token",
		})
		return
	}

	// 返回用户信息，登陆成功
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "success to login",
		"token":   token,
	})
}
