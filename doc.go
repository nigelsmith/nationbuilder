/*
Package Nationbuilder provides an API client for the Nationbuilder remote API.

In order to use it a client must be instantiated with a nation slug and remote API key.

Create a nationbuilder client

	client, err := nationbuilder.NewNationbuilderClient("myNation", "myAPIKey")
	if err != nil {
		log.Fatal(err.Error())
	}

Call an endpoint

	people, result := client.GetPeople(nil)
	if result.HasError() {
		log.Fatal(result.Error())
	}

	for _, person := range people.Results {
		fmt.Println(person)
	}

Endpoint options

Some endpoints have specific options available to them - for example the people search endpoint.  In that case a special
PeopleSearchOptions type can be passed to the client API call.  In other cases where results can be paginated then a standard Options
object can be provided to the api call specifying the max number of items to return (the limit) and providing the page token and nonce
used to return a paginated resource.

If a resource supports pagination then it provides a Next and Prev method which will return a suitably configured Options object or nil
in the event that there are no more results.

	options, err := people.Next()
	if err != nil {
		log.Fatal(err.Error())
	}

	if options != nil {
		people, result = client.GetPeople(options)
	}

Results

Alongside the resource requested, endpoints return a results object with information on the status code encountered and any error
that was encountered.

Make sure to call:

	results.HasError()

to check whether an error has occurred.
*/
package nationbuilder
