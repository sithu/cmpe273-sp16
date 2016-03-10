## 1. Automatic Class Attendance System

Automed time and attendance tracking system shall help save time and money by eliminating manual processes. With the system, teachers can more accurately and quickly see whether a student is in the class or not.

### Requirement

#### BLE Sensor
You need either one of the Bluetooth Low Energy (4.0) hardware for the project.
* 1. [Arduino BLE breakout](https://www.sparkfun.com/products/13632?_ga=1.256235395.9077466.1434957323)
* 2. Raspberry Pi with [BLE Adapter](  http://www.robotshop.com/en/bluetooth-40-console-adapter-raspberry-pi.html?gclid=Cj0KEQiAsP-2BRCFl4Lb2NTJttEBEiQAmj2tbR9CikBAc2FQlFk-BsD0JPNmiN_3svT176ZqvARI2eQaAi_G8P8HAQ)


#### Mobile App (Android or iOS)
You will also need to write some code in either iOS or Android to communite with BLE chip and your backend server.

#### Backend Server
The backend server will handle the business logic of tracking student attendance and will also have a set of REST APIs for moible app or Raspberry Pi client to communicate.

The backend server must use either NOSQL including key-value store or RDBMS to store the data.

One single web page that shows all student attendance for teachers is required.



