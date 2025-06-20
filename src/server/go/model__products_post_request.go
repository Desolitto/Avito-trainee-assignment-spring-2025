// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * backend service
 *
 * Сервис для управления ПВЗ и приемкой товаров
 *
 * API version: 1.0.0
 */

package openapi




type ProductsPostRequest struct {

	Type string `json:"type"`

	PvzId string `json:"pvzId"`
}

// AssertProductsPostRequestRequired checks if the required fields are not zero-ed
func AssertProductsPostRequestRequired(obj ProductsPostRequest) error {
	elements := map[string]interface{}{
		"type": obj.Type,
		"pvzId": obj.PvzId,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertProductsPostRequestConstraints checks if the values respects the defined constraints
func AssertProductsPostRequestConstraints(obj ProductsPostRequest) error {
	return nil
}
