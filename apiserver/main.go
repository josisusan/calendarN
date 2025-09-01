package apiserver

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/samit22/calendarN/api"
)

type CalendarService struct {
}

func (c CalendarService) CalendarToday(ctx context.Context) (*api.CalendarResponse, error) {
	return &api.CalendarResponse{
		Nepali:  api.OptCalendarResponseNepali{},
		English: api.OptCalendarResponseEnglish{},
	}, nil
}

type APISec struct {
	Token string
}

func (sec APISec) HandleApiKeyAuth(ctx context.Context, operationName api.OperationName, t api.ApiKeyAuth) (context.Context, error) {
	if t.APIKey == sec.Token {
		return ctx, nil
	}
	return nil, fmt.Errorf("Security breached")
}

func Run(port int, token string) {
	service := CalendarService{}
	secHandler := APISec{
		Token: token,
	}
	srv, err := api.NewServer(service, secHandler)
	if err != nil {
		fmt.Println("Error")
	}
	portString := fmt.Sprintf(":%d", port)
	fmt.Printf("Server running at port%s \n", portString)
	if err := http.ListenAndServe(portString, srv); err != nil {
		fmt.Println(err)
	}
}
