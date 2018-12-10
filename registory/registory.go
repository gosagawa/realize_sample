package registry

// Repository domain.repositoryの実装を取得するための構造体
type Repository interface{}

type repositoryImpl struct{}

// NewRepository Repository構造体を生成する
func NewRepository() Repository {
	return &repositoryImpl{}
}
