package main

import (
	"context"
	"fmt"
	"golang.org/x/oauth2/clientcredentials"
	"io"
	"net/http"
	"time"
)

func main() {
	ctx := context.Background()
	conf := clientcredentials.Config{
		ClientID:       "000000",
		ClientSecret:   "999999",
		TokenURL:       "http://localhost:9096/token",
		Scopes:         nil,
		EndpointParams: nil,
		AuthStyle:      0,
	}

	tok, _ := conf.Token(ctx)
	fmt.Println(tok.AccessToken, tok.Expiry)

	client := conf.Client(ctx)

	for {
		request, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:9096/hitme", nil)
		response, err := client.Do(request)
		if err != nil {
			fmt.Printf(err.Error())
			return
		}
		body, _ := io.ReadAll(response.Body)
		fmt.Println(string(body))
		time.Sleep(3 * time.Second)
	}
}

/*
	conf := &oauth2.Config{
		ClientID:     "000000",
		ClientSecret: "999999",
		Scopes:       []string{"my-scope"},
		Endpoint: oauth2.Endpoint{
			AuthURL:  "http://localhost:9096/authorize",
			TokenURL: "http://localhost:9096/token",
		},
		RedirectURL: "http://localhost:9095",
	}

	// Redirect user to consent page to ask for permission
	// for the scopes specified above.
	url := conf.AuthCodeURL("", oauth2.AccessTypeOffline)
	fmt.Println(url)

	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	resp, err := client.Get(url)
	if err != nil {
		fmt.Println("ERROR", err.Error())
		return
	}
	location, err := resp.Location()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	code := location.Query().Get("code")

	fmt.Println("code:", code)

	// Use the authorization code that is pushed to the redirect
	// URL. Exchange will do the handshake to retrieve the
	// initial access token. The HTTP Client returned by
	// conf.Client will refresh the token as necessary.
	tok, err := conf.Exchange(ctx, code)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(tok.AccessToken, tok.RefreshToken, tok.Expiry)
*/
