package mongoutil

var (
	UserCollection    = Database.Collection("users")
	CompanyCollection = Database.Collection("companies")
	PanCollection     = Database.Collection("pans")
)
