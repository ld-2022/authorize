package authorize

type Authorize interface {
	// CheckLogin 检查用户是否登录
	CheckLogin(parameter RequestParameter) (bool, error)

	// FindUserProjectTeamList 查询用户项目团队列表
	FindUserProjectTeamList(parameter RequestParameter) ([]ProjectTeam, error)
}
