package employerinfo

import (
	"database/sql"
	"fmt"
	"repo.inplayer.com/workshop/Unsolved_Problems/jsonExample/pkg/errorhandle"
)

type EmployerInfo struct {

	ID int `json:"id"`
	FullName string `json:"full_name"`
	Email string `json:"email"`
	Gender string `json:"gender"`
	BirthDate string `json:"birth_date"`
	City string `json:"city"`
	Country string `json:"country"`
	Contracts *[]Contract `json:"contracts"`
	Equipment *Equipment `json:"equipment"`

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

func inputChecks(e EmployerInfo,contracts []Contract) error {

	err:=stringChecks(e)

	if err != nil {
		return err
	}

	err= errorhandle.CheckSalary(contracts[0].Salary)

	if err!=nil {
		return err
	}

	return nil
}

func (e *EmployerInfo) Create(db *sql.DB)error {

	contracts:=*e.Contracts

	err:=inputChecks(*e,contracts)

	if err!=nil {
		return err
	}

	var count int
	query := fmt.Sprintf("SELECT COUNT(emp_position) FROM position_info WHERE emp_position='%s'",contracts[0].PositionName)
	err= db.QueryRow(query).Scan(&count)

	if err != nil {
		return err
	}

	if count == 0 {
		return sql.ErrNoRows
	}

	query=fmt.Sprintf("INSERT INTO employer_info(employer_id,fullname, email,gender,birth_date,city,country) VALUES(%d,'%s','%s','%s','%s','%s','%s')",e.ID,e.FullName,e.Email,e.Gender,e.BirthDate,e.City,e.Country)
	_,err=db.Exec(query)

	if err != nil {
		return err
	}

	query = fmt.Sprintf("INSERT INTO contract(employer_id,date_of_contract,expiring_date,salary,emp_position) VALUES(%d,'%s','%s','%s','%s')",e.ID,contracts[0].HiredDate,contracts[0].DueDate,contracts[0].Salary,contracts[0].PositionName)
	_,err= db.Exec(query)

	if err != nil {
		return err
	}

	eq:=*e.Equipment
	query = fmt.Sprintf("INSERT INTO equipment(employer_id,computer,monitor,mouse,keyboard,headset) VALUES(%d,%d,%d,%d,%d,%d)",e.ID,eq.Copmuters,eq.Monitors,eq.Mouses,eq.Keyboards,eq.Headsets)
	_,err= db.Exec(query)
	return err

	}

	func stringChecks(e EmployerInfo)error{
		err:= errorhandle.CheckString(&e.FullName)

		if err != nil {
			return err
		}

		err = errorhandle.CheckString(&e.Country)

		if err != nil {
			return err
		}
		err = errorhandle.CheckString(&e.City)

		if err != nil {
			return err
		}
		err = errorhandle.CheckEmail(e.Email)

		if err != nil {
			return err
		}
		return nil
	}

func (e *EmployerInfo) Update(db *sql.DB)error {
	err:=stringChecks(*e)

	if err != nil {
		return err
	}

	query:=fmt.Sprintf("UPDATE employer_info SET fullname='%s',email='%s',gender='%s',birth_date='%s',city='%s',country='%s' WHERE employer_id=%d",e.FullName,e.Email,e.Gender,e.BirthDate,e.City,e.Country,e.ID)
	_,err =db.Exec(query)
	return err

}

func (e *EmployerInfo) Delete(db *sql.DB)error {

	query:=fmt.Sprintf("DELETE FROM contract WHERE employer_id=%d",e.ID)
	_,err:=db.Exec(query)

	if err!= nil {
		//fmt.Println(err)
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
