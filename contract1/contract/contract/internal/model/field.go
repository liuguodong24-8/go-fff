package model

// Type nft类型
type Type int

const (
	// FFF ...
	FFF Type = 1
	// ETH ...
	ETH Type = 2
	// BSC ...
	BSC Type = 3
)

func (t Type) String() string {
	switch t {
	case FFF:
		return "FFF"
	case ETH:
		return "ETH"
	case BSC:
		return "BSC"
	default:
		return "UNKNOWN"
	}
}
