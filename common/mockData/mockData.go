package mockData

import (
	"app/matchingAppMatchingService/common/dataStructures"
)

var MatchData = []dataStructures.Match{
	{Id: "1", UserId1: 1, UserId2: 2},
	{Id: "2", UserId1: 1, UserId2: 3},
	{Id: "3", UserId1: 2, UserId2: 3},
}

var UserData = []dataStructures.User{
	{Id: 1, City: "Mannheim", Email: "babett.mueller@sap.com",
		FirstName: "Babett", LastName: "Müller", Password: "0000", Street: "Wörthfelder Weg", HouseNumber: "19",
		Username: "BettyxMue"},
	{Id: 2, City: "Homberg (Ohm)", Email: "jost-tomke-mueller@t-online.de",
		FirstName: "Tomke", LastName: "Müller", Password: "Test1234", Street: "Lichtenau", HouseNumber: "5",
		Username: "Seyna"},
	{Id: 3, City: "Mannheim", Email: "mathis.neunzig@sap.com",
		FirstName: "Mathis", LastName: "Neunzig", Password: "0000", Street: "Am Parkring", HouseNumber: "21",
		Username: "Flauuuschi"},
}
