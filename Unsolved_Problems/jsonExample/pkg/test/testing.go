package testing

import (
	"database/sql"
	"fmt"
)

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