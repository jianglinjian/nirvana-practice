package descriptors

import (
	"github.com/caicloud/nirvana/definition"

	"github.com/lichuan0620/nirvana-practice/pkg/apis/v1alpha1/handlers"
)

// ProductDescriptor builds and returns a Descriptor for all Product APIs.
func ProductDescriptor() definition.Descriptor {
	return definition.Descriptor{
		Path: "/products",
		Definitions: []definition.Definition{
			{
				Method:   definition.Create,
				Function: handlers.CreateProduct,
				Parameters: []definition.Parameter{
					definition.BodyParameterFor("product object to create"),
				},
				Results: definition.DataErrorResults("the create result (or error)"),
			},
			{
				Method:     definition.List,
				Function:   handlers.ListProducts,
				Parameters: []definition.Parameter{},
				Results:    definition.DataErrorResults("the list result (or error)"),
			},
		},
		Children: []definition.Descriptor{
			{
				Path: "/{product}",
				Definitions: []definition.Definition{
					{
						Method:   definition.Get,
						Function: handlers.GetProduct,
						Parameters: []definition.Parameter{
							productName(),
						},
						Results: definition.DataErrorResults("the get result (or error)"),
					},
					{
						Method:   definition.Update,
						Function: handlers.UpdateProduct,
						Parameters: []definition.Parameter{
							productName(),
							definition.BodyParameterFor(""),
						},
						Results: definition.DataErrorResults("the update result (or error)"),
					},
					{
						Method:   definition.Delete,
						Function: handlers.DeleteProduct,
						Parameters: []definition.Parameter{
							productName(),
						},
						Results: []definition.Result{definition.ErrorResult()},
					},
				},
				Description: "single-target Product APIs",
			},
		},
		Description: "all Product APIs",
	}
}

func productName() definition.Parameter {
	return definition.PathParameterFor("product", "name of the product")
}
