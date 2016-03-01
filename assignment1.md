## Assignment 1 

You will be building a simple RESTful service to manage personal perferences that can be shared with service providers like travel agencies so that they can provide personalization in their services.
> No data persistence to a permanent storage is required at the point but you need to store the application state in memory only.

### APIs
1. Create new user profile.

__URL: POST /profile__

***Request***
```json
{
    "email": "foo@gmail.com",
    "zip": "94112",
    "country": "U.S.A",
    "profession": "student",
    "favorite_color": "blue",
    "is_smoking": "yes|no",
    "favorite_sport": "hiking",
    "food": {
        "type": "vegetrian|meat_eater|eat_everything",
        "drink_alcohol": "yes|no"
    },
    "music": {
        "spotify_user_id": "wizzler"
    },
    "movie": {
        "tv_shows": ["x", "y", "z"],
        "movies": ["x", "y", "z"]
    },
    "travel": {
        "flight": {
            "seat": "aisle|window"            
        }
    }
}
``` 
***Response***
```sh
HTTP/1.1 201 Created
Date: Mon, 29 Feb 2016 19:55:15 GMT
....
``` 
---
2. Get user profile.

__URL: GET /profile/{email}__

***Response***
```sh
HTTP/1.1 200 OK
Date: Mon, 29 Feb 2016 19:55:15 GMT
....
``` 

```json
{
    "email": "foo@gmail.com",
    "zip": "94112",
    "country": "U.S.A",
    "profession": "student",
    "favorite_color": "blue",
    "is_smoking": "yes",
    "favorite_sport": "hiking",
    "food": {
        "type": "meat_eater",
        "drink_alcohol": "no"
    },
    "music": {
        "spotify_user_id": "wizzler"
    },
    "movie": {
        "tv_shows": ["x", "y", "z"],
        "movies": ["x", "y", "z"]
    },
    "travel": {
        "flight": {
            "seat": "aisle|window"            
        }
    }
}
``` 
---
3. Update existing user profile.

__URL: PUT /profile/{email}__

***Request***
```json
{
    "travel": {
        "flight": {
            "seat": "window"            
        }
    }
}
``` 

*** Response ***
```sh
HTTP/1.1 204 No Content
Date: Mon, 29 Feb 2016 19:55:15 GMT
....
``` 
---
4. Delete user profile.

__URL: DELETE /profile/{email}__

***Response***
```sh
HTTP/1.1 204 No Content
Date: Mon, 29 Feb 2016 19:55:15 GMT
....
``` 
---


