package controller

import (
	"errors"
	"fmt"

	"github.com/TwiN/logr"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/hc-psl-cs-group/backend/httputil"
	"github.com/hc-psl-cs-group/backend/model"
)

// Request handler to store response to questionnaire
//
// @Summary 	Store response to a qustionnaire
// @Description	Creates a new response for to a questionnaire for the current user
// @Tags		responses
// @Produce		json
// @Param 		payload body model.RecordResponseRequest true "Response to store"
// @Success		204	{object}		nil
// @Failure 	400 {object} 		nil
// @Failure		401	{object}		nil
// @Failure		500	{object}		httputil.HTTPError
// @Router		/responses/record [put]
func (m *Controller) RecordResponse(c *gin.Context) {
	var req model.RecordResponseRequest
	err := c.BindJSON(&req)
	if err != nil {
		logr.Debug(fmt.Sprint("Failed to parse body to record response, error: ", err))
		c.AbortWithStatus(400)
		return
	}

	uid, ok := sessions.Default(c).Get("userid").(uint)
	if !ok {
		c.AbortWithStatus(401)
		return
	}

	_, ok = m.db.RecordResponse(req, uid)
	if ok {
		c.JSON(204, nil)
	} else {
		httputil.NewError(c, 500, errors.New("ERR_CREATE_ASSET"))
		c.Abort()
	}
}
