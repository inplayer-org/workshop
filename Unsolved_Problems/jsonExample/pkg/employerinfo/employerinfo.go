package employerinfo

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

func (e *Equipment) Get(db *sql.DB) error {

	query := fmt.Sprintf("SELECT employer_id,computer,monitor,mouse,keyboard,headset FROM equipment WHERE employer_id=%d",e.EmployerID)
	return db.QueryRow(query).Scan(&e.EmployerID,&e.Copmuters,&e.Monitors,&e.Mouses,&e.Keyboards,&e.Headsets)

}

func (e *Equipment) Create(db *sql.DB) error {

	query := fmt.Sprintf("INSERT INTO equipment(employer_id,computer,monitor,mouse,keyboard,headset) VALUES(%d,%d,%d,%d,%d,%d)",e.EmployerID,e.Copmuters,e.Monitors,e.Mouses,e.Keyboards,e.Headsets)
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
	var equipments []Equipment

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


type Position struct {

	Name string `json:"name"`
	Description string `json:"description"`

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
	var positions  []Position

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

type EmployerInfo struct {

	ID int
	FullName string
	Email string
	Gender string
	BirthDate string
	City string
	Country string
	Contracts *[]Contract
	Equipment *Equipment

}

func (e *EmployerInfo) Get(db *sql.DB) error {

	query:=fmt.Sprintf("SELECT fullname,email,gender,birth_date,city,country FROM employer_info WHERE employer_id=%d",e.ID)
	err:=db.QueryRow(query).Scan(&e.FullName,&e.Email,&e.Gender,&e.BirthDate,&e.City,&e.Country)

	if err!= nil {
		fmt.Println(err)
		return err
	}

	query=fmt.Sprintf("    SELECT contract_number,employer_id,date_of_contract,expiring_date,salary,emp_position FROM contract WHERE employer_id=%d",e.ID)
	rows,err:=db.Query(query)


	if err!= nil {
		fmt.Println(err)

		return err
	}

	contracts,err:=contractsForEmployer(db,	rows)
	e.Contracts=&contracts

	if err!= nil {
		fmt.Println(err)

		return err
	}

    query=fmt.Sprintf("SELECT computer,monitor,mouse,keyboard,headset FROM equipment where employer_id=%d",e.ID)
  //  &e.Equipment.EmployerID=&e.ID
    var eq Equipment
    eq.EmployerID=e.ID
	err= db.QueryRow(query).Scan(&eq.Copmuters,&eq.Monitors,&eq.Mouses,&eq.Keyboards,&eq.Headsets)

	e.Equipment=&eq

	if err!= nil {
		if err == sql.ErrNoRows {
			eq.Copmuters = 0
			eq.Monitors = 0
			eq.Mouses = 0
			eq.Keyboards = 0
			eq.Headsets = 0
		} else {
			return  err
		}
	}
	return nil
}

func contractsForEmployer(db *sql.DB,rows *sql.Rows)([]Contract,error){
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

func (e *EmployerInfo) Create(db *sql.DB)error {

	query:=fmt.Sprintf("INSERT INTO employer_info(fullname, email,gender,birth_date,city,country) VALUES('%s','%s','%s','%s','%s','%s')",e.FullName,e.Email,e.Gender,e.BirthDate,e.City,e.Country)
	_,err:=db.Exec(query)



	return err

	}

func (e *EmployerInfo) Update(db *sql.DB)error {

	query:=fmt.Sprintf("UPDATE employer_info SET firstname='%s',email='%s',gender='%s',birth_date='%s',city='%s',country='%s' WHERE employer_id=%d",e.FullName,e.Email,e.Gender,e.BirthDate,e.City,e.Country,&e.ID)
	_,err:=db.Exec(query)
	return err

}

func (e *EmployerInfo) Delete(db *sql.DB)error {

	query:=fmt.Sprintf("DELETE FROM contract WHERE employer_id=%d",e.ID)
	_,err:=db.Exec(query)

	if err!= nil {
		return err
	}

	query=fmt.Sprintf("DELETE FROM equipment WHERE employer_id=%d",e.ID)
	_,err=db.Exec(query)

	if err!= nil {
		return err
	}


	query=fmt.Sprintf("DELETE FROM employer_info WHERE employer_id=%d",e.ID)
	_,err=db.Exec(query)

	if err!= nil {
		return err
	}

	return err

}

func GetAllEmployers(db *sql.DB)([]EmployerInfo,error){

	query:=fmt.Sprintf("SELECT employer_id,fullname,email,gender,birth_date,city,country FROM employer_info")
	rows,err:=db.Query(query)

	if err!=nil {
		fmt.Println(err)
		return nil,err
	}

	defer rows.Close()

	return rowsToEmployers(db,rows)

}

func rowsToEmployers(db *sql.DB,rows *sql.Rows) ([]EmployerInfo, error) {

	var employers []EmployerInfo
	for rows.Next() {
		var e EmployerInfo
		err:=rows.Scan(&e.ID,&e.FullName,&e.Email,&e.Gender,&e.BirthDate,&e.City,&e.Country)

		if err != nil {
			//fmt.Println(err)
			return nil,err
		}

		query:=fmt.Sprintf("SELECT contract_number,employer_id,date_of_contract,expiring_date,salary,emp_position FROM contract WHERE employer_id=%d",e.ID)
  		crows,err:=db.Query(query)

  		defer crows.Close()

  		if err!= nil {
			//fmt.Println(err)
			return nil, err
		}

		contracts,err:=rowsToContracts(db,crows)

		if err!= nil {
			//fmt.Println(err)
			return nil, err
		}

		e.Contracts=&contracts

		var eq Equipment
		eq.EmployerID=e.ID
		query=fmt.Sprintf("SELECT computer,monitor,mouse,keyboard,headset FROM equipment WHERE employer_id=%d",eq.EmployerID)
		err=db.QueryRow(query).Scan(&eq.Copmuters,&eq.Monitors,&eq.Mouses,&eq.Keyboards,&eq.Headsets)
 		/*fmt.Println(eq)
		e.Equipment=&eq
		fmt.Println(e)
		fmt.Println(employers)*/
		if err!= nil {
			if err == sql.ErrNoRows {
				eq.Copmuters=0
				eq.Monitors=0
				eq.Mouses=0
				eq.Keyboards=0
				eq.Headsets=0
			} else {
				return nil,err
			}
		}


    	e.Equipment=&eq

		employers=append(employers,e)

	}

	return employers,nil
}
