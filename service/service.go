package service

import (
	"context"

	"github.com/filmoteka/entities"
	"github.com/filmoteka/usecase"
)

type Service struct {
	UseCase *usecase.UseCase
}

func (s *Service) GetFilmInfo(ctx context.Context, title, order string) ([]entities.Film, error) {
	return s.UseCase.GetFilmInfo(ctx, title, order)
}

func (s *Service) GetActorInfo(ctx context.Context, fullname string) ([]entities.Actor, error) {
	return s.UseCase.GetActorInfo(ctx, fullname)
}

func (s *Service) AddFilm(ctx context.Context, film entities.Film) (int, error) {
	return s.UseCase.AddFilm(ctx, film)
}

func (s *Service) AddActor(ctx context.Context, actor entities.Actor) (int, error) {
	return s.UseCase.AddActor(ctx, actor)
}

func (s *Service) DeleteFilm(ctx context.Context, film entities.Film) bool {
	return s.UseCase.DeleteFilm(ctx, film)
}

func (s *Service) DeleteActor(ctx context.Context, actor entities.Actor) bool {
	return s.UseCase.DeleteActor(ctx, actor)
}

func (s *Service) SetFilmInfo(ctx context.Context, m map[string]interface{}) bool {
	return s.UseCase.SetFilmInfo(ctx, m)
}

func (s *Service) SetActorInfo(ctx context.Context, m map[string]interface{}) bool {
	return s.UseCase.SetActorInfo(ctx, m)
}
