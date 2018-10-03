package client

import (
	"database/sql"
	"fmt"
	"repo.inplayer.com/workshop/Unsolved_Problems/trello/pkg/members"
	"repo.inplayer.com/workshop/Unsolved_Problems/trello/pkg/interfaces"
)

func UpdateMember(c interfaces.ClientInterface, DB *sql.DB,memberID string)error{

	fmt.Println("start geting member")

	member,err:=c.GetMember(memberID)



	if err!=nil{
		return err
	}

	fmt.Println("finished geting member")
	fmt.Println(member)
	fmt.Println("start updating member")

	err=member.Update(DB)

	if err!=nil{
		return err
	}

	fmt.Println("finished updating member")

	fmt.Println("looping in boards id")

	ms:=members.DataStructureToMember(member)

	for _,IDboard:=range ms.IDboards{

		fmt.Println("start geting big boards")

		bb,err:=c.BigBoardRequest(IDboard)

		if err!=nil{
			return err
		}

		fmt.Println("finished geting big boards")
		fmt.Println(bb)
		fmt.Println("start updating big boards")

		bb.Update(DB)

		fmt.Println("finished updating big boards")

	}

	fmt.Println("finished")


	return nil

}
