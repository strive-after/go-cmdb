package module

import (
	"sync"
)

const  (
	userTableName string = "user"
	groupTableName string = "group"
)

var (
	ok bool
	once      sync.Once
	err       error
)


