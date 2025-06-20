// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * backend service
 *
 * Сервис для управления ПВЗ и приемкой товаров
 *
 * API version: 1.0.0
 */

package openapi




type PvzGet200ResponseInnerReceptionsInner struct {

	Reception Reception `json:"reception,omitempty"`

	Products []Product `json:"products,omitempty"`
}

// AssertPvzGet200ResponseInnerReceptionsInnerRequired checks if the required fields are not zero-ed
func AssertPvzGet200ResponseInnerReceptionsInnerRequired(obj PvzGet200ResponseInnerReceptionsInner) error {
	if err := AssertReceptionRequired(obj.Reception); err != nil {
		return err
	}
	for _, el := range obj.Products {
		if err := AssertProductRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertPvzGet200ResponseInnerReceptionsInnerConstraints checks if the values respects the defined constraints
func AssertPvzGet200ResponseInnerReceptionsInnerConstraints(obj PvzGet200ResponseInnerReceptionsInner) error {
	if err := AssertReceptionConstraints(obj.Reception); err != nil {
		return err
	}
	for _, el := range obj.Products {
		if err := AssertProductConstraints(el); err != nil {
			return err
		}
	}
	return nil
}
