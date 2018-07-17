// Package classification User API.
//
// The purpose of this service is to provide an application
// that is using plain go code to define an API
//
//      Host: localhost
//      Version: 0.0.1
//
// swagger:meta
package main

import "github.com/jamesblockk/Start-Go/HelloGo/Swagger/models"

// GetUsrReqWrapper Request Setup
//
// swagger:parameters GetUsrReqWrapper
type GetUsrReqWrapper struct {
	// in: body
	Body models.GetUsrReq
}

// GetUsrRespWrapper Response Setup
//
// swagger:response GetUsrRespWrapper
type GetUsrRespWrapper struct {
	// in: body
	Body models.GetUsrResp
}

// swagger:route GET /usr usr GetUsrRespWrapper
//
// GET Usr
//
// This will Get user info
//
//     Responses:
//       200: UpdateUserResponseWrapper
