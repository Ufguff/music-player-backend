package config

import (
	"fmt"
	"os"
)

type Config struct {
	PublicHost      string
	Port            string
	DBUser          string
	DBPassword      string
	DBAddress       string
	DBName          string
	StaticPathAudio string
	StaticPathImage string
}

var Envs = initConfig()

func initConfig() Config {
	return Config{
		PublicHost:      getEnv("PUBLIC_HOST", "http://localhost"),
		Port:            getEnv("PORT", "8080"),
		DBUser:          getEnv("DB_USER", "root"),
		DBPassword:      getEnv("DB_PASSWORD", "root"),
		DBAddress:       fmt.Sprintf("%s:%s", getEnv("DB_HOST", "127.0.0.1"), getEnv("DB_PORT", "3306")),
		DBName:          getEnv("DB_NAME", "MusicPlayer"),
		StaticPathAudio: getEnv("STATICPATHAUDIO", "static/audio.mp3"),  // maybe other formats??
		StaticPathImage: getEnv("STATICPATHIMAGE", "static/image.jpeg"), // maybe other formats??
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}
