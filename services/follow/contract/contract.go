package contract

type Following struct {
	Userids []string
}

type FollowReq struct {
	Userid    string
	Following string
}

type GetFollowersReq struct {
	Userid string
}
