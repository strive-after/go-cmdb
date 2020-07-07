package module

const  (
	userTableName string = "user"
	groupTableName string = "group"
)

var (
	ok bool
	err       error
)


type Operation interface {
	CRUD
	//返回一个表名  如果做关联查询好用
	TableName() string
	UserShare
}




type CRUD interface {
	Add(mold interface{}) error
	GetId(mold interface{})  error
	Update(mold interface{}) error
	Del(id uint) error
	Get(mold string,value interface{}) error
	GetAll(mold interface{}) (interface{},error)
}


type UserShare interface {
	ComparePass(passwd string) error
	ChangePass(id uint,oldpass ,newpass string) error
}


func NewOperation(mold interface{}) Operation{
	var operation  Operation
	switch mold.(type) {
	case *User:
		operation = mold.(*User)
		return operation
	}
	return nil
}