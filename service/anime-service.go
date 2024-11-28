package service

import "gilab.com/pragmaticreviews/golang-gin-poc/entity"

type AnimeService interface {
	Save(entity.Anime) entity.Anime
	FindAll() []entity.Anime
}

type animeService struct {
	animes []entity.Anime
}

func New() AnimeService {
	return &animeService{}
}

func (service *animeService) Save(entity.Anime) entity.Anime{
	
}

FindAll() []entity.Anime