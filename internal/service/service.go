package service

type TeamService interface {
}

type JobCategory interface {
}

type AuthService interface {
}

type UserService interface {
}

type VacancyService interface {
}

type VacancyDetailService interface {
}

type UserFileService interface {
}

type Service struct {
	TeamService
	JobCategory
	AuthService
	UserService
	VacancyService
	VacancyDetailService
	UserFileService
}
