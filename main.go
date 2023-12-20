package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {

	// Función que traemos de importar el paquete godotenv para que cargue automáticamente las variables en el archivo .env sin nececidad de exportarlas cada vez que las usamos
	godotenv.Load(".env")

	// os.Getenv nos permite traer la información que tenemos en el archivo .env
	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("port is not found in the enviroment")
	}

	// Enrutador donde manejamos todos los endpoints
	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	v1Router := chi.NewRouter()

	v1Router.Get("/healthz", handlerReadiness)
	v1Router.Get("/err", handlerErr)
	router.Mount("/v1", v1Router)

	// Creamos un servidor
	srv := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}

	log.Printf("Server starting on port %v", portString)
	// Para el caso de que falle la carga del servidor devolverá un error
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Port: ", portString)
}
