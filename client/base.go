package client

import (
	"strings"
)

// ACK statuses.
const (
	ACKSuccess = "Success"
	ACKFailure = "Failure"
)

// BaseRequest is base struct for api request
type BaseRequest struct {
	Method string `url:"METHOD"`
	Action string `url:"PAYMENTREQUEST_0_PAYMENTACTION,omitempty"`
}

// BaseResponse is base struct for api response
type BaseResponse struct {
	Timestamp     string `url:"TIMESTAMP"`
	ACK           string `url:"ACK"`
	Version       string `url:"VERSION"`
	Build         string `url:"BUILD"`
	CorrelationID string `url:"CORRELATIONID"`

	// error0
	ErrorCode    string `url:"L_ERRORCODE0"`
	ShortMessage string `url:"L_SHORTMESSAGE0"`
	LongMessage  string `url:"L_LONGMESSAGE0"`
	SeverityCode string `url:"L_SEVERITYCODE0"`

	// error1
	ErrorCode1    string `url:"L_ERRORCODE1"`
	ShortMessage1 string `url:"L_SHORTMESSAGE1"`
	LongMessage1  string `url:"L_LONGMESSAGE1"`
	SeverityCode1 string `url:"L_SEVERITYCODE1"`
}

// IsRequestSuccess checks the request is success or not.
func (r BaseResponse) IsRequestSuccess() bool {
	return r.ACK == ACKSuccess
}

func (r BaseResponse) baseError() []string {
	var errs []string
	if r.ShortMessage != "" {
		errs = append(errs, r.ShortMessage+" "+r.LongMessage)
	}
	if r.ShortMessage1 != "" {
		errs = append(errs, r.ShortMessage1+" "+r.LongMessage1)
	}
	return errs
}

// Error returns error message
func (r BaseResponse) Error() string {
	return parseErrors(r.baseError())
}

// Errors returns error messages
func (r BaseResponse) Errors(s []string) string {
	errs := r.baseError()
	errs = append(errs, s...)
	return parseErrors(errs)
}

func parseErrors(errs []string) string {
	return strings.Join(errs, ",")
}

// Response is interface of response from each API
type Response interface {
	IsSuccess() bool
	IsRequestSuccess() bool
	IsOperationSuccess() bool
	Error() string
}

// ShippingResponse is for shipping information
type ShippingResponse struct {
	ShipToName        string `url:"SHIPTONAME"`
	ShipToStreet      string `url:"SHIPTOSTREET"`
	ShipToCity        string `url:"SHIPTOCITY"`
	ShipToState       string `url:"SHIPTOSTATE"`
	ShipToZip         string `url:"SHIPTOZIP"`
	ShipToCountryCode string `url:"SHIPTOCOUNTRYCODE"`
	ShipToCountry     string `url:"SHIPTOCOUNTRY"`
	ShipToCountryName string `url:"SHIPTOCOUNTRYNAME"`
	ShipAddressOwner  string `url:"SHIPADDRESSOWNER"`
	ShipAddressStatus string `url:"SHIPADDRESSSTATUS"`
}
