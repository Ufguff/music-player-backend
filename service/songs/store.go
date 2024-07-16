package songs

import (
	"database/sql"
	"fmt"
	"strconv"

	"github.com/ufguff/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

//TODO сделать также отправку id

func (s *Store) GetPathImage(idStr string) (string, error) {
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return "", err
	}

	rows, err := s.db.Query("SELECT imagePath FROM tracks WHERE id = ?;", id)
	if err != nil {
		return "", nil
	}

	var path string
	for rows.Next() {
		err := rows.Scan(&path)

		if err != nil {
			return "", nil
		}
	}

	if path != "" {
		return path, nil
	} else {
		return "", fmt.Errorf("Not found that image!")
	}
}

func (s *Store) GetPathSong(idStr string) (string, error) {
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return "", err
	}

	rows, err := s.db.Query("SELECT recordPath FROM tracks WHERE id = ?;", id)
	if err != nil {
		return "", nil
	}

	var path string
	for rows.Next() {
		err := rows.Scan(&path)

		if err != nil {
			return "", nil
		}
	}

	if path != "" {
		return path, nil
	} else {
		return "", fmt.Errorf("Not found that song!")
	}

}

func (s *Store) GetInfoSongs() ([]types.SongForFront, error) {
	rows, err := s.db.Query("SELECT * FROM tracks;")

	if err != nil {
		return nil, err
	}

	tracks := make([]types.Song, 0)

	for rows.Next() {
		s, err := scanRowsIntoSong(rows)
		if err != nil {
			return nil, err
		}
		tracks = append(tracks, *s)
	}

	tracksForFront := make([]types.SongForFront, 0)
	for _, song := range tracks {
		author, err := s.GetAuthor(song.IdAuthor)
		if err != nil {
			return nil, err
		}

		songForFront := types.SongForFront{
			Id:     song.Id,
			Name:   song.Name,
			Author: author,
			Image:  song.Image,
		}
		tracksForFront = append(tracksForFront, songForFront)
	}

	return tracksForFront, nil
}

func (s *Store) GetAuthor(id int) (string, error) {
	rows, err := s.db.Query("SELECT name FROM authors WHERE id= ?;", id)

	if err != nil {
		return "", err
	}

	var author string
	for rows.Next() {
		err := rows.Scan(&author)

		if err != nil {
			return "", err
		}
	}

	return author, nil
}

func scanRowsIntoSong(rows *sql.Rows) (*types.Song, error) {
	song := new(types.Song)

	err := rows.Scan(
		&song.Id,
		&song.Name,
		&song.IdAuthor,
		&song.PathTrack,
		&song.Image,
	)

	if err != nil {
		return nil, err
	}
	return song, nil
}
