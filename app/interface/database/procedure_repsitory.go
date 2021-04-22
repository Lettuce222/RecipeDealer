package database

import "github.com/Lettuce222/RecipeDealer/entity"

type ProcedureRepository struct {
	DbHandler DbHandler
}

var procedureTableName = "procedure_records"

func (repo *ProcedureRepository) Create(p *entity.Procedure) error {
	err := repo.DbHandler.Create(NewProcedureRecordFromProcedure(p))
	return err
}

func (repo *ProcedureRepository) Update(ingredient *entity.Procedure) error {
	// err := repo.DbHandler.Update(
	// 	ingredient.ID,
	// 	map[string]interface{}{"name": ingredient.Name, "number": ingredient.Number, "unit": ingredient.Unit},
	// 	procedureTableName,
	// )
	// return err
	return nil
}

func (repo *ProcedureRepository) Delete(identifier uint) error {
	err := repo.DbHandler.Delete(identifier, procedureTableName)
	return err
}

func (repo *ProcedureRepository) Show() ([]*entity.Procedure, error) {
	// rows, err := repo.DbHandler.Show(procedureTableName, "id, name, number, unit")
	// if err != nil {
	// 	return nil, err
	// }
	// defer rows.Close()

	// ingredients := []*entity.Procedure{}
	// for rows.Next() {
	// 	var id uint
	// 	var name string
	// 	var number float32
	// 	var unit string

	// 	err = rows.Scan(&id, &name, &number, &unit)
	// 	if err != nil {
	// 		return nil, err
	// 	}

	// 	ingredients = append(
	// 		ingredients,
	// 		entity.NewProcedure(id, name, number, unit),
	// 	)
	// }

	return nil, nil
}

func (repo *ProcedureRepository) Find(identifier uint) (*entity.Procedure, error) {
	// row := repo.DbHandler.Find(identifier, procedureTableName, "id, name, number, unit")

	// var id uint
	// var name string
	// var number float32
	// var unit string

	// err := row.Scan(&id, &name, &number, &unit)
	// if err != nil {
	// 	return nil, err
	// }

	// ingredient := entity.NewProcedure(id, name, number, unit)

	return nil, nil
}

type ProcedureRecord struct {
	Model
	Action            string
	Ingredients       []IngredientRecord
	InputProcedures   []ProcedureRecord `gorm:"foreignkey:ProcedureRecordId"`
	ProcedureRecordId *uint
}

func NewProcedureRecord(action string, ingredients []IngredientRecord, inputProcedures []ProcedureRecord) *ProcedureRecord {
	return &ProcedureRecord{
		Action:          action,
		Ingredients:     ingredients,
		InputProcedures: inputProcedures,
	}
}

func NewProcedureRecordFromProcedure(p *entity.Procedure) *ProcedureRecord {
	ingredientRecords := []IngredientRecord{}
	for _, i := range p.Ingredients {
		ingredientRecords = append(ingredientRecords, *NewIngredientRecord(
			i.Name,
			i.Number,
			i.Unit,
		))
	}

	procedureRecords := []ProcedureRecord{}
	for _, p := range p.InputProcedures {
		procedureRecords = append(procedureRecords, *NewProcedureRecordFromProcedure(&p))
	}

	return &ProcedureRecord{
		Action:          p.Action,
		Ingredients:     ingredientRecords,
		InputProcedures: procedureRecords,
	}
}
