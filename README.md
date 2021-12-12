#Mikrus CLI

It's just simple cli for [mikr.us](https://mikr.us/)

###How to install
1. Install [GO](https://go.dev/doc/install)
2. Clone this repository
3. In same folder where you have main.go execute

``go build -o mikruscli``

4. (optional)Add file to PATH

###How to use
It's required to setup two environment variables. 

``MIKRUS_SRV`` - server name 

``MIKRUS_KEY`` - api key from [control panel](https://mikr.us/panel/?a=api)

With cli added to PATH

``mikruscli <module>``

Without cli added to PATH you have to be in same folder where your mikruscli file is located

``./mikruscli <module>``

###Available modules

**info**

Returns information about server

``mikruscli info``

**serwery**

Returns list of servers

``mikruscli serwery``

**restart**

Restarts your server

``mikruscli restart``

**logs**

Returns last 10 logs

``mikruscli logs``

**log by id**

Returns log by id

``mikruscli logs <id>``

**amfetamina**

Boost for your server

``mikruscli amfetamina``

**db**

Returns access data to databases 

``mikruscli db``

**exec**

Execute command

``mikruscli exec <command>``

**stats**

Returns server stats

``mikruscli stats``

**porty**

Returns available TCP/UDP ports

``mikruscli stats``