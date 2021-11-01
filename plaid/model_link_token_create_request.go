/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.44.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// LinkTokenCreateRequest LinkTokenCreateRequest defines the request schema for `/link/token/create`
type LinkTokenCreateRequest struct {
	// Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.
	ClientId *string `json:"client_id,omitempty"`
	// Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.
	Secret *string `json:"secret,omitempty"`
	// The name of your application, as it should be displayed in Link. Maximum length of 30 characters. If a value longer than 30 characters is provided, Link will display \"This Application\" instead.
	ClientName string `json:"client_name"`
	// The language that Link should be displayed in.  Supported languages are: - English (`'en'`) - French (`'fr'`) - Spanish (`'es'`) - Dutch (`'nl'`)  When using a Link customization, the language configured here must match the setting in the customization, or the customization will not be applied.
	Language string `json:"language"`
	// Specify an array of Plaid-supported country codes using the ISO-3166-1 alpha-2 country code standard. Institutions from all listed countries will be shown.  Supported country codes are: `US`, `CA`, `ES`, `FR`, `GB`, `IE`, `NL`.  If Link is launched with multiple country codes, only products that you are enabled for in all countries will be used by Link. Note that while all countries are enabled by default in Sandbox and Development, in Production only US and Canada are enabled by default. To gain access to European institutions in the Production environment, [file a product access Support ticket](https://dashboard.plaid.com/support/new/product-and-development/product-troubleshooting/request-product-access) via the Plaid dashboard. If you initialize with a European country code, your users will see the European consent panel during the Link flow.  If using a Link customization, make sure the country codes in the customization match those specified in `country_codes`. If both `country_codes` and a Link customization are used, the value in `country_codes` may override the value in the customization.  If using the Auth features Instant Match, Same-day Micro-deposits, or Automated Micro-deposits, `country_codes` must be set to `['US']`.
	CountryCodes []CountryCode `json:"country_codes"`
	User LinkTokenCreateRequestUser `json:"user"`
	// List of Plaid product(s) you wish to use. If launching Link in update mode, should be omitted; required otherwise. Valid products are:  `transactions`, `auth`, `identity`, `assets`, `investments`, `liabilities`, `payment_initiation`, `deposit_switch`, `income_verification`, `transfer`  `balance` is *not* a valid value, the Balance product does not require explicit initalization and will automatically be initialized when any other product is initialized.  Only institutions that support *all* requested products will be shown in Link; to maximize the number of institutions listed, it is recommended to initialize Link with the minimal product set required for your use case. Additional products can be added after Link initialization by calling the relevant endpoints. For details and exceptions, see [Choosing when to initialize products](https://plaid.com/docs/link/best-practices/#choosing-when-to-initialize-products).  Note that, unless you have opted to disable Instant Match support, institutions that support Instant Match will also be shown in Link if `auth` is specified as a product, even though these institutions do not contain `auth` in their product array.  In Production, you will be billed for each product that you specify when initializing Link. Note that a product cannot be removed from an Item once the Item has been initialized with that product. To stop billing on an Item for subscription-based products, such as Liabilities, Investments, and Transactions, remove the Item via `/item/remove`.
	Products *[]Products `json:"products,omitempty"`
	// The destination URL to which any webhooks should be sent.
	Webhook *string `json:"webhook,omitempty"`
	// The `access_token` associated with the Item to update, used when updating or modifying an existing `access_token`. Used when launching Link in update mode, when completing the Same-day (manual) Micro-deposit flow, or (optionally) when initializing Link as part of the Payment Initiation (UK and Europe) flow.
	AccessToken *string `json:"access_token,omitempty"`
	// The name of the Link customization from the Plaid Dashboard to be applied to Link. If not specified, the `default` customization will be used. When using a Link customization, the language in the customization must match the language selected via the `language` parameter, and the countries in the customization should match the country codes selected via `country_codes`.
	LinkCustomizationName *string `json:"link_customization_name,omitempty"`
	// A URI indicating the destination where a user should be forwarded after completing the Link flow; used to support OAuth authentication flows when launching Link in the browser or via a webview. The `redirect_uri` should not contain any query parameters. When used in Production or Development, must be an https URI. To specify any subdomain, use `*` as a wildcard character, e.g. `https://_*.example.com/oauth.html`. If `android_package_name` is specified, this field should be left blank.  Note that any redirect URI must also be added to the Allowed redirect URIs list in the [developer dashboard](https://dashboard.plaid.com/team/api).
	RedirectUri *string `json:"redirect_uri,omitempty"`
	// The name of your app's Android package. Required if using the `link_token` to initialize Link on Android. When creating a `link_token` for initializing Link on other platforms, this field must be left blank. Any package name specified here must also be added to the Allowed Android package names setting on the [developer dashboard](https://dashboard.plaid.com/team/api). 
	AndroidPackageName *string `json:"android_package_name,omitempty"`
	AccountFilters *LinkTokenAccountFilters `json:"account_filters,omitempty"`
	EuConfig *LinkTokenEUConfig `json:"eu_config,omitempty"`
	// Used for certain Europe-only configurations, as well as certain legacy use cases in other regions.
	InstitutionId *string `json:"institution_id,omitempty"`
	PaymentInitiation *LinkTokenCreateRequestPaymentInitiation `json:"payment_initiation,omitempty"`
	DepositSwitch *LinkTokenCreateRequestDepositSwitch `json:"deposit_switch,omitempty"`
	IncomeVerification *LinkTokenCreateRequestIncomeVerification `json:"income_verification,omitempty"`
	Auth *LinkTokenCreateRequestAuth `json:"auth,omitempty"`
	Update *LinkTokenCreateRequestUpdate `json:"update,omitempty"`
}

// NewLinkTokenCreateRequest instantiates a new LinkTokenCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLinkTokenCreateRequest(clientName string, language string, countryCodes []CountryCode, user LinkTokenCreateRequestUser) *LinkTokenCreateRequest {
	this := LinkTokenCreateRequest{}
	this.ClientName = clientName
	this.Language = language
	this.CountryCodes = countryCodes
	this.User = user
	return &this
}

// NewLinkTokenCreateRequestWithDefaults instantiates a new LinkTokenCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLinkTokenCreateRequestWithDefaults() *LinkTokenCreateRequest {
	this := LinkTokenCreateRequest{}
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *LinkTokenCreateRequest) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkTokenCreateRequest) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *LinkTokenCreateRequest) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *LinkTokenCreateRequest) SetClientId(v string) {
	o.ClientId = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *LinkTokenCreateRequest) GetSecret() string {
	if o == nil || o.Secret == nil {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkTokenCreateRequest) GetSecretOk() (*string, bool) {
	if o == nil || o.Secret == nil {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *LinkTokenCreateRequest) HasSecret() bool {
	if o != nil && o.Secret != nil {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *LinkTokenCreateRequest) SetSecret(v string) {
	o.Secret = &v
}

// GetClientName returns the ClientName field value
func (o *LinkTokenCreateRequest) GetClientName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientName
}

// GetClientNameOk returns a tuple with the ClientName field value
// and a boolean to check if the value has been set.
func (o *LinkTokenCreateRequest) GetClientNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ClientName, true
}

// SetClientName sets field value
func (o *LinkTokenCreateRequest) SetClientName(v string) {
	o.ClientName = v
}

// GetLanguage returns the Language field value
func (o *LinkTokenCreateRequest) GetLanguage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Language
}

// GetLanguageOk returns a tuple with the Language field value
// and a boolean to check if the value has been set.
func (o *LinkTokenCreateRequest) GetLanguageOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Language, true
}

// SetLanguage sets field value
func (o *LinkTokenCreateRequest) SetLanguage(v string) {
	o.Language = v
}

// GetCountryCodes returns the CountryCodes field value
func (o *LinkTokenCreateRequest) GetCountryCodes() []CountryCode {
	if o == nil {
		var ret []CountryCode
		return ret
	}

	return o.CountryCodes
}

// GetCountryCodesOk returns a tuple with the CountryCodes field value
// and a boolean to check if the value has been set.
func (o *LinkTokenCreateRequest) GetCountryCodesOk() (*[]CountryCode, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CountryCodes, true
}

// SetCountryCodes sets field value
func (o *LinkTokenCreateRequest) SetCountryCodes(v []CountryCode) {
	o.CountryCodes = v
}

// GetUser returns the User field value
func (o *LinkTokenCreateRequest) GetUser() LinkTokenCreateRequestUser {
	if o == nil {
		var ret LinkTokenCreateRequestUser
		return ret
	}

	return o.User
}

// GetUserOk returns a tuple with the User field value
// and a boolean to check if the value has been set.
func (o *LinkTokenCreateRequest) GetUserOk() (*LinkTokenCreateRequestUser, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.User, true
}

// SetUser sets field value
func (o *LinkTokenCreateRequest) SetUser(v LinkTokenCreateRequestUser) {
	o.User = v
}

// GetProducts returns the Products field value if set, zero value otherwise.
func (o *LinkTokenCreateRequest) GetProducts() []Products {
	if o == nil || o.Products == nil {
		var ret []Products
		return ret
	}
	return *o.Products
}

// GetProductsOk returns a tuple with the Products field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkTokenCreateRequest) GetProductsOk() (*[]Products, bool) {
	if o == nil || o.Products == nil {
		return nil, false
	}
	return o.Products, true
}

// HasProducts returns a boolean if a field has been set.
func (o *LinkTokenCreateRequest) HasProducts() bool {
	if o != nil && o.Products != nil {
		return true
	}

	return false
}

// SetProducts gets a reference to the given []Products and assigns it to the Products field.
func (o *LinkTokenCreateRequest) SetProducts(v []Products) {
	o.Products = &v
}

// GetWebhook returns the Webhook field value if set, zero value otherwise.
func (o *LinkTokenCreateRequest) GetWebhook() string {
	if o == nil || o.Webhook == nil {
		var ret string
		return ret
	}
	return *o.Webhook
}

// GetWebhookOk returns a tuple with the Webhook field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkTokenCreateRequest) GetWebhookOk() (*string, bool) {
	if o == nil || o.Webhook == nil {
		return nil, false
	}
	return o.Webhook, true
}

// HasWebhook returns a boolean if a field has been set.
func (o *LinkTokenCreateRequest) HasWebhook() bool {
	if o != nil && o.Webhook != nil {
		return true
	}

	return false
}

// SetWebhook gets a reference to the given string and assigns it to the Webhook field.
func (o *LinkTokenCreateRequest) SetWebhook(v string) {
	o.Webhook = &v
}

// GetAccessToken returns the AccessToken field value if set, zero value otherwise.
func (o *LinkTokenCreateRequest) GetAccessToken() string {
	if o == nil || o.AccessToken == nil {
		var ret string
		return ret
	}
	return *o.AccessToken
}

// GetAccessTokenOk returns a tuple with the AccessToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkTokenCreateRequest) GetAccessTokenOk() (*string, bool) {
	if o == nil || o.AccessToken == nil {
		return nil, false
	}
	return o.AccessToken, true
}

// HasAccessToken returns a boolean if a field has been set.
func (o *LinkTokenCreateRequest) HasAccessToken() bool {
	if o != nil && o.AccessToken != nil {
		return true
	}

	return false
}

// SetAccessToken gets a reference to the given string and assigns it to the AccessToken field.
func (o *LinkTokenCreateRequest) SetAccessToken(v string) {
	o.AccessToken = &v
}

// GetLinkCustomizationName returns the LinkCustomizationName field value if set, zero value otherwise.
func (o *LinkTokenCreateRequest) GetLinkCustomizationName() string {
	if o == nil || o.LinkCustomizationName == nil {
		var ret string
		return ret
	}
	return *o.LinkCustomizationName
}

// GetLinkCustomizationNameOk returns a tuple with the LinkCustomizationName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkTokenCreateRequest) GetLinkCustomizationNameOk() (*string, bool) {
	if o == nil || o.LinkCustomizationName == nil {
		return nil, false
	}
	return o.LinkCustomizationName, true
}

// HasLinkCustomizationName returns a boolean if a field has been set.
func (o *LinkTokenCreateRequest) HasLinkCustomizationName() bool {
	if o != nil && o.LinkCustomizationName != nil {
		return true
	}

	return false
}

// SetLinkCustomizationName gets a reference to the given string and assigns it to the LinkCustomizationName field.
func (o *LinkTokenCreateRequest) SetLinkCustomizationName(v string) {
	o.LinkCustomizationName = &v
}

// GetRedirectUri returns the RedirectUri field value if set, zero value otherwise.
func (o *LinkTokenCreateRequest) GetRedirectUri() string {
	if o == nil || o.RedirectUri == nil {
		var ret string
		return ret
	}
	return *o.RedirectUri
}

// GetRedirectUriOk returns a tuple with the RedirectUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkTokenCreateRequest) GetRedirectUriOk() (*string, bool) {
	if o == nil || o.RedirectUri == nil {
		return nil, false
	}
	return o.RedirectUri, true
}

// HasRedirectUri returns a boolean if a field has been set.
func (o *LinkTokenCreateRequest) HasRedirectUri() bool {
	if o != nil && o.RedirectUri != nil {
		return true
	}

	return false
}

// SetRedirectUri gets a reference to the given string and assigns it to the RedirectUri field.
func (o *LinkTokenCreateRequest) SetRedirectUri(v string) {
	o.RedirectUri = &v
}

// GetAndroidPackageName returns the AndroidPackageName field value if set, zero value otherwise.
func (o *LinkTokenCreateRequest) GetAndroidPackageName() string {
	if o == nil || o.AndroidPackageName == nil {
		var ret string
		return ret
	}
	return *o.AndroidPackageName
}

// GetAndroidPackageNameOk returns a tuple with the AndroidPackageName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkTokenCreateRequest) GetAndroidPackageNameOk() (*string, bool) {
	if o == nil || o.AndroidPackageName == nil {
		return nil, false
	}
	return o.AndroidPackageName, true
}

// HasAndroidPackageName returns a boolean if a field has been set.
func (o *LinkTokenCreateRequest) HasAndroidPackageName() bool {
	if o != nil && o.AndroidPackageName != nil {
		return true
	}

	return false
}

// SetAndroidPackageName gets a reference to the given string and assigns it to the AndroidPackageName field.
func (o *LinkTokenCreateRequest) SetAndroidPackageName(v string) {
	o.AndroidPackageName = &v
}

// GetAccountFilters returns the AccountFilters field value if set, zero value otherwise.
func (o *LinkTokenCreateRequest) GetAccountFilters() LinkTokenAccountFilters {
	if o == nil || o.AccountFilters == nil {
		var ret LinkTokenAccountFilters
		return ret
	}
	return *o.AccountFilters
}

// GetAccountFiltersOk returns a tuple with the AccountFilters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkTokenCreateRequest) GetAccountFiltersOk() (*LinkTokenAccountFilters, bool) {
	if o == nil || o.AccountFilters == nil {
		return nil, false
	}
	return o.AccountFilters, true
}

// HasAccountFilters returns a boolean if a field has been set.
func (o *LinkTokenCreateRequest) HasAccountFilters() bool {
	if o != nil && o.AccountFilters != nil {
		return true
	}

	return false
}

// SetAccountFilters gets a reference to the given LinkTokenAccountFilters and assigns it to the AccountFilters field.
func (o *LinkTokenCreateRequest) SetAccountFilters(v LinkTokenAccountFilters) {
	o.AccountFilters = &v
}

// GetEuConfig returns the EuConfig field value if set, zero value otherwise.
func (o *LinkTokenCreateRequest) GetEuConfig() LinkTokenEUConfig {
	if o == nil || o.EuConfig == nil {
		var ret LinkTokenEUConfig
		return ret
	}
	return *o.EuConfig
}

// GetEuConfigOk returns a tuple with the EuConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkTokenCreateRequest) GetEuConfigOk() (*LinkTokenEUConfig, bool) {
	if o == nil || o.EuConfig == nil {
		return nil, false
	}
	return o.EuConfig, true
}

// HasEuConfig returns a boolean if a field has been set.
func (o *LinkTokenCreateRequest) HasEuConfig() bool {
	if o != nil && o.EuConfig != nil {
		return true
	}

	return false
}

// SetEuConfig gets a reference to the given LinkTokenEUConfig and assigns it to the EuConfig field.
func (o *LinkTokenCreateRequest) SetEuConfig(v LinkTokenEUConfig) {
	o.EuConfig = &v
}

// GetInstitutionId returns the InstitutionId field value if set, zero value otherwise.
func (o *LinkTokenCreateRequest) GetInstitutionId() string {
	if o == nil || o.InstitutionId == nil {
		var ret string
		return ret
	}
	return *o.InstitutionId
}

// GetInstitutionIdOk returns a tuple with the InstitutionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkTokenCreateRequest) GetInstitutionIdOk() (*string, bool) {
	if o == nil || o.InstitutionId == nil {
		return nil, false
	}
	return o.InstitutionId, true
}

// HasInstitutionId returns a boolean if a field has been set.
func (o *LinkTokenCreateRequest) HasInstitutionId() bool {
	if o != nil && o.InstitutionId != nil {
		return true
	}

	return false
}

// SetInstitutionId gets a reference to the given string and assigns it to the InstitutionId field.
func (o *LinkTokenCreateRequest) SetInstitutionId(v string) {
	o.InstitutionId = &v
}

// GetPaymentInitiation returns the PaymentInitiation field value if set, zero value otherwise.
func (o *LinkTokenCreateRequest) GetPaymentInitiation() LinkTokenCreateRequestPaymentInitiation {
	if o == nil || o.PaymentInitiation == nil {
		var ret LinkTokenCreateRequestPaymentInitiation
		return ret
	}
	return *o.PaymentInitiation
}

// GetPaymentInitiationOk returns a tuple with the PaymentInitiation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkTokenCreateRequest) GetPaymentInitiationOk() (*LinkTokenCreateRequestPaymentInitiation, bool) {
	if o == nil || o.PaymentInitiation == nil {
		return nil, false
	}
	return o.PaymentInitiation, true
}

// HasPaymentInitiation returns a boolean if a field has been set.
func (o *LinkTokenCreateRequest) HasPaymentInitiation() bool {
	if o != nil && o.PaymentInitiation != nil {
		return true
	}

	return false
}

// SetPaymentInitiation gets a reference to the given LinkTokenCreateRequestPaymentInitiation and assigns it to the PaymentInitiation field.
func (o *LinkTokenCreateRequest) SetPaymentInitiation(v LinkTokenCreateRequestPaymentInitiation) {
	o.PaymentInitiation = &v
}

// GetDepositSwitch returns the DepositSwitch field value if set, zero value otherwise.
func (o *LinkTokenCreateRequest) GetDepositSwitch() LinkTokenCreateRequestDepositSwitch {
	if o == nil || o.DepositSwitch == nil {
		var ret LinkTokenCreateRequestDepositSwitch
		return ret
	}
	return *o.DepositSwitch
}

// GetDepositSwitchOk returns a tuple with the DepositSwitch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkTokenCreateRequest) GetDepositSwitchOk() (*LinkTokenCreateRequestDepositSwitch, bool) {
	if o == nil || o.DepositSwitch == nil {
		return nil, false
	}
	return o.DepositSwitch, true
}

// HasDepositSwitch returns a boolean if a field has been set.
func (o *LinkTokenCreateRequest) HasDepositSwitch() bool {
	if o != nil && o.DepositSwitch != nil {
		return true
	}

	return false
}

// SetDepositSwitch gets a reference to the given LinkTokenCreateRequestDepositSwitch and assigns it to the DepositSwitch field.
func (o *LinkTokenCreateRequest) SetDepositSwitch(v LinkTokenCreateRequestDepositSwitch) {
	o.DepositSwitch = &v
}

// GetIncomeVerification returns the IncomeVerification field value if set, zero value otherwise.
func (o *LinkTokenCreateRequest) GetIncomeVerification() LinkTokenCreateRequestIncomeVerification {
	if o == nil || o.IncomeVerification == nil {
		var ret LinkTokenCreateRequestIncomeVerification
		return ret
	}
	return *o.IncomeVerification
}

// GetIncomeVerificationOk returns a tuple with the IncomeVerification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkTokenCreateRequest) GetIncomeVerificationOk() (*LinkTokenCreateRequestIncomeVerification, bool) {
	if o == nil || o.IncomeVerification == nil {
		return nil, false
	}
	return o.IncomeVerification, true
}

// HasIncomeVerification returns a boolean if a field has been set.
func (o *LinkTokenCreateRequest) HasIncomeVerification() bool {
	if o != nil && o.IncomeVerification != nil {
		return true
	}

	return false
}

// SetIncomeVerification gets a reference to the given LinkTokenCreateRequestIncomeVerification and assigns it to the IncomeVerification field.
func (o *LinkTokenCreateRequest) SetIncomeVerification(v LinkTokenCreateRequestIncomeVerification) {
	o.IncomeVerification = &v
}

// GetAuth returns the Auth field value if set, zero value otherwise.
func (o *LinkTokenCreateRequest) GetAuth() LinkTokenCreateRequestAuth {
	if o == nil || o.Auth == nil {
		var ret LinkTokenCreateRequestAuth
		return ret
	}
	return *o.Auth
}

// GetAuthOk returns a tuple with the Auth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkTokenCreateRequest) GetAuthOk() (*LinkTokenCreateRequestAuth, bool) {
	if o == nil || o.Auth == nil {
		return nil, false
	}
	return o.Auth, true
}

// HasAuth returns a boolean if a field has been set.
func (o *LinkTokenCreateRequest) HasAuth() bool {
	if o != nil && o.Auth != nil {
		return true
	}

	return false
}

// SetAuth gets a reference to the given LinkTokenCreateRequestAuth and assigns it to the Auth field.
func (o *LinkTokenCreateRequest) SetAuth(v LinkTokenCreateRequestAuth) {
	o.Auth = &v
}

// GetUpdate returns the Update field value if set, zero value otherwise.
func (o *LinkTokenCreateRequest) GetUpdate() LinkTokenCreateRequestUpdate {
	if o == nil || o.Update == nil {
		var ret LinkTokenCreateRequestUpdate
		return ret
	}
	return *o.Update
}

// GetUpdateOk returns a tuple with the Update field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkTokenCreateRequest) GetUpdateOk() (*LinkTokenCreateRequestUpdate, bool) {
	if o == nil || o.Update == nil {
		return nil, false
	}
	return o.Update, true
}

// HasUpdate returns a boolean if a field has been set.
func (o *LinkTokenCreateRequest) HasUpdate() bool {
	if o != nil && o.Update != nil {
		return true
	}

	return false
}

// SetUpdate gets a reference to the given LinkTokenCreateRequestUpdate and assigns it to the Update field.
func (o *LinkTokenCreateRequest) SetUpdate(v LinkTokenCreateRequestUpdate) {
	o.Update = &v
}

func (o LinkTokenCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientId != nil {
		toSerialize["client_id"] = o.ClientId
	}
	if o.Secret != nil {
		toSerialize["secret"] = o.Secret
	}
	if true {
		toSerialize["client_name"] = o.ClientName
	}
	if true {
		toSerialize["language"] = o.Language
	}
	if true {
		toSerialize["country_codes"] = o.CountryCodes
	}
	if true {
		toSerialize["user"] = o.User
	}
	if o.Products != nil {
		toSerialize["products"] = o.Products
	}
	if o.Webhook != nil {
		toSerialize["webhook"] = o.Webhook
	}
	if o.AccessToken != nil {
		toSerialize["access_token"] = o.AccessToken
	}
	if o.LinkCustomizationName != nil {
		toSerialize["link_customization_name"] = o.LinkCustomizationName
	}
	if o.RedirectUri != nil {
		toSerialize["redirect_uri"] = o.RedirectUri
	}
	if o.AndroidPackageName != nil {
		toSerialize["android_package_name"] = o.AndroidPackageName
	}
	if o.AccountFilters != nil {
		toSerialize["account_filters"] = o.AccountFilters
	}
	if o.EuConfig != nil {
		toSerialize["eu_config"] = o.EuConfig
	}
	if o.InstitutionId != nil {
		toSerialize["institution_id"] = o.InstitutionId
	}
	if o.PaymentInitiation != nil {
		toSerialize["payment_initiation"] = o.PaymentInitiation
	}
	if o.DepositSwitch != nil {
		toSerialize["deposit_switch"] = o.DepositSwitch
	}
	if o.IncomeVerification != nil {
		toSerialize["income_verification"] = o.IncomeVerification
	}
	if o.Auth != nil {
		toSerialize["auth"] = o.Auth
	}
	if o.Update != nil {
		toSerialize["update"] = o.Update
	}
	return json.Marshal(toSerialize)
}

type NullableLinkTokenCreateRequest struct {
	value *LinkTokenCreateRequest
	isSet bool
}

func (v NullableLinkTokenCreateRequest) Get() *LinkTokenCreateRequest {
	return v.value
}

func (v *NullableLinkTokenCreateRequest) Set(val *LinkTokenCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableLinkTokenCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableLinkTokenCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinkTokenCreateRequest(val *LinkTokenCreateRequest) *NullableLinkTokenCreateRequest {
	return &NullableLinkTokenCreateRequest{value: val, isSet: true}
}

func (v NullableLinkTokenCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinkTokenCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


