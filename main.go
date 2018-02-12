package main

import (
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/napalm684/cleandemo/delivery/http/handlers"
	episodeDomain "github.com/napalm684/cleandemo/domain/episode"
	inmemRepo "github.com/napalm684/cleandemo/infrastructure/repository/inmemory"
	"github.com/napalm684/cleandemo/usecase"
)

func main() {
	srv := setupServer()
	setupGracefulShutdown(srv)
	srv.ListenAndServe()
}

func setupServer() *http.Server {
	srv := &http.Server{Addr: ":8888"}

	setupEpisodeHandlers()

	return srv
}

func setupGracefulShutdown(srv *http.Server) {
	var stopChannel = make(chan os.Signal)
	signal.Notify(stopChannel, syscall.SIGTERM)
	signal.Notify(stopChannel, syscall.SIGINT)
	go func() {
		<-stopChannel

		println("Shutting down server...")
		if err := srv.Shutdown(nil); err != nil {
			panic(err) //Failure/timeout shutting down server gracefully
		}

		os.Exit(0)
	}()
}

func setupEpisodeHandlers() {
	repository := inmemRepo.NewEpisodeRepository(generateInMemoryData())
	episodeService := usecase.NewEpisodeService(repository)
	episodeHandler := handlers.NewEpisodeHandler(*episodeService)
	http.HandleFunc("/episodes/", episodeHandler.GetEpisodeByName)
}

func generateInMemoryData() map[string]*episodeDomain.Episode {
	var m map[string]*episodeDomain.Episode
	m = make(map[string]*episodeDomain.Episode)
	m["Kitten"] = &episodeDomain.Episode{
		Name:    "Kitten",
		Season:  11,
		Episode: 6,
	}

	return m
}
