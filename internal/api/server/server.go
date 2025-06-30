package server

import (
	"tugas-cli/internal/api"
	"tugas-cli/internal/config"
	"tugas-cli/internal/utils"
	"tugas-cli/internal/sensor"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

// For your project move this variable to yaml and load through NewConfig
const PORT = ":8080"
const TAG = "SERVER"

type Server struct {
	config *config.Config
	mux    *chi.Mux
	log    *utils.Logger
}

func NewServer(config *config.Config) *Server {
	log := utils.NewLogger(10)
	mux := chi.NewRouter()
	mux.Use(middleware.Logger)
	mux.Use(middleware.Recoverer)
	
	mux.Get("/", func(w http.ResponseWriter, r *http.Request){
		_, err := w.Write([]byte("Endpoint utama"))
		if err != nil {
			fmt.Printf("Error writing response: %v\n", err)
		}
	})
	
	mux.Get("/test", func(w http.ResponseWriter, r *http.Request){
		_, err := w.Write([]byte("Testing"))
		if err != nil {
			fmt.Printf("Error writing response: %v\n", err)
		}
	})

	mux.Get("/weather", func(w http.ResponseWriter, r *http.Request) {
		apiClient := api.NewAPI(config)
		weather, err := apiClient.GetWeather()
		if err != nil {
			http.Error(w, "Gagal ambil cuaca", http.StatusInternalServerError)
			return
		}
		result := fmt.Sprintf("Cuaca di %s: %.1f°C", weather.Location.Name, weather.Current.TempC)
		_, err2 := w.Write([]byte(result))
		if err2 != nil {
			fmt.Printf("Error writing response: %v\n", err2)
		}
	})

	mux.Get("/sys/cpu", func(w http.ResponseWriter, r *http.Request){
		cpuUsage, err := sensor.GetCPUUsage()
		if err != nil{
			http.Error(w, "Gagal baca CPU", http.StatusInternalServerError)
			return
		}
		w.Write([]byte(fmt.Sprintf("CPU Usage: %.2f%%", cpuUsage)))
	})

	mux.Get("/sys/memory", func(w http.ResponseWriter, r *http.Request){
		memUsage, total, err := sensor.GetMemoryUsage()
		if err != nil{
			http.Error(w, "Gagal baca memori", http.StatusInternalServerError)
			return
		}
		w.Write([]byte(fmt.Sprintf("RAM: %.2f%% digunakan dari total %d bytes", memUsage, total)))
	})

	mux.Get("/sys/disk", func(w http.ResponseWriter, r *http.Request){
		diskUsage, total, err := sensor.GetDiskUsage("/")
		if err != nil{
			http.Error(w, "Gagal baca disk", http.StatusInternalServerError)
			return
		}
		w.Write([]byte(fmt.Sprintf("Disk: %.2f%% digunakan dari total %d bytes", diskUsage, total)))
	})

	return &Server{log: log, config: config, mux: mux}
}

func (s *Server) Run() error {

	server := &http.Server{
		Addr:         PORT,
		Handler:      s.mux,
		WriteTimeout: time.Second * 20,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Minute,
	}

	s.log.Add(TAG, fmt.Sprintf("Server running on http://localhost%s", PORT))

	return server.ListenAndServe()
}
