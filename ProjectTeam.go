package authorize

type ProjectTeam struct {
	TeamId      int       `json:"teamId"`
	TeamName    string    `json:"teamName"`
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
