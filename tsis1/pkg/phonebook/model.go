package api

type Person struct {
	ID       string `json:""`
	Name     string `json:""`
	LastName string `json:""`
	Phone    string `json:""`
}

var People = []Person{
	{"1", "Artur", "Kultyshev", "87770230743"},
	{"2", "Egor", "Petiahin", "87773187115"},
	{"3", "Nikita", "Nikitov", "87776284554"},
	{"4", "Artem", "Artemov", "87779428529"},
	{"5", "Artem", "Ni", "87773975405"},
	{"6", "Ivan", "Ivanov", "87770803713"},
	{"7", "Alexsander", "Alex", "87779134865"},
	{"8", "Leha", "Chi", "87776374059"},
	{"9", "Petr", "Perviy", "87770378158"},
	{"10", "Masha", "Pirogova", "87777836784"},
	{"11", "Sasha", "Bistrov", "87771823084"},
	{"12", "Vladislav", "baby don't hurt me", "87777751284"},
	{"13", "Natailya", "Morskaya pehota", "87770046654"},
	{"14", "Ira", "Bobr", "87776445109"},
	{"15", "Sveta", "Ivanova", "87777590505"},
	{"16", "Ksusha", "Baranova", "87771553212"},
}
