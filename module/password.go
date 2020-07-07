package module

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func ChangePasswordHash(oldpass ,newpass string,user *User) error {
	operation :=  NewOperation(&User{})
	err := operation.Get("email",user)
	if err != nil {
		return fmt.Errorf("获取用户失败 %v\n",err)
	}
	if oldpass == "" {
		user.Password ,err = HashPassword(newpass)
		if err != nil {
			return fmt.Errorf("获取新密码失败 %v\n",err)
		}
		err = operation.UpdateMold(user)
		if err != nil {
			return fmt.Errorf("更新失败 %v\n",err)
		}
	}else {
		err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(oldpass))
		if err != nil {
			return fmt.Errorf("密码错误 %v\n",err)
		}
		user.Password ,err = HashPassword(newpass)
		if err != nil {
			return fmt.Errorf("获取新密码失败 %v\n",err)
		}
		err = operation.UpdateMold(user)
		if err != nil {
			return fmt.Errorf("更新失败 %v\n",err)
		}
	}
	return nil
}


func CheckPassword(password string,user *User) error{
	operation :=  NewOperation(&User{})
	err := operation.Get("email",user)
	if err != nil {
		return  fmt.Errorf("获取用户失败 %v\n",err)
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return fmt.Errorf("密码错误 %v\n",err)
	}
	return nil
}



func GetHashingCost(hashedPassword []byte) int {
	cost, _ := bcrypt.Cost(hashedPassword) // 为了简单忽略错误处理
	return cost
}
