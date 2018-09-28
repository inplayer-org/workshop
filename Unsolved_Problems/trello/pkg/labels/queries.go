package labels

import (
	"database/sql"
	"repo.inplayer.com/workshop/Unsolved_Problems/trello/pkg/validators"

)

func (label *Label) Insert(DB *sql.DB)error{

	_,err := DB.Exec("INSERT INTO Trello.Labels (ID,IDboard,nameLabel,color) VALUES (?,?,?,?);",label.ID,label.IDboard,label.nameLabel,label.color)


	return err

}



func(label *Label) Update(DB *sql.DB)error{

	exists,err := validators.ExistsElementInColumn(DB,"Labels",label.ID,"ID")

	if err!= nil{
		return err
	}

	if exists{
		return label.updatelabel(DB)
	}


	return label.Insert(DB)

}


func (label *Label) updatelabel(DB *sql.DB)error{

	_,err := DB.Exec("UPDATE Trello.Labels SET IDboard=?,nameLabel=?,color=? WHERE ID=?",label.IDboard,label.nameLabel,label.color,label.ID)

	return err
}






// Returning string(label name from DB table Labels)
func GetLabelName(db *sql.DB,boardID string)(string,error){


	var labelName string

	err := db.QueryRow("SELECT nameLabel FROM Labels WHERE ID=?",boardID).Scan(&labelName)

	if err!=nil{
		return labelName,err
	}

	return labelName,nil

}

// Returns slice of all labels present in the database
func GetAllLabels(db *sql.DB) ([]Label, error) {

	var labels []Label
	var label Label

	rows, err := db.Query("SELECT ID,IDboard,nameLabel,color FROM Labels;")

	if err != nil {
		return labels, err
	}

	for rows.Next() {
		err := rows.Scan(&label.ID, &label.IDboard,&label.NameLabel,&label.Color)

		if err != nil {
			return labels, err
		}

		labels = append(labels, label)
	}

	return labels, nil
}

// Returning slice of label structure with boardid u get labelname and labelid
func GetLabelsFromBoard(db *sql.DB,boardID string)([]Label,error){

	var labels []Label

	rows,err:=db.Query("SELECT ID,nameLabel,color FROM Labels Where idBoard Like (?)","%"+boardID+"%")

	if err !=nil {
		return nil,err
	}

	for rows.Next(){


		var l Label
		err = rows.Scan(&l.ID,&l.NameLabel,&l.Color)

		if err !=nil {
			return nil,err
		}

		labels = append(labels,l)
	}


	return labels,nil
}