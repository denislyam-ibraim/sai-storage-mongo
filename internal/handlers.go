package internal

import (
	"encoding/json"

	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
	"go.uber.org/zap"

	"github.com/saiset-co/sai-storage-mongo/internal/actions"
	"github.com/saiset-co/sai-storage-mongo/logger"
	"github.com/saiset-co/sai-storage-mongo/types"
	"github.com/saiset-co/saiService"
)

func (is InternalService) NewHandler() saiService.Handler {
	return saiService.Handler{
		"create": saiService.HandlerElement{
			Name:        "Create documents",
			Description: "Create documents",
			Function: func(data interface{}, metadata interface{}) (interface{}, int, error) {
				request, err := is.convertRequest(data, "create")
				if err != nil {
					return nil, 500, err
				}

				return actions.NewSaveAction(is.Client).Handle(request)
			},
		},
		"read": saiService.HandlerElement{
			Name:        "Read documents",
			Description: "Read documents",
			Function: func(data interface{}, metadata interface{}) (interface{}, int, error) {
				request, err := is.convertRequest(data, "read")
				if err != nil {
					return nil, 500, err
				}

				return actions.NewGetAction(is.Client).Handle(request)
			},
		},
		"update": saiService.HandlerElement{
			Name:        "Update documents",
			Description: "Update documents",
			Function: func(data interface{}, metadata interface{}) (interface{}, int, error) {
				request, err := is.convertRequest(data, "update")
				if err != nil {
					return nil, 500, err
				}

				return actions.NewUpdateAction(is.Client).Handle(request)
			},
		},
		"upsert": saiService.HandlerElement{
			Name:        "Upsert documents",
			Description: "Upsert documents",
			Function: func(data interface{}, metadata interface{}) (interface{}, int, error) {
				request, err := is.convertRequest(data, "upsert")
				if err != nil {
					return nil, 500, err
				}

				return actions.NewUpsertAction(is.Client).Handle(request)
			},
		},
		"delete": saiService.HandlerElement{
			Name:        "Delete documents",
			Description: "Delete documents",
			Function: func(data interface{}, metadata interface{}) (interface{}, int, error) {
				request, err := is.convertRequest(data, "delete")
				if err != nil {
					return nil, 500, err
				}

				return actions.NewDeleteAction(is.Client).Handle(request)
			},
		},
		"aggregate": saiService.HandlerElement{
			Name:        "Aggregate documents",
			Description: "Aggregate documents",
			Function: func(data interface{}, metadata interface{}) (interface{}, int, error) {
				request, err := is.convertRequest(data, "aggregate")
				if err != nil {
					return nil, 500, err
				}

				return actions.NewAggregateAction(is.Client).Handle(request)
			},
		},
	}
}

func (is InternalService) convertRequest(data interface{}, requestType string) (types.IRequest, error) {
	switch requestType {
	case "create":
		request := types.CreateRequest{}
		dataJson, err := json.Marshal(data)
		if err != nil {
			logger.Logger.Error("convertRequest", zap.Error(err))
			return nil, errors.Wrap(err, "convertRequest - marshaling - save")
		}

		err = json.Unmarshal(dataJson, &request)
		if err != nil {
			logger.Logger.Error("convertRequest", zap.Error(err))
			return nil, errors.Wrap(err, "convertRequest - unmarshaling - save")
		}

		err = validator.New().Struct(request)
		if err != nil {
			logger.Logger.Error("convertRequest", zap.Error(err))
			return nil, errors.Wrap(err, "convertRequest - validation - save")
		}

		return request, nil
	case "read":
		request := types.ReadRequest{}
		dataJson, err := json.Marshal(data)
		if err != nil {
			logger.Logger.Error("convertRequest", zap.Error(err))
			return nil, errors.Wrap(err, "convertRequest - marshaling - get")
		}

		err = json.Unmarshal(dataJson, &request)
		if err != nil {
			logger.Logger.Error("convertRequest", zap.Error(err))
			return nil, errors.Wrap(err, "convertRequest - unmarshaling - get")
		}

		err = validator.New().Struct(request)
		if err != nil {
			logger.Logger.Error("convertRequest", zap.Error(err))
			return nil, errors.Wrap(err, "convertRequest - validation - get")
		}

		return request, nil
	case "update":
		request := types.UpdateRequest{}
		dataJson, err := json.Marshal(data)
		if err != nil {
			logger.Logger.Error("convertRequest", zap.Error(err))
			return nil, errors.Wrap(err, "convertRequest - marshaling - update")
		}

		err = json.Unmarshal(dataJson, &request)
		if err != nil {
			logger.Logger.Error("convertRequest", zap.Error(err))
			return nil, errors.Wrap(err, "convertRequest - unmarshaling - update")
		}

		err = validator.New().Struct(request)
		if err != nil {
			logger.Logger.Error("convertRequest", zap.Error(err))
			return nil, errors.Wrap(err, "convertRequest - validation - update")
		}

		return request, nil
	case "upsert":
		request := types.UpsertRequest{}
		dataJson, err := json.Marshal(data)
		if err != nil {
			logger.Logger.Error("convertRequest", zap.Error(err))
			return nil, errors.Wrap(err, "convertRequest - marshaling - update")
		}

		err = json.Unmarshal(dataJson, &request)
		if err != nil {
			logger.Logger.Error("convertRequest", zap.Error(err))
			return nil, errors.Wrap(err, "convertRequest - unmarshaling - update")
		}

		err = validator.New().Struct(request)
		if err != nil {
			logger.Logger.Error("convertRequest", zap.Error(err))
			return nil, errors.Wrap(err, "convertRequest - validation - update")
		}

		return request, nil
	case "delete":
		request := types.DeleteRequest{}
		dataJson, err := json.Marshal(data)
		if err != nil {
			logger.Logger.Error("convertRequest", zap.Error(err))
			return nil, errors.Wrap(err, "convertRequest - marshaling - delete")
		}

		err = json.Unmarshal(dataJson, &request)
		if err != nil {
			logger.Logger.Error("convertRequest", zap.Error(err))
			return nil, errors.Wrap(err, "convertRequest - unmarshaling - delete")
		}

		err = validator.New().Struct(request)
		if err != nil {
			logger.Logger.Error("convertRequest", zap.Error(err))
			return nil, errors.Wrap(err, "convertRequest - validation - delete")
		}

		return request, nil
	case "aggregate":
		request := types.AggregateRequest{}
		dataJson, err := json.Marshal(data)
		if err != nil {
			logger.Logger.Error("convertRequest", zap.Error(err))
			return nil, errors.Wrap(err, "convertRequest - marshaling - aggregate")
		}

		err = json.Unmarshal(dataJson, &request)
		if err != nil {
			logger.Logger.Error("convertRequest", zap.Error(err))
			return nil, errors.Wrap(err, "convertRequest - unmarshaling - aggregate")
		}

		err = validator.New().Struct(request)
		if err != nil {
			logger.Logger.Error("convertRequest", zap.Error(err))
			return nil, errors.Wrap(err, "convertRequest - validation - aggregate")
		}

		return request, nil
	}

	return nil, errors.Wrap(errors.New("no variable type"), "convertRequest")
}

func (is InternalService) convertData(data interface{}, requestType string) (types.IRequest, error) {
	switch requestType {
	case "create":
		request := types.CreateRequest{}
		dataJson, err := json.Marshal(data)
		if err != nil {
			logger.Logger.Error("convertData", zap.Error(err))
			return nil, errors.Wrap(err, "convertRequest - marshaling - save")
		}

		err = json.Unmarshal(dataJson, &request)
		if err != nil {
			logger.Logger.Error("convertData", zap.Error(err))
			return nil, errors.Wrap(err, "convertRequest - unmarshaling - save")
		}

		err = validator.New().Struct(request)
		if err != nil {
			logger.Logger.Error("convertData", zap.Error(err))
			return nil, errors.Wrap(err, "convertRequest - validation - save")
		}

		return request, nil
	case "read":
		request := types.ReadRequest{}
		dataJson, err := json.Marshal(data)
		if err != nil {
			logger.Logger.Error("convertData", zap.Error(err))
			return nil, errors.Wrap(err, "convertRequest - marshaling - get")
		}

		err = json.Unmarshal(dataJson, &request)
		if err != nil {
			logger.Logger.Error("convertData", zap.Error(err))
			return nil, errors.Wrap(err, "convertRequest - unmarshaling - get")
		}

		err = validator.New().Struct(request)
		if err != nil {
			logger.Logger.Error("convertData", zap.Error(err))
			return nil, errors.Wrap(err, "convertRequest - validation - get")
		}

		return request, nil
	case "update":
		request := types.UpdateRequest{}
		dataJson, err := json.Marshal(data)
		if err != nil {
			logger.Logger.Error("convertData", zap.Error(err))
			return nil, errors.Wrap(err, "convertRequest - marshaling - update")
		}

		err = json.Unmarshal(dataJson, &request)
		if err != nil {
			logger.Logger.Error("convertData", zap.Error(err))
			return nil, errors.Wrap(err, "convertRequest - unmarshaling - update")
		}

		err = validator.New().Struct(request)
		if err != nil {
			logger.Logger.Error("convertData", zap.Error(err))
			return nil, errors.Wrap(err, "convertRequest - validation - update")
		}

		return request, nil
	case "delete":
		request := types.DeleteRequest{}
		dataJson, err := json.Marshal(data)
		if err != nil {
			logger.Logger.Error("convertData", zap.Error(err))
			return nil, errors.Wrap(err, "convertRequest - marshaling - delete")
		}

		err = json.Unmarshal(dataJson, &request)
		if err != nil {
			logger.Logger.Error("convertData", zap.Error(err))
			return nil, errors.Wrap(err, "convertRequest - unmarshaling - delete")
		}

		err = validator.New().Struct(request)
		if err != nil {
			logger.Logger.Error("convertData", zap.Error(err))
			return nil, errors.Wrap(err, "convertRequest - validation - delete")
		}

		return request, nil
	}

	return nil, errors.Wrap(errors.New("no variable type"), "convertRequest")
}
