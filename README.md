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
## Tools used to create this project
* Built in golang
* The CLI is built using cobra-cli
* TUI is bubble tea and lipgloss
