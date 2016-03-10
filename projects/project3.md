## 3. Server-driven SLA 

In general, server defines a set of SLA and client decides (HTTP) connection settings based on its own required. Instead, the server will tell each client about what the ideal connection settings should be for a given time window. 

### Requirement

#### Server side
You will be building a server side component that monitors the system health and client load. Based on that data, the server will suggest the best configuration setting for clients in the HTTP response. 

The solution must support HTTP and the other protocols are optional.

You need a set of web pages to render the status of the system for monitoring purpose meaning some REST APIs to consume from JavaScript loaded from the pages.

#### Client side
You will need a demo code that uses the configration settings returned from the server.



