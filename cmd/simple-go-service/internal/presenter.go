package internal

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Presenter struct {
}

func NewPresenter() *Presenter {
	return &Presenter{}
}

type SuccessfulResponse struct {
	Message Message `json:"success"`
}

type Message struct {
	Message string `json:"message"`
}

type ErrorRespose struct {
	Message Message `json:"error"`
}

func NewBadRequestResponse() ErrorRespose {
	return ErrorRespose{
		Message: Message{http.StatusText(http.StatusBadRequest)},
	}
}

func NewSuccessfulResponse() SuccessfulResponse {
	return SuccessfulResponse{
		Message: Message{http.StatusText(http.StatusOK)},
	}
}

func (p *Presenter) GetData(ctx *gin.Context) {
    // Log the value of the If-Modified-Since header
    ifModifiedSince := ctx.GetHeader("If-Modified-Since")
    fmt.Println("HEADER: ", ifModifiedSince)

    // Attempt to parse the If-Modified-Since header
    var lastModified time.Time
    var err error
    if ifModifiedSince != "" {
        lastModified, err = http.ParseTime(ifModifiedSince)
        if err != nil {
            logrus.Errorf("Problem parsing If-Modified-Since header: %v", err)
            ctx.JSON(http.StatusBadRequest, NewBadRequestResponse())
            return
        }
    }

    // For debugging purposes
    if !lastModified.IsZero() {
        fmt.Println(lastModified)
    }

    // Set Last-Modified header with the current time
    ctx.Header("Last-Modified", time.Now().UTC().Format(http.TimeFormat))

    // Respond with a successful response
    ctx.JSON(http.StatusOK, NewSuccessfulResponse())
}
