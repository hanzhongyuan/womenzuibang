import re, os, json, time
import threading, multiprocessing
from urllib.parse import quote_plus
import requests
import pymongo

headers={
    "Content-Type": "application/x-www-form-urlencoded; charset=UTF-8",
    "User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/76.0.3809.100 Safari/537.36",
    "X-Requested-With": "XMLHttpRequest",
    "Connection": "keep-alive",
    "Accept-Encoding": "gzip, deflate"
}

param = {
    "Param":"案件类型:{}案件,裁判日期:1900-01-01 TO 2019-08-10".format("刑事"),   
    "Index":1,
    "Page":20,
    "Order":"法院层级",
    "Direction":"asc"
}

url = "http://ws.gxcourt.gov.cn:28040/List/ListContent"

uri = "mongodb://{}:{}@localhost/gxwenshu".format(quote_plus("wenshu"), quote_plus("wenshu"))
client = pymongo.MongoClient(uri)

def getList(param, ):
    collection=client["gxwenshu"]["list"]       #mongodb client
    #session=requests.Session()

    response=requests.post(url,headers=headers,data=param)

    data=json.loads(re.sub('\t', '\\t', response.json()))

    listNum=int(data.pop(0)["Count"])
    indexL = int(listNum / 20)

    if indexL*20 < listNum:
        indexL += 1

    for index in range(1, indexL+1):      #if index > 100 -->>  data =[]
        param["Index"] = index

        try:
            response=requests.post(url,headers=headers,data=param)
        except Exception as e:
            print("index = {}, param = {}".format(index, param["Param"]))
            raise ValueError("Error: Network") from e

        data=json.loads(re.sub('\t','\\t',response.json()))
        
        if len(data)<1:
            print("index = {}, param = {}. No data!".format(index, param["Param"]))
        
        for x in range(1, len(data)):
            collection.insert_one(data[x])          #  db.list.insert()

def genParam():
    pass


if __name__ == "__main__":
    getList(param)