package types

type Song struct {
	Id        int
	Name      string
	IdAuthor  int
	PathTrack string
	Image     string
}

type SongForFront struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Author string `json:"author"`
	Image  string `json:"image"`
}

type ISong interface {
	GetInfoSongs() ([]SongForFront, error)
	GetAuthor(int) (string, error)
	GetPathSong(string) (string, error)
	GetPathImage(string) (string, error)
}
