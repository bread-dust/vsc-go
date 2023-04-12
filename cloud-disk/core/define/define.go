package define

import "github.com/dgrijalva/jwt-go"

// jwt token
type UserCLaim struct {
	Id       uint64
	Identity string
	Name     string
	jwt.StandardClaims
}


const (
	// jwt 秘钥
	JwtKey = "cloud-disk-key"

	// 邮箱验证token
	MailPassToken = "NELNPMNEBPAETNTZ"

	// jwt token 过期时间(s)
	JwtTokenExpired = 3600 * 24 

	// jwt 刷新token 过期时间(s)
	JwtRefreshTokenExpired = 3600 * 24 * 7

	// 验证码长度
	CodeLenth = 6
	

	// 验证码过期时间(s)
	CodeExpire=3600

	// tencent cos appid
	TencentSecretID = "AKIDjNnzwfAQSFn6KSPBk8KK6f097uUr1yjI"

	// tencent cos 秘钥
	TencentSecretKey = "TkpiNq8D8tVv98wYyFSTIsDOL7h5HH4d"

	// tencent cos url
	TencentUrl = "https://cloud-disk-1310840816.cos.ap-nanjing.myqcloud.com"

	// 分页默认参数
	PageSize =20

	// datetime 格式化
	DateTime = "2006-01-02 15:04:05"
	
	// 分片大小	
	ChunkSize = 1024 * 1024 * 5 // 5M

)
