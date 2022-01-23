select Person.FirstName as firstName, Person.LastName as lastName, Address.City as city, Address.State as state
from Address
         right join Person on Address.PersonId = Person.PersonId;
