package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	errInitConfig := initConfig()

	if errInitConfig != nil {
		logrus.Fatalf("Error load config file server: %s", errInitConfig.Error())
	}

	remote, err := url.Parse(viper.GetString("ip_auth_service"))
	if err != nil {
		panic(err)
	}

	handler := func(p *httputil.ReverseProxy) func(http.ResponseWriter, *http.Request) {
		return func(w http.ResponseWriter, r *http.Request) {
			log.Println(r.URL)
			r.Host = remote.Host
			w.Header().Set("X-Ben", "Rad")
			p.ServeHTTP(w, r)
		}
	}

	proxy := httputil.NewSingleHostReverseProxy(remote)
	http.HandleFunc("/auth/", handler(proxy))

	remote_url, err := url.Parse(viper.GetString("ip_url_service"))
	if err != nil {
		panic(err)
	}

	handler_url := func(p *httputil.ReverseProxy) func(http.ResponseWriter, *http.Request) {
		return func(w http.ResponseWriter, r *http.Request) {
			log.Println(r.URL)
			r.Host = remote_url.Host
			w.Header().Set("X-Ben", "Rad")
			p.ServeHTTP(w, r)
		}
	}

	proxy_url := httputil.NewSingleHostReverseProxy(remote_url)

	http.HandleFunc("/api/urls/", handler_url(proxy_url))

	//Redirect-----------------------------------------------------------------------------
	handler_red := func(p *httputil.ReverseProxy) func(http.ResponseWriter, *http.Request) {
		return func(w http.ResponseWriter, r *http.Request) {
			log.Println(r.URL)
			w.Header().Set("X-Ben", "Rad")
			p.ServeHTTP(w, r)
		}
	}

	proxy_red := httputil.NewSingleHostReverseProxy(remote_url)

	http.HandleFunc("/:hashURL", handler_red(proxy_red))

	err = http.ListenAndServe(":8011", nil)
	if err != nil {
		panic(err)
	}
}

func initConfig() error {
	viper.AddConfigPath("/home/andrey/go/api_gateway/configs/")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
