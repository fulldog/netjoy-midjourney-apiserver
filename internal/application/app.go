package application

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/hongliang5316/midjourney-apiserver/internal/common"
	"github.com/hongliang5316/midjourney-apiserver/internal/config"
	"github.com/hongliang5316/midjourney-apiserver/internal/service"
	"github.com/hongliang5316/midjourney-apiserver/pkg/api"
	"github.com/hongliang5316/midjourney-apiserver/pkg/store"
	"github.com/hongliang5316/midjourney-go/midjourney"
	"google.golang.org/grpc"
)

type Application struct {
	*common.Base
}

func New() *Application {
	cfg := config.Load()

	dg, err := discordgo.New(cfg.Midjourney.UserToken)
	if err != nil {
		log.Fatal(err)
	}

	cli := midjourney.NewClient(&midjourney.Config{
		UserToken: cfg.Midjourney.UserToken,
	})

	stor := store.NewStore(&store.Config{
		Redis: store.Redis{
			Address:  cfg.Redis.Address,
			Password: cfg.Redis.Password,
		},
	})

	app := &Application{Base: &common.Base{Session: dg, Store: stor, MJClient: cli, Config: cfg}}

	dg.AddHandler(app.messageCreate)
	dg.AddHandler(app.messageUpdate)
	dg.AddHandler(app.messageDelete)

	dg.Identify.Intents = discordgo.IntentsAll

	return app
}

func NewCfg(cfg *config.Config) *Application {
	dg, err := discordgo.New(cfg.Midjourney.UserToken)
	if err != nil {
		log.Fatal(err)
	}

	cli := midjourney.NewClient(&midjourney.Config{
		UserToken: cfg.Midjourney.UserToken,
	})
	if cfg.HttpProxy != "" {
		proxy, _ := url.Parse(cfg.HttpProxy)
		dg.Client = &http.Client{
			Transport: &http.Transport{
				Proxy: http.ProxyURL(proxy),
			},
		}
		dg.Dialer = &websocket.Dialer{
			Proxy: http.ProxyURL(proxy),
		}
		cli.Client = &http.Client{
			Transport: &http.Transport{
				Proxy: http.ProxyURL(proxy),
			},
		}
	}
	stor := store.NewStore(&store.Config{
		Redis: store.Redis{
			Address:  cfg.Redis.Address,
			Password: cfg.Redis.Password,
		},
	})

	app := &Application{Base: &common.Base{Session: dg, Store: stor, MJClient: cli, Config: cfg}}

	dg.AddHandler(app.messageCreate)
	dg.AddHandler(app.messageUpdate)
	dg.AddHandler(app.messageDelete)

	dg.Identify.Intents = discordgo.IntentsAll

	return app
}

func (app *Application) Run() error {
	go func(app *Application) {
		lis, err := net.Listen("tcp", fmt.Sprintf(":%d", app.Config.ListenPort))
		if err != nil {
			log.Fatalf("failed to listen: %+v", err)
		}

		s := grpc.NewServer()
		api.RegisterAPIServiceServer(s, service.New(app.Base))

		if err := s.Serve(lis); err != nil {
			log.Fatal(err)
		}
	}(app)

	err := app.Open()
	if err != nil {
		return fmt.Errorf("Call app.Open failed, err: %w", err)
	}

	log.Printf("Start...")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	return app.Close()
}
