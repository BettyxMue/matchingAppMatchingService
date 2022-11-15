package mockData

import (
	"app/matchingAppMatchingService/common/dataStructures"
)

var MatchData = []dataStructures.Match{
	{Id: "1", User1: UserData[0], User2: UserData[1]},
	{Id: "2", User1: UserData[1], User2: UserData[2]},
	{Id: "3", User1: UserData[0], User2: UserData[2]},
}

var UserData = []dataStructures.User{
	{Id: 1, City: "Mannheim", Email: "babett.mueller@sap.com",
		First_name: "Babett", Name: "Müller", Password: "0000", Street: "Wörthfelder Weg", HouseNumber: "19",
		Username: "BettyxMue"},
	{Id: 2, City: "Homberg (Ohm)", Email: "jost-tomke-mueller@t-online.de",
		First_name: "Tomke", Name: "Müller", Password: "Test1234", Street: "Lichtenau", HouseNumber: "5",
		Username: "Seyna"},
	{Id: 3, City: "Mannheim", Email: "mathis.neunzig@sap.com",
		First_name: "Mathis", Name: "Neunzig", Password: "0000", Street: "Am Parkring", HouseNumber: "21",
		Username: "Flauuuschi"},
}
