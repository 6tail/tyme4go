package tyme

// IChildLimitProvider 童限计算接口
type IChildLimitProvider interface {

	// GetInfo 根据出生公历时刻和节令计算童限信息
	GetInfo(birthTime SolarTime, term SolarTerm) ChildLimitInfo
}
