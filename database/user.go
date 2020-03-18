package database

import (
	"golang.org/x/crypto/bcrypt"
	"good/debug"
)

func makePassword(p string) (string,bool) {
	hash, e := bcrypt.GenerateFromPassword([]byte(p), bcrypt.DefaultCost)
	if e != nil {
		return "",false
	}
	return string(hash),true
}

func checkPassword(p string, encoded string) bool {
	e := bcrypt.CompareHashAndPassword([]byte(encoded), []byte(p))
	return e != nil
}

//定义用户模型
type User struct {
	ModelCom
	Username     string `gorm:"primary_key"`
	Nickname     *string
	Password     string `json:"-"`
	Sex          *int
	Email        *string
	Website      *string
	Education    *string
	Collage      *string
	Introduction *string
	Github       *string
	Avatar       *string
	Coin         int
}

func (User) TableName() string {
	return "users"
}

//用户名注册 insert into user (`username`,`password`) values(xxx,xxx)
func (this *User)CreateByPassword() bool {
	var pw,b = makePassword(this.Password)
	if b {
		this.Password = pw
		return database.Create(this).Error == nil
	}else{
		return false
	}
}

func (this *User) Exist() bool {
	return database.Model(this).Where("username = ?", this.Username).Count(nil).Error == nil
}

//登录  select id from user   check user.ID
func (this *User) Login() bool {
	var find User
	if database.Where(this).Select("id").First(&find); find.ID > 0 {
		return true
	}
	return false
}

//获取用户信息 select * from user where id = ?
func UserGetInfo(id string) (user *User, err error) {
	err = database.Where("id = ?", id).First(&user).Error
	if err != nil {
		debug.LogDebug(err)
	}
	return
}

//更新用户信息 update user set `aa` = xx ,`bb`=xx
func (user *User)Update() bool {
	if err := database.Model(user).Updates(user).Error; err == nil {
		return true
	}
	return false
}


func CheckUserAuth(username, password string) (int,bool) {
	var user User
	var err=database.Select("id, password").Where(User{Username: username}).First(&user).Error
	if err == nil{
		return user.ID, checkPassword(password, user.Password)
	}else{
		return 0, false
	}
}

