package query

import (
	"github.com/graphql-go/graphql"
	"github.com/jackwind9000/graphql/datatype"
	"github.com/jackwind9000/graphql/model"
)

// Employee get employee information by employee
var Employee = &graphql.Field{
	Type:        datatype.Employee,
	Description: "Get employee information",
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		id, _ := p.Args["id"].(string)

		employee, err := getEmployeeByID(id)
		if err != nil {
			return model.Employee{}, err
		}

		address, err := getAddressByEmployeeID(id)
		if err != nil {
			return model.Employee{}, err
		}

		employee.Address = address
		return employee, nil
	},
}

func getEmployeeByID(id string) (model.Employee, error) {
	return model.Employee{
		ID:       id,
		Name:     "Jack Wind",
		Age:      26,
		Position: "Sale",
	}, nil
	// return model.Employee{}, errors.New("employee not found")
}

func getAddressByEmployeeID(id string) (model.Address, error) {
	return model.Address{
		Ward:     "1",
		District: "2",
		Province: "HCMC",
		Country:  "VN",
	}, nil
}
