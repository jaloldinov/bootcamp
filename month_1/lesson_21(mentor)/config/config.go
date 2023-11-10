package config

type Config struct {
	Limit   int
	Page    int
	Methods []string
	Objects []string
}

const (
	SuccessStatus = iota + 1
	CancelStatus
)

const (
	Fixed = iota + 1
	Percent
)

func Load() *Config {
	return &Config{
		Limit: 10,
		Page:  1,
		Methods: []string{
			"create", "update", "get", "getAll", "update", "delete", "getTopStaff",
		},
		Objects: []string{
			"branch", "staff", "sale", "transaction", "tariff",
		},
	}
}
