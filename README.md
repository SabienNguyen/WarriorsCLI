# WarriorsCLI
WarriorsCli is an unofficial command line interface for the Warriors basketball team
## Features
* Get the game schedule for the warriors
* Get the updated roster of the warriors
* More to be added in future updates!
## Installation
```
git clone https://github.com/SabienNguyen/WarriorsCLI.git && warriors-cli

#install
go install

#run
warriors-cli roster
```
> When building make sure $GOPATH/bin is in your path environment variables
## Usage
```
warriors-cli roster
```
Use the arrows keys to naviage the tables
## Examples
![image](https://github.com/user-attachments/assets/1fe4f060-6aa1-4329-b9dd-dfa54834b3a3)
![image](https://github.com/user-attachments/assets/8159c713-f4c6-499b-a6cd-353ea3a14cfb)

## Tools used to create this project
* Built in golang
* The CLI is built using cobra-cli
* TUI is bubble tea and lipgloss
* My custom [WAPI](https://github.com/SabienNguyen/WAPI) is used to get team information!
