package user

type Service interface {
}

type service struct {
}

func InitService() Service {
	return &service{}
}
