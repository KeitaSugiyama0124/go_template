this is docker go template.
<br>
For people who want to touch for the time being
<br>
<br>
can use mac only.
<br>
It's a hassle, so all the packages are included.

# require.
・ Docker Desktop

# usage procedure
```
git clone https://github.com/KeitaSugiyama0124/go_template.git
cd go_template
docker-compose build
docker-compose up -d
```
please open browser and access 
<br>
http://localhost
<br>
<br>
could you can't access? check logs if building end.
<br>
```
docker logs -f [container id]
```

# using framework
gin

# Constitution
nginx -> fastCGI -> go -> gin -> sql

# if you add route?
go_work/default/config/routes/routes.go

# go.mod is where.
default project root directory only.
