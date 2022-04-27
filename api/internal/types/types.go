// Code generated by goctl. DO NOT EDIT.
package types

type ReqLogin struct {
	ValiCode string `json:"vali_code"`
}

type RespLogin struct {
	ID   string `json:"_id"`
	Name string `json:"name"`
	JwtToken
}

type JwtToken struct {
	AccessToken  string `json:"access_token,omitempty"`
	AccessExpire int64  `json:"access_expire,omitempty"`
	RefreshAfter int64  `json:"refresh_after,omitempty"`
}

type ReqUserId struct {
	ID string `path:"_id"`
}

type RespUser struct {
	ID        string   `json:"_id"`
	Avatar    string   `json:"avatar"`
	Birthday  int64    `json:"birthday"`
	Blog      string   `json:"blog"`
	City      string   `json:"city"`
	Email     string   `json:"email"`
	Followers int64    `json:"followers"`
	Following int64    `json:"following"`
	Name      string   `json:"name"`
	Phone     int64    `json:"phone"`
	RealName  string   `json:"real_name"`
	Skills    []string `json:"skills"`
	Summary   string   `json:"summary"`
}

type ReqUser struct {
	ID       string   `path:"_id"`
	Avatar   *string  `json:"avatar,optional"`
	Birthday *int64   `json:"birthday,optional"`
	City     *string  `json:"city,optional"`
	Email    *string  `json:"email,optional"`
	Name     *string  `json:"name,optional"`
	Phone    *int64   `json:"phone,optional"`
	RealName *string  `json:"real_name,optional"`
	Summary  *string  `json:"summary,optional"`
	Skills   []string `json:"skills,optional"`
	Blog     *string  `json:"blog,optional"`
}

type InterviewHardStatus struct {
	Difficulty string `json:"difficulty"`
	Count      int64  `json:"count"`
	Total      int64  `json:"total"`
}

type ReqUsers struct {
	ID       string `path:"_id"`
	Search   string `form:"search,optional"`
	PageNo   int    `form:"page_no"`
	PageSize int    `form:"page_size"`
}

type OtherUser struct {
	ID      string `json:"_id"`
	Avatar  string `json:"avatar"`
	Name    string `json:"name"`
	Summary string `json:"summary"`
}

type ReqStarInterviews struct {
	ID       string   `path:"_id"`
	Tags     []string `form:"tags,optional"`
	Search   string   `form:"search,optional"`
	PageNo   int      `form:"page_no"`
	PageSize int      `form:"page_size"`
}

type RespStarInterviews struct {
	ID          string `json:"_id"`
	Title       string `json:"title"`
	UpdatedTime int64  `json:"updated_time"`
}

type ReqStarInterviewsID struct {
	ID string `path:"_id"`
}

type ReqInterviewId struct {
	ID string `path:"_id"`
}

type Author struct {
	ID   string `json:"_id"`
	Name string `json:"name"`
}

type Interview struct {
	ID          string   `json:"_id"`
	Author      Author   `json:"author"`
	ClickNum    int64    `json:"click_num,default=0"`
	Good        int64    `json:"good,default=0"`
	HardStatus  string   `json:"hard_status,options=hard,medium,easy"`
	HotNum      int64    `json:"hot_num,default=0"`
	Summary     string   `json:"summary"`
	Tags        []string `json:"tags"`
	Title       string   `json:"title"`
	UpdatedTime int64    `json:"updated_time"`
	Status      bool     `json:"status,default=false"`
}

type Interview_detail struct {
	Interview
	StarNum     int64  `json:"star_num,default=0"`
	Bad         int64  `json:"bad,default=0"`
	Content     string `json:"content"`
	CreatedTime int64  `json:"created_time,omitempty"`
}

type ReqInterviewUpdate struct {
	ID         string   `path:"_id"`
	HardStatus *string  `json:"hard_status,optional,options=hard,medium,easy"`
	Summary    *string  `json:"summary,optional"`
	Tags       []string `json:"tags,optional"`
	Title      *string  `json:"title,optional"`
	Content    *string  `json:"content,optional"`
}

type ReqInterviewAdd struct {
	HardStatus string   `json:"hard_status,options=hard,medium,easy"`
	Summary    string   `json:"summary,omitempty"`
	Tags       []string `json:"tags"`
	Title      string   `json:"title"`
	Content    string   `json:"content"`
}

type ReqInterviews struct {
	Tags   []string `form:"tags,optional"`
	Search string   `form:"search,optional"`
	CommonPage
}

type CommInterviewsResp struct {
	Ok    bool   `json:"ok"`
	Error string `json:"error"`
}

type CommonPage struct {
	PageNo   int `form:"page_no"`
	PageSize int `form:"page_size"`
}

type RespInterviewsTags struct {
	ID      string   `json:"_id"`
	Name    string   `json:"name"`
	SubTags []string `json:"sub_tags"`
}

type MailNotice struct {
	Comments    bool `json:"comments"`
	Follow      bool `json:"follow"`
	GoodAndStar bool `json:"good_and_star"`
}

type MessageConfig struct {
	Comments    bool       `json:"comments"`
	Follow      bool       `json:"follow"`
	GoodAndStar bool       `json:"good_and_star"`
	MailNotice  MailNotice `json:"mail_notice"`
}
