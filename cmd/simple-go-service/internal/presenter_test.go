package internal_test

import (
    "net/http"
    "net/http/httptest"
    "time"

    internal "github.com/nenov92/simple-go-service/cmd/simple-go-service/internal"
    "github.com/gin-gonic/gin"
    . "github.com/onsi/ginkgo"
    . "github.com/onsi/gomega"
)

var _ = Describe("Presenter", func() {

    var presenter *internal.Presenter
    var recorder *httptest.ResponseRecorder
    var mockContext *gin.Context

    BeforeEach(func() {
        presenter = internal.NewPresenter()
        recorder = httptest.NewRecorder()
        mockContext, _ = gin.CreateTestContext(recorder)
        mockContext.Request = httptest.NewRequest("GET", "/v1/data", nil)
    })

    Describe("Get Data", func() {
        When("Get Data is called with proper If-Modified-Since Header", func() {
            It("should return a data list with JSON Content-Type and Status OK", func() {
                // Set a valid If-Modified-Since header
                mockContext.Request.Header.Set("If-Modified-Since", time.Now().UTC().Format(http.TimeFormat))
                presenter.GetData(mockContext)

                Expect(mockContext.Writer.Status()).To(Equal(http.StatusOK))
                Expect(mockContext.Writer.Header().Get("Content-Type")).To(Equal("application/json; charset=utf-8"))
            })
        })

        When("Get Data is called without If-Modified-Since Header", func() {
            It("should return a data list with JSON Content-Type and Status OK", func() {
                // No If-Modified-Since header set
                presenter.GetData(mockContext)

                Expect(mockContext.Writer.Status()).To(Equal(http.StatusOK))
                Expect(mockContext.Writer.Header().Get("Content-Type")).To(Equal("application/json; charset=utf-8"))
            })
        })

        When("Get Data is called with an invalid If-Modified-Since Header", func() {
            It("should return Status BadRequest", func() {
                // Set an invalid If-Modified-Since header
                mockContext.Request.Header.Set("If-Modified-Since", "InvalidHeaderValue")
                presenter.GetData(mockContext)

                Expect(mockContext.Writer.Status()).To(Equal(http.StatusBadRequest))
                Expect(mockContext.Writer.Header().Get("Content-Type")).To(Equal("application/json; charset=utf-8"))
            })
        })
    })

})
