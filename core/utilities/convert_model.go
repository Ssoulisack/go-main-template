package utilities

import (
	"github.com/jinzhu/copier"
)

// ConvertModelToEntity converts a model to an entity using generics
// func ConvertModelToEntity[T any, U any](model T, entity *U) error {
// 	err := copier.Copy(entity, model)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// ConvertModelToEntity converts a single model to an entity
func ConvertModelToEntity[T any, U any](model T) U {
	var entity U
	copier.Copy(&entity, model)
	return entity
}

// ConvertEntityToModel converts a single entity to a model
func ConvertEntityToModel[T any, U any](entity T) U {
	var model U
	copier.Copy(&model, entity)
	return model
}

// ConvertModelToEntities converts a slice of models to a slice of entities
func ConvertModelToEntities[T any, U any](models []T) []U {
	var entities []U
	for _, model := range models {
		var entity U
		copier.Copy(&entity, model)
		entities = append(entities, entity)
	}
	return entities
}

// ConvertEntitiesToModels converts a slice of entities to a slice of models
func ConvertEntitiesToModels[T any, U any](entities []T) []U {
	var models []U
	for _, entity := range entities {
		var model U
		copier.Copy(&model, entity)
		models = append(models, model)
	}
	return models
}


