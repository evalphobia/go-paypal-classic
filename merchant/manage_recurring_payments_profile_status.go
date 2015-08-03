package merchant

// ManageRecurringPaymentsProfileStatus is struct for ManageRecurringPaymentsProfileStatus API
// see: https://developer.paypal.com/docs/classic/api/merchant/ManageRecurringPaymentsProfileStatus_API_Operation_NVP/
type ManageRecurringPaymentsProfileStatus struct {
	Merchant    `url:",squash"`
	BaseRequest `url:",squash"`

	ProfileID string `url:"PROFILEID"`
	Action    string `url:"ACTION"`
	Note      string `url:"NOTE"`
}

// SetMerchant sets Merchant
func (svc *ManageRecurringPaymentsProfileStatus) SetMerchant(m Merchant) {
	svc.Merchant = m
}

// SetAsCancel sets action as `Cancel`
func (svc *ManageRecurringPaymentsProfileStatus) SetAsCancel(desc string) *ManageRecurringPaymentsProfileStatus {
	svc.Action = "Cancel"
	svc.Note = desc
	return svc
}

// SetAsSuspend		 sets action as `Suspend`
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
func (svc *ManageRecurringPaymentsProfileStatus) Do(m Merchant) (*ManageRecurringPaymentsProfileStatusResponse, error) {
	const method = "ManageRecurringPaymentsProfileStatus"
	svc.BaseRequest.Method = method

	result := &ManageRecurringPaymentsProfileStatusResponse{}
	err := m.call(svc, result)
	return result, err
}

// ManageRecurringPaymentsProfileStatusResponse is struct for response of ManageRecurringPaymentsProfileStatus API
type ManageRecurringPaymentsProfileStatusResponse struct {
	BaseResponse `url:",squash"`

	// success
	ProfileID string `url:"PROFILEID"`
}

// IsSuccess checks the request is success or not
func (r *ManageRecurringPaymentsProfileStatusResponse) IsSuccess() bool {
	return r.ACK == ackSuccess
}

// Error returns error text
func (r *ManageRecurringPaymentsProfileStatusResponse) Error() string {
	return r.BaseResponse.Error()
}
