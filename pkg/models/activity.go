package models

// SendActivity ... 发送活动的结构体
/**
	Data 经过 Marshal 的请求数据
	EndInbox 终点的收件箱，发送请求的终点
	Method 请求的方法
	Name 当前用户的用户名
	UserAddress 当前用户的 Activity 地址
	EndActor 终点用户的 Actor 地址
 */
type SendActivity struct {
	Data []byte
	EndInbox string
	Method string
	Name string
	UserAddress string
	EndActor string
}

func NewSendActivity(data []byte, endInbox, method, name, userAddress, endActor string) *SendActivity {
	nsa := &SendActivity{
		Data:     data,
		EndInbox: endInbox,
		Method:   method,
		Name:     name,
		UserAddress:  userAddress,
		EndActor: endActor,
	}
	return nsa
}