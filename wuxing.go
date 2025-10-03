package wuxing

type WuXing int8

var WuXingMap = map[string]WuXing{
	"metal": METAL,
	"wood":  WOOD,
	"earth": EARTH,
	"water": WATER,
	"fire":  FIRE,
}

var WuXingHanMap = map[string]WuXing{
	"金": METAL,
	"木": WOOD,
	"土": EARTH,
	"水": WATER,
	"火": FIRE,
	"空": NONE,
	"无": NONE,
}

const (
	NONE  WuXing = 0  // 无/空，非下面标记的数值都为None
	METAL        = 1  // 金
	WOOD         = 2  // 木
	EARTH        = 4  // 土
	WATER        = 8  // 水
	FIRE         = 16 // 火
)

func (l WuXing) String() string {
	switch l {
	case METAL:
		return "金"
	case WOOD:
		return "木"
	case EARTH:
		return "土"
	case WATER:
		return "水"
	case FIRE:
		return "火"
	default:
		return "无"
	}
}

// Restrain 相克
func (l WuXing) Restrain(d WuXing) bool {
	if l == 0 || d == 0 {
		return false
	}
	return (l<<1 == d) || (l == FIRE && d == METAL)
}

// Promote 相生
func (l WuXing) Promote(d WuXing) bool {
	if l == 0 || d == 0 {
		return false
	}
	return l<<3 == d || l>>2 == d
}
