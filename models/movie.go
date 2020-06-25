package models

import (
	"errors"
)

func GetMovie(movieId string) (m *Movie, err error) {
	m, err = getNodeMovie(movieId)
	if err == nil {
		return m, nil
	}

	return nil, errors.New("Movie not exists")
}

func GetAllMovies() map[string]*Movie {
	return getAllNodeMovies()
}

func AddMovie(reqmovie *ReqMovie) string {
	_, err := getNodeMovie(reqmovie.Id)

	if err != nil {
		return "MovieId exists"
	}

	err = addNodeMovie(reqmovie)
	if err != nil {
		return reqmovie.Id
	}

	return reqmovie.Id
}

func UpdateMovie(movieId string, m *Movie) (movie *Movie, err error) {
	status, e := editNodeMovie(movieId, m)

	if e != nil {
		return nil, e
	}

	if status == "true" {
		return m, nil
	}

	return nil, errors.New("Movie not exists")
}

func DeleteMovie(movieId string) string {
	_, err := getNodeMovie(movieId)

	if err != nil {
		return "MovieId not exists"
	}

	err = deleteNodeMovie(movieId)
	//fmt.Println(err, "DeleteMovie")
	return "Delete movie by movieId = " + movieId + " successed"
}
