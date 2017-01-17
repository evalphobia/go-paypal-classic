package merchant

import "github.com/evalphobia/go-paypal-classic/client"

// ManageRecurringPaymentsProfileStatus is struct for ManageRecurringPaymentsProfileStatus API
// see: https://developer.paypal.com/docs/classic/api/merchant/ManageRecurringPaymentsProfileStatus_API_Operation_NVP/
type ManageRecurringPaymentsProfileStatus struct {
	client.BaseRequest `url:",squash"`

	ProfileID string `url:"PROFILEID"`
	Action    string `url:"ACTION"`
	Note      string `url:"NOTE"`
}

// SetAsCancel sets action as `Cancel`
func (svc *ManageRecurringPaymentsProfileStatus) SetAsCancel(desc string) *ManageRecurringPaymentsProfileStatus {
	svc.Action = "Cancel"
	svc.Note = desc
	return svc
}

// SetAsSuspend sets action as `Suspend`
func (svc *ManageRecurringPaymentsProfileStatus) SetAsSuspend(desc string) *ManageRecurringPaymentsProfileStatus {
	svc.Action = "Suspend"
	svc.Note = desc
	return svc
}

// SetAsReactivate sets action as `Reactivate`
func (svc *ManageRecurringPaymentsProfileStatus) SetAsReactivate(desc string) *ManageRecurringPaymentsProfileStatus {
	svc.Action = "Reactivate"
	svc.Note = desc
	return svc
}

// Do executes ManageRecurringPaymentsProfileStatus operation
func (svc *ManageRecurringPaymentsProfileStatus) Do(cli client.Client) (*ManageRecurringPaymentsProfileStatusResponse, error) {
	const method = "ManageRecurringPaymentsProfileStatus"
	svc.BaseRequest.Method = method

	result := &ManageRecurringPaymentsProfileStatusResponse{}
	err := cli.Call(svc, result)
	return result, err
}

// ManageRecurringPaymentsProfileStatusResponse is struct for response of ManageRecurringPaymentsProfileStatus API
type ManageRecurringPaymentsProfileStatusResponse struct {
	client.BaseResponse `url:",squash"`

	// success
	ProfileID string `url:"PROFILEID"`
}

// IsSuccess checks the request is success or not
func (r *ManageRecurringPaymentsProfileStatusResponse) IsSuccess() bool {
	return r.IsRequestSuccess()
}

// IsOperationSuccess checks the request is success or not
func (r *ManageRecurringPaymentsProfileStatusResponse) IsOperationSuccess() bool {
	return r.IsRequestSuccess()
}

// Error returns error text
func (r *ManageRecurringPaymentsProfileStatusResponse) Error() string {
	return r.BaseResponse.Error()
}
