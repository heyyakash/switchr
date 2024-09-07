
# Switchr Introduction
**Switchr** is a **Dynamic Feature Flag Management Tool**.\
It is a powerful system built with NextJS and Golang, designed to streamline feature flag implementation in your development process. This tool allows for real-time feature flag management, project-based organization, and role-based access control (RBAC).

### Key features:

 - Project-based feature flag management
 - User management and RBAC
 - Real-time feature flag updates
 - REST API access
 - High-performance flag retrieval using Redis


# Getting Started
 
Switchr is fairly easy to use.

### Prerequisites
- Switchr Account

### Creating an account
In order to create a Switchr account, head to the [Signup page](/login).\
After successfull account creation you'll be redirected to the Dashboard.

**Before using the features of switchr you are required to verify your email once you sign up**

### Logging into your account
In order to login to your account, head to the [Login Page](/login).\
You can login through 2 methods - 
- Email / Password : Login directly using your email and password.
- Magic Links (Passwordless): Submit your email and if you have a switchr account you'll directly receive a login link valid for 5 minutes.


# Creating a new project

Creating a new project is fairly simple.\
\
**In order to create a new project, your email is needed to be verified first**

### Steps to create a new project
- Navigate to the [Dashboard Page](/dashboard).
- Click on the **Create New** button.
- Enter the name of your project. 
- Click on **Create Project** button.
- There you go. Your new project is created.

# Accessing Flags through REST API.
One of the main features of Switchr is accessing the feature flags through REST APIs.\
For security purposes in order to get the feature flag value from api, a **Token** has to be generated and sent through the headers.

### Steps to generate a auth token
- Navigate to [Dashboard Page](/dashboard).
- Open the project that has feature flag that needs to be accessed.
- Click on the **API** Button. 
- A dialog box will appear, Click on **Generate Token** Button.
- A token with validity of 120 days will be generated.\
Example Token - 
```
eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MzYwNzI3MDQsIlVpZCI6IiIsIlR5cGUiOiJhcGktdG9rZW4iLCJFbWFpbCI6IiIsIlJvbGUiOjAsIldasdasI6ImRkNjFmOTlkLWU4MGEtNDhlYS04ODA2LWNmN2VjNmNjNTM5NyJ9.WGpNELEtE-RswhiGAdBBmHO3Yh1IyOW8ykMMRWiDj_E
```

### Fetching data through API.
- In order to fetch data, following endpoint has to be used - 
```
GET https://dev-server.live:8020/api/get/<FLAG_NAME>
``` 

#### Params
- **FLAG_NAME** : Name of the feature flag, **Required**

#### Headers
- **token** : Authorization token generated above, **Required**

#### Example Request
Assume, requesting a flag named "create_enable" with its value as "true"
```
curl -X GET  http://localhost:8020/api/get/create_enable \
--header 'token:eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MzYwNzI3MDQsIlVpZCI6IiIsIlR5cGUiOiJhcGktdG9rZW4iLCdaskdjkasIiIsIlJvbGUiOjAsIlBpZCI6ImRkNjFmOTlkLWU4MGEtNDhlYS04ODA2LWNmN2VjNmNjNTM5NyJ9.WGpNELEtE-RswhiGAdBBmHO3Yh1IyOW8ykMMRWiDj_E'
```

### Example Success Response
``` json
{"flag":"true"} 
```
- Response body has a field called "flag" which has the value of the flag

### Example Failure Response
```
{"message":"Record not found","success":false}
```