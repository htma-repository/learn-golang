package dependency_injection

type Smartphone struct {
	Name string
}

type Samsung struct {
	*Smartphone
}

func NewSamsung() *Samsung {
	return &Samsung{
		Smartphone: &Smartphone{
			Name: "Galaxy S24",
		},
	}
}
