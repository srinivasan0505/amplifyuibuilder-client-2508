package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// FieldConfigposition represents the FieldConfigposition schema from the OpenAPI specification
type FieldConfigposition struct {
	Below interface{} `json:"below,omitempty"`
	Fixed interface{} `json:"fixed,omitempty"`
	Rightof interface{} `json:"rightOf,omitempty"`
}

// FormStyleConfig represents the FormStyleConfig schema from the OpenAPI specification
type FormStyleConfig struct {
	Tokenreference interface{} `json:"tokenReference,omitempty"`
	Value interface{} `json:"value,omitempty"`
}

// PutMetadataFlagRequestbody represents the PutMetadataFlagRequestbody schema from the OpenAPI specification
type PutMetadataFlagRequestbody struct {
	Newvalue interface{} `json:"newValue"`
}

// GetCodegenJobRequest represents the GetCodegenJobRequest schema from the OpenAPI specification
type GetCodegenJobRequest struct {
}

// FormDataTypeConfig represents the FormDataTypeConfig schema from the OpenAPI specification
type FormDataTypeConfig struct {
	Datatypename interface{} `json:"dataTypeName"`
	Datasourcetype interface{} `json:"dataSourceType"`
}

// FormCTA represents the FormCTA schema from the OpenAPI specification
type FormCTA struct {
	Cancel FormCTAcancel `json:"cancel,omitempty"`
	Clear FormCTAclear `json:"clear,omitempty"`
	Position interface{} `json:"position,omitempty"`
	Submit FormCTAsubmit `json:"submit,omitempty"`
}

// CreateComponentRequestcomponentToCreate represents the CreateComponentRequestcomponentToCreate schema from the OpenAPI specification
type CreateComponentRequestcomponentToCreate struct {
	Componenttype interface{} `json:"componentType"`
	Events interface{} `json:"events,omitempty"`
	Bindingproperties interface{} `json:"bindingProperties"`
	Sourceid interface{} `json:"sourceId,omitempty"`
	Name interface{} `json:"name"`
	Overrides interface{} `json:"overrides"`
	Schemaversion interface{} `json:"schemaVersion,omitempty"`
	Tags interface{} `json:"tags,omitempty"`
	Children interface{} `json:"children,omitempty"`
	Properties interface{} `json:"properties"` // Describes the component's properties.
	Variants interface{} `json:"variants"`
	Collectionproperties interface{} `json:"collectionProperties,omitempty"`
}

// UpdateThemeResponseentity represents the UpdateThemeResponseentity schema from the OpenAPI specification
type UpdateThemeResponseentity struct {
	Values interface{} `json:"values"`
	Id interface{} `json:"id"`
	Environmentname interface{} `json:"environmentName"`
	Overrides interface{} `json:"overrides,omitempty"`
	Appid interface{} `json:"appId"`
	Createdat interface{} `json:"createdAt"`
	Modifiedat interface{} `json:"modifiedAt,omitempty"`
	Name interface{} `json:"name"`
	Tags interface{} `json:"tags,omitempty"`
}

// FormInputBindingProperties represents the FormInputBindingProperties schema from the OpenAPI specification
type FormInputBindingProperties struct {
}

// CreateComponentResponse represents the CreateComponentResponse schema from the OpenAPI specification
type CreateComponentResponse struct {
	Entity CreateComponentResponseentity `json:"entity,omitempty"`
}

// DeleteComponentRequest represents the DeleteComponentRequest schema from the OpenAPI specification
type DeleteComponentRequest struct {
}

// FieldConfiginputType represents the FieldConfiginputType schema from the OpenAPI specification
type FieldConfiginputType struct {
	Maxvalue interface{} `json:"maxValue,omitempty"`
	Placeholder interface{} `json:"placeholder,omitempty"`
	Defaultvalue interface{} `json:"defaultValue,omitempty"`
	Minvalue interface{} `json:"minValue,omitempty"`
	Readonly interface{} `json:"readOnly,omitempty"`
	Value interface{} `json:"value,omitempty"`
	Required interface{} `json:"required,omitempty"`
	Defaultchecked interface{} `json:"defaultChecked,omitempty"`
	Name interface{} `json:"name,omitempty"`
	Fileuploaderconfig FieldInputConfigfileUploaderConfig `json:"fileUploaderConfig,omitempty"`
	Valuemappings FieldInputConfigvalueMappings `json:"valueMappings,omitempty"`
	Step interface{} `json:"step,omitempty"`
	Defaultcountrycode interface{} `json:"defaultCountryCode,omitempty"`
	Isarray interface{} `json:"isArray,omitempty"`
	Descriptivetext interface{} `json:"descriptiveText,omitempty"`
	TypeField interface{} `json:"type"`
}

// ComponentPropertybindingProperties represents the ComponentPropertybindingProperties schema from the OpenAPI specification
type ComponentPropertybindingProperties struct {
	Field interface{} `json:"field,omitempty"`
	Property interface{} `json:"property"`
}

// ComponentOverrides represents the ComponentOverrides schema from the OpenAPI specification
type ComponentOverrides struct {
}

// CreateThemeResponseentity represents the CreateThemeResponseentity schema from the OpenAPI specification
type CreateThemeResponseentity struct {
	Environmentname interface{} `json:"environmentName"`
	Overrides interface{} `json:"overrides,omitempty"`
	Appid interface{} `json:"appId"`
	Createdat interface{} `json:"createdAt"`
	Modifiedat interface{} `json:"modifiedAt,omitempty"`
	Name interface{} `json:"name"`
	Tags interface{} `json:"tags,omitempty"`
	Values interface{} `json:"values"`
	Id interface{} `json:"id"`
}

// ExportThemesRequest represents the ExportThemesRequest schema from the OpenAPI specification
type ExportThemesRequest struct {
}

// ValueMappingvalue represents the ValueMappingvalue schema from the OpenAPI specification
type ValueMappingvalue struct {
	Bindingproperties FormInputValuePropertybindingProperties `json:"bindingProperties,omitempty"`
	Concat interface{} `json:"concat,omitempty"`
	Value interface{} `json:"value,omitempty"`
}

// CodegenJobSummary represents the CodegenJobSummary schema from the OpenAPI specification
type CodegenJobSummary struct {
	Appid interface{} `json:"appId"`
	Createdat interface{} `json:"createdAt,omitempty"`
	Environmentname interface{} `json:"environmentName"`
	Id interface{} `json:"id"`
	Modifiedat interface{} `json:"modifiedAt,omitempty"`
}

// ActionParameterstype represents the ActionParameterstype schema from the OpenAPI specification
type ActionParameterstype struct {
	TypeField interface{} `json:"type,omitempty"`
	Bindings interface{} `json:"bindings,omitempty"`
	Defaultvalue interface{} `json:"defaultValue,omitempty"`
	Componentname interface{} `json:"componentName,omitempty"`
	Value interface{} `json:"value,omitempty"`
	Condition ComponentPropertycondition `json:"condition,omitempty"`
	Event interface{} `json:"event,omitempty"`
	Collectionbindingproperties ComponentPropertycollectionBindingProperties `json:"collectionBindingProperties,omitempty"`
	Importedvalue interface{} `json:"importedValue,omitempty"`
	Property interface{} `json:"property,omitempty"`
	Userattribute interface{} `json:"userAttribute,omitempty"`
	Concat interface{} `json:"concat,omitempty"`
	Configured interface{} `json:"configured,omitempty"`
	Bindingproperties ComponentPropertybindingProperties `json:"bindingProperties,omitempty"`
	Model interface{} `json:"model,omitempty"`
}

// FormInputValueProperty represents the FormInputValueProperty schema from the OpenAPI specification
type FormInputValueProperty struct {
	Value interface{} `json:"value,omitempty"`
	Bindingproperties FormInputValuePropertybindingProperties `json:"bindingProperties,omitempty"`
	Concat interface{} `json:"concat,omitempty"`
}

// UpdateThemeRequest represents the UpdateThemeRequest schema from the OpenAPI specification
type UpdateThemeRequest struct {
	Updatedtheme UpdateThemeRequestupdatedTheme `json:"updatedTheme"`
}

// CodegenGenericDataEnums represents the CodegenGenericDataEnums schema from the OpenAPI specification
type CodegenGenericDataEnums struct {
}

// FormInputValuePropertyBindingProperties represents the FormInputValuePropertyBindingProperties schema from the OpenAPI specification
type FormInputValuePropertyBindingProperties struct {
	Property interface{} `json:"property"`
	Field interface{} `json:"field,omitempty"`
}

// UpdateFormResponse represents the UpdateFormResponse schema from the OpenAPI specification
type UpdateFormResponse struct {
	Entity UpdateFormResponseentity `json:"entity,omitempty"`
}

// ActionParametersanchor represents the ActionParametersanchor schema from the OpenAPI specification
type ActionParametersanchor struct {
	Concat interface{} `json:"concat,omitempty"`
	Configured interface{} `json:"configured,omitempty"`
	Bindingproperties ComponentPropertybindingProperties `json:"bindingProperties,omitempty"`
	Model interface{} `json:"model,omitempty"`
	TypeField interface{} `json:"type,omitempty"`
	Bindings interface{} `json:"bindings,omitempty"`
	Defaultvalue interface{} `json:"defaultValue,omitempty"`
	Componentname interface{} `json:"componentName,omitempty"`
	Value interface{} `json:"value,omitempty"`
	Condition ComponentPropertycondition `json:"condition,omitempty"`
	Event interface{} `json:"event,omitempty"`
	Collectionbindingproperties ComponentPropertycollectionBindingProperties `json:"collectionBindingProperties,omitempty"`
	Importedvalue interface{} `json:"importedValue,omitempty"`
	Property interface{} `json:"property,omitempty"`
	Userattribute interface{} `json:"userAttribute,omitempty"`
}

// CreateThemerequest represents the CreateThemerequest schema from the OpenAPI specification
type CreateThemerequest struct {
	Themetocreate CreateThemerequestthemeToCreate `json:"themeToCreate"` // Represents all of the information that is required to create a theme.
}

// GetMetadataResponse represents the GetMetadataResponse schema from the OpenAPI specification
type GetMetadataResponse struct {
	Features interface{} `json:"features"`
}

// ComponentBindingProperties represents the ComponentBindingProperties schema from the OpenAPI specification
type ComponentBindingProperties struct {
}

// ComponentEvent represents the ComponentEvent schema from the OpenAPI specification
type ComponentEvent struct {
	Action interface{} `json:"action,omitempty"`
	Bindingevent interface{} `json:"bindingEvent,omitempty"`
	Parameters ComponentEventparameters `json:"parameters,omitempty"`
}

// NoApiRenderConfig represents the NoApiRenderConfig schema from the OpenAPI specification
type NoApiRenderConfig struct {
}

// UpdateThemeResponse represents the UpdateThemeResponse schema from the OpenAPI specification
type UpdateThemeResponse struct {
	Entity UpdateThemeResponseentity `json:"entity,omitempty"`
}

// GraphQLRenderConfig represents the GraphQLRenderConfig schema from the OpenAPI specification
type GraphQLRenderConfig struct {
	Fragmentsfilepath interface{} `json:"fragmentsFilePath"`
	Mutationsfilepath interface{} `json:"mutationsFilePath"`
	Queriesfilepath interface{} `json:"queriesFilePath"`
	Subscriptionsfilepath interface{} `json:"subscriptionsFilePath"`
	Typesfilepath interface{} `json:"typesFilePath"`
}

// UpdateFormResponseentity represents the UpdateFormResponseentity schema from the OpenAPI specification
type UpdateFormResponseentity struct {
	Appid interface{} `json:"appId"`
	Id interface{} `json:"id"`
	Labeldecorator interface{} `json:"labelDecorator,omitempty"`
	Schemaversion interface{} `json:"schemaVersion"`
	Sectionalelements interface{} `json:"sectionalElements"`
	Style Formstyle `json:"style"`
	Cta Formcta `json:"cta,omitempty"`
	Datatype CreateFormrequestformToCreatedataType `json:"dataType"`
	Fields interface{} `json:"fields"`
	Tags interface{} `json:"tags,omitempty"`
	Environmentname interface{} `json:"environmentName"`
	Formactiontype interface{} `json:"formActionType"`
	Name interface{} `json:"name"`
}

// CreateFormrequestformToCreatestyle represents the CreateFormrequestformToCreatestyle schema from the OpenAPI specification
type CreateFormrequestformToCreatestyle struct {
	Outerpadding FormStyleouterPadding `json:"outerPadding,omitempty"`
	Verticalgap FormStyleverticalGap `json:"verticalGap,omitempty"`
	Horizontalgap FormStylehorizontalGap `json:"horizontalGap,omitempty"`
}

// FieldInputConfigfileUploaderConfig represents the FieldInputConfigfileUploaderConfig schema from the OpenAPI specification
type FieldInputConfigfileUploaderConfig struct {
	Maxfilecount interface{} `json:"maxFileCount,omitempty"`
	Maxsize interface{} `json:"maxSize,omitempty"`
	Showthumbnails interface{} `json:"showThumbnails,omitempty"`
	Acceptedfiletypes interface{} `json:"acceptedFileTypes"`
	Accesslevel interface{} `json:"accessLevel"`
	Isresumable interface{} `json:"isResumable,omitempty"`
}

// UpdateThemeRequestupdatedTheme represents the UpdateThemeRequestupdatedTheme schema from the OpenAPI specification
type UpdateThemeRequestupdatedTheme struct {
	Values interface{} `json:"values"`
	Id interface{} `json:"id,omitempty"`
	Name interface{} `json:"name,omitempty"`
	Overrides interface{} `json:"overrides,omitempty"`
}

// CodegenJobGenericDataSchema represents the CodegenJobGenericDataSchema schema from the OpenAPI specification
type CodegenJobGenericDataSchema struct {
	Datasourcetype interface{} `json:"dataSourceType"`
	Enums interface{} `json:"enums"`
	Models interface{} `json:"models"`
	Nonmodels interface{} `json:"nonModels"`
}

// DeleteFormRequest represents the DeleteFormRequest schema from the OpenAPI specification
type DeleteFormRequest struct {
}

// CreateComponentResponseentity represents the CreateComponentResponseentity schema from the OpenAPI specification
type CreateComponentResponseentity struct {
	Appid interface{} `json:"appId"`
	Createdat interface{} `json:"createdAt"`
	Sourceid interface{} `json:"sourceId,omitempty"`
	Children interface{} `json:"children,omitempty"`
	Name interface{} `json:"name"`
	Componenttype interface{} `json:"componentType"`
	Modifiedat interface{} `json:"modifiedAt,omitempty"`
	Overrides interface{} `json:"overrides"`
	Schemaversion interface{} `json:"schemaVersion,omitempty"`
	Tags interface{} `json:"tags,omitempty"`
	Bindingproperties interface{} `json:"bindingProperties"`
	Variants interface{} `json:"variants"`
	Events interface{} `json:"events,omitempty"`
	Id interface{} `json:"id"`
	Environmentname interface{} `json:"environmentName"`
	Collectionproperties interface{} `json:"collectionProperties,omitempty"`
	Properties interface{} `json:"properties"` // Describes the component's properties. You can't specify <code>tags</code> as a valid property for <code>properties</code>.
}

// CodegenGenericDataModel represents the CodegenGenericDataModel schema from the OpenAPI specification
type CodegenGenericDataModel struct {
	Isjointable interface{} `json:"isJoinTable,omitempty"`
	Primarykeys interface{} `json:"primaryKeys"`
	Fields interface{} `json:"fields"`
}

// FieldsMap represents the FieldsMap schema from the OpenAPI specification
type FieldsMap struct {
}

// ListThemesResponse represents the ListThemesResponse schema from the OpenAPI specification
type ListThemesResponse struct {
	Entities interface{} `json:"entities"`
	Nexttoken interface{} `json:"nextToken,omitempty"`
}

// FormInputBindingPropertiesValueProperties represents the FormInputBindingPropertiesValueProperties schema from the OpenAPI specification
type FormInputBindingPropertiesValueProperties struct {
	Model interface{} `json:"model,omitempty"`
}

// UpdateComponentRequest represents the UpdateComponentRequest schema from the OpenAPI specification
type UpdateComponentRequest struct {
	Updatedcomponent UpdateComponentRequestupdatedComponent `json:"updatedComponent"`
}

// StartCodegenJobrequestcodegenJobToCreatefeatures represents the StartCodegenJobrequestcodegenJobToCreatefeatures schema from the OpenAPI specification
type StartCodegenJobrequestcodegenJobToCreatefeatures struct {
	Isnonmodelsupported interface{} `json:"isNonModelSupported,omitempty"`
	Isrelationshipsupported interface{} `json:"isRelationshipSupported,omitempty"`
}

// Predicate represents the Predicate schema from the OpenAPI specification
type Predicate struct {
	Field interface{} `json:"field,omitempty"`
	Operand interface{} `json:"operand,omitempty"`
	Operandtype interface{} `json:"operandType,omitempty"`
	Operator interface{} `json:"operator,omitempty"`
	Or interface{} `json:"or,omitempty"`
	And interface{} `json:"and,omitempty"`
}

// CreateFormrequest represents the CreateFormrequest schema from the OpenAPI specification
type CreateFormrequest struct {
	Formtocreate CreateFormrequestformToCreate `json:"formToCreate"` // Represents all of the information that is required to create a form.
}

// ExchangeCodeForTokenRequest represents the ExchangeCodeForTokenRequest schema from the OpenAPI specification
type ExchangeCodeForTokenRequest struct {
	Request ExchangeCodeForTokenRequestrequest `json:"request"`
}

// ListThemesRequest represents the ListThemesRequest schema from the OpenAPI specification
type ListThemesRequest struct {
}

// SortProperty represents the SortProperty schema from the OpenAPI specification
type SortProperty struct {
	Direction interface{} `json:"direction"`
	Field interface{} `json:"field"`
}

// StartCodegenJobRequest represents the StartCodegenJobRequest schema from the OpenAPI specification
type StartCodegenJobRequest struct {
	Codegenjobtocreate StartCodegenJobRequestcodegenJobToCreate `json:"codegenJobToCreate"`
}

// UpdateComponentResponseentity represents the UpdateComponentResponseentity schema from the OpenAPI specification
type UpdateComponentResponseentity struct {
	Properties interface{} `json:"properties"` // Describes the component's properties. You can't specify <code>tags</code> as a valid property for <code>properties</code>.
	Appid interface{} `json:"appId"`
	Createdat interface{} `json:"createdAt"`
	Sourceid interface{} `json:"sourceId,omitempty"`
	Children interface{} `json:"children,omitempty"`
	Name interface{} `json:"name"`
	Componenttype interface{} `json:"componentType"`
	Modifiedat interface{} `json:"modifiedAt,omitempty"`
	Overrides interface{} `json:"overrides"`
	Schemaversion interface{} `json:"schemaVersion,omitempty"`
	Tags interface{} `json:"tags,omitempty"`
	Bindingproperties interface{} `json:"bindingProperties"`
	Variants interface{} `json:"variants"`
	Events interface{} `json:"events,omitempty"`
	Id interface{} `json:"id"`
	Environmentname interface{} `json:"environmentName"`
	Collectionproperties interface{} `json:"collectionProperties,omitempty"`
}

// Component represents the Component schema from the OpenAPI specification
type Component struct {
	Events interface{} `json:"events,omitempty"`
	Id interface{} `json:"id"`
	Environmentname interface{} `json:"environmentName"`
	Collectionproperties interface{} `json:"collectionProperties,omitempty"`
	Properties interface{} `json:"properties"` // Describes the component's properties. You can't specify <code>tags</code> as a valid property for <code>properties</code>.
	Appid interface{} `json:"appId"`
	Createdat interface{} `json:"createdAt"`
	Sourceid interface{} `json:"sourceId,omitempty"`
	Children interface{} `json:"children,omitempty"`
	Name interface{} `json:"name"`
	Componenttype interface{} `json:"componentType"`
	Modifiedat interface{} `json:"modifiedAt,omitempty"`
	Overrides interface{} `json:"overrides"`
	Schemaversion interface{} `json:"schemaVersion,omitempty"`
	Tags interface{} `json:"tags,omitempty"`
	Bindingproperties interface{} `json:"bindingProperties"`
	Variants interface{} `json:"variants"`
}

// CreateFormResponseentity represents the CreateFormResponseentity schema from the OpenAPI specification
type CreateFormResponseentity struct {
	Fields interface{} `json:"fields"`
	Tags interface{} `json:"tags,omitempty"`
	Environmentname interface{} `json:"environmentName"`
	Formactiontype interface{} `json:"formActionType"`
	Name interface{} `json:"name"`
	Appid interface{} `json:"appId"`
	Id interface{} `json:"id"`
	Labeldecorator interface{} `json:"labelDecorator,omitempty"`
	Schemaversion interface{} `json:"schemaVersion"`
	Sectionalelements interface{} `json:"sectionalElements"`
	Style Formstyle `json:"style"`
	Cta Formcta `json:"cta,omitempty"`
	Datatype CreateFormrequestformToCreatedataType `json:"dataType"`
}

// ValueMappingdisplayValue represents the ValueMappingdisplayValue schema from the OpenAPI specification
type ValueMappingdisplayValue struct {
	Concat interface{} `json:"concat,omitempty"`
	Value interface{} `json:"value,omitempty"`
	Bindingproperties FormInputValuePropertybindingProperties `json:"bindingProperties,omitempty"`
}

// StartCodegenJobResponse represents the StartCodegenJobResponse schema from the OpenAPI specification
type StartCodegenJobResponse struct {
	Entity StartCodegenJobResponseentity `json:"entity,omitempty"`
}

// ExportComponentsRequest represents the ExportComponentsRequest schema from the OpenAPI specification
type ExportComponentsRequest struct {
}

// UpdateComponentrequestupdatedComponent represents the UpdateComponentrequestupdatedComponent schema from the OpenAPI specification
type UpdateComponentrequestupdatedComponent struct {
	Overrides interface{} `json:"overrides,omitempty"`
	Sourceid interface{} `json:"sourceId,omitempty"`
	Bindingproperties interface{} `json:"bindingProperties,omitempty"`
	Collectionproperties interface{} `json:"collectionProperties,omitempty"`
	Properties interface{} `json:"properties,omitempty"` // Describes the component's properties.
	Schemaversion interface{} `json:"schemaVersion,omitempty"`
	Componenttype interface{} `json:"componentType,omitempty"`
	Variants interface{} `json:"variants,omitempty"`
	Children interface{} `json:"children,omitempty"`
	Events interface{} `json:"events,omitempty"`
	Id interface{} `json:"id,omitempty"`
	Name interface{} `json:"name,omitempty"`
}

// CodegenJobRenderConfigreact represents the CodegenJobRenderConfigreact schema from the OpenAPI specification
type CodegenJobRenderConfigreact struct {
	Inlinesourcemap interface{} `json:"inlineSourceMap,omitempty"`
	Module interface{} `json:"module,omitempty"`
	Rendertypedeclarations interface{} `json:"renderTypeDeclarations,omitempty"`
	Script interface{} `json:"script,omitempty"`
	Target interface{} `json:"target,omitempty"`
	Apiconfiguration ReactStartCodegenJobDataapiConfiguration `json:"apiConfiguration,omitempty"`
}

// ComponentBindingPropertiesValuebindingProperties represents the ComponentBindingPropertiesValuebindingProperties schema from the OpenAPI specification
type ComponentBindingPropertiesValuebindingProperties struct {
	Slotname interface{} `json:"slotName,omitempty"`
	Userattribute interface{} `json:"userAttribute,omitempty"`
	Bucket interface{} `json:"bucket,omitempty"`
	Defaultvalue interface{} `json:"defaultValue,omitempty"`
	Field interface{} `json:"field,omitempty"`
	Key interface{} `json:"key,omitempty"`
	Model interface{} `json:"model,omitempty"`
	Predicates interface{} `json:"predicates,omitempty"`
}

// ActionParametersurl represents the ActionParametersurl schema from the OpenAPI specification
type ActionParametersurl struct {
	Concat interface{} `json:"concat,omitempty"`
	Configured interface{} `json:"configured,omitempty"`
	Bindingproperties ComponentPropertybindingProperties `json:"bindingProperties,omitempty"`
	Model interface{} `json:"model,omitempty"`
	TypeField interface{} `json:"type,omitempty"`
	Bindings interface{} `json:"bindings,omitempty"`
	Defaultvalue interface{} `json:"defaultValue,omitempty"`
	Componentname interface{} `json:"componentName,omitempty"`
	Value interface{} `json:"value,omitempty"`
	Condition ComponentPropertycondition `json:"condition,omitempty"`
	Event interface{} `json:"event,omitempty"`
	Collectionbindingproperties ComponentPropertycollectionBindingProperties `json:"collectionBindingProperties,omitempty"`
	Importedvalue interface{} `json:"importedValue,omitempty"`
	Property interface{} `json:"property,omitempty"`
	Userattribute interface{} `json:"userAttribute,omitempty"`
}

// FormBindingElement represents the FormBindingElement schema from the OpenAPI specification
type FormBindingElement struct {
	Element interface{} `json:"element"`
	Property interface{} `json:"property"`
}

// CodegenJobAsset represents the CodegenJobAsset schema from the OpenAPI specification
type CodegenJobAsset struct {
	Downloadurl interface{} `json:"downloadUrl,omitempty"`
}

// CreateFormrequestformToCreatedataType represents the CreateFormrequestformToCreatedataType schema from the OpenAPI specification
type CreateFormrequestformToCreatedataType struct {
	Datasourcetype interface{} `json:"dataSourceType"`
	Datatypename interface{} `json:"dataTypeName"`
}

// RefreshTokenResponse represents the RefreshTokenResponse schema from the OpenAPI specification
type RefreshTokenResponse struct {
	Expiresin interface{} `json:"expiresIn"`
	Accesstoken interface{} `json:"accessToken"`
}

// UpdateFormrequestupdatedForm represents the UpdateFormrequestupdatedForm schema from the OpenAPI specification
type UpdateFormrequestupdatedForm struct {
	Schemaversion interface{} `json:"schemaVersion,omitempty"`
	Style CreateFormrequestformToCreatestyle `json:"style,omitempty"`
	Labeldecorator interface{} `json:"labelDecorator,omitempty"`
	Fields interface{} `json:"fields,omitempty"`
	Formactiontype interface{} `json:"formActionType,omitempty"`
	Sectionalelements interface{} `json:"sectionalElements,omitempty"`
	Cta CreateFormrequestformToCreatecta `json:"cta,omitempty"`
	Datatype CreateFormrequestformToCreatedataType `json:"dataType,omitempty"`
	Name interface{} `json:"name,omitempty"`
}

// ComponentProperties represents the ComponentProperties schema from the OpenAPI specification
type ComponentProperties struct {
}

// UpdateComponentRequestupdatedComponent represents the UpdateComponentRequestupdatedComponent schema from the OpenAPI specification
type UpdateComponentRequestupdatedComponent struct {
	Schemaversion interface{} `json:"schemaVersion,omitempty"`
	Events interface{} `json:"events,omitempty"`
	Id interface{} `json:"id,omitempty"`
	Overrides interface{} `json:"overrides,omitempty"`
	Properties interface{} `json:"properties,omitempty"` // Describes the component's properties.
	Sourceid interface{} `json:"sourceId,omitempty"`
	Children interface{} `json:"children,omitempty"`
	Collectionproperties interface{} `json:"collectionProperties,omitempty"`
	Componenttype interface{} `json:"componentType,omitempty"`
	Name interface{} `json:"name,omitempty"`
	Variants interface{} `json:"variants,omitempty"`
	Bindingproperties interface{} `json:"bindingProperties,omitempty"`
}

// RefreshTokenRequestBody represents the RefreshTokenRequestBody schema from the OpenAPI specification
type RefreshTokenRequestBody struct {
	Clientid interface{} `json:"clientId,omitempty"`
	Token interface{} `json:"token"`
}

// FieldConfig represents the FieldConfig schema from the OpenAPI specification
type FieldConfig struct {
	Excluded interface{} `json:"excluded,omitempty"`
	Inputtype FieldConfiginputType `json:"inputType,omitempty"`
	Label interface{} `json:"label,omitempty"`
	Position FieldConfigposition `json:"position,omitempty"`
	Validations interface{} `json:"validations,omitempty"`
}

// FormCTAclear represents the FormCTAclear schema from the OpenAPI specification
type FormCTAclear struct {
	Children interface{} `json:"children,omitempty"`
	Excluded interface{} `json:"excluded,omitempty"`
	Position FormButtonposition `json:"position,omitempty"`
}

// CreateThemeData represents the CreateThemeData schema from the OpenAPI specification
type CreateThemeData struct {
	Values interface{} `json:"values"`
	Name interface{} `json:"name"`
	Overrides interface{} `json:"overrides,omitempty"`
	Tags interface{} `json:"tags,omitempty"`
}

// CreateThemerequestthemeToCreate represents the CreateThemerequestthemeToCreate schema from the OpenAPI specification
type CreateThemerequestthemeToCreate struct {
	Name interface{} `json:"name,omitempty"`
	Overrides interface{} `json:"overrides,omitempty"`
	Tags interface{} `json:"tags,omitempty"`
	Values interface{} `json:"values,omitempty"`
}

// CodegenGenericDataModels represents the CodegenGenericDataModels schema from the OpenAPI specification
type CodegenGenericDataModels struct {
}

// ApiConfigurationgraphQLConfig represents the ApiConfigurationgraphQLConfig schema from the OpenAPI specification
type ApiConfigurationgraphQLConfig struct {
	Fragmentsfilepath interface{} `json:"fragmentsFilePath"`
	Mutationsfilepath interface{} `json:"mutationsFilePath"`
	Queriesfilepath interface{} `json:"queriesFilePath"`
	Subscriptionsfilepath interface{} `json:"subscriptionsFilePath"`
	Typesfilepath interface{} `json:"typesFilePath"`
}

// FormCTAsubmit represents the FormCTAsubmit schema from the OpenAPI specification
type FormCTAsubmit struct {
	Children interface{} `json:"children,omitempty"`
	Excluded interface{} `json:"excluded,omitempty"`
	Position FormButtonposition `json:"position,omitempty"`
}

// UpdateFormData represents the UpdateFormData schema from the OpenAPI specification
type UpdateFormData struct {
	Style CreateFormrequestformToCreatestyle `json:"style,omitempty"`
	Schemaversion interface{} `json:"schemaVersion,omitempty"`
	Fields interface{} `json:"fields,omitempty"`
	Sectionalelements interface{} `json:"sectionalElements,omitempty"`
	Cta CreateFormrequestformToCreatecta `json:"cta,omitempty"`
	Formactiontype interface{} `json:"formActionType,omitempty"`
	Datatype CreateFormrequestformToCreatedataType `json:"dataType,omitempty"`
	Labeldecorator interface{} `json:"labelDecorator,omitempty"`
	Name interface{} `json:"name,omitempty"`
}

// ComponentPropertycondition represents the ComponentPropertycondition schema from the OpenAPI specification
type ComponentPropertycondition struct {
	ElseField ComponentConditionPropertyelse `json:"else,omitempty"`
	Field interface{} `json:"field,omitempty"`
	Operand interface{} `json:"operand,omitempty"`
	Operandtype interface{} `json:"operandType,omitempty"`
	Operator interface{} `json:"operator,omitempty"`
	Property interface{} `json:"property,omitempty"`
	Then ComponentConditionPropertythen `json:"then,omitempty"`
}

// CreateFormRequestformToCreate represents the CreateFormRequestformToCreate schema from the OpenAPI specification
type CreateFormRequestformToCreate struct {
	Sectionalelements interface{} `json:"sectionalElements"`
	Schemaversion interface{} `json:"schemaVersion"`
	Cta CreateFormrequestformToCreatecta `json:"cta,omitempty"`
	Labeldecorator interface{} `json:"labelDecorator,omitempty"`
	Style CreateFormrequestformToCreatestyle `json:"style"`
	Tags interface{} `json:"tags,omitempty"`
	Datatype CreateFormrequestformToCreatedataType `json:"dataType"`
	Formactiontype interface{} `json:"formActionType"`
	Fields interface{} `json:"fields"`
	Name interface{} `json:"name"`
}

// Form represents the Form schema from the OpenAPI specification
type Form struct {
	Labeldecorator interface{} `json:"labelDecorator,omitempty"`
	Schemaversion interface{} `json:"schemaVersion"`
	Sectionalelements interface{} `json:"sectionalElements"`
	Style Formstyle `json:"style"`
	Cta Formcta `json:"cta,omitempty"`
	Datatype CreateFormrequestformToCreatedataType `json:"dataType"`
	Fields interface{} `json:"fields"`
	Tags interface{} `json:"tags,omitempty"`
	Environmentname interface{} `json:"environmentName"`
	Formactiontype interface{} `json:"formActionType"`
	Name interface{} `json:"name"`
	Appid interface{} `json:"appId"`
	Id interface{} `json:"id"`
}

// FeaturesMap represents the FeaturesMap schema from the OpenAPI specification
type FeaturesMap struct {
}

// CodegenJobasset represents the CodegenJobasset schema from the OpenAPI specification
type CodegenJobasset struct {
	Downloadurl interface{} `json:"downloadUrl,omitempty"`
}

// ActionParametersid represents the ActionParametersid schema from the OpenAPI specification
type ActionParametersid struct {
	Event interface{} `json:"event,omitempty"`
	Collectionbindingproperties ComponentPropertycollectionBindingProperties `json:"collectionBindingProperties,omitempty"`
	Importedvalue interface{} `json:"importedValue,omitempty"`
	Property interface{} `json:"property,omitempty"`
	Userattribute interface{} `json:"userAttribute,omitempty"`
	Concat interface{} `json:"concat,omitempty"`
	Configured interface{} `json:"configured,omitempty"`
	Bindingproperties ComponentPropertybindingProperties `json:"bindingProperties,omitempty"`
	Model interface{} `json:"model,omitempty"`
	TypeField interface{} `json:"type,omitempty"`
	Bindings interface{} `json:"bindings,omitempty"`
	Defaultvalue interface{} `json:"defaultValue,omitempty"`
	Componentname interface{} `json:"componentName,omitempty"`
	Value interface{} `json:"value,omitempty"`
	Condition ComponentPropertycondition `json:"condition,omitempty"`
}

// ListFormsRequest represents the ListFormsRequest schema from the OpenAPI specification
type ListFormsRequest struct {
}

// UpdateFormRequest represents the UpdateFormRequest schema from the OpenAPI specification
type UpdateFormRequest struct {
	Updatedform UpdateFormRequestupdatedForm `json:"updatedForm"`
}

// SectionalElementposition represents the SectionalElementposition schema from the OpenAPI specification
type SectionalElementposition struct {
	Below interface{} `json:"below,omitempty"`
	Fixed interface{} `json:"fixed,omitempty"`
	Rightof interface{} `json:"rightOf,omitempty"`
}

// StartCodegenJobrequest represents the StartCodegenJobrequest schema from the OpenAPI specification
type StartCodegenJobrequest struct {
	Codegenjobtocreate StartCodegenJobrequestcodegenJobToCreate `json:"codegenJobToCreate"` // The code generation job resource configuration.
}

// UpdateComponentResponse represents the UpdateComponentResponse schema from the OpenAPI specification
type UpdateComponentResponse struct {
	Entity UpdateComponentResponseentity `json:"entity,omitempty"`
}

// FieldInputConfigvalueMappings represents the FieldInputConfigvalueMappings schema from the OpenAPI specification
type FieldInputConfigvalueMappings struct {
	Values interface{} `json:"values"`
	Bindingproperties interface{} `json:"bindingProperties,omitempty"`
}

// SectionalElementMap represents the SectionalElementMap schema from the OpenAPI specification
type SectionalElementMap struct {
}

// GetCodegenJobResponse represents the GetCodegenJobResponse schema from the OpenAPI specification
type GetCodegenJobResponse struct {
	Job GetCodegenJobResponsejob `json:"job,omitempty"`
}

// FormCTAcancel represents the FormCTAcancel schema from the OpenAPI specification
type FormCTAcancel struct {
	Children interface{} `json:"children,omitempty"`
	Excluded interface{} `json:"excluded,omitempty"`
	Position FormButtonposition `json:"position,omitempty"`
}

// ComponentProperty represents the ComponentProperty schema from the OpenAPI specification
type ComponentProperty struct {
	Concat interface{} `json:"concat,omitempty"`
	Configured interface{} `json:"configured,omitempty"`
	Bindingproperties ComponentPropertybindingProperties `json:"bindingProperties,omitempty"`
	Model interface{} `json:"model,omitempty"`
	TypeField interface{} `json:"type,omitempty"`
	Bindings interface{} `json:"bindings,omitempty"`
	Defaultvalue interface{} `json:"defaultValue,omitempty"`
	Componentname interface{} `json:"componentName,omitempty"`
	Value interface{} `json:"value,omitempty"`
	Condition ComponentPropertycondition `json:"condition,omitempty"`
	Event interface{} `json:"event,omitempty"`
	Collectionbindingproperties ComponentPropertycollectionBindingProperties `json:"collectionBindingProperties,omitempty"`
	Importedvalue interface{} `json:"importedValue,omitempty"`
	Property interface{} `json:"property,omitempty"`
	Userattribute interface{} `json:"userAttribute,omitempty"`
}

// FieldPosition represents the FieldPosition schema from the OpenAPI specification
type FieldPosition struct {
	Rightof interface{} `json:"rightOf,omitempty"`
	Below interface{} `json:"below,omitempty"`
	Fixed interface{} `json:"fixed,omitempty"`
}

// ExchangeCodeForTokenRequestrequest represents the ExchangeCodeForTokenRequestrequest schema from the OpenAPI specification
type ExchangeCodeForTokenRequestrequest struct {
	Clientid interface{} `json:"clientId,omitempty"`
	Code interface{} `json:"code"`
	Redirecturi interface{} `json:"redirectUri"`
}

// FormSummary represents the FormSummary schema from the OpenAPI specification
type FormSummary struct {
	Datatype FormSummarydataType `json:"dataType"`
	Environmentname interface{} `json:"environmentName"`
	Formactiontype interface{} `json:"formActionType"`
	Id interface{} `json:"id"`
	Name interface{} `json:"name"`
	Appid interface{} `json:"appId"`
}

// StartCodegenJobRequestcodegenJobToCreate represents the StartCodegenJobRequestcodegenJobToCreate schema from the OpenAPI specification
type StartCodegenJobRequestcodegenJobToCreate struct {
	Autogenerateforms interface{} `json:"autoGenerateForms,omitempty"`
	Features StartCodegenJobrequestcodegenJobToCreatefeatures `json:"features,omitempty"`
	Genericdataschema StartCodegenJobrequestcodegenJobToCreategenericDataSchema `json:"genericDataSchema,omitempty"`
	Renderconfig StartCodegenJobrequestcodegenJobToCreaterenderConfig `json:"renderConfig"`
	Tags interface{} `json:"tags,omitempty"`
}

// CodegenGenericDataFieldrelationship represents the CodegenGenericDataFieldrelationship schema from the OpenAPI specification
type CodegenGenericDataFieldrelationship struct {
	Relatedmodelfields interface{} `json:"relatedModelFields,omitempty"`
	Relatedmodelname interface{} `json:"relatedModelName"`
	Ishasmanyindex interface{} `json:"isHasManyIndex,omitempty"`
	Relatedjointablename interface{} `json:"relatedJoinTableName,omitempty"`
	Belongstofieldonrelatedmodel interface{} `json:"belongsToFieldOnRelatedModel,omitempty"`
	TypeField interface{} `json:"type"`
	Associatedfields interface{} `json:"associatedFields,omitempty"`
	Canunlinkassociatedmodel interface{} `json:"canUnlinkAssociatedModel,omitempty"`
	Relatedjoinfieldname interface{} `json:"relatedJoinFieldName,omitempty"`
}

// StartCodegenJobrequestcodegenJobToCreategenericDataSchema represents the StartCodegenJobrequestcodegenJobToCreategenericDataSchema schema from the OpenAPI specification
type StartCodegenJobrequestcodegenJobToCreategenericDataSchema struct {
	Datasourcetype interface{} `json:"dataSourceType"`
	Enums interface{} `json:"enums"`
	Models interface{} `json:"models"`
	Nonmodels interface{} `json:"nonModels"`
}

// FormBindings represents the FormBindings schema from the OpenAPI specification
type FormBindings struct {
}

// ComponentConditionProperty represents the ComponentConditionProperty schema from the OpenAPI specification
type ComponentConditionProperty struct {
	Property interface{} `json:"property,omitempty"`
	Then ComponentConditionPropertythen `json:"then,omitempty"`
	ElseField ComponentConditionPropertyelse `json:"else,omitempty"`
	Field interface{} `json:"field,omitempty"`
	Operand interface{} `json:"operand,omitempty"`
	Operandtype interface{} `json:"operandType,omitempty"`
	Operator interface{} `json:"operator,omitempty"`
}

// ExchangeCodeForTokenrequestrequest represents the ExchangeCodeForTokenrequestrequest schema from the OpenAPI specification
type ExchangeCodeForTokenrequestrequest struct {
	Clientid interface{} `json:"clientId,omitempty"`
	Code interface{} `json:"code,omitempty"`
	Redirecturi interface{} `json:"redirectUri,omitempty"`
}

// UpdateComponentData represents the UpdateComponentData schema from the OpenAPI specification
type UpdateComponentData struct {
	Properties interface{} `json:"properties,omitempty"` // Describes the component's properties.
	Sourceid interface{} `json:"sourceId,omitempty"`
	Children interface{} `json:"children,omitempty"`
	Collectionproperties interface{} `json:"collectionProperties,omitempty"`
	Componenttype interface{} `json:"componentType,omitempty"`
	Name interface{} `json:"name,omitempty"`
	Variants interface{} `json:"variants,omitempty"`
	Bindingproperties interface{} `json:"bindingProperties,omitempty"`
	Schemaversion interface{} `json:"schemaVersion,omitempty"`
	Events interface{} `json:"events,omitempty"`
	Id interface{} `json:"id,omitempty"`
	Overrides interface{} `json:"overrides,omitempty"`
}

// StartCodegenJobResponseentity represents the StartCodegenJobResponseentity schema from the OpenAPI specification
type StartCodegenJobResponseentity struct {
	Createdat interface{} `json:"createdAt,omitempty"`
	Features CodegenFeatureFlags `json:"features,omitempty"` // Describes the feature flags that you can specify for a code generation job.
	Status interface{} `json:"status,omitempty"`
	Appid interface{} `json:"appId"`
	Genericdataschema CodegenJobGenericDataSchema `json:"genericDataSchema,omitempty"` // Describes the data schema for a code generation job.
	Modifiedat interface{} `json:"modifiedAt,omitempty"`
	Environmentname interface{} `json:"environmentName"`
	Statusmessage interface{} `json:"statusMessage,omitempty"`
	Autogenerateforms interface{} `json:"autoGenerateForms,omitempty"`
	Id interface{} `json:"id"`
	Renderconfig CodegenJobRenderConfig `json:"renderConfig,omitempty"` // Describes the configuration information for rendering the UI component associated with the code generation job.
	Tags interface{} `json:"tags,omitempty"`
	Asset CodegenJobasset `json:"asset,omitempty"`
}

// PutMetadataFlagBody represents the PutMetadataFlagBody schema from the OpenAPI specification
type PutMetadataFlagBody struct {
	Newvalue interface{} `json:"newValue"`
}

// ActionParameters represents the ActionParameters schema from the OpenAPI specification
type ActionParameters struct {
	Url ActionParametersurl `json:"url,omitempty"`
	Fields interface{} `json:"fields,omitempty"`
	Target ActionParameterstarget `json:"target,omitempty"`
	Model interface{} `json:"model,omitempty"`
	Anchor ActionParametersanchor `json:"anchor,omitempty"`
	Global ActionParametersglobal `json:"global,omitempty"`
	Id ActionParametersid `json:"id,omitempty"`
	State ActionParametersstate `json:"state,omitempty"`
	TypeField ActionParameterstype `json:"type,omitempty"`
}

// GetCodegenJobResponsejob represents the GetCodegenJobResponsejob schema from the OpenAPI specification
type GetCodegenJobResponsejob struct {
	Statusmessage interface{} `json:"statusMessage,omitempty"`
	Autogenerateforms interface{} `json:"autoGenerateForms,omitempty"`
	Id interface{} `json:"id"`
	Renderconfig CodegenJobRenderConfig `json:"renderConfig,omitempty"` // Describes the configuration information for rendering the UI component associated with the code generation job.
	Tags interface{} `json:"tags,omitempty"`
	Asset CodegenJobasset `json:"asset,omitempty"`
	Createdat interface{} `json:"createdAt,omitempty"`
	Features CodegenFeatureFlags `json:"features,omitempty"` // Describes the feature flags that you can specify for a code generation job.
	Status interface{} `json:"status,omitempty"`
	Appid interface{} `json:"appId"`
	Genericdataschema CodegenJobGenericDataSchema `json:"genericDataSchema,omitempty"` // Describes the data schema for a code generation job.
	Modifiedat interface{} `json:"modifiedAt,omitempty"`
	Environmentname interface{} `json:"environmentName"`
}

// ValueMappings represents the ValueMappings schema from the OpenAPI specification
type ValueMappings struct {
	Values interface{} `json:"values"`
	Bindingproperties interface{} `json:"bindingProperties,omitempty"`
}

// CodegenGenericDataNonModelFields represents the CodegenGenericDataNonModelFields schema from the OpenAPI specification
type CodegenGenericDataNonModelFields struct {
}

// ExportComponentsResponse represents the ExportComponentsResponse schema from the OpenAPI specification
type ExportComponentsResponse struct {
	Nexttoken interface{} `json:"nextToken,omitempty"`
	Entities interface{} `json:"entities"`
}

// ThemeSummary represents the ThemeSummary schema from the OpenAPI specification
type ThemeSummary struct {
	Appid interface{} `json:"appId"`
	Environmentname interface{} `json:"environmentName"`
	Id interface{} `json:"id"`
	Name interface{} `json:"name"`
}

// CreateComponentRequest represents the CreateComponentRequest schema from the OpenAPI specification
type CreateComponentRequest struct {
	Componenttocreate CreateComponentRequestcomponentToCreate `json:"componentToCreate"`
}

// CodegenGenericDataField represents the CodegenGenericDataField schema from the OpenAPI specification
type CodegenGenericDataField struct {
	Relationship CodegenGenericDataFieldrelationship `json:"relationship,omitempty"`
	Required interface{} `json:"required"`
	Datatype interface{} `json:"dataType"`
	Datatypevalue interface{} `json:"dataTypeValue"`
	Isarray interface{} `json:"isArray"`
	Readonly interface{} `json:"readOnly"`
}

// GetFormResponseform represents the GetFormResponseform schema from the OpenAPI specification
type GetFormResponseform struct {
	Environmentname interface{} `json:"environmentName"`
	Formactiontype interface{} `json:"formActionType"`
	Name interface{} `json:"name"`
	Appid interface{} `json:"appId"`
	Id interface{} `json:"id"`
	Labeldecorator interface{} `json:"labelDecorator,omitempty"`
	Schemaversion interface{} `json:"schemaVersion"`
	Sectionalelements interface{} `json:"sectionalElements"`
	Style Formstyle `json:"style"`
	Cta Formcta `json:"cta,omitempty"`
	Datatype CreateFormrequestformToCreatedataType `json:"dataType"`
	Fields interface{} `json:"fields"`
	Tags interface{} `json:"tags,omitempty"`
}

// CreateFormResponse represents the CreateFormResponse schema from the OpenAPI specification
type CreateFormResponse struct {
	Entity CreateFormResponseentity `json:"entity,omitempty"`
}

// StartCodegenJobrequestcodegenJobToCreaterenderConfig represents the StartCodegenJobrequestcodegenJobToCreaterenderConfig schema from the OpenAPI specification
type StartCodegenJobrequestcodegenJobToCreaterenderConfig struct {
	React CodegenJobRenderConfigreact `json:"react,omitempty"`
}

// ExportThemesResponse represents the ExportThemesResponse schema from the OpenAPI specification
type ExportThemesResponse struct {
	Entities interface{} `json:"entities"`
	Nexttoken interface{} `json:"nextToken,omitempty"`
}

// ExportFormsResponse represents the ExportFormsResponse schema from the OpenAPI specification
type ExportFormsResponse struct {
	Nexttoken interface{} `json:"nextToken,omitempty"`
	Entities interface{} `json:"entities"`
}

// FileUploaderFieldConfig represents the FileUploaderFieldConfig schema from the OpenAPI specification
type FileUploaderFieldConfig struct {
	Maxsize interface{} `json:"maxSize,omitempty"`
	Showthumbnails interface{} `json:"showThumbnails,omitempty"`
	Acceptedfiletypes interface{} `json:"acceptedFileTypes"`
	Accesslevel interface{} `json:"accessLevel"`
	Isresumable interface{} `json:"isResumable,omitempty"`
	Maxfilecount interface{} `json:"maxFileCount,omitempty"`
}

// GetFormRequest represents the GetFormRequest schema from the OpenAPI specification
type GetFormRequest struct {
}

// GetThemeRequest represents the GetThemeRequest schema from the OpenAPI specification
type GetThemeRequest struct {
}

// ComponentVariant represents the ComponentVariant schema from the OpenAPI specification
type ComponentVariant struct {
	Overrides interface{} `json:"overrides,omitempty"`
	Variantvalues interface{} `json:"variantValues,omitempty"`
}

// GetComponentResponse represents the GetComponentResponse schema from the OpenAPI specification
type GetComponentResponse struct {
	Component GetComponentResponsecomponent `json:"component,omitempty"`
}

// CreateThemeResponse represents the CreateThemeResponse schema from the OpenAPI specification
type CreateThemeResponse struct {
	Entity CreateThemeResponseentity `json:"entity,omitempty"`
}

// FormInputValuePropertybindingProperties represents the FormInputValuePropertybindingProperties schema from the OpenAPI specification
type FormInputValuePropertybindingProperties struct {
	Property interface{} `json:"property"`
	Field interface{} `json:"field,omitempty"`
}

// ListComponentsRequest represents the ListComponentsRequest schema from the OpenAPI specification
type ListComponentsRequest struct {
}

// CreateComponentrequest represents the CreateComponentrequest schema from the OpenAPI specification
type CreateComponentrequest struct {
	Componenttocreate CreateComponentrequestcomponentToCreate `json:"componentToCreate"` // Represents all of the information that is required to create a component.
}

// Theme represents the Theme schema from the OpenAPI specification
type Theme struct {
	Appid interface{} `json:"appId"`
	Createdat interface{} `json:"createdAt"`
	Modifiedat interface{} `json:"modifiedAt,omitempty"`
	Name interface{} `json:"name"`
	Tags interface{} `json:"tags,omitempty"`
	Values interface{} `json:"values"`
	Id interface{} `json:"id"`
	Environmentname interface{} `json:"environmentName"`
	Overrides interface{} `json:"overrides,omitempty"`
}

// FormButtonposition represents the FormButtonposition schema from the OpenAPI specification
type FormButtonposition struct {
	Below interface{} `json:"below,omitempty"`
	Fixed interface{} `json:"fixed,omitempty"`
	Rightof interface{} `json:"rightOf,omitempty"`
}

// ComponentDataConfiguration represents the ComponentDataConfiguration schema from the OpenAPI specification
type ComponentDataConfiguration struct {
	Model interface{} `json:"model"`
	Predicate ComponentDataConfigurationpredicate `json:"predicate,omitempty"`
	Sort interface{} `json:"sort,omitempty"`
	Identifiers interface{} `json:"identifiers,omitempty"`
}

// FormButton represents the FormButton schema from the OpenAPI specification
type FormButton struct {
	Children interface{} `json:"children,omitempty"`
	Excluded interface{} `json:"excluded,omitempty"`
	Position FormButtonposition `json:"position,omitempty"`
}

// RefreshTokenrequest represents the RefreshTokenrequest schema from the OpenAPI specification
type RefreshTokenrequest struct {
	Refreshtokenbody RefreshTokenrequestrefreshTokenBody `json:"refreshTokenBody"` // Describes a refresh token.
}

// ReactStartCodegenJobDataapiConfiguration represents the ReactStartCodegenJobDataapiConfiguration schema from the OpenAPI specification
type ReactStartCodegenJobDataapiConfiguration struct {
	Noapiconfig interface{} `json:"noApiConfig,omitempty"`
	Datastoreconfig interface{} `json:"dataStoreConfig,omitempty"`
	Graphqlconfig ApiConfigurationgraphQLConfig `json:"graphQLConfig,omitempty"`
}

// UpdateThemerequestupdatedTheme represents the UpdateThemerequestupdatedTheme schema from the OpenAPI specification
type UpdateThemerequestupdatedTheme struct {
	Overrides interface{} `json:"overrides,omitempty"`
	Values interface{} `json:"values,omitempty"`
	Id interface{} `json:"id,omitempty"`
	Name interface{} `json:"name,omitempty"`
}

// ExchangeCodeForTokenResponse represents the ExchangeCodeForTokenResponse schema from the OpenAPI specification
type ExchangeCodeForTokenResponse struct {
	Accesstoken interface{} `json:"accessToken"`
	Expiresin interface{} `json:"expiresIn"`
	Refreshtoken interface{} `json:"refreshToken"`
}

// GetMetadataRequest represents the GetMetadataRequest schema from the OpenAPI specification
type GetMetadataRequest struct {
}

// Formstyle represents the Formstyle schema from the OpenAPI specification
type Formstyle struct {
	Horizontalgap FormStylehorizontalGap `json:"horizontalGap,omitempty"`
	Outerpadding FormStyleouterPadding `json:"outerPadding,omitempty"`
	Verticalgap FormStyleverticalGap `json:"verticalGap,omitempty"`
}

// UpdateThemerequest represents the UpdateThemerequest schema from the OpenAPI specification
type UpdateThemerequest struct {
	Updatedtheme UpdateThemerequestupdatedTheme `json:"updatedTheme"` // Saves the data binding information for a theme.
}

// FormSummarydataType represents the FormSummarydataType schema from the OpenAPI specification
type FormSummarydataType struct {
	Datasourcetype interface{} `json:"dataSourceType"`
	Datatypename interface{} `json:"dataTypeName"`
}

// CreateThemeRequest represents the CreateThemeRequest schema from the OpenAPI specification
type CreateThemeRequest struct {
	Themetocreate CreateThemeRequestthemeToCreate `json:"themeToCreate"`
}

// ReactStartCodegenJobData represents the ReactStartCodegenJobData schema from the OpenAPI specification
type ReactStartCodegenJobData struct {
	Rendertypedeclarations interface{} `json:"renderTypeDeclarations,omitempty"`
	Script interface{} `json:"script,omitempty"`
	Target interface{} `json:"target,omitempty"`
	Apiconfiguration ReactStartCodegenJobDataapiConfiguration `json:"apiConfiguration,omitempty"`
	Inlinesourcemap interface{} `json:"inlineSourceMap,omitempty"`
	Module interface{} `json:"module,omitempty"`
}

// GetThemeResponsetheme represents the GetThemeResponsetheme schema from the OpenAPI specification
type GetThemeResponsetheme struct {
	Name interface{} `json:"name"`
	Tags interface{} `json:"tags,omitempty"`
	Values interface{} `json:"values"`
	Id interface{} `json:"id"`
	Environmentname interface{} `json:"environmentName"`
	Overrides interface{} `json:"overrides,omitempty"`
	Appid interface{} `json:"appId"`
	Createdat interface{} `json:"createdAt"`
	Modifiedat interface{} `json:"modifiedAt,omitempty"`
}

// ComponentBindingPropertiesValueProperties represents the ComponentBindingPropertiesValueProperties schema from the OpenAPI specification
type ComponentBindingPropertiesValueProperties struct {
	Key interface{} `json:"key,omitempty"`
	Model interface{} `json:"model,omitempty"`
	Predicates interface{} `json:"predicates,omitempty"`
	Slotname interface{} `json:"slotName,omitempty"`
	Userattribute interface{} `json:"userAttribute,omitempty"`
	Bucket interface{} `json:"bucket,omitempty"`
	Defaultvalue interface{} `json:"defaultValue,omitempty"`
	Field interface{} `json:"field,omitempty"`
}

// ListFormsResponse represents the ListFormsResponse schema from the OpenAPI specification
type ListFormsResponse struct {
	Entities interface{} `json:"entities"`
	Nexttoken interface{} `json:"nextToken,omitempty"`
}

// CodegenGenericDataEnum represents the CodegenGenericDataEnum schema from the OpenAPI specification
type CodegenGenericDataEnum struct {
	Values interface{} `json:"values"`
}

// PutMetadataFlagrequest represents the PutMetadataFlagrequest schema from the OpenAPI specification
type PutMetadataFlagrequest struct {
	Body PutMetadataFlagrequestbody `json:"body"` // Stores the metadata information about a feature on a form.
}

// UpdateFormRequestupdatedForm represents the UpdateFormRequestupdatedForm schema from the OpenAPI specification
type UpdateFormRequestupdatedForm struct {
	Fields interface{} `json:"fields,omitempty"`
	Sectionalelements interface{} `json:"sectionalElements,omitempty"`
	Cta CreateFormrequestformToCreatecta `json:"cta,omitempty"`
	Formactiontype interface{} `json:"formActionType,omitempty"`
	Datatype CreateFormrequestformToCreatedataType `json:"dataType,omitempty"`
	Labeldecorator interface{} `json:"labelDecorator,omitempty"`
	Name interface{} `json:"name,omitempty"`
	Style CreateFormrequestformToCreatestyle `json:"style,omitempty"`
	Schemaversion interface{} `json:"schemaVersion,omitempty"`
}

// StartCodegenJobData represents the StartCodegenJobData schema from the OpenAPI specification
type StartCodegenJobData struct {
	Features StartCodegenJobrequestcodegenJobToCreatefeatures `json:"features,omitempty"`
	Genericdataschema StartCodegenJobrequestcodegenJobToCreategenericDataSchema `json:"genericDataSchema,omitempty"`
	Renderconfig StartCodegenJobrequestcodegenJobToCreaterenderConfig `json:"renderConfig"`
	Tags interface{} `json:"tags,omitempty"`
	Autogenerateforms interface{} `json:"autoGenerateForms,omitempty"`
}

// ComponentOverridesValue represents the ComponentOverridesValue schema from the OpenAPI specification
type ComponentOverridesValue struct {
}

// ListCodegenJobsResponse represents the ListCodegenJobsResponse schema from the OpenAPI specification
type ListCodegenJobsResponse struct {
	Entities interface{} `json:"entities"`
	Nexttoken interface{} `json:"nextToken,omitempty"`
}

// ComponentDataConfigurationpredicate represents the ComponentDataConfigurationpredicate schema from the OpenAPI specification
type ComponentDataConfigurationpredicate struct {
	Operand interface{} `json:"operand,omitempty"`
	Operandtype interface{} `json:"operandType,omitempty"`
	Operator interface{} `json:"operator,omitempty"`
	Or interface{} `json:"or,omitempty"`
	And interface{} `json:"and,omitempty"`
	Field interface{} `json:"field,omitempty"`
}

// ActionParameterstarget represents the ActionParameterstarget schema from the OpenAPI specification
type ActionParameterstarget struct {
	TypeField interface{} `json:"type,omitempty"`
	Bindings interface{} `json:"bindings,omitempty"`
	Defaultvalue interface{} `json:"defaultValue,omitempty"`
	Componentname interface{} `json:"componentName,omitempty"`
	Value interface{} `json:"value,omitempty"`
	Condition ComponentPropertycondition `json:"condition,omitempty"`
	Event interface{} `json:"event,omitempty"`
	Collectionbindingproperties ComponentPropertycollectionBindingProperties `json:"collectionBindingProperties,omitempty"`
	Importedvalue interface{} `json:"importedValue,omitempty"`
	Property interface{} `json:"property,omitempty"`
	Userattribute interface{} `json:"userAttribute,omitempty"`
	Concat interface{} `json:"concat,omitempty"`
	Configured interface{} `json:"configured,omitempty"`
	Bindingproperties ComponentPropertybindingProperties `json:"bindingProperties,omitempty"`
	Model interface{} `json:"model,omitempty"`
}

// SectionalElement represents the SectionalElement schema from the OpenAPI specification
type SectionalElement struct {
	Text interface{} `json:"text,omitempty"`
	TypeField interface{} `json:"type"`
	Excluded interface{} `json:"excluded,omitempty"`
	Level interface{} `json:"level,omitempty"`
	Orientation interface{} `json:"orientation,omitempty"`
	Position SectionalElementposition `json:"position,omitempty"`
}

// MutationActionSetStateParameter represents the MutationActionSetStateParameter schema from the OpenAPI specification
type MutationActionSetStateParameter struct {
	Componentname interface{} `json:"componentName"`
	Property interface{} `json:"property"`
	Set MutationActionSetStateParameterset `json:"set"`
}

// CodegenFeatureFlags represents the CodegenFeatureFlags schema from the OpenAPI specification
type CodegenFeatureFlags struct {
	Isnonmodelsupported interface{} `json:"isNonModelSupported,omitempty"`
	Isrelationshipsupported interface{} `json:"isRelationshipSupported,omitempty"`
}

// CodegenJobRenderConfig represents the CodegenJobRenderConfig schema from the OpenAPI specification
type CodegenJobRenderConfig struct {
	React CodegenJobRenderConfigreact `json:"react,omitempty"`
}

// GetFormResponse represents the GetFormResponse schema from the OpenAPI specification
type GetFormResponse struct {
	Form GetFormResponseform `json:"form,omitempty"`
}

// ComponentEvents represents the ComponentEvents schema from the OpenAPI specification
type ComponentEvents struct {
}

// CodegenGenericDataNonModel represents the CodegenGenericDataNonModel schema from the OpenAPI specification
type CodegenGenericDataNonModel struct {
	Fields interface{} `json:"fields"`
}

// GetThemeResponse represents the GetThemeResponse schema from the OpenAPI specification
type GetThemeResponse struct {
	Theme GetThemeResponsetheme `json:"theme,omitempty"`
}

// CodegenGenericDataNonModels represents the CodegenGenericDataNonModels schema from the OpenAPI specification
type CodegenGenericDataNonModels struct {
}

// ThemeValues represents the ThemeValues schema from the OpenAPI specification
type ThemeValues struct {
	Key interface{} `json:"key,omitempty"`
	Value ThemeValuesvalue `json:"value,omitempty"`
}

// ComponentEventparameters represents the ComponentEventparameters schema from the OpenAPI specification
type ComponentEventparameters struct {
	Model interface{} `json:"model,omitempty"`
	Anchor ActionParametersanchor `json:"anchor,omitempty"`
	Global ActionParametersglobal `json:"global,omitempty"`
	Id ActionParametersid `json:"id,omitempty"`
	State ActionParametersstate `json:"state,omitempty"`
	TypeField ActionParameterstype `json:"type,omitempty"`
	Url ActionParametersurl `json:"url,omitempty"`
	Fields interface{} `json:"fields,omitempty"`
	Target ActionParameterstarget `json:"target,omitempty"`
}

// ApiConfiguration represents the ApiConfiguration schema from the OpenAPI specification
type ApiConfiguration struct {
	Datastoreconfig interface{} `json:"dataStoreConfig,omitempty"`
	Graphqlconfig ApiConfigurationgraphQLConfig `json:"graphQLConfig,omitempty"`
	Noapiconfig interface{} `json:"noApiConfig,omitempty"`
}

// CreateFormrequestformToCreatecta represents the CreateFormrequestformToCreatecta schema from the OpenAPI specification
type CreateFormrequestformToCreatecta struct {
	Position interface{} `json:"position,omitempty"`
	Submit FormCTAsubmit `json:"submit,omitempty"`
	Cancel FormCTAcancel `json:"cancel,omitempty"`
	Clear FormCTAclear `json:"clear,omitempty"`
}

// UpdateFormrequest represents the UpdateFormrequest schema from the OpenAPI specification
type UpdateFormrequest struct {
	Updatedform UpdateFormrequestupdatedForm `json:"updatedForm"` // Updates and saves all of the information about a form, based on form ID.
}

// PutMetadataFlagrequestbody represents the PutMetadataFlagrequestbody schema from the OpenAPI specification
type PutMetadataFlagrequestbody struct {
	Newvalue interface{} `json:"newValue,omitempty"`
}

// ActionParametersglobal represents the ActionParametersglobal schema from the OpenAPI specification
type ActionParametersglobal struct {
	Event interface{} `json:"event,omitempty"`
	Collectionbindingproperties ComponentPropertycollectionBindingProperties `json:"collectionBindingProperties,omitempty"`
	Importedvalue interface{} `json:"importedValue,omitempty"`
	Property interface{} `json:"property,omitempty"`
	Userattribute interface{} `json:"userAttribute,omitempty"`
	Concat interface{} `json:"concat,omitempty"`
	Configured interface{} `json:"configured,omitempty"`
	Bindingproperties ComponentPropertybindingProperties `json:"bindingProperties,omitempty"`
	Model interface{} `json:"model,omitempty"`
	TypeField interface{} `json:"type,omitempty"`
	Bindings interface{} `json:"bindings,omitempty"`
	Defaultvalue interface{} `json:"defaultValue,omitempty"`
	Componentname interface{} `json:"componentName,omitempty"`
	Value interface{} `json:"value,omitempty"`
	Condition ComponentPropertycondition `json:"condition,omitempty"`
}

// ComponentBindingPropertiesValue represents the ComponentBindingPropertiesValue schema from the OpenAPI specification
type ComponentBindingPropertiesValue struct {
	TypeField interface{} `json:"type,omitempty"`
	Bindingproperties ComponentBindingPropertiesValuebindingProperties `json:"bindingProperties,omitempty"`
	Defaultvalue interface{} `json:"defaultValue,omitempty"`
}

// UpdateThemeData represents the UpdateThemeData schema from the OpenAPI specification
type UpdateThemeData struct {
	Id interface{} `json:"id,omitempty"`
	Name interface{} `json:"name,omitempty"`
	Overrides interface{} `json:"overrides,omitempty"`
	Values interface{} `json:"values"`
}

// CodegenGenericDataRelationshipType represents the CodegenGenericDataRelationshipType schema from the OpenAPI specification
type CodegenGenericDataRelationshipType struct {
	TypeField interface{} `json:"type"`
	Associatedfields interface{} `json:"associatedFields,omitempty"`
	Canunlinkassociatedmodel interface{} `json:"canUnlinkAssociatedModel,omitempty"`
	Relatedjoinfieldname interface{} `json:"relatedJoinFieldName,omitempty"`
	Relatedmodelfields interface{} `json:"relatedModelFields,omitempty"`
	Relatedmodelname interface{} `json:"relatedModelName"`
	Ishasmanyindex interface{} `json:"isHasManyIndex,omitempty"`
	Relatedjointablename interface{} `json:"relatedJoinTableName,omitempty"`
	Belongstofieldonrelatedmodel interface{} `json:"belongsToFieldOnRelatedModel,omitempty"`
}

// ListCodegenJobsRequest represents the ListCodegenJobsRequest schema from the OpenAPI specification
type ListCodegenJobsRequest struct {
}

// PutMetadataFlagRequest represents the PutMetadataFlagRequest schema from the OpenAPI specification
type PutMetadataFlagRequest struct {
	Body PutMetadataFlagRequestbody `json:"body"`
}

// RefreshTokenRequest represents the RefreshTokenRequest schema from the OpenAPI specification
type RefreshTokenRequest struct {
	Refreshtokenbody RefreshTokenRequestrefreshTokenBody `json:"refreshTokenBody"`
}

// FormInputBindingPropertiesValue represents the FormInputBindingPropertiesValue schema from the OpenAPI specification
type FormInputBindingPropertiesValue struct {
	TypeField interface{} `json:"type,omitempty"`
	Bindingproperties FormInputBindingPropertiesValuebindingProperties `json:"bindingProperties,omitempty"`
}

// UpdateComponentrequest represents the UpdateComponentrequest schema from the OpenAPI specification
type UpdateComponentrequest struct {
	Updatedcomponent UpdateComponentrequestupdatedComponent `json:"updatedComponent"` // Updates and saves all of the information about a component, based on component ID.
}

// DeleteThemeRequest represents the DeleteThemeRequest schema from the OpenAPI specification
type DeleteThemeRequest struct {
}

// ValueMapping represents the ValueMapping schema from the OpenAPI specification
type ValueMapping struct {
	Value ValueMappingvalue `json:"value"`
	Displayvalue ValueMappingdisplayValue `json:"displayValue,omitempty"`
}

// RefreshTokenrequestrefreshTokenBody represents the RefreshTokenrequestrefreshTokenBody schema from the OpenAPI specification
type RefreshTokenrequestrefreshTokenBody struct {
	Token interface{} `json:"token,omitempty"`
	Clientid interface{} `json:"clientId,omitempty"`
}

// StartCodegenJobrequestcodegenJobToCreate represents the StartCodegenJobrequestcodegenJobToCreate schema from the OpenAPI specification
type StartCodegenJobrequestcodegenJobToCreate struct {
	Tags interface{} `json:"tags,omitempty"`
	Autogenerateforms interface{} `json:"autoGenerateForms,omitempty"`
	Features StartCodegenJobrequestcodegenJobToCreatefeatures `json:"features,omitempty"`
	Genericdataschema StartCodegenJobrequestcodegenJobToCreategenericDataSchema `json:"genericDataSchema,omitempty"`
	Renderconfig StartCodegenJobrequestcodegenJobToCreaterenderConfig `json:"renderConfig,omitempty"`
}

// GetComponentResponsecomponent represents the GetComponentResponsecomponent schema from the OpenAPI specification
type GetComponentResponsecomponent struct {
	Appid interface{} `json:"appId"`
	Createdat interface{} `json:"createdAt"`
	Sourceid interface{} `json:"sourceId,omitempty"`
	Children interface{} `json:"children,omitempty"`
	Name interface{} `json:"name"`
	Componenttype interface{} `json:"componentType"`
	Modifiedat interface{} `json:"modifiedAt,omitempty"`
	Overrides interface{} `json:"overrides"`
	Schemaversion interface{} `json:"schemaVersion,omitempty"`
	Tags interface{} `json:"tags,omitempty"`
	Bindingproperties interface{} `json:"bindingProperties"`
	Variants interface{} `json:"variants"`
	Events interface{} `json:"events,omitempty"`
	Id interface{} `json:"id"`
	Environmentname interface{} `json:"environmentName"`
	Collectionproperties interface{} `json:"collectionProperties,omitempty"`
	Properties interface{} `json:"properties"` // Describes the component's properties. You can't specify <code>tags</code> as a valid property for <code>properties</code>.
}

// FormInputBindingPropertiesValuebindingProperties represents the FormInputBindingPropertiesValuebindingProperties schema from the OpenAPI specification
type FormInputBindingPropertiesValuebindingProperties struct {
	Model interface{} `json:"model,omitempty"`
}

// CodegenGenericDataFields represents the CodegenGenericDataFields schema from the OpenAPI specification
type CodegenGenericDataFields struct {
}

// Formcta represents the Formcta schema from the OpenAPI specification
type Formcta struct {
	Cancel FormCTAcancel `json:"cancel,omitempty"`
	Clear FormCTAclear `json:"clear,omitempty"`
	Position interface{} `json:"position,omitempty"`
	Submit FormCTAsubmit `json:"submit,omitempty"`
}

// ExchangeCodeForTokenrequest represents the ExchangeCodeForTokenrequest schema from the OpenAPI specification
type ExchangeCodeForTokenrequest struct {
	Request ExchangeCodeForTokenrequestrequest `json:"request"` // Describes the configuration of a request to exchange an access code for a token.
}

// CreateFormrequestformToCreate represents the CreateFormrequestformToCreate schema from the OpenAPI specification
type CreateFormrequestformToCreate struct {
	Fields interface{} `json:"fields,omitempty"`
	Name interface{} `json:"name,omitempty"`
	Labeldecorator interface{} `json:"labelDecorator,omitempty"`
	Schemaversion interface{} `json:"schemaVersion,omitempty"`
	Style CreateFormrequestformToCreatestyle `json:"style,omitempty"`
	Cta CreateFormrequestformToCreatecta `json:"cta,omitempty"`
	Datatype CreateFormrequestformToCreatedataType `json:"dataType,omitempty"`
	Formactiontype interface{} `json:"formActionType,omitempty"`
	Sectionalelements interface{} `json:"sectionalElements,omitempty"`
	Tags interface{} `json:"tags,omitempty"`
}

// ExchangeCodeForTokenRequestBody represents the ExchangeCodeForTokenRequestBody schema from the OpenAPI specification
type ExchangeCodeForTokenRequestBody struct {
	Clientid interface{} `json:"clientId,omitempty"`
	Code interface{} `json:"code"`
	Redirecturi interface{} `json:"redirectUri"`
}

// DataStoreRenderConfig represents the DataStoreRenderConfig schema from the OpenAPI specification
type DataStoreRenderConfig struct {
}

// FormStyleouterPadding represents the FormStyleouterPadding schema from the OpenAPI specification
type FormStyleouterPadding struct {
	Tokenreference interface{} `json:"tokenReference,omitempty"`
	Value interface{} `json:"value,omitempty"`
}

// ComponentVariantValues represents the ComponentVariantValues schema from the OpenAPI specification
type ComponentVariantValues struct {
}

// MutationActionSetStateParameterset represents the MutationActionSetStateParameterset schema from the OpenAPI specification
type MutationActionSetStateParameterset struct {
	Model interface{} `json:"model,omitempty"`
	TypeField interface{} `json:"type,omitempty"`
	Bindings interface{} `json:"bindings,omitempty"`
	Defaultvalue interface{} `json:"defaultValue,omitempty"`
	Componentname interface{} `json:"componentName,omitempty"`
	Value interface{} `json:"value,omitempty"`
	Condition ComponentPropertycondition `json:"condition,omitempty"`
	Event interface{} `json:"event,omitempty"`
	Collectionbindingproperties ComponentPropertycollectionBindingProperties `json:"collectionBindingProperties,omitempty"`
	Importedvalue interface{} `json:"importedValue,omitempty"`
	Property interface{} `json:"property,omitempty"`
	Userattribute interface{} `json:"userAttribute,omitempty"`
	Concat interface{} `json:"concat,omitempty"`
	Configured interface{} `json:"configured,omitempty"`
	Bindingproperties ComponentPropertybindingProperties `json:"bindingProperties,omitempty"`
}

// GetComponentRequest represents the GetComponentRequest schema from the OpenAPI specification
type GetComponentRequest struct {
}

// ActionParametersstate represents the ActionParametersstate schema from the OpenAPI specification
type ActionParametersstate struct {
	Componentname interface{} `json:"componentName"`
	Property interface{} `json:"property"`
	Set MutationActionSetStateParameterset `json:"set"`
}

// ComponentCollectionProperties represents the ComponentCollectionProperties schema from the OpenAPI specification
type ComponentCollectionProperties struct {
}

// ListComponentsResponse represents the ListComponentsResponse schema from the OpenAPI specification
type ListComponentsResponse struct {
	Entities interface{} `json:"entities"`
	Nexttoken interface{} `json:"nextToken,omitempty"`
}

// ComponentConditionPropertyelse represents the ComponentConditionPropertyelse schema from the OpenAPI specification
type ComponentConditionPropertyelse struct {
	TypeField interface{} `json:"type,omitempty"`
	Bindings interface{} `json:"bindings,omitempty"`
	Defaultvalue interface{} `json:"defaultValue,omitempty"`
	Componentname interface{} `json:"componentName,omitempty"`
	Value interface{} `json:"value,omitempty"`
	Condition ComponentPropertycondition `json:"condition,omitempty"`
	Event interface{} `json:"event,omitempty"`
	Collectionbindingproperties ComponentPropertycollectionBindingProperties `json:"collectionBindingProperties,omitempty"`
	Importedvalue interface{} `json:"importedValue,omitempty"`
	Property interface{} `json:"property,omitempty"`
	Userattribute interface{} `json:"userAttribute,omitempty"`
	Concat interface{} `json:"concat,omitempty"`
	Configured interface{} `json:"configured,omitempty"`
	Bindingproperties ComponentPropertybindingProperties `json:"bindingProperties,omitempty"`
	Model interface{} `json:"model,omitempty"`
}

// CodegenJob represents the CodegenJob schema from the OpenAPI specification
type CodegenJob struct {
	Appid interface{} `json:"appId"`
	Genericdataschema CodegenJobGenericDataSchema `json:"genericDataSchema,omitempty"` // Describes the data schema for a code generation job.
	Modifiedat interface{} `json:"modifiedAt,omitempty"`
	Environmentname interface{} `json:"environmentName"`
	Statusmessage interface{} `json:"statusMessage,omitempty"`
	Autogenerateforms interface{} `json:"autoGenerateForms,omitempty"`
	Id interface{} `json:"id"`
	Renderconfig CodegenJobRenderConfig `json:"renderConfig,omitempty"` // Describes the configuration information for rendering the UI component associated with the code generation job.
	Tags interface{} `json:"tags,omitempty"`
	Asset CodegenJobasset `json:"asset,omitempty"`
	Createdat interface{} `json:"createdAt,omitempty"`
	Features CodegenFeatureFlags `json:"features,omitempty"` // Describes the feature flags that you can specify for a code generation job.
	Status interface{} `json:"status,omitempty"`
}

// CreateComponentData represents the CreateComponentData schema from the OpenAPI specification
type CreateComponentData struct {
	Children interface{} `json:"children,omitempty"`
	Properties interface{} `json:"properties"` // Describes the component's properties.
	Variants interface{} `json:"variants"`
	Collectionproperties interface{} `json:"collectionProperties,omitempty"`
	Componenttype interface{} `json:"componentType"`
	Events interface{} `json:"events,omitempty"`
	Bindingproperties interface{} `json:"bindingProperties"`
	Sourceid interface{} `json:"sourceId,omitempty"`
	Name interface{} `json:"name"`
	Overrides interface{} `json:"overrides"`
	Schemaversion interface{} `json:"schemaVersion,omitempty"`
	Tags interface{} `json:"tags,omitempty"`
}

// ComponentPropertyBindingProperties represents the ComponentPropertyBindingProperties schema from the OpenAPI specification
type ComponentPropertyBindingProperties struct {
	Field interface{} `json:"field,omitempty"`
	Property interface{} `json:"property"`
}

// FormStylehorizontalGap represents the FormStylehorizontalGap schema from the OpenAPI specification
type FormStylehorizontalGap struct {
	Tokenreference interface{} `json:"tokenReference,omitempty"`
	Value interface{} `json:"value,omitempty"`
}

// FieldInputConfig represents the FieldInputConfig schema from the OpenAPI specification
type FieldInputConfig struct {
	Minvalue interface{} `json:"minValue,omitempty"`
	Readonly interface{} `json:"readOnly,omitempty"`
	Value interface{} `json:"value,omitempty"`
	Required interface{} `json:"required,omitempty"`
	Defaultchecked interface{} `json:"defaultChecked,omitempty"`
	Name interface{} `json:"name,omitempty"`
	Fileuploaderconfig FieldInputConfigfileUploaderConfig `json:"fileUploaderConfig,omitempty"`
	Valuemappings FieldInputConfigvalueMappings `json:"valueMappings,omitempty"`
	Step interface{} `json:"step,omitempty"`
	Defaultcountrycode interface{} `json:"defaultCountryCode,omitempty"`
	Isarray interface{} `json:"isArray,omitempty"`
	Descriptivetext interface{} `json:"descriptiveText,omitempty"`
	TypeField interface{} `json:"type"`
	Maxvalue interface{} `json:"maxValue,omitempty"`
	Placeholder interface{} `json:"placeholder,omitempty"`
	Defaultvalue interface{} `json:"defaultValue,omitempty"`
}

// FormStyleverticalGap represents the FormStyleverticalGap schema from the OpenAPI specification
type FormStyleverticalGap struct {
	Tokenreference interface{} `json:"tokenReference,omitempty"`
	Value interface{} `json:"value,omitempty"`
}

// CreateFormRequest represents the CreateFormRequest schema from the OpenAPI specification
type CreateFormRequest struct {
	Formtocreate CreateFormRequestformToCreate `json:"formToCreate"`
}

// ThemeValue represents the ThemeValue schema from the OpenAPI specification
type ThemeValue struct {
	Children interface{} `json:"children,omitempty"`
	Value interface{} `json:"value,omitempty"`
}

// FieldValidationConfiguration represents the FieldValidationConfiguration schema from the OpenAPI specification
type FieldValidationConfiguration struct {
	Numvalues interface{} `json:"numValues,omitempty"`
	Strvalues interface{} `json:"strValues,omitempty"`
	TypeField interface{} `json:"type"`
	Validationmessage interface{} `json:"validationMessage,omitempty"`
}

// RefreshTokenRequestrefreshTokenBody represents the RefreshTokenRequestrefreshTokenBody schema from the OpenAPI specification
type RefreshTokenRequestrefreshTokenBody struct {
	Clientid interface{} `json:"clientId,omitempty"`
	Token interface{} `json:"token"`
}

// CreateFormData represents the CreateFormData schema from the OpenAPI specification
type CreateFormData struct {
	Style CreateFormrequestformToCreatestyle `json:"style"`
	Tags interface{} `json:"tags,omitempty"`
	Datatype CreateFormrequestformToCreatedataType `json:"dataType"`
	Formactiontype interface{} `json:"formActionType"`
	Fields interface{} `json:"fields"`
	Name interface{} `json:"name"`
	Sectionalelements interface{} `json:"sectionalElements"`
	Schemaversion interface{} `json:"schemaVersion"`
	Cta CreateFormrequestformToCreatecta `json:"cta,omitempty"`
	Labeldecorator interface{} `json:"labelDecorator,omitempty"`
}

// CreateComponentrequestcomponentToCreate represents the CreateComponentrequestcomponentToCreate schema from the OpenAPI specification
type CreateComponentrequestcomponentToCreate struct {
	Sourceid interface{} `json:"sourceId,omitempty"`
	Componenttype interface{} `json:"componentType,omitempty"`
	Properties interface{} `json:"properties,omitempty"` // Describes the component's properties.
	Variants interface{} `json:"variants,omitempty"`
	Children interface{} `json:"children,omitempty"`
	Collectionproperties interface{} `json:"collectionProperties,omitempty"`
	Schemaversion interface{} `json:"schemaVersion,omitempty"`
	Events interface{} `json:"events,omitempty"`
	Name interface{} `json:"name,omitempty"`
	Tags interface{} `json:"tags,omitempty"`
	Bindingproperties interface{} `json:"bindingProperties,omitempty"`
	Overrides interface{} `json:"overrides,omitempty"`
}

// CreateThemeRequestthemeToCreate represents the CreateThemeRequestthemeToCreate schema from the OpenAPI specification
type CreateThemeRequestthemeToCreate struct {
	Tags interface{} `json:"tags,omitempty"`
	Values interface{} `json:"values"`
	Name interface{} `json:"name"`
	Overrides interface{} `json:"overrides,omitempty"`
}

// ComponentConditionPropertythen represents the ComponentConditionPropertythen schema from the OpenAPI specification
type ComponentConditionPropertythen struct {
	Property interface{} `json:"property,omitempty"`
	Userattribute interface{} `json:"userAttribute,omitempty"`
	Concat interface{} `json:"concat,omitempty"`
	Configured interface{} `json:"configured,omitempty"`
	Bindingproperties ComponentPropertybindingProperties `json:"bindingProperties,omitempty"`
	Model interface{} `json:"model,omitempty"`
	TypeField interface{} `json:"type,omitempty"`
	Bindings interface{} `json:"bindings,omitempty"`
	Defaultvalue interface{} `json:"defaultValue,omitempty"`
	Componentname interface{} `json:"componentName,omitempty"`
	Value interface{} `json:"value,omitempty"`
	Condition ComponentPropertycondition `json:"condition,omitempty"`
	Event interface{} `json:"event,omitempty"`
	Collectionbindingproperties ComponentPropertycollectionBindingProperties `json:"collectionBindingProperties,omitempty"`
	Importedvalue interface{} `json:"importedValue,omitempty"`
}

// Tags represents the Tags schema from the OpenAPI specification
type Tags struct {
}

// ComponentSummary represents the ComponentSummary schema from the OpenAPI specification
type ComponentSummary struct {
	Componenttype interface{} `json:"componentType"`
	Environmentname interface{} `json:"environmentName"`
	Id interface{} `json:"id"`
	Name interface{} `json:"name"`
	Appid interface{} `json:"appId"`
}

// FormStyle represents the FormStyle schema from the OpenAPI specification
type FormStyle struct {
	Outerpadding FormStyleouterPadding `json:"outerPadding,omitempty"`
	Verticalgap FormStyleverticalGap `json:"verticalGap,omitempty"`
	Horizontalgap FormStylehorizontalGap `json:"horizontalGap,omitempty"`
}

// ThemeValuesvalue represents the ThemeValuesvalue schema from the OpenAPI specification
type ThemeValuesvalue struct {
	Children interface{} `json:"children,omitempty"`
	Value interface{} `json:"value,omitempty"`
}

// ComponentChild represents the ComponentChild schema from the OpenAPI specification
type ComponentChild struct {
	Name interface{} `json:"name"`
	Properties interface{} `json:"properties"` // Describes the properties of the child component. You can't specify <code>tags</code> as a valid property for <code>properties</code>.
	Sourceid interface{} `json:"sourceId,omitempty"`
	Children interface{} `json:"children,omitempty"`
	Componenttype interface{} `json:"componentType"`
	Events interface{} `json:"events,omitempty"`
}

// ComponentPropertycollectionBindingProperties represents the ComponentPropertycollectionBindingProperties schema from the OpenAPI specification
type ComponentPropertycollectionBindingProperties struct {
	Field interface{} `json:"field,omitempty"`
	Property interface{} `json:"property"`
}

// ExportFormsRequest represents the ExportFormsRequest schema from the OpenAPI specification
type ExportFormsRequest struct {
}
