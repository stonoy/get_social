package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/stonoy/get_social/internal"
)

type apiConfig struct {
	db         *internal.Queries
	jwt_secret string
}

func main() {
	// load env files
	err := godotenv.Load()
	if err != nil {
		log.Panicf("error inloading env variable -> %v", err)
		return
	}

	port := os.Getenv("PORT")
	if port == "" {
		log.Panicln("no port in env file")
		return
	}

	secret := os.Getenv("JWT_SECRET")

	apiCfg := &apiConfig{
		jwt_secret: secret,
	}

	conn_url := os.Getenv("CONN_URL")
	if conn_url != "" {
		// open the driver
		db, err := sql.Open("postgres", conn_url)
		if err != nil {
			log.Panicf("can not connect to the server via go -> %v", err)
		}

		db_quries := internal.New(db)

		apiCfg.db = db_quries
	} else {
		log.Println("server started without database connection")
	}

	// create a main router
	mainRouter := chi.NewRouter()

	// create a sub router
	apiRouter := chi.NewRouter()

	// user
	apiRouter.Post("/register", apiCfg.register)
	apiRouter.Post("/login", apiCfg.login)
	apiRouter.Get("/getusers", apiCfg.GetUsers)
	apiRouter.Put("/updateusers", apiCfg.authChecker(apiCfg.UpdateUsers))
	apiRouter.Get("/getusersdetails/{userID}", apiCfg.GetSingleUserDetails)

	// posts
	apiRouter.Post("/createposts", apiCfg.authChecker(apiCfg.CreatePosts))
	apiRouter.Get("/getsinglepost/{postID}", apiCfg.GetSinglePost)
	apiRouter.Delete("/deleteposts/{postID}", apiCfg.authChecker(apiCfg.DeletePost))
	apiRouter.Put("/updateposts/{postID}", apiCfg.authChecker(apiCfg.UpdatePost))
	apiRouter.Get("/getpostsuggestions", apiCfg.authChecker(apiCfg.PostSuggestion))
	apiRouter.Get("/getpostsbyuser", apiCfg.authChecker(apiCfg.GetPostsByUser))

	// likes
	apiRouter.Get("/addlike/{postID}", apiCfg.authChecker(apiCfg.LikeAPost))
	apiRouter.Get("/removelike/{postID}", apiCfg.authChecker(apiCfg.RemoveLike))

	// comments
	apiRouter.Post("/createcomments", apiCfg.authChecker(apiCfg.CreateComment))
	apiRouter.Delete("/deletecomments/{commentID}", apiCfg.authChecker(apiCfg.DeleteComment))
	apiRouter.Put("/updatecomments/{commentID}", apiCfg.authChecker(apiCfg.UpdateComment))
	apiRouter.Get("/getpostcomments/{postID}", apiCfg.GetPostComments)

	// follows
	apiRouter.Post("/followpersons", apiCfg.authChecker(apiCfg.Follow))
	apiRouter.Delete("/unfollowpersons/{personID}", apiCfg.authChecker(apiCfg.Unfollow))
	apiRouter.Get("/followsuggestions", apiCfg.authChecker(apiCfg.FollowSuggestions))

	// mount
	mainRouter.Mount("/api/v1", apiRouter)

	// create a new server from http.Server type
	the_server := &http.Server{
		Addr:    ":" + port,
		Handler: mainRouter,
	}

	log.Printf("Server is listenning on port %v", port)

	// starts the server
	log.Panic(the_server.ListenAndServe())
}
