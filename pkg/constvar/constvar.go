package constvar

var (
	DefaultLimit = 50
)

const (
	// 正常状态
	normalStatus = iota
	// 不可用状态
	disableStatus
	// 删除状态
	delStatus
)

const (
	// 未审核
	UnreviewRoleStatus = iota
	// 管理员
	AdminRoleStatus
	// 游客
	TouristRoleStatus
)
