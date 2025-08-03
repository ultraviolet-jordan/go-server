package cache

import (
	"awesomeProject/io"
	"fmt"
	"strings"
)

var (
	objNames   map[string]int32
	objConfigs []*ObjType
)

// ----

type ObjType struct {
	ID        int32    // id
	Model     int32    // model (0)
	Name      string   // name (null)
	Desc      string   // desc (null)
	RecolS    []int32  // recol_s (null)
	RecolD    []int32  // recol_d (null)
	Zoom2D    int32    // zoom2d (2000)
	Xan2D     int32    // xan2d (0)
	Yan2D     int32    // yan2d (0)
	Zan2D     int32    // zan2d (0)
	Xof2D     int32    // xof2d (0)
	Yof2D     int32    // yof2d (0)
	Code9     bool     // code9 (false)
	Code10    int32    // code10 (-1)
	Stackable bool     // stackable (false)
	Cost      int32    // cost (1)
	Members   bool     // members (false)
	OP        []string // op (null)
	IOP       []string // iop (null)

	ManWear          int32 // manwear (-1)
	ManWear2         int32 // manwear2  (-1)
	ManWearOffsetY   int8  // manwearOffsetY (0)
	WomanWear        int32 // womanwear  (-1)
	WomanWear2       int32 // womanwear2  (-1)
	WomanWearOffsetY int8  // womanwearOffsetY (0)
	ManWear3         int32 // manwear3  (-1)
	WomanWear3       int32 // womanwear3 (-1)
	ManHead          int32 // manhead (-1)
	ManHead2         int32 // manhead2 (-1)
	WomanHead        int32 // womanhead (-1)
	WomanHead2       int32 // womanhead2 (-1)

	CountObj     []int32 // countobj (null)
	CountCo      []int32 // countco (null)
	CertLink     int32   // certlink (-1)
	CertTemplate int32   // certtemplate (-1)

	// server-side
	WearPos     int32  // wearpos  (-1)
	WearPos2    int32  //wearpos2 (-1)
	WearPos3    int32  // wearpos3 (-1)
	Weight      int32  // weight (0) (in grams)
	Category    int32  // category (-1)
	DummyItem   int32  // dummyitem (0)
	Tradeable   bool   // tradeable (true)
	RespawnRate int32  // respawnrate (100) (default to 1-minute)
	DebugName   string // debugname (null)

	Params ParamMap // params (null)
}

func newObjType(id int32) *ObjType {
	return &ObjType{
		ID:           id,
		Zoom2D:       2000,
		Code10:       -1,
		Cost:         1,
		ManWear:      -1,
		ManWear2:     -1,
		WomanWear:    -1,
		WomanWear2:   -1,
		ManWear3:     -1,
		WomanWear3:   -1,
		ManHead:      -1,
		ManHead2:     -1,
		WomanHead:    -1,
		WomanHead2:   -1,
		CertLink:     -1,
		CertTemplate: -1,
		WearPos:      -1,
		WearPos2:     -1,
		WearPos3:     -1,
		Category:     -1,
		Tradeable:    true,
		RespawnRate:  100,
	}
}

func LoadObjs(members bool, dir string) {
	server := io.FromIO(dir + "/server/obj.dat")
	jag, err := io.NewJagFile(io.FromIO(dir + "/client/config"))

	if err != nil {
		panic(err)
	}

	count := server.G2()

	objConfigs = make([]*ObjType, count)
	objNames = make(map[string]int32, count)

	client, err := jag.Read("obj.dat")

	if err != nil {
		panic(err)
	}

	client.Pos = 2

	for id := range count {
		config := newObjType(id)
		DecodeType(server, config)
		DecodeType(client, config)

		objConfigs[id] = config

		if len(config.DebugName) > 0 {
			objNames[config.DebugName] = id
		}
	}

	for id := range count {
		config := objConfigs[id]

		if config.CertTemplate != -1 {
			config.cert()
		}

		if config.DummyItem != 0 {
			config.Tradeable = false
		}

		if !members && config.Members {
			config.Tradeable = false
			config.OP = nil
			config.IOP = nil

			if config.Params != nil {
				for range config.Params {
					// TODO: autodisable params
				}
			}
		}
	}
}

func Get(id int32) *ObjType {
	if id < 0 || id >= int32(len(objConfigs)) {
		return nil
	}
	return objConfigs[id]
}

func GetByName(name string) *ObjType {
	id := GetId(name)
	if id == -1 {
		return nil
	}
	return Get(id)
}

func GetId(name string) int32 {
	if id, ok := objNames[name]; ok {
		return id
	}
	return -1
}

func Count() int32 {
	return int32(len(objConfigs))
}

// ----

func (o *ObjType) decode(buf *io.Packet, code int32) {
	switch code {
	case 1:
		o.Model = buf.G2()
	case 2:
		o.Name = buf.GSTR(10)
	case 3:
		o.Desc = buf.GSTR(10)
	case 4:
		o.Zoom2D = buf.G2()
	case 5:
		o.Xan2D = buf.G2()
	case 6:
		o.Yan2D = buf.G2()
	case 7:
		o.Xof2D = buf.G2S()
	case 8:
		o.Yof2D = buf.G2S()
	case 9:
		o.Code9 = true
	case 10:
		o.Code10 = buf.G2()
	case 11:
		o.Stackable = true
	case 12:
		o.Cost = buf.G4()
	case 13:
		o.WearPos = buf.G1()
	case 14:
		o.WearPos2 = buf.G1()
	case 15:
		o.Tradeable = false
	case 16:
		o.Members = true
	case 23:
		o.ManWear = buf.G2()
		o.ManWearOffsetY = buf.G1S()
	case 24:
		o.ManWear2 = buf.G2()
	case 25:
		o.WomanWear = buf.G2()
		o.WomanWearOffsetY = buf.G1S()
	case 26:
		o.WomanWear2 = buf.G2()
	case 27:
		o.WearPos3 = buf.G1()
	case 30, 31, 32, 33, 34:
		if o.OP == nil {
			o.OP = make([]string, 5)
		}
		o.OP[code-30] = buf.GSTR(10)
	case 35, 36, 37, 38, 39:
		if o.IOP == nil {
			o.IOP = make([]string, 5)
		}
		o.IOP[code-35] = buf.GSTR(10)
	case 40:
		count := buf.G1()
		o.RecolS = make([]int32, count)
		o.RecolD = make([]int32, count)
		for i := range count {
			o.RecolS[i] = buf.G2()
			o.RecolD[i] = buf.G2()
		}
	case 75:
		o.Weight = buf.G2S()
	case 78:
		o.ManWear3 = buf.G2()
	case 79:
		o.WomanWear3 = buf.G2()
	case 90:
		o.ManHead = buf.G2()
	case 91:
		o.WomanHead = buf.G2()
	case 92:
		o.ManHead2 = buf.G2()
	case 93:
		o.WomanHead2 = buf.G2()
	case 94:
		o.Category = buf.G2()
	case 95:
		o.Zan2D = buf.G2()
	case 96:
		o.DummyItem = buf.G1()
	case 97:
		o.CertLink = buf.G2()
	case 98:
		o.CertTemplate = buf.G2()
	case 100, 101, 102, 103, 104, 105, 106, 107, 108, 109:
		if o.CountObj == nil || o.CountCo == nil {
			o.CountObj = make([]int32, 10)
			o.CountCo = make([]int32, 10)
		}
		o.CountObj[code-100] = buf.G2()
		o.CountCo[code-100] = buf.G2()
	case 201:
		o.RespawnRate = buf.G2()
	case 249:
		o.Params = DecodeParams(buf)
	case 250:
		o.DebugName = buf.GSTR(10)
	default:
		panic(fmt.Sprintf("Unrecognized obj config code: %d", code))
	}
}

func (o *ObjType) cert() {
	template := Get(o.CertTemplate)
	link := Get(o.CertLink)
	if template == nil || link == nil {
		return
	}

	o.Model = template.Model
	o.Zoom2D = template.Zoom2D
	o.Xan2D = template.Xan2D
	o.Yan2D = template.Yan2D
	o.Zan2D = template.Zan2D
	o.Xof2D = template.Xof2D
	o.Yof2D = template.Yof2D
	o.RecolS = template.RecolS
	o.RecolD = template.RecolD

	o.Name = link.Name
	o.Members = link.Members
	o.Cost = link.Cost
	o.Tradeable = link.Tradeable

	article := "a"
	first := strings.ToLower(link.Name)
	if len(first) > 0 {
		switch first[0] {
		case 'a', 'e', 'i', 'o', 'u':
			article = "an"
		}
	}
	o.Desc = fmt.Sprintf("Swap this note at any bank for %s %s.", article, link.Name)
	o.Stackable = true
}
