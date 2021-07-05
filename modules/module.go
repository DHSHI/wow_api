package modules

type ApiForGet struct {
	ID     int    `json:"id" db:"id"`
	Name   string `json:"name" db:"name"`
	NameCn string `json:"nameCn" db:"name_cn"`
	Desc   string `json:"desc" db:"desc"`
}

type Api struct {
	ID         int    `json:"id" db:"id"`
	ParentID   int    `json:"parentID" db:"parent_id"`
	Name       string `json:"name" db:"name"`
	NameCn     string `json:"nameCn" db:"name_cn"`
	Desc       string `json:"desc" db:"desc"`
	Enabled    int    `json:"enabled" db:"enabled"`
	CreateTime string `json:"createTime" db:"create_time"`
	UpdateTime string `json:"updateTime" db:"update_time"`
}

type LoginForGet struct {
	ID        int    `json:"id" db:"id"`
	IP        string `json:"ip" db:"ip"`
	Method    string `json:"method" db:"method"`
	LoginDate string `json:"loginDate" db:"login_date"`
	Count     int    `json:"count" db:"count"`
}

type UnVerifyApi struct {
	ApiID      int    `json:"apiID" db:"api_id"`
	Type       string `json:"type" db:"type"`
	Name       string `json:"name" db:"name"`
	NameCn     string `json:"nameCn" db:"name_cn"`
	Desc       string `json:"desc" db:"desc"`
	InfoDesc   string `json:"infoDesc" db:"info_desc"`
	CreateTime string `json:"createTime" db:"create_time"`
}

type ChartData struct {
	Name string `json:"name"`
	Data []int  `json:"data"`
}
