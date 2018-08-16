package testing
import (
	"database/sql"
	"fmt"
)

type Equipment struct {

	EmployerID int `json:"employer_id"`
	Copmuters int `json:"copmuters"`
	Monitors int `json:"monitors"`
	Mouses int `json:"mouses"`
	Keyboards int `json:"keyboards"`
	Headsets int `json:"headsets"`

}
type Position struct {

	Name string `json:"name"`
	Description string `json:"description"`

}

type Contract struct {

	ContractNumber int `json:"contract_number"`
	EmployerID int `json:"employer_id"`
	HiredDate string `json:"hired_date"`
	DueDate string `json:"due_date"`
	Salary string `json:"salary"`
	PositionName string `json:"position_name"`
	Position *Position `json:"position"`

}

func (c *Contract) Get(db *sql.DB) error {

	query := fmt.Sprintf("SELECT contract.contract_number,contract.employer_id,contract.date_of_contract,contract.expiring_date,contract.salary,contract.emp_position,position_info.emp_position,position_info.description FROM contract INNER JOIN position_info ON contract.emp_position=position_info.emp_position AND contract.employer_id=%d",c.EmployerID)
	return db.QueryRow(query).Scan(&c.ContractNumber,&c.EmployerID,&c.HiredDate,&c.DueDate,&c.Salary,&c.PositionName,&c.Position.Name,&c.Position.Description)

}

func (c *Contract) Create(db *sql.DB) error {

	query := fmt.Sprintf("INSERT INTO contract(employer_id,date_of_contract,expiring_date,salary,emp_position) VALUES(%d,'%s','%s','%s','%s')",&c.EmployerID,c.HiredDate,c.DueDate,c.Salary,c.PositionName)
	_,err:= db.Exec(query)
	return err

}

func (c *Contract) Update(db *sql.DB) error {

	query := fmt.Sprintf("UPDATE contract SET employer_id=%d,date_of_contract='%s',expiring_date='%s',salary='%s',emp_position='%s'",&c.EmployerID,c.HiredDate,c.DueDate,c.Salary,c.PositionName)
	_,err:= db.Exec(query)
	return err

}

func (c *Contract) Delete(db *sql.DB) error{

	query:=fmt.Sprintf("DELETE FROM contract WHERE contract_number=%d",c.ContractNumber)
	_,err:=db.Exec(query)
	return err

}

func GetAllContracts(db *sql.DB) ([]Contract,error){
fmt.Println("test")
	query:=fmt.Sprintf("    SELECT contract_number,employer_id,date_of_contract,expiring_date,salary,emp_position FROM contract")
	rows,err:=db.Query(query)

fmt.Println(rows)
	if err!=nil{
		fmt.Println(err)
		return nil,err
	}

	defer rows.Close()

	return rowsToContracts(db,rows)

}

func rowsToContracts(db *sql.DB,rows *sql.Rows) ([]Contract, error) {

	var contracts []Contract

	for rows.Next() {
		var c Contract
		err:=rows.Scan(&c.ContractNumber,&c.EmployerID,&c.HiredDate,&c.DueDate,&c.Salary,&c.PositionName)

		if err!=nil {
			return nil,err
		}
		var p Position
		query:=fmt.Sprintf("SELECT emp_position,description FROM position_info WHERE emp_position='%s'",c.PositionName)
		err=db.QueryRow(query).Scan(&p.Name,&p.Description)

		c.Position=&p

		contracts=append(contracts,c)
	}

	return contracts,nil
}
func (p *Position) Get(db *sql.DB) error {

	query := fmt.Sprintf("SELECT emp_position, description FROM position_info WHERE emp_position='%s'",p.Name)
	return db.QueryRow(query).Scan(&p.Name,&p.Description)

}

func (p *Position) Create(db *sql.DB) error{

	query:=fmt.Sprintf("INSERT INTO position_info(emp_position,description) VALUES('%s','%s')",p.Name,p.Description)
	_,err:=db.Exec(query)
	return err

}

func (p *Position) Update(db *sql.DB) error{

	query:=fmt.Sprintf("UPDATE position_info SET description='%s' WHERE emp_position='%s'",p.Description,p.Name)
	_,err:=db.Exec(query)
	return err

}

func (p *Position) Delete(db *sql.DB) error{

	query:=fmt.Sprintf("DELETE FROM position_info WHERE emp_position='%s'",p.Name)
	_,err:=db.Exec(query)
	return err

}

func GetAllPositions(db *sql.DB)([]Position,error){

	query:=fmt.Sprintf("SELECT emp_position,description FROM position_info")
	rows,err:=db.Query(query)

	if err!=nil {
		return nil,err
	}

	defer rows.Close()

	return rowsToPositions(rows)
}

func rowsToPositions(rows *sql.Rows)([]Position,error){
	positions := []Position{}

	for rows.Next() {
		var p Position
		err:=rows.Scan(&p.Name,&p.Description)

		if err!=nil {
			return nil,err
		}

		positions=append(positions,p)
	}

	return positions,nil
}

func (e *Equipment) Get(db *sql.DB) error {

	query := fmt.Sprintf("SELECT employer_id,computer,monitor,mouse,keyboard,headset FROM equipment WHERE employer_id=%d",e.EmployerID)
	return db.QueryRow(query).Scan(&e.EmployerID,&e.Copmuters,&e.Monitors,&e.Mouses,&e.Keyboards,&e.Headsets)

}

func (e *Equipment) Create(db *sql.DB) error {

	query := fmt.Sprintf("INSERT INTO equipment(employer_id,computer,monitor,mouse,keyboard,headset) VALUES(%d,%d,%d,%d,%d,%d)",&e.EmployerID,&e.Copmuters,&e.Monitors,&e.Mouses,&e.Keyboards,&e.Headsets)
	_,err:= db.Exec(query)
	return err

}

func (e *Equipment) Update(db *sql.DB) error {
	//fmt.Println(e.Copmuters)
	query := fmt.Sprintf("UPDATE equipment SET computer=%d,monitor=%d,mouse=%d,keyboard=%d,headset=%d WHERE employer_id=%d",e.Copmuters,e.Monitors,e.Mouses,e.Keyboards,e.Headsets,e.EmployerID)
	_,err:= db.Exec(query)
	//fmt.Println(err)
	return err

}

func (e *Equipment) Delete(db *sql.DB) error{

	query:=fmt.Sprintf("DELETE FROM equipment WHERE employer_id=%d",e.EmployerID)
	_,err:=db.Exec(query)
	return err

}

func GetAllEquipments(db *sql.DB) ([]Equipment,error){

	query:=fmt.Sprintf("	SELECT employer_id,computer,monitor,mouse,keyboard,headset FROM equipment")
	rows,err:=db.Query(query)

	if err!=nil{
		return nil,err
	}

	defer rows.Close()

	return rowsToEquipments(rows)
}

func rowsToEquipments(rows *sql.Rows) ([]Equipment, error) {
	equipments := []Equipment{}

	for rows.Next() {
		var e Equipment
		err:=rows.Scan(&e.EmployerID,&e.Copmuters,&e.Monitors,&e.Mouses,&e.Keyboards,&e.Headsets)

		if err!=nil {
			return nil,err
		}

		equipments=append(equipments,e)
	}

	return equipments,nil
}

