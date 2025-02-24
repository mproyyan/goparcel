// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package generated

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"sync/atomic"

	"github.com/99designs/gqlgen/graphql"
	"github.com/mproyyan/goparcel/internal/graphql/graph/model"
	"github.com/vektah/gqlparser/v2/ast"
)

// region    ************************** generated!.gotpl **************************

// endregion ************************** generated!.gotpl **************************

// region    ***************************** args.gotpl *****************************

// endregion ***************************** args.gotpl *****************************

// region    ************************** directives.gotpl **************************

// endregion ************************** directives.gotpl **************************

// region    **************************** field.gotpl *****************************

func (ec *executionContext) _UserEntity_id(ctx context.Context, field graphql.CollectedField, obj *model.UserEntity) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_UserEntity_id(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (any, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.ID, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalNID2string(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_UserEntity_id(_ context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "UserEntity",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type ID does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _UserEntity_user_id(ctx context.Context, field graphql.CollectedField, obj *model.UserEntity) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_UserEntity_user_id(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (any, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.UserID, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalNID2string(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_UserEntity_user_id(_ context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "UserEntity",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type ID does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _UserEntity_name(ctx context.Context, field graphql.CollectedField, obj *model.UserEntity) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_UserEntity_name(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (any, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Name, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalNString2string(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_UserEntity_name(_ context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "UserEntity",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type String does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _UserEntity_email(ctx context.Context, field graphql.CollectedField, obj *model.UserEntity) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_UserEntity_email(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (any, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Email, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalNString2string(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_UserEntity_email(_ context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "UserEntity",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type String does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _UserEntity_location(ctx context.Context, field graphql.CollectedField, obj *model.UserEntity) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_UserEntity_location(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (any, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Location, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*model.Location)
	fc.Result = res
	return ec.marshalOLocation2ᚖgithubᚗcomᚋmproyyanᚋgoparcelᚋinternalᚋgraphqlᚋgraphᚋmodelᚐLocation(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_UserEntity_location(_ context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "UserEntity",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			switch field.Name {
			case "id":
				return ec.fieldContext_Location_id(ctx, field)
			case "name":
				return ec.fieldContext_Location_name(ctx, field)
			case "type":
				return ec.fieldContext_Location_type(ctx, field)
			case "warehouse":
				return ec.fieldContext_Location_warehouse(ctx, field)
			case "address":
				return ec.fieldContext_Location_address(ctx, field)
			}
			return nil, fmt.Errorf("no field named %q was found under type Location", field.Name)
		},
	}
	return fc, nil
}

// endregion **************************** field.gotpl *****************************

// region    **************************** input.gotpl *****************************

func (ec *executionContext) unmarshalInputRegisterAsCarrierInput(ctx context.Context, obj any) (model.RegisterAsCarrierInput, error) {
	var it model.RegisterAsCarrierInput
	asMap := map[string]any{}
	for k, v := range obj.(map[string]any) {
		asMap[k] = v
	}

	fieldsInOrder := [...]string{"name", "email", "password", "location_id"}
	for _, k := range fieldsInOrder {
		v, ok := asMap[k]
		if !ok {
			continue
		}
		switch k {
		case "name":
			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("name"))
			data, err := ec.unmarshalNString2string(ctx, v)
			if err != nil {
				return it, err
			}
			it.Name = data
		case "email":
			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("email"))
			data, err := ec.unmarshalNString2string(ctx, v)
			if err != nil {
				return it, err
			}
			it.Email = data
		case "password":
			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("password"))
			data, err := ec.unmarshalNString2string(ctx, v)
			if err != nil {
				return it, err
			}
			it.Password = data
		case "location_id":
			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("location_id"))
			data, err := ec.unmarshalNString2string(ctx, v)
			if err != nil {
				return it, err
			}
			it.LocationID = data
		}
	}

	return it, nil
}

func (ec *executionContext) unmarshalInputRegisterAsCourierInput(ctx context.Context, obj any) (model.RegisterAsCourierInput, error) {
	var it model.RegisterAsCourierInput
	asMap := map[string]any{}
	for k, v := range obj.(map[string]any) {
		asMap[k] = v
	}

	fieldsInOrder := [...]string{"name", "email", "password", "location_id"}
	for _, k := range fieldsInOrder {
		v, ok := asMap[k]
		if !ok {
			continue
		}
		switch k {
		case "name":
			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("name"))
			data, err := ec.unmarshalNString2string(ctx, v)
			if err != nil {
				return it, err
			}
			it.Name = data
		case "email":
			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("email"))
			data, err := ec.unmarshalNString2string(ctx, v)
			if err != nil {
				return it, err
			}
			it.Email = data
		case "password":
			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("password"))
			data, err := ec.unmarshalNString2string(ctx, v)
			if err != nil {
				return it, err
			}
			it.Password = data
		case "location_id":
			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("location_id"))
			data, err := ec.unmarshalNString2string(ctx, v)
			if err != nil {
				return it, err
			}
			it.LocationID = data
		}
	}

	return it, nil
}

func (ec *executionContext) unmarshalInputRegisterAsOperatorInput(ctx context.Context, obj any) (model.RegisterAsOperatorInput, error) {
	var it model.RegisterAsOperatorInput
	asMap := map[string]any{}
	for k, v := range obj.(map[string]any) {
		asMap[k] = v
	}

	fieldsInOrder := [...]string{"name", "email", "password", "location_id", "type"}
	for _, k := range fieldsInOrder {
		v, ok := asMap[k]
		if !ok {
			continue
		}
		switch k {
		case "name":
			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("name"))
			data, err := ec.unmarshalNString2string(ctx, v)
			if err != nil {
				return it, err
			}
			it.Name = data
		case "email":
			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("email"))
			data, err := ec.unmarshalNString2string(ctx, v)
			if err != nil {
				return it, err
			}
			it.Email = data
		case "password":
			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("password"))
			data, err := ec.unmarshalNString2string(ctx, v)
			if err != nil {
				return it, err
			}
			it.Password = data
		case "location_id":
			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("location_id"))
			data, err := ec.unmarshalNString2string(ctx, v)
			if err != nil {
				return it, err
			}
			it.LocationID = data
		case "type":
			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("type"))
			data, err := ec.unmarshalNOperatorType2githubᚗcomᚋmproyyanᚋgoparcelᚋinternalᚋgraphqlᚋgraphᚋmodelᚐOperatorType(ctx, v)
			if err != nil {
				return it, err
			}
			it.Type = data
		}
	}

	return it, nil
}

// endregion **************************** input.gotpl *****************************

// region    ************************** interface.gotpl ***************************

// endregion ************************** interface.gotpl ***************************

// region    **************************** object.gotpl ****************************

var userEntityImplementors = []string{"UserEntity"}

func (ec *executionContext) _UserEntity(ctx context.Context, sel ast.SelectionSet, obj *model.UserEntity) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, userEntityImplementors)

	out := graphql.NewFieldSet(fields)
	deferred := make(map[string]*graphql.FieldSet)
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("UserEntity")
		case "id":
			out.Values[i] = ec._UserEntity_id(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				out.Invalids++
			}
		case "user_id":
			out.Values[i] = ec._UserEntity_user_id(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				out.Invalids++
			}
		case "name":
			out.Values[i] = ec._UserEntity_name(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				out.Invalids++
			}
		case "email":
			out.Values[i] = ec._UserEntity_email(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				out.Invalids++
			}
		case "location":
			out.Values[i] = ec._UserEntity_location(ctx, field, obj)
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch(ctx)
	if out.Invalids > 0 {
		return graphql.Null
	}

	atomic.AddInt32(&ec.deferred, int32(len(deferred)))

	for label, dfs := range deferred {
		ec.processDeferredGroup(graphql.DeferredGroup{
			Label:    label,
			Path:     graphql.GetPath(ctx),
			FieldSet: dfs,
			Context:  ctx,
		})
	}

	return out
}

// endregion **************************** object.gotpl ****************************

// region    ***************************** type.gotpl *****************************

func (ec *executionContext) unmarshalNOperatorType2githubᚗcomᚋmproyyanᚋgoparcelᚋinternalᚋgraphqlᚋgraphᚋmodelᚐOperatorType(ctx context.Context, v any) (model.OperatorType, error) {
	var res model.OperatorType
	err := res.UnmarshalGQL(v)
	return res, graphql.ErrorOnPath(ctx, err)
}

func (ec *executionContext) marshalNOperatorType2githubᚗcomᚋmproyyanᚋgoparcelᚋinternalᚋgraphqlᚋgraphᚋmodelᚐOperatorType(ctx context.Context, sel ast.SelectionSet, v model.OperatorType) graphql.Marshaler {
	return v
}

func (ec *executionContext) unmarshalNRegisterAsCarrierInput2githubᚗcomᚋmproyyanᚋgoparcelᚋinternalᚋgraphqlᚋgraphᚋmodelᚐRegisterAsCarrierInput(ctx context.Context, v any) (model.RegisterAsCarrierInput, error) {
	res, err := ec.unmarshalInputRegisterAsCarrierInput(ctx, v)
	return res, graphql.ErrorOnPath(ctx, err)
}

func (ec *executionContext) unmarshalNRegisterAsCourierInput2githubᚗcomᚋmproyyanᚋgoparcelᚋinternalᚋgraphqlᚋgraphᚋmodelᚐRegisterAsCourierInput(ctx context.Context, v any) (model.RegisterAsCourierInput, error) {
	res, err := ec.unmarshalInputRegisterAsCourierInput(ctx, v)
	return res, graphql.ErrorOnPath(ctx, err)
}

func (ec *executionContext) unmarshalNRegisterAsOperatorInput2githubᚗcomᚋmproyyanᚋgoparcelᚋinternalᚋgraphqlᚋgraphᚋmodelᚐRegisterAsOperatorInput(ctx context.Context, v any) (model.RegisterAsOperatorInput, error) {
	res, err := ec.unmarshalInputRegisterAsOperatorInput(ctx, v)
	return res, graphql.ErrorOnPath(ctx, err)
}

// endregion ***************************** type.gotpl *****************************
