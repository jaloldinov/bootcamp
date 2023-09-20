package config

const (
	FixDiscountType     = "FIX"
	PercentDiscountType = "PERCENT"
)

var OrderStatus = map[string]string{
	"0": "in_process",
	"1": "success",
}

type Config struct {
	DefaultOffset int
	DefaultLimit  int

	Path             string
	UserFileName     string
	CategoryFileName string
	ProductFileName  string
	OrderFileName    string
	BranchFileName   string
}

func Load() Config {

	cfg := Config{}

	cfg.DefaultOffset = 0
	cfg.DefaultLimit = 10

	cfg.Path = "./data"
	cfg.UserFileName = "/user.json"
	cfg.CategoryFileName = "/category.json"
	cfg.ProductFileName = "/product.json"
	cfg.OrderFileName = "/order.json"
	cfg.BranchFileName = "/branch.json"

	return cfg
}
