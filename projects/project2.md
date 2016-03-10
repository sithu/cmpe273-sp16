## 2. Pi  Cloud

Pi Cloud is a Raspberry Pi based PaaS service that can add or remove the nodes in real-time without restarting the server.

### Requirement

#### Raspberry Pi
You will need at least two Raspberry Pi nodes to build the system. One is for the main (Master) controller and the other is for a worker node.

#### The Main Controller
The whole project is to build this main controller which is the brain of the Pi Cloud. Its features are:

1. Register the nodes
2. Download the application from Github
2. Install runtime and dependency of the application packages (E.g. go get xxxxx, for GO)
3. Launch the application
4. Add the new route to the routing table.
5. Route the HTTP traffic to the node based on the URL path (E.g. http://x.y.z/app1 -> http://1.2.3.4:5000)




