package constant

const (
	Dashboard  = 1
	MobileApp  = 2
	PartnerApp = 3
)

var ConvertRoleID = map[string]uint64{
	"dashboard":   Dashboard,
	"mobile-app":  MobileApp,
	"partner-app": PartnerApp,
}
