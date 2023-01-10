package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
)

func main() {
	cfg := jaegercfg.Configuration{
		Sampler: &jaegercfg.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 10,
		},
		Reporter: &jaegercfg.ReporterConfig{
			LogSpans:           true,
			LocalAgentHostPort: "127.0.0.1:6831",
		},
	}

	closer, err := cfg.InitGlobalTracer(
		"main-service",
	)
	defer closer.Close()

	if err != nil {
		log.Printf("Could not initialize jaeger tracer: %s", err.Error())
		return
	}

	http.HandleFunc("/get-product", handleGetProduct)

	log.Println("Starting service in port 8000")
	http.ListenAndServe(":8000", nil)
}

func handleGetProduct(w http.ResponseWriter, req *http.Request) {
	trace, ctx := opentracing.StartSpanFromContext(req.Context(), "GET /get-product")
	time.Sleep(250 * time.Millisecond)
	defer trace.Finish()

	// check login
	if isLogin(ctx) {
		log.Println("User Logged In")
	}

	go sendNotifEmail(ctx)
	go sendNotifSms(ctx)

	// Get Product
	product := getProduct(ctx)
	log.Println("Product: ", product)

	fmt.Fprint(w, "Product: ", product)
}

func sendNotifEmail(ctx context.Context) {
	trace, ctx := opentracing.StartSpanFromContext(ctx, "func sendNotifEmail")
	time.Sleep(2000 * time.Millisecond)
	defer trace.Finish()
}

func sendNotifSms(ctx context.Context) {
	trace, ctx := opentracing.StartSpanFromContext(ctx, "func sendNotifSms")
	time.Sleep(1000 * time.Millisecond)
	defer trace.Finish()
}

func isLogin(ctx context.Context) bool {
	trace, ctx := opentracing.StartSpanFromContext(ctx, "func isLogin")
	time.Sleep(450 * time.Millisecond)
	defer trace.Finish()

	return true
}

func getProduct(ctx context.Context) map[string]string {
	trace, ctx := opentracing.StartSpanFromContext(ctx, "func getProduct")
	time.Sleep(750 * time.Millisecond)
	defer trace.Finish()

	return map[string]string{
		"P001": "Sabun",
		"P002": "Handuk",
	}
}
