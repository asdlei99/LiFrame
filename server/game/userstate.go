package game

import "github.com/llr104/LiFrame/core/liFace"

type GUserState int
const (
	GUserStateOnline = iota
	GUserStateOffLine
	GUserStateLeave
)

type userState struct {
	UserId  uint32
	State   int
	SceneId int
	Conn    liFace.IConnection
}
