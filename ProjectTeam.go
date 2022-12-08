package authorize

type ProjectTeam struct {
	// 项目组id
	TeamId int `json:"teamId"`
	// 项目组名称
	TeamName string `json:"teamName"`
	// 项目组类型 0:通用 1:定制
	TeamType int `json:"teamType"`
	// 项目列表
	ProjectList []Project `json:"projectList"`
}

type Project struct {
	// 项目id
	ProjectId int `json:"projectId"`
	// 项目名称
	ProjectName string `json:"projectName"`
	// 项目git地址
	ProjectGit string `json:"projectGit"`
	// 项目唯一值
	ProjectUUID string `json:"projectUUID"`
}
