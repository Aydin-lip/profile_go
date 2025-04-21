package constant

type update struct {
	as string
	to string
}

func Schemas() []string {
	return []string{
		Security(),
	}
}

func UpdateSchema() update {
	new := update{as: "Security", to: "Profile"}
	return new
}

func Security() string {
	return "Security"
}
