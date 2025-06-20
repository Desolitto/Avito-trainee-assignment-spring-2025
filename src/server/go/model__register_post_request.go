// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * backend service
 *
 * Сервис для управления ПВЗ и приемкой товаров
 *
 * API version: 1.0.0
 */

package openapi




type RegisterPostRequest struct {

	Email string `json:"email"`

	Password string `json:"password"`

	Role string `json:"role"`
}

// AssertRegisterPostRequestRequired checks if the required fields are not zero-ed
func AssertRegisterPostRequestRequired(obj RegisterPostRequest) error {
	elements := map[string]interface{}{
		"email": obj.Email,
		"password": obj.Password,
		"role": obj.Role,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertRegisterPostRequestConstraints checks if the values respects the defined constraints
func AssertRegisterPostRequestConstraints(obj RegisterPostRequest) error {
	return nil
}
