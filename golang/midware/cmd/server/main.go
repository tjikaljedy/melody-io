package main

import (
	"context"

	"flag"
	"melody-io/core/pkg/config"
	"melody-io/core/pkg/log"
	route "melody-io/midware/pkg/rsocket-route"
	"net/http"
	"os"

	"github.com/rsocket/rsocket-go"
	"github.com/rsocket/rsocket-go/core/transport"
	"github.com/rsocket/rsocket-go/payload"
	"github.com/rsocket/rsocket-go/rx"
	"github.com/rsocket/rsocket-go/rx/flux"
	"github.com/rsocket/rsocket-go/rx/mono"
)

var tp transport.ServerTransporter
var flagConfig = flag.String("config", "../../config/local.yml", "config file")
var logger = log.New().With(context.Background())
var Routings = make(map[string]interface{})

func init() {
	flag.Parse()
	cfg, err := config.Load(*flagConfig, logger)
	if err != nil {
		logger.Errorf("failed to load application configuration: %s", err)
		os.Exit(-1)
	}

	Routings["/request-response"] = echoRR
	Routings["/request-stream"] = echoRS
	Routings["/request-channel"] = echoRC

	err = route.Add(Routings)
	if err != nil {
		panic(err)
	}

	tp = rsocket.TCPServer().SetHostAndPort(cfg.RSocketHost, cfg.RSocketPort).Build()
	go func() {
		logger.Info(http.ListenAndServe(cfg.GatewayServe, nil))
	}()
}

func main() {
	err := rsocket.Receive().
		OnStart(func() {
			logger.Info("melody-io middleware run...")
		}).
		Acceptor(func(ctx context.Context, setup payload.SetupPayload, sendingSocket rsocket.CloseableRSocket) (rsocket.RSocket, error) {
			sendingSocket.OnClose(func(err error) {
				logger.Info("*** socket disconnected ***")
			})
			//return responder(), nil
			return rsocket.NewAbstractSocket(route.GetHandlers()...), nil
		}).
		Transport(tp).
		Serve(context.Background())
	if err != nil {
		panic(err)
	}
}
func echoRR(msg payload.Payload) mono.Mono {
	return mono.Just(msg)
}

func echoRS(msg payload.Payload) flux.Flux {
	return flux.Just(msg)
}

func echoRC(msgs rx.Publisher) flux.Flux {
	return msgs.(flux.Flux)
}

/*
func responder() rsocket.RSocket {
	return rsocket.NewAbstractSocket(
		rsocket.MetadataPush(func(item payload.Payload) {
			logger.Info("GOT METADATA_PUSH:", item)
		}),
		rsocket.FireAndForget(func(elem payload.Payload) {
			logger.Info("GOT FNF:", elem)
		}),
		rsocket.RequestResponse(func(pl payload.Payload) mono.Mono {
			fmt.Print(">>>>>>>>> ", pl.DataUTF8())
			fmt.Print(pl.MetadataUTF8())
			return mono.JustOneshot(pl)
		}),
		rsocket.RequestStream(func(pl payload.Payload) flux.Flux {
			s := pl.DataUTF8()
			m, _ := pl.MetadataUTF8()
			fmt.Print("data:", s, "metadata:", m)
			totals := 10
			if n, err := strconv.Atoi(m); err == nil {
				totals = n
			}
			return flux.Create(func(ctx context.Context, emitter flux.Sink) {
				for i := 0; i < totals; i++ {
					// You can use context for graceful coroutine shutdown, stop produce.
					select {
					case <-ctx.Done():
						fmt.Print("ctx done:", ctx.Err())
						return
					default:
						emitter.Next(payload.NewString(fmt.Sprintf("%s_%d", s, i), m))
					}
				}
				emitter.Complete()
			})
		}),
		rsocket.RequestChannel(func(payloads flux.Flux) flux.Flux {
			//return payloads.(flux.Flux)
			payloads.
				//LimitRate(1).
				DoOnNext(func(next payload.Payload) error {
					fmt.Print("receiving:", next.DataUTF8())
					return nil
				}).
				Subscribe(context.Background())
			return flux.Create(func(i context.Context, sink flux.Sink) {
				for i := 0; i < 3; i++ {
					sink.Next(payload.NewString("world", fmt.Sprintf("%d", i)))
				}
				sink.Complete()
			})
		}),
	)
}
*/
