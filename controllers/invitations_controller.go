package controllers

import (
	"dashcode/security"
	"dashcode/services"
	"dashcode/services/invitations"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func FetchInvitations(c echo.Context) error {
	jwt := c.Request().Header.Get("Authorization")

	if len(jwt) == 0 {
		return c.JSON(services.ErrorJWT.Code, services.ErrorJWT)
	}

	id, err := security.IDFromJWT(jwt)

	if err != nil {
		return c.JSON(services.ErrorJWT.Code, services.ErrorJWT)
	}

	invs, srvErr := invitations.FetchInvitations(id)

	if srvErr != nil {
		return c.JSON(srvErr.Code, srvErr)
	}

	return c.JSON(http.StatusOK, invs)
}

func FetchInvitationsByGroup(c echo.Context) error {
	jwt := c.Request().Header.Get("Authorization")

	id, err := security.IDFromJWT(jwt)

	if err != nil {
		return c.JSON(services.ErrorJWT.Code, services.ErrorJWT)
	}

	groupIdParam := c.Param("groupId")

	groupId, err := strconv.Atoi(groupIdParam)

	if err != nil {
		return c.JSON(services.ErrorPathParam.Code, services.ErrorPathParam)
	}

	t := c.QueryParam("state")

	var state *invitations.InvitationState

	switch t {
	case "accepted":
		state = invitations.Accepted
	case "rejected":
		state = invitations.Rejected
	default:
		state = nil
	}

	invs, srvErr := invitations.FetchInvitationsWithState(id, int64(groupId), state)

	if srvErr != nil {
		return c.JSON(srvErr.Code, srvErr)
	}

	return c.JSON(http.StatusOK, invs)
}
