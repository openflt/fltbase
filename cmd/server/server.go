package main

import (
	"context"
	"net/http"

	"entgo.io/contrib/entgql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/debug"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/alecthomas/kong"
	"github.com/go-chi/chi"
	"github.com/gorilla/websocket"
	_ "github.com/mattn/go-sqlite3"
	"github.com/rs/cors"
	"go.uber.org/zap"

	"github.com/openflt/fltbase"
	"github.com/openflt/fltbase/ent"
	_ "github.com/openflt/fltbase/ent/runtime"
)

func main() {
	var cli struct {
		Addr  string `name:"address" default:":8081" help:"Address to listen on."`
		Debug bool   `name:"debug" help:"Enable debugging mode."`
	}
	kong.Parse(&cli)

	log, _ := zap.NewDevelopment()
	client, err := ent.Open(
		"sqlite3",
		"file:ent?mode=memory&cache=shared&_fk=1",
	)
	if err != nil {
		log.Fatal("opening ent client", zap.Error(err))
	}
	if err := client.Schema.Create(
		context.Background(),
		//migrate.WithGlobalUniqueID(true),
	); err != nil {
		log.Fatal("running schema migration", zap.Error(err))
	}

	router := chi.NewRouter()

	router.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		Debug:            true,
	}).Handler)

	srv := handler.NewDefaultServer(fltbase.NewSchema(client))
	srv.Use(entgql.Transactioner{TxOpener: client})
	srv.AddTransport(&transport.Websocket{
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
		},
	})

	if cli.Debug {
		srv.Use(&debug.Tracer{})
	}

	router.Handle("/", playground.Handler("Todos", "/query"))
	router.Handle("/query", srv)

	//
	// http.Handle("/",
	//
	//	playground.Handler("Todo", "/query"),
	//
	// )
	// http.Handle("/query", cors(srv))
	//
	log.Info("listening on", zap.String("address", cli.Addr))
	if err := http.ListenAndServe(cli.Addr, router); err != nil {
		log.Error("http server terminated", zap.Error(err))
	}
}
