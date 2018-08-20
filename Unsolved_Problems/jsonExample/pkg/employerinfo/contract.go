package employerinfo

import (
	"fmt"
	"database/sql"
)

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

	query:=fmt.Sprintf("    SELECT employer_id,date_of_contract,expiring_date,salary,emp_position FROM contract WHERE contract_number=%d",c.ContractNumber)
	err := db.QueryRow(query).Scan(&c.EmployerID,&c.HiredDate,&c.DueDate,&c.Salary,&c.PositionName)

	if err!=nil {
		return err
	}

	var p Position
	query=fmt.Sprintf("SELECT emp_position,description FROM position_info WHERE emp_position='%s'",c.PositionName)
	err= db.QueryRow(query).Scan(&p.Name,&p.Description)

	if err!=nil {
		return err
	}

	c.Position=&p

	return nil
}

func (c *Contract) Create(db *sql.DB) error {

	query := fmt.Sprintf("INSERT INTO contract(employer_id,date_of_contract,expiring_date,salary,emp_position) VALUES(%d,'%s','%s','%s','%s')",c.EmployerID,c.HiredDate,c.DueDate,c.Salary,c.PositionName)
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
	//fmt.Println("test")
	query:=fmt.Sprintf("    SELECT contract_number,employer_id,date_of_contract,expiring_date,salary,emp_position FROM contract")
	rows,err:=db.Query(query)

	//fmt.Println(rows)
	if err!=nil{
		//fmt.Println(err)
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