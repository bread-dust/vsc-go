package helper

import (
	"bytes"
	"context"
	"core/define"
	"errors"
	"io"
	"math"
	"net/http"
	"net/smtp"
	"net/url"
	"os"
	"path"
	"strconv"
	"strings"
	"time"

	"crypto/md5"
	"crypto/tls"
	"fmt"
	"math/rand"

	"github.com/dgrijalva/jwt-go"
	"github.com/jordan-wright/email"
	"github.com/satori/go.uuid"
	"github.com/tencentyun/cos-go-sdk-v5"
)

// Md5 加密
func Md5(s string) string {
	return fmt.Sprintf("%x",md5.Sum([]byte(s)))
}

// GenerateToken jwtToken生成
func GenerateToken(id uint64,identity,name string,seconds int64)(string,error){
	// id 
	// identity
	// name
	uc := define.UserCLaim{
		Id:             id,
		Identity:       identity,
		Name:           name,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Second*time.Duration(seconds)).Unix(), 
		 },
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,uc)
	tokenString,err:=token.SignedString([]byte(define.JwtKey))
	if err!=nil{
		return "",err
	}
	return tokenString,nil

}

// MailSendCode 邮箱验证码发送
func MailSendCode(mail,code string)error{
	e:=email.NewEmail()
	e.From = "Get <dust82d@163.com>" //发送者邮箱
	e.To = []string{"dust82d@163.com"} // 接受者邮箱
	e.Subject = "验证码发送" // 主题
	e.HTML = []byte("<h1>"+code+"</h1>")
	err:= e.SendWithTLS("smtp.163.com:465",smtp.PlainAuth("","dust82d@163.com",define.MailPassToken,"smtp.163.com"),&tls.Config{InsecureSkipVerify: true,ServerName: "smtp.163.com"})
	if err!=nil{
		return err
	}	
	return nil
}

// RandCode 随机验证码
func RandCode()string{
	s:= "1234567890"
	code := ""
	// 随机数种子
	rand.NewSource(time.Now().UnixNano())
	// rand.Seed(time.Now().UnixNano())
	for i:=0;i<define.CodeLenth;i++{
		code += string(s[rand.Intn(len(s))])
	}
	return code
}

// 生成uuid
func GetUUID()string{
	return uuid.NewV4().String()
}

// CosUpload 上传文件到腾讯云
func CosUpload(r *http.Request)(string,error){
	u, _ := url.Parse(define.TencentUrl)
	b := &cos.BaseURL{BucketURL: u}
	c := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID: define.TencentSecretID,  // 替换为
			SecretKey: define.TencentSecretKey, // 替换
	},
	})

	file,fileHeader,_:= r.FormFile("file") //获取上传的文件对象
	name := "cloud-disk/"+GetUUID()+path.Ext(fileHeader.Filename) 


	_, err := c.Object.Put(context.Background(), name, file, nil)
	if err != nil {
		panic(err)
	}
	return define.TencentUrl + "/" + name,nil
}


	// AnalyzeToken 解析token
func AnalyzeToken(token string)(*define.UserCLaim,error){
	uc := new(define.UserCLaim)
	claims,err:=jwt.ParseWithClaims(token,uc,func(t *jwt.Token) (interface{}, error) {
		return []byte(define.JwtKey),nil
	})
	if err!=nil{
		return nil,err
	}
	if !claims.Valid{
		return uc,errors.New("token is invalid")
	}
	return uc,nil
}

// 文件分片 os.File.Read() -> []byte -> os.File.Write()
func FileChunk(file *os.File,chunkSize int64)(error){
	fileInfo,err:=os.Stat(file.Name()) // 获取文件信息
	if err!=nil{
		return err
	}

	chunkNum := math.Ceil(float64(fileInfo.Size())/float64(chunkSize)) // 分片数量

	myFile,err:=os.OpenFile(file.Name(),os.O_RDONLY,os.ModePerm) // 读取文件
	if err!=nil{
		return err
	}
	defer myFile.Close() // 关闭文件
	

	chunkArr := make([]byte,chunkSize) // 分片文件数组

	for i := 0; i < int(chunkNum); i++ {
		myFile.Seek(int64(i*int(chunkSize)),0) // 文件指针移动到指定位置

		if chunkSize > fileInfo.Size() - int64(i*int(chunkSize)){ // 判断是否是最后一片
			chunkArr = make([]byte,fileInfo.Size() - int64(i*int(chunkSize))) // 最后一片大小
		}

		myFile.Read(chunkArr) // 读取文件

		// 创建分片文件
		f,err:= os.OpenFile("../../static/chunk/"+strconv.Itoa(i)+".chunk",os.O_CREATE|os.O_WRONLY,os.ModePerm) 
		
		if err!=nil{
			return err
		}

		f.Write(chunkArr) // 写入分片文件
		f.Close() // 关闭文件
	}

	return nil // 返回

}

// 分片文件的合并 os.File.ReadALl() -> []byte -> os.File.Write()
func FileMerge(oldFile *os.File,fileName string,chunkSize int64)(error){
	myFile,err:=os.OpenFile(fileName,os.O_CREATE|os.O_WRONLY|os.O_APPEND,os.ModePerm) // 创建文件
	if err!=nil{
		return err
	}
	defer myFile.Close() // 关闭文件


	fileInfo,err:=os.Stat(oldFile.Name()) // 获取文件信息
	if err!=nil{
		return err
	}
	chunkNum := math.Ceil(float64(fileInfo.Size())/float64(chunkSize)) // 分片数量

	for i := 0; i < int(chunkNum); i++ {
		f,err:= os.OpenFile("../../static/chunk/"+strconv.Itoa(i)+".chunk",os.O_RDONLY,os.ModePerm) // 打开分片文件
		if err!=nil{
			return err
		}
		b,err:=io.ReadAll(f) // 读取分片文件
		if err!=nil{
			return err
		}

		myFile.Write(b) // 写入文件
		f.Close() // 关闭文件
	}
	return nil
}

// 文件一致性校验
func FileCheck(fileName1 string,fileName2 string,chunkSize int64)(error){
	// 获取一个文件的信息
	file1,err:=os.OpenFile(fileName1,os.O_RDONLY,os.ModePerm)
	if err!=nil{
		return err
	}
	b1,err:=io.ReadAll(file1) // 读取文件
	if err!=nil{
		return err
	}


	// 获取另一个文件的信息
	fiel2 ,err := os.OpenFile(fileName2,os.O_RDONLY,os.ModePerm)
	if err!=nil{
		return err
	}
	b2,err:=io.ReadAll(fiel2) // 读取文件
	if err!=nil{
		return err
	}

	// 判断两个文件的校验和是否一致
	s1 := fmt.Sprintf("%x",md5.Sum(b1)) // 获取文件1的校验和
	s2 := fmt.Sprintf("%x",md5.Sum(b2)) // 获取文件2的校验和
	if s1 != s2{
		return errors.New("文件不一致")
	}

	return errors.New("文件一致")
}

// CosInitPart 分片上传初始化
func CosInitPart(ext string)(string,string,error){
	u, _ := url.Parse(define.TencentUrl)
	b := &cos.BaseURL{BucketURL: u}
	c := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID: define.TencentSecretID,  // 替换为
			SecretKey: define.TencentSecretKey, // 替换
		},
	})
   name := "cloud-disk/"+GetUUID()+ext
    // 可选 opt,如果不是必要操作，建议上传文件时不要给单个文件设置权限，避免达到限制。若不设置默认继承桶的权限。
    v, _, err := c.Object.InitiateMultipartUpload(context.Background(), name, nil)
    if err != nil {
        panic(err)
    }
    UploadID := v.UploadID
    fmt.Println(UploadID)
	return name,UploadID,nil
}

// 分片上传
func CosPartUpload(r *http.Request)(string,error){
	u, _ := url.Parse(define.TencentUrl)
	b := &cos.BaseURL{BucketURL: u}
	c := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID: define.TencentSecretID,  
			SecretKey: define.TencentSecretKey, 
		},
	})
   name := r.Header.Get("name")
    UploadID :=r.Header.Get("upload_id")
	partNumber,err:=strconv.Atoi(r.PostForm.Get("part_number"))
	if err!=nil{
		return "",err
	}
    fmt.Println(UploadID)
	f,_,err:=r.FormFile("file")
	if err!=nil{
		return "",err
	}

	buf := bytes.NewBuffer(nil)
	io.Copy(buf,f)

    // opt 可选
    resp, err := c.Object.UploadPart(
        context.Background(), name, UploadID, partNumber,bytes.NewReader(buf.Bytes()), nil,
    )
    if err != nil {
        panic(err)
    }

    PartETag := resp.Header.Get("ETag")
    fmt.Println(PartETag)
	return strings.Trim(resp.Header.Get("ETag"),"\""),nil
}

// 分片上传完成
func CosPartUploadComplete(name,uploadID string,cs []cos.Object)error{
	u, _ := url.Parse(define.TencentUrl)
	b := &cos.BaseURL{BucketURL: u}
	c := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID: define.TencentSecretID,  
			SecretKey: define.TencentSecretKey, 
		},
	})
  
    // opt可选
    opt := &cos.CompleteMultipartUploadOptions{}
    opt.Parts = append(opt.Parts, cs...)
    _, _, err := c.Object.CompleteMultipartUpload(
        context.Background(), name, uploadID, opt,
    )
    if err != nil {
        panic(err)
    }
	return nil
}