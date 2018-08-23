package update


/*
func GetAllClans(db sql.DB)(string, error) {

	var clans structures.Clan
	rows,err:=db.Query("SELECT (clanTag,clanName) FROM clans;")

	if err !=nil {
		return queries.ClansTable,err
	}

	count:=0
	for rows.Next(){
		err:=rows.Scan(&clans.structures.Clan[count].Tag,&clans.structures.Clan[count].Name)

		if err!=nil {
			return queries.ClansTable, err
		}

		count++

	}

	return queries.ClansTable,nil
}



func DailyUpdateClans(db *sql.DB)(Clan string){


	clans,err:=get.GetClans()
	if err!=nil{
		return clans
	}

	err=UpdateClans(db,clans)

	if err!=nil{
		return clans
	}

	return clans

}
*/