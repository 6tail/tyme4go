package tyme

// TerrainNames 地势(长生十二神)名称
var TerrainNames = []string{"长生", "沐浴", "冠带", "临官", "帝旺", "衰", "病", "死", "墓", "绝", "胎", "养"}

// Terrain 地势(长生十二神)
type Terrain struct {
	LoopTyme
}

func (Terrain) FromIndex(index int) Terrain {
	return Terrain{LoopTyme{}.FromIndex(TerrainNames, index)}
}

func (Terrain) FromName(name string) (*Terrain, error) {
	p, err := LoopTyme{}.FromName(TerrainNames, name)
	if err != nil {
		return nil, err
	}
	return &Terrain{*p}, nil
}

func (o Terrain) Next(n int) Terrain {
	return o.FromIndex(o.nextIndex(n))
}
