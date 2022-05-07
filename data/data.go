package data

const (
	Hour      = "0"
	Day       = "1"
	Week      = "2"
	HalfMonth = "3"
	Month     = "4"
)

type TmpCode struct {
	Code string `json:"code"`
	Type string `json:"type"`
}
