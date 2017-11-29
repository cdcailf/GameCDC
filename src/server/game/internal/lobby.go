package internal

import (
	"fmt"
	// "strconv"
	"time"
)

// 游戏大厅的数据结构
type GameLobby struct {
	lobbyConf                         //大厅相关基本参数的结构体
	AnnouncementInfo                  //大厅公告结构体
	LoginUserArray   []*LoginUserInfo //登录用户映射的array，对应login_user_info_map实例中的地址
}

// 大厅相关的基本参数信息
type lobbyConf struct {
	MaxConnectNum int       //最大连接数
	CurConnectNum int       //当前连接数
	StartTime     time.Time //大厅启动时间
	Version       string    //大厅版本当前版本（或许需要加上版本信息对应的描述）
}

// 大厅中公告的结构定义（公告，保留）
type AnnouncementInfo struct {
	Theme      string    //公告主题
	Content    string    //公告内容（可以是一段html页面）
	CreateTime time.Time //公告发布时间
}

// 大厅中当前登录用户的映射表
type LoginUserInfo struct {
	UserId       string        //用户id
	UserName     string        //用户名
	LoginTime    string        //用户登录时间
	OnlineStatus bool          //用户在线状况（需要心跳维持与检测）
	GameData     *UserGameData //用户实例的地址（指向的事UserGameData）
}

// 登录的用户的具体游戏数据(需要和得才对接)
type UserGameData struct {
}

func CreateLobby() *GameLobby {
	lobby := newLobby()
	fmt.Println(lobby)
	return lobby
}

func newLobby() *GameLobby {
	current := time.Now()
	gamelobby := GameLobby{}
	gamelobby.MaxConnectNum = 100
	gamelobby.CurConnectNum = 1
	gamelobby.StartTime = current
	gamelobby.Version = "v1.0"
	gamelobby.Theme = "welcome come here...this is 公告"
	gamelobby.Content = "公告内容，游戏世界"
	gamelobby.CreateTime = current
	gamelobby.LoginUserArray = make([]*LoginUserInfo, 5)

	return &gamelobby
}

func InitGameLobby() {
	fmt.Println("this is game lobby init....")
	CreateLobby()
	//test

}
