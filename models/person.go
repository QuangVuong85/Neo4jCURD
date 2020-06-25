package models

import "errors"

func GetPerson(personId string) (per *Person, err error)  {
	per, err = getNodePerson(personId)
	if err == nil {
		return per, nil
	}

	return per, errors.New("Person not exists")
}

func GetAllPersons() map[string]*Person {
	return getAllNodePersons()
}

func UpdatePerson(personId string, p *Person) (person *Person, err error) {
	status, e := editNodePerson(personId, p)

	if e != nil {
		return nil, e
	}

	if status == "true" {
		return p, nil
	}

	return nil, errors.New("Person not exists")
}

func DeletePerson(personId string) string {
	_, err := getNodePerson(personId)
	if err != nil {
		return "personId not exists"
	}

	err = deleteNodePerson(personId)
	return "Delete Person by personId = " + personId + " successed"
}

func AddPerson(reqperson *ReqPerson) string  {
	_, err := getNodePerson(reqperson.Id)

	if err != nil {
		return "PersonId exists"
	}

	err = addNodePerson(reqperson)
	if err != nil {
		return reqperson.Id
	}

	return reqperson.Id
}
