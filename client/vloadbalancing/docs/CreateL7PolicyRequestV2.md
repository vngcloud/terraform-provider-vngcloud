# CreateL7PolicyRequestV2

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | **string** | Action. This indicates how the listener will route traffic. The value can be REDIRECT_TO_POOL or REDIRECT_TO_URL. | [default to null]
**CompareType** | **string** | Compare operation. The value can be CONTAINS or EQUAL_TO | [default to null]
**KeepQueryString** | **bool** | Keep the query string or not. | [optional] [default to null]
**Name** | **string** | Policy name. Only letters (a-z, A-Z, 0-9, &#39;_&#39;, &#39;.&#39;) are allowed and your input data must be between 6 and 20 characters. | [default to null]
**RedirectHttpCode** | **int32** | Redirect HTTP code for redirecting to other URL. | [optional] [default to null]
**RedirectPoolId** | **string** | Pool for forwarding. | [optional] [default to null]
**RedirectUrl** | **string** | URL for forwarding. | [optional] [default to null]
**Type_** | **string** | Which attribute to compare. The value can be PATH or HOST_NAME | [default to null]
**Value** | **string** | The value to compare with attribute. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


