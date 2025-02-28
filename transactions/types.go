package transactions

type Customer struct {
	ID string `json:"id"`
}

type UserField struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type Setting struct {
	Name  string `json:"settingName"`
	Value string `json:"settingValue"`
}

type CodeText struct {
	Code string `json:"code"`
	Text string `json:"text"`
}

type CodeDescription struct {
	Code        string `json:"code"`
	Description string `json:"description"`
}

type ProcessingOptions struct {
	IsSubsequentAuth string `json:"isSubsequentAuth"`
}

type SubsequentAuthInformation struct {
	OriginalNetworkTransID string `json:"originalNetworkTransId"`
	OriginalAuthAmount     string `json:"originalAuthAmount"`
	Reason                 string `json:"reason"`
}

type AuthorizationIndicatorType struct {
	AuthorizationIndicator string `json:"authorizationIndicator"`
}

type AddressInfo struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Company   string `json:"company"`
	Address   string `json:"address"`
	City      string `json:"city"`
	State     string `json:"state"`
	Zip       string `json:"zip"`
	Country   string `json:"country"`
}

type Fee struct {
	Amount      string `json:"amount"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type CreditCard struct {
	CardNumber     string `json:"cardNumber"`
	ExpirationDate string `json:"expirationDate"`
	CardCode       string `json:"cardCode"`
}

type Payment struct {
	CreditCard *CreditCard `json:"creditCard,omitempty"`
}

type LineItem struct {
	ItemID      string `json:"itemId"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Quantity    string `json:"quantity"`
	UnitPrice   string `json:"unitPrice"`
}

type TransactionType string

const (
	TransactionTypeAuthCapture TransactionType = "authCaptureTransaction"
)

type TransactionRequest struct {
	TransactionType TransactionType     `json:"transactionType"`
	Amount          string              `json:"amount"`
	Payment         Payment             `json:"payment"`
	LineItems       map[string]LineItem `json:"lineItems"`
	Tax             *Fee                `json:"tax,omitempty"`
	Duty            *Fee                `json:"duty,omitempty"`
	Shipping        *Fee                `json:"shipping,omitempty"`

	PONumber string      `json:"poNumber"`
	Customer Customer    `json:"customer"`
	BillTo   AddressInfo `json:"billTo"`
	ShipTo   AddressInfo `json:"shipTo"`

	CustomerIP                 string                     `json:"customerIP"`
	TransactionSettings        map[string]Setting         `json:"transactionSettings"`
	UserFields                 map[string][]UserField     `json:"userFields"`
	ProcessingOptions          ProcessingOptions          `json:"processingOptions"`
	SubsequentAuthInformation  *SubsequentAuthInformation `json:"subsequentAuthInformation,omitempty"`
	AuthorizationIndicatorType AuthorizationIndicatorType `json:"authorizationIndicatorType"`
}

type TransactionResponse struct {
	ResponseCode   string            `json:"responseCode"`
	AuthCode       string            `json:"authCode"`
	AVSResultCode  string            `json:"avsResultCode"`
	CVVResultCode  string            `json:"cvvResultCode"`
	CAVVResultCode string            `json:"cavvResultCode"`
	TransID        string            `json:"transId"`
	RefTransID     string            `json:"refTransID"`
	TransHash      string            `json:"transHash"`
	TestRequest    string            `json:"testRequest"`
	AccountNumber  string            `json:"accountNumber"`
	AccountType    string            `json:"accountType"`
	Messages       []CodeDescription `json:"messages"`

	UserFields                             []UserField `json:"userFields"`
	TransHashSha2                          string      `json:"transHashSha2"`
	SupplementalDataQualificationIndicator int         `json:"SupplementalDataQualificationIndicator"`
	NetworkTransID                         string      `json:"networkTransId"`
}

type CreateTransactionResponseMessages struct {
	ResultCode string     `json:"resultCode"`
	Message    []CodeText `json:"message"`
}
