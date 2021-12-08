package main

import (
	"fmt"
	"log"
	"net/http"
)

type Logger interface {
	Info(str string)
	Error(str string)
}

type Server interface {
	Serve() error
	SetLogger(Logger)
}

type HttpServer struct {
	httpPort int

	logger Logger
}

func (s *HttpServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.logger.Info("Http Handler start")
}

func (s *HttpServer) Serve() error {
	return http.ListenAndServe(fmt.Sprintf(":%d", s.httpPort), s)
}

func (s *HttpServer) SetLogger(l Logger) {
	s.logger = l
}

func NewHttpServer(port int) Server {
	return &HttpServer{
		httpPort: port,
	}
}

type FasthttpServer struct {
	fasthttpPort int

	logger Logger
}

func (s *FasthttpServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.logger.Info("Типа Fasthttp Handler start")
}

func (s *FasthttpServer) Serve() error {
	return http.ListenAndServe(fmt.Sprintf(":%d", s.fasthttpPort), s)
}

func (s *FasthttpServer) SetLogger(l Logger) {
	s.logger = l
}

func NewFasthttpServer(port int) Server {
	return &FasthttpServer{
		fasthttpPort: port,
	}
}

type BlablaLogger struct{}

func (b *BlablaLogger) Info(str string) {
	log.Println("blabla", str)
}

func (b *BlablaLogger) Error(str string) {
	log.Println("blablabla!!!", str)
}

func NewBlablaLogger() Logger {
	return &BlablaLogger{}
}

type NormalLogger struct{}

func (b *NormalLogger) Info(str string) {
	log.Println("INFO", str)
}

func (b *NormalLogger) Error(str string) {
	log.Println("ERROR", str)
}

func NewNormalLogger() Logger {
	return &NormalLogger{}
}

func main() {
	httpServer := NewHttpServer(8080)

	fasthttpServer := NewFasthttpServer(8081)

	blablaLogger := NewBlablaLogger()

	normLogger := NewNormalLogger()

	httpServer.SetLogger(blablaLogger)

	fasthttpServer.SetLogger(normLogger)

	go func() {
		httpServer.Serve()
	}()

	fasthttpServer.Serve()
}
