package strategy

type StrategyInterface interface {
}

type StrategyRepository struct {
}

func New() *StrategyRepository {
	return &StrategyRepository{}
}
