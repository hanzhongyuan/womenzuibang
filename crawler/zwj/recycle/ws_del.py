import requests
import threading
import re, json
import pprint
from utils.docid.decrypt import decrypt_doc_id
from utils.docid.runeval import decrypt_runeval
from utils.document.parse import parse_detail
from utils.vl5x.args import Vjkl5, Vl5x, Number, Guid
from utils.wzws.decrypt import decrypt_wzws

data = {
  "Param": "案件类型:刑事案件",
  "Index": 2,
  "Page": 10,
  "Order": "法院层级",
  "Direction": "asc",
  "vl5x":None,
  "number":None,
  "guid":None
}


class Spider:
    def __init__(self, proxy=""):   #http://127.0.0.1:10809
        self.session = requests.Session()
        self.session.headers.update({
            "User-Agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/75.0.3770.100 Safari/537.36",
        })
        self.proxies = {
          "http": proxy,
          "https": proxy,
        }
        self.docIds =[]

    def list_page(self, data=data):
        url = "http://wenshu.court.gov.cn/List/ListContent"
        data["vl5x"] = Vl5x(self.session.cookies.setdefault("vjkl5", Vjkl5()))
        data["number"] = Number()
        data["guid"] = Guid()
        response = self.session.post(url, data=data, proxies=self.proxies)  # 请求1
        text = response.content.decode()

        if "请开启JavaScript并刷新该页" in text:
            dynamic_url = decrypt_wzws(text)
            response = self.session.post(dynamic_url, data=data, proxies = self.proxies)  # 请求2
        
        if response.status_code != 200:
            raise ValueError("HTTP {}".format(response.status_code))

        try:
            json_data = json.loads(response.json())
        except Exception as e:
            raise ValueError("获取List失败") from e

        runeval = json_data.pop(0)["RunEval"]
        try:
            key = decrypt_runeval(runeval)
        except ValueError as e:
            raise ValueError("返回脏数据") from e

        key = key.encode()
        for item in json_data:
            cipher_text = item["文书ID"]
            plain_text = decrypt_doc_id(doc_id=cipher_text, key=key)
            self.docIds.append(plain_text)

    def detail_page(self):
        url = "http://wenshu.court.gov.cn/CreateContentJS/CreateContentJS.aspx"
        for DocID in self.docIds:
            params = {
                "DocID": DocID,
            }
            response = self.session.get(url, params=params, proxies = self.proxies)  # 请求1
            text = response.content.decode()

            if "请开启JavaScript并刷新该页" in text:
                dynamic_url = decrypt_wzws(text)
                response = self.session.get(dynamic_url, proxies = self.proxies)  # 请求2

            group_dict = parse_detail(response.text)
            pprint(group_dict)
    
    def run(self):
        pass


if __name__ == "__main__":
    a=Spider()
    a.list_page()
    a.detail_page()