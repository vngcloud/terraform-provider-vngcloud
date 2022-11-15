# UpdateL7PolicyRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | **string** | Action. This indicates how the listener will route traffic. The value can be REDIRECT_TO_POOL or REDIRECT_TO_URL. | [default to null]
**CompareType** | **string** | Compare operation. The value can be CONTAINS or EQUAL_TO | [default to null]
**KeepQueryString** | **bool** | Keep the query string or not. | [optional] [default to null]
**L7policyId** | **string** | Policy&#39;s id to be updated. | [default to null]
**Position** | **int64** | Position of the policy | [optional] [default to null]
**RedirectHttpCode** | **int32** | Redirect HTTP code for redirecting to other URL. | [optional] [default to null]
**RedirectPoolId** | **string** | Pool for forwarding. | [optional] [default to null]
**RedirectUrl** | **string** | URL for forwarding. | [optional] [default to null]
**Type_** | **string** | Which attribute to compare. The value can be PATH or HOST_NAME | [default to null]
**Value** | **string** | The value to compare with attribute. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


