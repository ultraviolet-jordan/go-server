package cache

import "awesomeProject/io"

type ParamMap map[int32]interface{}

func DecodeParams(dat *io.Packet) ParamMap {
	count := dat.G1()
	params := make(ParamMap, count)

	for range count {
		key := dat.G3()
		if dat.G1() == 1 {
			params[key] = dat.GSTR(10)
		} else {
			params[key] = dat.G4()
		}
	}
	return params
}
