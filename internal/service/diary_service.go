package service

import (
	"context"
	model "moviediary/internal/model"
	model_dto "moviediary/internal/model/dto"
	provider "moviediary/internal/provider/tmdb"
	repository "moviediary/internal/repository"
	apperror "moviediary/pkg/apperror"
	"time"
)

type DiaryService struct {
	diaryRepository *repository.DiaryRepository
	movieRepository *repository.MovieRepository
	tmdbClient      *provider.Client
}

func NewDiaryService(
	diaryRepository *repository.DiaryRepository,
	movieRepository *repository.MovieRepository,
	tmdbClient *provider.Client,
) *DiaryService {
	return &DiaryService{
		diaryRepository: diaryRepository,
		movieRepository: movieRepository,
		tmdbClient:      tmdbClient,
	}
}

func (s *DiaryService) AddDiary(
	ctx context.Context,
	userID uint,
	tmdbID int,
	comment string,
	rating float64,
	watchedAt time.Time,
) (*model_dto.AddDiaryResponse, error) {

	movieInDB, err := s.movieRepository.GetByTmdbID(ctx, tmdbID)
	if err != nil {
		return nil, err
	}

	if movieInDB == nil {
		tmdbMovie, err := s.tmdbClient.GetMovieDetail(ctx, tmdbID)
		if err != nil {
			return nil, err
		}
		if tmdbMovie == nil {
			return nil, apperror.ErrMovieNotFound
		}

		movieInDB, err = s.movieRepository.Create(ctx, &model.Movie{
			TmdbID:    tmdbID,
			Title:     tmdbMovie.Title,
			Overview:  tmdbMovie.Overview,
			PosterURL: tmdbMovie.PosterPath,
		})
		if err != nil {
			return nil, err
		}
	}

	existingDiary, err := s.diaryRepository.GetByUserIDAndMovieID(ctx, userID, int(movieInDB.ID))
	if err != nil {
		return nil, err
	}
	if existingDiary != nil {
		return nil, apperror.ErrDiaryAlreadyExists
	}

	diary, err := s.diaryRepository.AddDiary(
		ctx,
		userID,
		int(movieInDB.ID),
		comment,
		rating,
		watchedAt,
	)
	if err != nil {
		return nil, err
	}

	return &model_dto.AddDiaryResponse{
		Diary:   diary,
		Message: "Diary created successfully",
	}, nil
}

func (s *DiaryService) RemoveDiary(ctx context.Context, userID uint, movieId int) error {
	diary, err := s.diaryRepository.GetByUserIDAndMovieID(ctx, userID, movieId)
	if err != nil {
		return err
	}
	if diary == nil {
		return apperror.ErrDiaryNotFound
	}
	return s.diaryRepository.RemoveDiary(ctx, userID, movieId)
}

func (s *DiaryService) GetDiaryList(ctx context.Context, userID uint, page int, pageSize int) (*model_dto.GetDiaryListResponse, error) {
	diaryList, err := s.diaryRepository.GetByUserID(ctx, userID, page, pageSize)
	if err != nil {
		return nil, err
	}
	totalPages := len(diaryList) / pageSize
	totalItems := len(diaryList)
	items := make([]*model.UserMovie, len(diaryList))
	for i, diary := range diaryList {
		items[i] = &model.UserMovie{
			MovieID:    diary.MovieID,
			Rating:     diary.Rating,
			Comment:    diary.Comment,
			WatchedAt:  diary.WatchedAt,
			IsWatched:  diary.IsWatched,
			IsFavorite: diary.IsFavorite,
		}
	}
	return &model_dto.GetDiaryListResponse{
		Page:       page,
		Items:      items,
		TotalPages: totalPages,
		TotalItems: totalItems,
		Message:    "Diary listed successfully",
	}, nil
}
