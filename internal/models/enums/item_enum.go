package enums

type ItemType string
type ItemStatus string

const (
	ADMISSION  ItemType = "ADMISSION"
	SUBMISSION ItemType = "SUBMISSION"
	REVERSAL   ItemType = "REVERSAL"
)

const (
	ACCEPTED ItemStatus = "ACCEPTED"
	DECLINED ItemStatus = "DECLINED"
)
