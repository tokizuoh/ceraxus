# ceraxus
  
Get and summarize iOS app reviews.
  

## Docker
  
### Version
  
```bash
> docker --version
Docker version 20.10.8, build 3967b7d

> docker-compose --version
docker-compose version 1.29.2, build 5becea4c
```
  
### Build (docker-compose build)
  
```bash
> make dc-build
```
  
### Run
  
```bash
> docker-compose exec app go run main.go 340368403
2022/01/08 08:48:07 SUCCESSED!

> cat tmp.csv | head -n 3
id,version,name,date,title,content
8196587123,2021.51.0,{USER_NAME},2022-01-02T10:01:57-07:00,趣...て,き...う
8187019395,2021.51.0,{USER_NAME},2021-12-30T20:27:10-07:00,初...け,ク...め
```