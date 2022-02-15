# ms-user-portal

# Error Codes for JSON Response- 
    -----------------------------------------------------------------------------------------
    | Error Codes               |  Description                                              |
    -----------------------------------------------------------------------------------------
    | 400(Bad Request)          |  If JSON request is not able to bind with object          |
    -----------------------------------------------------------------------------------------
    | 422(Unprocessable entity) |  If data inside JSON request is empty or is not valid     |
    -----------------------------------------------------------------------------------------
    | 200(status ok)            |  Response is successfull                                  |
    -----------------------------------------------------------------------------------------
    | 201(status created)       |  account created successfully                             |
    -----------------------------------------------------------------------------------------

# go run main.go -> this command will start the application locally

# Application Running Port - 4376

# Base URL - http://localhost:4376/ms-user-portal

# EndPoints - 
    # EndPoints - 
    ---------------------------------------------------------------------------------------------------------------------------------
    | EndPoints          | Request Type | Description       | URL                                                   | Headers       |
    ---------------------------------------------------------------------------------------------------------------------------------
    | /ping              |    GET       | test application  | http://localhost:4376/ms-user-portal/ping             | Not Applicable|
    ---------------------------------------------------------------------------------------------------------------------------------
    | /createaccount     |    POST      | create account    | http://localhost:4376/ms-user-portal/createaccount    | Not Applicable|
    ---------------------------------------------------------------------------------------------------------------------------------
    | /login             |    POST      | login into account| http://localhost:4376/ms-user-portal/login            | Not Applicable|
    ---------------------------------------------------------------------------------------------------------------------------------
    | /updateaccount/:id |    PATCH     | update account    | http://localhost:4376/ms-user-portal/updateaccount/:id| Authorization |
    ---------------------------------------------------------------------------------------------------------------------------------
    | /deleteaccount/:id |    DELETE    | delete account    | http://localhost:4376/ms-user-portal/deleteaccount/:id| Authorization |
    ---------------------------------------------------------------------------------------------------------------------------------
    | /logout/:id        |    POST      | logout account    | http://localhost:4376/ms-user-portal/logout/:id       | Authorization |
    ---------------------------------------------------------------------------------------------------------------------------------

# Json Request Object for endpoint -> /createaccount 
    {
        name     string 
        location    string 
        pan         string
        address     string
        contact     string
        sex         string
        nationality string
        userName    string
        password    string
    }
    Example - 
    {
        "name":"shubham bansal",
        "location": "Agra",
        "pan":"BKLPB2994Q",
        "address":"paramount floraville",
        "contact": "8448610134",
        "sex": "M",
        "nationality": "Indian",
        "userName": "shubham3869",
        "password":"test1234"
    }
    
# Json Response object for endpoint - > /createaccount 
        {
            data  interface{} 
            error {
                code int
                message string
            }      
        }

    Example
    1. If there is no error it means response is successful
        {
            "data": "account created successfully",
            "error": null
        }
    2. If there is any error
            {
                "data": null,
                "error": {
                    "code": 422,
                    "message": "Either pan, contact or user name already exists"
                }
            }
                                            
# END