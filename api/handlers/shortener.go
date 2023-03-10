package handlers

import (
	"fmt"
	pb "go_auth_api_gateway/genproto/auth_service"
	"go_auth_api_gateway/pkg/utils"

	"net/http"

	http_status "go_auth_api_gateway/api/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/skip2/go-qrcode"
)

// @Security ApiKeyAuth
// CreateShortUrl godoc
// @ID create_short_url
// @Router /short-url [POST]
// @Summary Create ShortUrl
// @Description Create ShortUrl
// @Tags urls
// @Accept json
// @Produce json
// @Param body body auth_service.CreateShortUrlRequest true "Request body"
// @Success 201 {object} http.Response{data=auth_service.CreateShortUrlResponse} "Response Body"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) CreateShortUrl(c *gin.Context) {
	var req pb.CreateShortUrlRequest

	if err := c.ShouldBindJSON(&req); err != nil {

		h.handleResponse(c, http_status.BadRequest, err.Error())
		return
	}

	if !utils.IsLongCorrect(string(req.GetLongUrl())) {
		err := fmt.Errorf(utils.InvalidURLError, req.GetLongUrl())
		h.handleResponse(c, http_status.BadRequest, err.Error())
		return
	}

	resp, err := h.services.ShortenerService().CreateShortUrl(c, &req)
	if err != nil {
		h.handleResponse(c, http_status.InternalServerError, err.Error())
		return
	}

	h.handleResponse(c, http_status.OK, resp)
}

// @Security ApiKeyAuth
// GetShortUrl godoc
// @ID get_short_url
// @Router /short-url/{hash} [GET]
// @Summary Get ShortUrlData
// @Description Get ShortUrl
// @Tags urls
// @Accept json
// @Produce json
// @Param hash path string true "short url hash"
// @Success 201 {object} http.Response{data=auth_service.GetShortUrlResponse} "Response Body"
func (h *Handler) GetShortUrlData(c *gin.Context) {

	var req pb.GetShortUrlRequest

	hash := c.Param("hash")
	if !utils.IsShortCorrect(hash) {
		err := fmt.Errorf(utils.InvalidHashError, hash)
		h.handleResponse(c, http_status.BadRequest, err.Error())
		return
	}
	req.ShortUrl = hash

	req.ShortUrl = hash
	resp, err := h.services.ShortenerService().GetShortUrl(c, &req)
	if err != nil {
		h.handleResponse(c, http_status.InternalServerError, err.Error())
		return
	}

	h.handleResponse(c, http_status.OK, resp)
}

// @Security ApiKeyAuth
// HandleLonger godoc
// @ID handle_longer
// @Router /sigma/{hash} [GET]
// @Summary Handle Longer
// @Description Handle Longer
// @Tags urls
// @Param hash path string true "short url hash"
// @Success 201 {object} http.Response{data=string} "Response Body"
func (h *Handler) HandleLonger(c *gin.Context) {

	url := c.Param("hash")
	if !utils.IsShortCorrect(url) {
		err := fmt.Errorf(utils.InvalidHashError, url)
		h.handleResponse(c, http_status.BadRequest, err.Error())
		return
	}

	resp, err := h.services.ShortenerService().HandleLongUrl(
		c.Request.Context(),
		&pb.HandleLongUrlRequest{
			ShortUrl: url,
		},
	)
	if err != nil {
		h.handleResponse(c, http_status.InternalServerError, err.Error())
		return
	}

	_, err = h.services.ShortenerService().IncClickCount(
		c.Request.Context(),
		&pb.IncClickCountRequest{
			ShortUrl: url,
		},
	)
	if err != nil {
		h.handleResponse(c, http_status.InternalServerError, err.Error())
		return
	}

	c.Redirect(http.StatusMovedPermanently, resp.GetLongUrl())
}

// @Security ApiKeyAuth
// GetAllUserUrls godoc
// @ID get_all_user_urls
// @Router /user/short-url/{user-id} [GET]
// @Summary Get All User Urls
// @Description Get All User Urls
// @Tags urls
// @Accept json
// @Produce json
// @Param user-id path string true "user id"
// @Success 200 {object} http.Response{data=auth_service.GetAllUserUrlsResponse} "Response Body"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetAllUserUrls(c *gin.Context) {

	userId := c.Param("user-id")
	if _, err := uuid.Parse(userId); err != nil {
		err := fmt.Errorf(utils.InvalidUUIDError, userId)
		h.handleResponse(c, http_status.BadRequest, err.Error())
		return
	}

	resp, err := h.services.ShortenerService().GetAllUserUrls(c, &pb.GetAllUserUrlsRequest{
		UserId: userId,
	})
	if err != nil {
		h.handleResponse(c, http_status.InternalServerError, err.Error())
		return
	}

	h.handleResponse(c, http_status.OK, resp)
}

// @Security ApiKeyAuth
// UrlToQrcode godoc
// @ID qr_code
// @Router /url-qrcode [PUT]
// @Summary Url Convert Qrcode
// @Description Convert Qrcode
// @Tags Qrcode
// @Accept json
// @Produce image/png
// @Param tiny-url query string true "tiny-url"
// @Success 200 {string} string "QR code image"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) UrlToQrcode(c *gin.Context) {
	stringUrl := c.Query("tiny-url")

	// Generate QR code image from the URL string
	qrCode, err := qrcode.Encode(stringUrl, qrcode.Medium, 256)
	if err != nil {
		h.handleResponse(c, http_status.InternalServerError, "Error generating QR code")
		return
	}

	// Set response header to indicate image type
	c.Header("Content-Type", "image/png")

	// Write image data to response body
	c.Writer.Write(qrCode)

	// Set response status to OK
	c.Status(http_status.OK.Code)
}

// @Security ApiKeyAuth
// UpdateShortUrl godoc
// @ID update_short_url
// @Router /short-url [PUT]
// @Summary Update ShortUrl
// @Description Update ShortUrl
// @Tags urls
// @Accept json
// @Produce json
// @Param body body auth_service.CreateShortUrlRequest true "Request body"
// @Success 201 {object} http.Response{data=string} "Response Body"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) UpdateShortUrl(c *gin.Context) {
	var req pb.CreateShortUrlRequest

	if err := c.ShouldBindJSON(&req); err != nil {

		h.handleResponse(c, http_status.BadRequest, err.Error())
		return
	}

	if !utils.IsLongCorrect(string(req.GetLongUrl())) {
		err := fmt.Errorf(utils.InvalidURLError, req.GetLongUrl())
		h.handleResponse(c, http_status.BadRequest, err.Error())
		return
	}

	resp, err := h.services.ShortenerService().UpdateShortUrl(c, &req)
	if err != nil {
		h.handleResponse(c, http_status.InternalServerError, err.Error())
		return
	}

	h.handleResponse(c, http_status.OK, resp)
}