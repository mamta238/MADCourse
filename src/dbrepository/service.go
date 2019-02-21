package dbrepository

import (
	"domain"
	"gopkg.in/mgo.v2/bson"
)

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Get(ID bson.ObjectId) (*domain.Restaurant, error) {
	return s.repo.Get(ID)
}

func (s *Service) GetAll() ([]*domain.Restaurant, error) {
	return s.repo.GetAll()
}


func (s *Service) FindByName(name string) ([] *domain.Restaurant, error){
	return s.repo.FindByName(name)
}


func (s *Service) FindByTypeOfFood(foodType string) ([] *domain.Restaurant,error){
	return s.repo.FindByTypeOfFood(foodType)
}


func (s *Service) FindByTypeOfPostCode(postcode string) ([] *domain.Restaurant,error){
	return s.repo.FindByTypeOfPostCode(postcode)
}


func (s *Service) Store(u *domain.Restaurant) (bson.ObjectId, error) {
	return s.repo.Store(u)
}

func (s *Service) Delete(inp bson.ObjectId) error{
	return s.repo.Delete(bson.ObjectId(inp))
}

func (s *Service) Search(query string) ([] *domain.Restaurant,error){
		return s.repo.Search(query)

}
	

