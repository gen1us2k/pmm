// Code generated by go-swagger; DO NOT EDIT.

package nodes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GetNodeReader is a Reader for the GetNode structure.
type GetNodeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNodeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNodeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetNodeDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetNodeOK creates a GetNodeOK with default headers values
func NewGetNodeOK() *GetNodeOK {
	return &GetNodeOK{}
}

/* GetNodeOK describes a response with status code 200, with default header values.

A successful response.
*/
type GetNodeOK struct {
	Payload *GetNodeOKBody
}

func (o *GetNodeOK) Error() string {
	return fmt.Sprintf("[POST /v1/inventory/Nodes/Get][%d] getNodeOk  %+v", 200, o.Payload)
}

func (o *GetNodeOK) GetPayload() *GetNodeOKBody {
	return o.Payload
}

func (o *GetNodeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(GetNodeOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNodeDefault creates a GetNodeDefault with default headers values
func NewGetNodeDefault(code int) *GetNodeDefault {
	return &GetNodeDefault{
		_statusCode: code,
	}
}

/* GetNodeDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type GetNodeDefault struct {
	_statusCode int

	Payload *GetNodeDefaultBody
}

// Code gets the status code for the get node default response
func (o *GetNodeDefault) Code() int {
	return o._statusCode
}

func (o *GetNodeDefault) Error() string {
	return fmt.Sprintf("[POST /v1/inventory/Nodes/Get][%d] GetNode default  %+v", o._statusCode, o.Payload)
}

func (o *GetNodeDefault) GetPayload() *GetNodeDefaultBody {
	return o.Payload
}

func (o *GetNodeDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(GetNodeDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetNodeBody get node body
swagger:model GetNodeBody
*/
type GetNodeBody struct {
	// Unique randomly generated instance identifier.
	NodeID string `json:"node_id,omitempty"`
}

// Validate validates this get node body
func (o *GetNodeBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get node body based on context it is used
func (o *GetNodeBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetNodeBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetNodeBody) UnmarshalBinary(b []byte) error {
	var res GetNodeBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetNodeDefaultBody get node default body
swagger:model GetNodeDefaultBody
*/
type GetNodeDefaultBody struct {
	// code
	Code int32 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// details
	Details []*GetNodeDefaultBodyDetailsItems0 `json:"details"`
}

// Validate validates this get node default body
func (o *GetNodeDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetNodeDefaultBody) validateDetails(formats strfmt.Registry) error {
	if swag.IsZero(o.Details) { // not required
		return nil
	}

	for i := 0; i < len(o.Details); i++ {
		if swag.IsZero(o.Details[i]) { // not required
			continue
		}

		if o.Details[i] != nil {
			if err := o.Details[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("GetNode default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("GetNode default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this get node default body based on the context it is used
func (o *GetNodeDefaultBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetNodeDefaultBody) contextValidateDetails(ctx context.Context, formats strfmt.Registry) error {
	for i := 0; i < len(o.Details); i++ {
		if o.Details[i] != nil {
			if err := o.Details[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("GetNode default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("GetNode default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetNodeDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetNodeDefaultBody) UnmarshalBinary(b []byte) error {
	var res GetNodeDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetNodeDefaultBodyDetailsItems0 get node default body details items0
swagger:model GetNodeDefaultBodyDetailsItems0
*/
type GetNodeDefaultBodyDetailsItems0 struct {
	// at type
	AtType string `json:"@type,omitempty"`
}

// Validate validates this get node default body details items0
func (o *GetNodeDefaultBodyDetailsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get node default body details items0 based on context it is used
func (o *GetNodeDefaultBodyDetailsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetNodeDefaultBodyDetailsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetNodeDefaultBodyDetailsItems0) UnmarshalBinary(b []byte) error {
	var res GetNodeDefaultBodyDetailsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetNodeOKBody get node OK body
swagger:model GetNodeOKBody
*/
type GetNodeOKBody struct {
	// container
	Container *GetNodeOKBodyContainer `json:"container,omitempty"`

	// generic
	Generic *GetNodeOKBodyGeneric `json:"generic,omitempty"`

	// remote
	Remote *GetNodeOKBodyRemote `json:"remote,omitempty"`

	// remote azure database
	RemoteAzureDatabase *GetNodeOKBodyRemoteAzureDatabase `json:"remote_azure_database,omitempty"`

	// remote rds
	RemoteRDS *GetNodeOKBodyRemoteRDS `json:"remote_rds,omitempty"`
}

// Validate validates this get node OK body
func (o *GetNodeOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateContainer(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateGeneric(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateRemote(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateRemoteAzureDatabase(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateRemoteRDS(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetNodeOKBody) validateContainer(formats strfmt.Registry) error {
	if swag.IsZero(o.Container) { // not required
		return nil
	}

	if o.Container != nil {
		if err := o.Container.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getNodeOk" + "." + "container")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getNodeOk" + "." + "container")
			}
			return err
		}
	}

	return nil
}

func (o *GetNodeOKBody) validateGeneric(formats strfmt.Registry) error {
	if swag.IsZero(o.Generic) { // not required
		return nil
	}

	if o.Generic != nil {
		if err := o.Generic.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getNodeOk" + "." + "generic")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getNodeOk" + "." + "generic")
			}
			return err
		}
	}

	return nil
}

func (o *GetNodeOKBody) validateRemote(formats strfmt.Registry) error {
	if swag.IsZero(o.Remote) { // not required
		return nil
	}

	if o.Remote != nil {
		if err := o.Remote.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getNodeOk" + "." + "remote")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getNodeOk" + "." + "remote")
			}
			return err
		}
	}

	return nil
}

func (o *GetNodeOKBody) validateRemoteAzureDatabase(formats strfmt.Registry) error {
	if swag.IsZero(o.RemoteAzureDatabase) { // not required
		return nil
	}

	if o.RemoteAzureDatabase != nil {
		if err := o.RemoteAzureDatabase.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getNodeOk" + "." + "remote_azure_database")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getNodeOk" + "." + "remote_azure_database")
			}
			return err
		}
	}

	return nil
}

func (o *GetNodeOKBody) validateRemoteRDS(formats strfmt.Registry) error {
	if swag.IsZero(o.RemoteRDS) { // not required
		return nil
	}

	if o.RemoteRDS != nil {
		if err := o.RemoteRDS.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getNodeOk" + "." + "remote_rds")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getNodeOk" + "." + "remote_rds")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this get node OK body based on the context it is used
func (o *GetNodeOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateContainer(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateGeneric(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateRemote(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateRemoteAzureDatabase(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateRemoteRDS(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetNodeOKBody) contextValidateContainer(ctx context.Context, formats strfmt.Registry) error {
	if o.Container != nil {
		if err := o.Container.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getNodeOk" + "." + "container")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getNodeOk" + "." + "container")
			}
			return err
		}
	}

	return nil
}

func (o *GetNodeOKBody) contextValidateGeneric(ctx context.Context, formats strfmt.Registry) error {
	if o.Generic != nil {
		if err := o.Generic.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getNodeOk" + "." + "generic")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getNodeOk" + "." + "generic")
			}
			return err
		}
	}

	return nil
}

func (o *GetNodeOKBody) contextValidateRemote(ctx context.Context, formats strfmt.Registry) error {
	if o.Remote != nil {
		if err := o.Remote.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getNodeOk" + "." + "remote")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getNodeOk" + "." + "remote")
			}
			return err
		}
	}

	return nil
}

func (o *GetNodeOKBody) contextValidateRemoteAzureDatabase(ctx context.Context, formats strfmt.Registry) error {
	if o.RemoteAzureDatabase != nil {
		if err := o.RemoteAzureDatabase.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getNodeOk" + "." + "remote_azure_database")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getNodeOk" + "." + "remote_azure_database")
			}
			return err
		}
	}

	return nil
}

func (o *GetNodeOKBody) contextValidateRemoteRDS(ctx context.Context, formats strfmt.Registry) error {
	if o.RemoteRDS != nil {
		if err := o.RemoteRDS.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getNodeOk" + "." + "remote_rds")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getNodeOk" + "." + "remote_rds")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetNodeOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetNodeOKBody) UnmarshalBinary(b []byte) error {
	var res GetNodeOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetNodeOKBodyContainer ContainerNode represents a Docker container.
swagger:model GetNodeOKBodyContainer
*/
type GetNodeOKBodyContainer struct {
	// Unique randomly generated instance identifier.
	NodeID string `json:"node_id,omitempty"`

	// Unique across all Nodes user-defined name.
	NodeName string `json:"node_name,omitempty"`

	// Node address (DNS name or IP).
	Address string `json:"address,omitempty"`

	// Linux machine-id of the Generic Node where this Container Node runs.
	MachineID string `json:"machine_id,omitempty"`

	// Container identifier. If specified, must be a unique Docker container identifier.
	ContainerID string `json:"container_id,omitempty"`

	// Container name.
	ContainerName string `json:"container_name,omitempty"`

	// Node model.
	NodeModel string `json:"node_model,omitempty"`

	// Node region.
	Region string `json:"region,omitempty"`

	// Node availability zone.
	Az string `json:"az,omitempty"`

	// Custom user-assigned labels.
	CustomLabels map[string]string `json:"custom_labels,omitempty"`
}

// Validate validates this get node OK body container
func (o *GetNodeOKBodyContainer) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get node OK body container based on context it is used
func (o *GetNodeOKBodyContainer) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetNodeOKBodyContainer) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetNodeOKBodyContainer) UnmarshalBinary(b []byte) error {
	var res GetNodeOKBodyContainer
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetNodeOKBodyGeneric GenericNode represents a bare metal server or virtual machine.
swagger:model GetNodeOKBodyGeneric
*/
type GetNodeOKBodyGeneric struct {
	// Unique randomly generated instance identifier.
	NodeID string `json:"node_id,omitempty"`

	// Unique across all Nodes user-defined name.
	NodeName string `json:"node_name,omitempty"`

	// Node address (DNS name or IP).
	Address string `json:"address,omitempty"`

	// Linux machine-id.
	MachineID string `json:"machine_id,omitempty"`

	// Linux distribution name and version.
	Distro string `json:"distro,omitempty"`

	// Node model.
	NodeModel string `json:"node_model,omitempty"`

	// Node region.
	Region string `json:"region,omitempty"`

	// Node availability zone.
	Az string `json:"az,omitempty"`

	// Custom user-assigned labels.
	CustomLabels map[string]string `json:"custom_labels,omitempty"`
}

// Validate validates this get node OK body generic
func (o *GetNodeOKBodyGeneric) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get node OK body generic based on context it is used
func (o *GetNodeOKBodyGeneric) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetNodeOKBodyGeneric) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetNodeOKBodyGeneric) UnmarshalBinary(b []byte) error {
	var res GetNodeOKBodyGeneric
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetNodeOKBodyRemote RemoteNode represents generic remote Node. It's a node where we don't run pmm-agents. Only external exporters can run on Remote Nodes.
swagger:model GetNodeOKBodyRemote
*/
type GetNodeOKBodyRemote struct {
	// Unique randomly generated instance identifier.
	NodeID string `json:"node_id,omitempty"`

	// Unique across all Nodes user-defined name.
	NodeName string `json:"node_name,omitempty"`

	// Node address (DNS name or IP).
	Address string `json:"address,omitempty"`

	// Node model.
	NodeModel string `json:"node_model,omitempty"`

	// Node region.
	Region string `json:"region,omitempty"`

	// Node availability zone.
	Az string `json:"az,omitempty"`

	// Custom user-assigned labels.
	CustomLabels map[string]string `json:"custom_labels,omitempty"`
}

// Validate validates this get node OK body remote
func (o *GetNodeOKBodyRemote) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get node OK body remote based on context it is used
func (o *GetNodeOKBodyRemote) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetNodeOKBodyRemote) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetNodeOKBodyRemote) UnmarshalBinary(b []byte) error {
	var res GetNodeOKBodyRemote
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetNodeOKBodyRemoteAzureDatabase RemoteAzureDatabaseNode represents remote AzureDatabase Node. Agents can't run on Remote AzureDatabase Nodes.
swagger:model GetNodeOKBodyRemoteAzureDatabase
*/
type GetNodeOKBodyRemoteAzureDatabase struct {
	// Unique randomly generated instance identifier.
	NodeID string `json:"node_id,omitempty"`

	// Unique across all Nodes user-defined name.
	NodeName string `json:"node_name,omitempty"`

	// DB instance identifier.
	Address string `json:"address,omitempty"`

	// Node model.
	NodeModel string `json:"node_model,omitempty"`

	// Node region.
	Region string `json:"region,omitempty"`

	// Node availability zone.
	Az string `json:"az,omitempty"`

	// Custom user-assigned labels.
	CustomLabels map[string]string `json:"custom_labels,omitempty"`
}

// Validate validates this get node OK body remote azure database
func (o *GetNodeOKBodyRemoteAzureDatabase) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get node OK body remote azure database based on context it is used
func (o *GetNodeOKBodyRemoteAzureDatabase) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetNodeOKBodyRemoteAzureDatabase) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetNodeOKBodyRemoteAzureDatabase) UnmarshalBinary(b []byte) error {
	var res GetNodeOKBodyRemoteAzureDatabase
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetNodeOKBodyRemoteRDS RemoteRDSNode represents remote RDS Node. Agents can't run on Remote RDS Nodes.
swagger:model GetNodeOKBodyRemoteRDS
*/
type GetNodeOKBodyRemoteRDS struct {
	// Unique randomly generated instance identifier.
	NodeID string `json:"node_id,omitempty"`

	// Unique across all Nodes user-defined name.
	NodeName string `json:"node_name,omitempty"`

	// DB instance identifier.
	Address string `json:"address,omitempty"`

	// Node model.
	NodeModel string `json:"node_model,omitempty"`

	// Node region.
	Region string `json:"region,omitempty"`

	// Node availability zone.
	Az string `json:"az,omitempty"`

	// Custom user-assigned labels.
	CustomLabels map[string]string `json:"custom_labels,omitempty"`
}

// Validate validates this get node OK body remote RDS
func (o *GetNodeOKBodyRemoteRDS) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get node OK body remote RDS based on context it is used
func (o *GetNodeOKBodyRemoteRDS) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetNodeOKBodyRemoteRDS) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetNodeOKBodyRemoteRDS) UnmarshalBinary(b []byte) error {
	var res GetNodeOKBodyRemoteRDS
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
